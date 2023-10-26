package worker

import (
	"consumer/models"
	"consumer/rmq"
	"consumer/utils"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	// "io"
	"bytes"
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
			msg := fmt.Sprintf("[SpinMQConsumer] Criteria channel is unknown | ok: %v", ok)
			log.Println(msg)
			return
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
			msg := fmt.Sprintf("[SpinMQConsumer] Error consuming from channel | err: %v", err)
			log.Println(msg)
			return
		}

		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		resultChan, err := rmq.Conn.Channel()
		if err != nil {
			msg := fmt.Sprintf("[SpinMQConsumer] Error creating unique result channel | err: %v", err)
			log.Println(msg)
			return
		}
		defer resultChan.Close()

		//matchMakingBufferMap := map[string]models.MessageQueueRequestPacket{}
		matchMakingBuffer := []models.MessageQueueRequestPacket{}

		for message := range messages {
			res := models.MessageQueueRequestPacket{}
			err := json.Unmarshal(message.Body, &res)
			if err != nil {
				msg := fmt.Sprintf("[SpinMQConsumer] Error unmarshalling request packet | err: %v", err)
				log.Println(msg)
				return
			}

			isReplaced := false

			// Check if user is currently stale data in map
			for index, bufferedUser := range matchMakingBuffer {
				if bufferedUser.RequestBody.Username == res.RequestBody.Username {
					matchMakingBuffer[index] = res
					isReplaced = true
				}
			}

			if !isReplaced {
				matchMakingBuffer = append(matchMakingBuffer, res)
			}

			// Just for assignment 5 PDF buffer illustration
			log.Printf(" > Received message: %s with buffer size %d\n", res, len(matchMakingBuffer))
			log.Printf("----- Current Queue (%d) -----\n", len(matchMakingBuffer))
			for _, v := range matchMakingBuffer {
				log.Printf("%s\n", v.RequestBody.Username)
			}
			log.Printf("----- End of Queue -----\n")

			// Match found. Safe to do since single threaded within this goroutine
			if len(matchMakingBuffer) == 2 {
				fmt.Println("Found a match!")
				// TODO: create room here
				reqBody, err := json.Marshal(map[string]string{
					"user1": matchMakingBuffer[0].RequestBody.Username,
					"user2": matchMakingBuffer[1].RequestBody.Username,
					"complexity": string(criteria),
				})
				if err != nil {
					log.Fatal(err)
				}
				resp, err := http.Post(os.Getenv("COLLAB_URL") + "/room", "application/json", bytes.NewBuffer(reqBody))
				if err != nil {
					log.Fatal(err)
					return
				}
				defer resp.Body.Close()
				body := map[string]interface{}{}
				json.NewDecoder(resp.Body).Decode(&body)
				log.Println(body["Id"])

				// body, err := io.ReadAll(resp.Body)
				// if err != nil {
				// 	log.Fatal(err)
				// }
				// log.Println(string(body))
				
				for index, user := range matchMakingBuffer {
					pubMsg := models.MessageQueueResponsePacket{
						ResponseBody: models.MatchResponse{
							MatchUser:    matchMakingBuffer[(index+1)%2].RequestBody.Username,
							MatchStatus:  1,
							RoomId:  body["Id"].(string),
							ErrorMessage: "",
						},
					}

					responsePacket, err := json.Marshal(pubMsg)
					if err != nil {
						msg := fmt.Sprintf("[SpinMQConsumer] Error marshalling response packet | err: %v", err)
						log.Println(msg)
						return
					}

					// Declare unique result queue
					resultQueue, err := resultChan.QueueDeclare(
						utils.ConstructResultChanIdentifier(user.RequestBody.Username),
						false,
						false,
						false,
						false,
						nil,
					)

					err = resultChan.PublishWithContext(
						ctx,
						"",               // exchange
						resultQueue.Name, // queue name
						false,            // mandatory
						false,            // immediate
						amqp.Publishing{
							ContentType: "text/plain",
							Body:        responsePacket, // message to publish
						},
					)
					if err != nil {
						msg := fmt.Sprintf("[SpinMQConsumer] Error publishing message to results channel | err: %v", err)
						log.Println(msg)
						return
					}
				}
				matchMakingBuffer = nil
			}
		}
	}()
}
