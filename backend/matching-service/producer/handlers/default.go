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
	"strings"
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

	// Retrieves the appropriate channel to publish the user into
	var criteriaChannel *amqp.Channel
	if curr, ok := rmq.OpenChannelsMap[utils.MatchCriteria(strings.ToLower(requestBody.MatchCriteria))]; ok {
		criteriaChannel = curr
	} else {
		msg := fmt.Sprintf("[MatchHandler] Criteria to match is unknown | ok: %v", ok)
		log.Println(msg)
		return c.JSON(http.StatusBadRequest, "Unknown matching criteria")
	}

	// Removes the user from our cancel buffer if they have previously tried to match and got cancelled
	utils.ResetUser(requestBody.Username, requestBody.MatchCriteria)

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	var lengthChannel *amqp.Channel
	if curr, ok := rmq.LengthChannelsMap[utils.MatchCriteria(strings.ToLower(requestBody.MatchCriteria))]; ok {
		lengthChannel = curr
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
	err = criteriaChannel.PublishWithContext(
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

	lengthExchangeName := "length" + string(requestBody.MatchCriteria)

	lengthMsgPacket := models.MessageQueueLengthRequest{
		Increment:     1,
		MatchCriteria: requestBody.MatchCriteria,
	}

	serialLengthPkt, err := json.Marshal(lengthMsgPacket)
	if err != nil {
		msg := fmt.Sprintf("[MatchHandler] Error marshalling length packet | err: %v", err)
		log.Println(msg)
		return err
	}

	// Publishes size of 1 into criteria length channel for consumer to see
	err = lengthChannel.PublishWithContext(
		ctx,
		lengthExchangeName,
		"",
		false,
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        serialLengthPkt,
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

	// Counts 30 seconds
	ctxTimer, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	syncChan := make(chan int)                 // Used to break consumer goroutine once timeout hits
	resChan := make(chan models.MatchResponse) // Used to pass result from consumer goroutine to main thread

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
				if utils.IsUserCancelled(matchedUser, requestBody.MatchCriteria) {
					log.Printf("Matched User %s is already out of queue, publishing %s back into %s queue\n", matchedUser, requestBody.Username, requestBody.MatchCriteria)
					// Publishes the user request into the selected MQ
					err = criteriaChannel.PublishWithContext(
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
				} else if utils.IsUserCancelled(requestBody.Username, requestBody.MatchCriteria) {
					// If user is already cancelled, cancel the timer
					log.Printf("User is already cancelled: %s\n", requestBody.Username)
					cancel()
				} else {
					// If matched user is valid, return matched user
					resChan <- packetResponse.ResponseBody
					return
				}
			}
		}
	}(syncChan)

	shouldBreak := false // Flag to aid breaking out of for-select loop

	var matchResponseBody models.MatchResponse

	// Loops infinitely until context timer is hit, or result is returned from consumer, whichever occurs first
	userCancelChan := make(chan bool)
	UserToChanMap[requestBody.Username] = userCancelChan
	for {
		select {
		// User cancelled, so terminate listener
		case <-userCancelChan:
			log.Printf("User %s manually cancelled on producer side\n", requestBody.Username)
			// Remove user from queue
			utils.CancelUser(requestBody.Username, requestBody.MatchCriteria)
			<-syncChan // Reads from sync channel to allow goroutine listening to result to break out of loop
			shouldBreak = true
			break
		// 30 seconds timer hit
		case <-ctxTimer.Done():
			log.Printf("30 seconds timer hit on producer side for %s\n", requestBody.Username)
			// Remove user from queue
			lengthMsgPacket := models.MessageQueueLengthRequest{
				Increment:     -1,
				MatchCriteria: requestBody.MatchCriteria,
			}
			serialLengthPkt, err := json.Marshal(lengthMsgPacket)
			if err != nil {
				msg := fmt.Sprintf("[Init | Cancel] Error marshalling length packet | err: %v", err)
				log.Println(msg)
				panic(err)
			}
			// Publishes size of 1 into criteria length channel for consumer to see
			err = lengthChannel.PublishWithContext(
				context.Background(),
				lengthExchangeName,
				"",
				false,
				false,
				amqp.Publishing{
					ContentType: "text/plain",
					Body:        serialLengthPkt,
				},
			)
			if err != nil {
				msg := fmt.Sprintf("[MatchHandler] Error publishing message | err: %v", err)
				fmt.Println(msg)
				panic(err)
			}
			utils.CancelUser(requestBody.Username, requestBody.MatchCriteria)
			<-syncChan // Reads from sync channel to allow goroutine listening to result to break out of loop
			shouldBreak = true
			break
		case res := <-resChan:
			log.Printf("Found a match for current user with: %s\n", res.MatchUser)
			matchResponseBody = models.MatchResponse{
				MatchUser:    res.MatchUser,
				MatchStatus:  1,
				RoomId:       res.RoomId,
				ErrorMessage: "",
			}
			lengthMsgPacket := models.MessageQueueLengthRequest{
				Increment:     -1,
				MatchCriteria: requestBody.MatchCriteria,
			}
			serialLengthPkt, err := json.Marshal(lengthMsgPacket)
			if err != nil {
				msg := fmt.Sprintf("[Init | Cancel] Error marshalling length packet | err: %v", err)
				log.Println(msg)
				panic(err)
			}
			// Publishes size of 1 into criteria length channel for consumer to see
			err = lengthChannel.PublishWithContext(
				context.Background(),
				lengthExchangeName,
				"",
				false,
				false,
				amqp.Publishing{
					ContentType: "text/plain",
					Body:        serialLengthPkt,
				},
			)
			if err != nil {
				msg := fmt.Sprintf("[MatchHandler] Error publishing message | err: %v", err)
				fmt.Println(msg)
				panic(err)
			}
			return c.JSON(http.StatusOK, matchResponseBody)
		}
		if shouldBreak {
			break
		}
	}

	matchResponseBody = models.MatchResponse{
		MatchUser:    "",
		MatchStatus:  0,
		RoomId:       "",
		ErrorMessage: "Match not found within 30 seconds.",
	}

	return c.JSON(http.StatusOK, matchResponseBody)
}
