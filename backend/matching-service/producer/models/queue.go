package models

import "producer/utils"

type MessageQueueInQueuePacket struct {
	Username string            `json:"username"`
	Config   utils.QueueConfig `json:"config"`
}
