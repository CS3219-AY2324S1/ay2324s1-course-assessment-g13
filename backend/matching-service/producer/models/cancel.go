package models

type MessageQueueCancelRequestPacket struct {
	RequestBody CancelRequest `json:"request_body"`
}
