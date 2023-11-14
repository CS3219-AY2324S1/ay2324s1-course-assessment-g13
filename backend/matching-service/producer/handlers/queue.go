package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"producer/models"
	"producer/rmq"
	"producer/utils"
	"time"
)

var userInQueueMap map[string]bool

func SetUserInQueue(user string) {
	userInQueueMap[user] = true
}

func RemoveUserFromQueueImmediate(user string) {
	delete(userInQueueMap, user)
}

func RemoveUserFromQueueDelay(user string) {
	time.Sleep(time.Second * 1) // Ensure match occurs first before removing from queue
	delete(userInQueueMap, user)
}

func IsUserInQueue(user string) bool {
	if v, ok := userInQueueMap[user]; ok {
		return v
	}
	return false
}

func InitQueueTracker() {
	userInQueueMap = make(map[string]bool)
	go func() {
		messages, err := rmq.InQueueChannel.Consume(
			rmq.InQueueName, // queue
			"",              // consumer
			true,            // auto-ack
			false,           // exclusive
			false,           // no-local
			false,           // no-wait
			nil,             // args
		)
		if err != nil {
			msg := fmt.Sprintf("[Init] Error consuming from cancel queue | err: %v", err)
			log.Println(msg)
			panic(err)
		}

		// Starts consuming from cancel channel
		for msg := range messages {
			// If message is received, means cancel the user
			var recvdPkt models.MessageQueueInQueuePacket
			err := json.Unmarshal(msg.Body, &recvdPkt)
			if err != nil {
				msg := fmt.Sprintf("[Init] Error unmarshalling inQueue packet into struct | err: %v", err)
				log.Println(msg)
				panic(err)
			}
			if recvdPkt.Config == utils.Queue {
				SetUserInQueue(recvdPkt.Username)
			} else if recvdPkt.Config == utils.Immediate {
				RemoveUserFromQueueImmediate(recvdPkt.Username)
			} else {
				// Delayed
				RemoveUserFromQueueDelay(recvdPkt.Username)
			}
		}
	}()
}
