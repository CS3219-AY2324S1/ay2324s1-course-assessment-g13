package handlers

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/labstack/echo/v4"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"net/http"
	"producer/models"
	"producer/rmq"
	"producer/utils"
	"time"
)

func MatchHandler(c echo.Context) error {
	// Destructures the API request body into our model
	requestBody := models.MatchRequest{}
	err := json.NewDecoder(c.Request().Body).Decode(&requestBody)
	if err != nil {
		msg := fmt.Sprintf("[MatchHandler] Error decoding request body | err: %v", err)
		log.Println(msg)
		return err
	}
	utils.ResetUser(requestBody.Username)
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	// Retrieves the appropriate channel to publish the user into
	var channel *amqp.Channel
	if curr, ok := rmq.OpenChannelsMap[utils.MatchCriteria(requestBody.MatchCriteria)]; ok {
		channel = curr
	} else {
		msg := fmt.Sprintf("[MatchHandler] Criteria to match is unknown | ok: %v", ok)
		log.Println(msg)
		return c.JSON(http.StatusBadRequest, "Unknown matching criteria")
	}

	msgPacket := models.MessageQueueRequestPacket{
		RequestBody: requestBody,
	}

	// Marshals the packet to be published into the MQ
	serialPkt, err := json.Marshal(msgPacket)
	if err != nil {
		msg := fmt.Sprintf("[MatchHandler] Error marshalling message packet | err: %v", err)
		log.Println(msg)
		return err
	}

	// Publishes the user request into the selected MQ
	err = channel.PublishWithContext(
		ctx,
		"",                        // exchange
		requestBody.MatchCriteria, // queue name
		false,                     // mandatory
		false,                     // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        serialPkt, // message to publish
		},
	)
	if err != nil {
		msg := fmt.Sprintf("[MatchHandler] Error publishing message | err: %v", err)
		fmt.Println(msg)
		return err
	}

	// Producer goroutine now consumes from their own results channel
	resultChan, err := rmq.Conn.Channel()
	if err != nil {
		msg := fmt.Sprintf("[MatchHandler] Error creating unique result channel | err: %v", err)
		log.Println(msg)
		return err
	}
	defer resultChan.Close()

	// Declare unique result queue
	resultQueue, err := resultChan.QueueDeclare(
		utils.ConstructResultChanIdentifier(requestBody.Username),
		false,
		false,
		false,
		false,
		nil,
	)

	// Starts consuming from the unique result queue
	messages, err := resultChan.Consume(
		resultQueue.Name, // queue name
		"",               // consumer
		true,             // auto-ack
		false,            // exclusive
		false,            // no local
		false,            // no wait
		nil,              // arguments
	)
	if err != nil {
		msg := fmt.Sprintf("[MatchHandler] Error consuming from results channel | err: %v", err)
		log.Println(msg)
		return err
	}

	// TODO set to 3 seconds for testing, reset to 30 for actual implementation
	// Counts 30 seconds
	ctxTimer, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	syncChan := make(chan int)   // Used to break consumer goroutine once timeout hits
	resChan := make(chan string) // Used to pass result from consumer goroutine to main thread

	// Spins off consumer goroutine to listen to results channel
	go func(syncChan chan int) {
		for {
			select {
			case syncChan <- 1:
				return
			case message := <-messages:
				packetResponse := models.MessageQueueResponsePacket{}
				err := json.Unmarshal(message.Body, &packetResponse)
				if err != nil {
					msg := fmt.Sprintf("[MatchHandler] Error unmarshalling response packet | err: %v", err)
					log.Println(msg)
					return
				}
				matchedUser := packetResponse.ResponseBody.MatchUser
				// Check if matched user is already out of queue
				if utils.IsUserCancelled(matchedUser) {
					// Publishes the user request into the selected MQ
					err = channel.PublishWithContext(
						ctx,
						"",                        // exchange
						requestBody.MatchCriteria, // queue name
						false,                     // mandatory
						false,                     // immediate
						amqp.Publishing{
							ContentType: "text/plain",
							Body:        serialPkt, // message to publish
						},
					)
					if err != nil {
						msg := fmt.Sprintf("[MatchHandler] Error publishing message | err: %v", err)
						fmt.Println(msg)
						return
					}
					continue
				} else {
					resChan <- packetResponse.ResponseBody.MatchUser
					return
				}
			}
		}
	}(syncChan)

	shouldBreak := false // Flag to aid breaking out of for-select loop

	// Loops infinitely until context timer is hit, or result is returned from consumer, whichever occurs first
	for {
		select {
		case <-ctxTimer.Done():
			utils.CancelUser(requestBody.Username)
			<-syncChan
			shouldBreak = true
			break
		case res := <-resChan:
			utils.PrintCancelledUsers()
			return c.JSON(http.StatusOK, res)
		}
		if shouldBreak {
			break
		}
	}

	return c.JSON(http.StatusOK, "Match not found within 3 seconds.")
}
