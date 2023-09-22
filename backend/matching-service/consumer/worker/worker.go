package worker

import (
	"consumer/models"
	"consumer/rmq"
	"consumer/utils"
	"context"
	"encoding/json"
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"time"
)

func SpinMQConsumer(criteria utils.MatchCriteria) {
	go func() {
		var channel *amqp.Channel
		// Declare message queues to ensure it exists
		if curr, ok := rmq.OpenChannelsMap[criteria]; ok {
			channel = curr
		} else {
			fmt.Println("[SpinMQConsumer] Error getting open channel")

		}

		messages, err := channel.Consume(
			string(criteria), // queue name
			"",               // consumer
			true,             // auto-ack
			false,            // exclusive
			false,            // no local
			false,            // no wait
			nil,              // arguments
		)
		if err != nil {
			fmt.Println("[SpinMQConsumer] Error consuming from channel")
		}

		ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
		defer cancel()

		for message := range messages {
			// For example, show received message in a console.
			res := models.MessageQueueRequestPacket{}
			err := json.Unmarshal(message.Body, &res)
			if err != nil {
				fmt.Println(err)
			}
			log.Printf(" > Received message: %s\n", res)

			pubMsg := fmt.Sprintf("Consumer received msg from %s", res.RequestBody.Username)

			err = channel.PublishWithContext(
				ctx,
				"",        // exchange
				"results", // queue name
				false,     // mandatory
				false,     // immediate
				amqp.Publishing{
					ContentType: "text/plain",
					Body:        []byte(pubMsg), // message to publish
				},
			)
			if err != nil {
				msg := fmt.Sprintf("[MatchHandler] Error publishing message | err: %v", err)
				fmt.Println(msg)
			}
		}
	}()
}
