package rmq

import (
	"consumer/models"
	"consumer/utils"
	"encoding/json"
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
)

func RunQueueListeners() {
	for _, channelType := range utils.MatchCriterias {
		currCh := LengthChannelsMap[channelType]
		go func(currCh *amqp.Channel, channelType utils.MatchCriteria) {
			exchangeName := "length" + string(channelType)
			declaredLengthQueue, err := currCh.QueueDeclare(
				"",    // name
				false, // durable
				false, // delete when unused
				true,  // exclusive
				false, // no-wait
				nil,   // arguments
			)
			if err != nil {
				msg := fmt.Sprintf("[Init] Error declaring length queue for %s | err: %v", err, string(channelType))
				log.Println(msg)
				panic(err)
			}
			err = currCh.QueueBind(
				declaredLengthQueue.Name, // queue name
				"",                       // routing key
				exchangeName,             // exchange
				false,
				nil,
			)

			msgs, err := currCh.Consume(
				declaredLengthQueue.Name, // queue
				"",                       // consumer
				true,                     // auto-ack
				false,                    // exclusive
				false,                    // no-local
				false,                    // no-wait
				nil,                      // args
			)
			for msg := range msgs {
				var msgResponse models.MessageQueueLengthRequest
				err := json.Unmarshal(msg.Body, &msgResponse)
				if err != nil {
					msg := fmt.Sprintf("[RunQueueListeners] Error unmarshalling request packet | err: %v", err)
					log.Fatal(msg)
					return
				}
				QueueLengthMap[channelType] += msgResponse.Increment
			}
		}(currCh, channelType)
	}
}
