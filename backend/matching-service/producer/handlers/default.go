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
	requestBody := models.MatchRequest{}
	err := json.NewDecoder(c.Request().Body).Decode(&requestBody)
	if err != nil {
		fmt.Println("[MatchHandler] Error decoding request body")
		return err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()
	var channel *amqp.Channel
	if curr, ok := rmq.OpenChannelsMap[utils.MatchCriteria(requestBody.MatchCriteria)]; ok {
		channel = curr
	} else {
		return c.JSON(http.StatusBadRequest, "Unknown matching criteria")
	}

	msgPacket := models.MessageQueueRequestPacket{
		RequestBody: requestBody,
	}

	serialPkt, err := json.Marshal(msgPacket)
	if err != nil {
		fmt.Println("[MatchHandler] Error marshalling message packet")
		return err
	}

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

	messages, err := rmq.ResultChannel.Consume(
		"results",            // queue name
		requestBody.Username, // consumer
		true,                 // auto-ack
		false,                // exclusive
		false,                // no local
		false,                // no wait
		nil,                  // arguments
	)
	if err != nil {
		log.Println(err)
	}

	// Counts 30 seconds
	// TODO set to 3 seconds for testing
	ctxTimer, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	syncChan := make(chan int)
	resChan := make(chan string)

	go func(syncChan chan int) {
		for {
			select {
			case syncChan <- 1:
				fmt.Println("Hoho time out")
				return
			case message := <-messages:
				err := rmq.ResultChannel.Cancel(requestBody.Username, true)
				if err != nil {
					return
				}
				if err != nil {
					fmt.Println(err)
				}

				log.Printf(" > Received message: %s\n", message.Body)
				resChan <- string(message.Body)
				return
			}
		}
	}(syncChan)

	shouldBreak := false

	for {
		select {
		case <-ctxTimer.Done():
			<-syncChan
			shouldBreak = true
			break
		case res := <-resChan:
			return c.JSON(http.StatusOK, res)
		}
		if shouldBreak {
			break
		}
	}

	return c.JSON(http.StatusOK, "End of 3 seconds!")
}
