package worker

import (
	"bytes"
	"consumer/models"
	"consumer/rmq"
	"consumer/utils"
	"context"
	"encoding/json"
	"fmt"
	amqp "github.com/rabbitmq/amqp091-go"
	"log"
	"net/http"
	"os"
	// "io"
	"strings"
	"time"
)

func SpinMQConsumer(criteria utils.MatchCriteria) {
	go func() {
		var reqChannel *amqp.Channel
		// Declare message queues to ensure it exists
		if curr, ok := rmq.OpenChannelsMap[criteria]; ok {
			reqChannel = curr
		} else {
			msg := fmt.Sprintf("[SpinMQConsumer] Criteria channel is unknown | ok: %v", ok)
			log.Fatal(msg)
			return
		}

		messages, err := reqChannel.Consume(
			string(criteria), // queue name
			"",               // consumer
			false,            // auto-ack
			false,            // exclusive
			false,            // no local
			false,            // no wait
			nil,              // arguments
		)
		if err != nil {
			msg := fmt.Sprintf("[SpinMQConsumer] Error consuming from channel | err: %v", err)
			log.Fatal(msg)
			return
		}

		var syncChannel *amqp.Channel
		// Declare sync MQ to ensure it exists
		if curr, ok := rmq.SyncChannelsMap[criteria]; ok {
			syncChannel = curr
		} else {
			msg := fmt.Sprintf("[SpinMQConsumer] Sync criteria channel is unknown | ok: %v", ok)
			log.Fatal(msg)
			return
		}
		syncQueueName := string(criteria) + "sync"

		syncMsges, err := syncChannel.Consume(
			syncQueueName, // queue name
			"",            // consumer
			false,         // auto-ack
			false,         // exclusive
			false,         // no local
			false,         // no wait
			nil,           // arguments
		)
		if err != nil {
			msg := fmt.Sprintf("[SpinMQConsumer] Error consuming from sync channel | err: %v", err)
			log.Fatal(msg)
			return
		}

		ctx, cancel := context.WithCancel(context.Background())
		defer cancel()

		resultChan, err := rmq.Conn.Channel()
		if err != nil {
			msg := fmt.Sprintf("[SpinMQConsumer] Error creating unique result channel | err: %v", err)
			log.Fatal(msg)
			return
		}
		defer resultChan.Close()

		// Worker polls queue every 1 second
		for {
			hasJustBroke := false
			// Get current MQ length
			queueSize := rmq.GetLocalQueueSize(string(criteria))
			log.Printf("Actual size of %s queue: %d\n", string(criteria), queueSize)
			// If queue has sufficient people queued up
			if queueSize >= 2 {
				// Check if request channel is being consumed via sync channel
				syncQueueSize := rmq.GetQueueSize(syncQueueName)
				if syncQueueSize == 1 {
					// If request channel is being consumed, re-loop
					continue
				} else {
					// If not, add into sync channel that this worker is currently consuming from request queue
					err = syncChannel.PublishWithContext(
						ctx,
						"",            // exchange
						syncQueueName, // queue name
						false,         // mandatory
						false,         // immediate
						amqp.Publishing{
							ContentType: "text/plain",
							Body:        []byte{1}, // message to publish
						})
					if err != nil {
						msg := fmt.Sprintf("[SpinMQConsumer] Error publishing message to sync channel | err: %v", err)
						log.Fatal(msg)
						return
					}
				}
				// Sequentially dequeue them and match them
				matchMakingBuffer := []models.MessageQueueRequestPacket{}
				for message := range messages {
					// Ack the message that is consumed
					err := reqChannel.Ack(message.DeliveryTag, false)
					if err != nil {
						msg := fmt.Sprintf("[SpinMQConsumer] Error ACKing request packet | err: %v", err)
						log.Fatal(msg)
						return
					}

					res := models.MessageQueueRequestPacket{}
					err = json.Unmarshal(message.Body, &res)
					if err != nil {
						msg := fmt.Sprintf("[SpinMQConsumer] Error unmarshalling request packet | err: %v", err)
						log.Fatal(msg)
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

					// Once 2 users are grabbed from the MQ, match them
					if len(matchMakingBuffer) == 2 {
						fmt.Println("Found a match!")
						reqBody, err := json.Marshal(map[string]string{
							"user1":      matchMakingBuffer[0].RequestBody.Username,
							"user2":      matchMakingBuffer[1].RequestBody.Username,
							"complexity": strings.Title(string(criteria)),
						})
						if err != nil {
							log.Fatal(err)
						}

						collabLink := os.Getenv("COLLAB_URL")

						resp, err := http.Post(collabLink+"/room", "application/json", bytes.NewBuffer(reqBody))
						if err != nil {
							log.Fatal(err)
							return
						}
						defer resp.Body.Close()
						body := map[string]interface{}{}
						json.NewDecoder(resp.Body).Decode(&body)

						for index, user := range matchMakingBuffer {
							pubMsg := models.MessageQueueResponsePacket{
								ResponseBody: models.MatchResponse{
									MatchUser:    matchMakingBuffer[(index+1)%2].RequestBody.Username,
									MatchStatus:  1,
									RoomId:       body["Id"].(string),
									ErrorMessage: "",
								},
							}

							responsePacket, err := json.Marshal(pubMsg)
							if err != nil {
								msg := fmt.Sprintf("[SpinMQConsumer] Error marshalling response packet | err: %v", err)
								log.Fatal(msg)
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
								log.Fatal(msg)
								return
							}
						}
						matchMakingBuffer = nil
						// Remove sync msg from sync channel
						for syncMsg := range syncMsges {
							log.Printf("Sync messaged retrieved | %v\n", syncMsg)
							err = syncChannel.Ack(syncMsg.DeliveryTag, false)
							if err != nil {
								msg := fmt.Sprintf("[SpinMQConsumer] Error ACKing sync channel content | err: %v", err)
								log.Fatal(msg)
								return
							}
							log.Printf("Worker (%s) breaking out of syncMsg queue\n", criteria)
							break
						}
						log.Printf("Worker (%s) breaking out of msg:Messages queue\n", criteria)
						hasJustBroke = true
						break // Break out of the message consumption loop
					}
				}
			}
			if hasJustBroke {
				log.Printf("Worker (%s) broke out using just_test\n", criteria)
				time.Sleep(time.Second * 1)
			} else {
				// 1 second interval polling
				time.Sleep(time.Second * 1)
			}
		}
	}()
}
