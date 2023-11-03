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
)

var UserToChanMap map[string]chan bool

func UserCancelHandler(c echo.Context) error {
	// Destructures the API request body into our model
	requestBody := models.CancelRequest{}
	err := json.NewDecoder(c.Request().Body).Decode(&requestBody)
	if err != nil {
		msg := fmt.Sprintf("[UserCancelHandler] Error decoding request body | err: %v", err)
		log.Println(msg)
		return err
	}

	msgToFanout := models.MessageQueueCancelRequestPacket{
		RequestBody: requestBody,
	}

	// Marshals the packet to be fanned out into the various producer cancel MQs
	serialPkt, err := json.Marshal(msgToFanout)
	if err != nil {
		msg := fmt.Sprintf("[MatchHandler] Error marshalling fanout message packet | err: %v", err)
		log.Println(msg)
		return err
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	err = rmq.CancelChannel.PublishWithContext(
		ctx,
		"cancels",
		"",    // routing key
		false, // mandatory
		false, // immediate
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        serialPkt,
		},
	)
	if err != nil {
		msg := fmt.Sprintf("[MatchHandler] Error fanning out message | err: %v", err)
		log.Println(msg)
		return err
	}

	// Tells producer that this user is cancelled
	UserToChanMap[requestBody.Username] <- true

	// On the own producer, indicate user has cancelled
	utils.CancelUser(requestBody.Username)

	cancelResponseBody := models.CancelResponse{CancelStatus: true}
	return c.JSON(http.StatusOK, cancelResponseBody)
}

func Init() {
	UserToChanMap = make(map[string]chan bool)
	go func() {
		messages, err := rmq.CancelChannel.Consume(
			"cancelQueue", // queue
			"",            // consumer
			true,          // auto-ack
			false,         // exclusive
			false,         // no-local
			false,         // no-wait
			nil,           // args
		)
		if err != nil {
			msg := fmt.Sprintf("[Init] Error consuming from cancel queue | err: %v", err)
			log.Println(msg)
			panic(err)
		}

		// Starts consuming from cancel channel
		for msg := range messages {
			// If message is received, means cancel the user
			var recvdPkt models.MessageQueueCancelRequestPacket
			err := json.Unmarshal(msg.Body, &recvdPkt)
			if err != nil {
				msg := fmt.Sprintf("[Init] Error unmarshalling cancel packet into struct | err: %v", err)
				log.Println(msg)
				panic(err)
			}
			log.Println("Cancelling user after message consumption!!!")
			utils.CancelUser(recvdPkt.RequestBody.Username)
		}
	}()
}
