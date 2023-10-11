package models

type MessageQueueRequestPacket struct {
	RequestBody MatchRequest `json:"request_body"`
}

type MessageQueueResponsePacket struct {
	ResponseBody MatchResponse `json:"response_body"`
}
