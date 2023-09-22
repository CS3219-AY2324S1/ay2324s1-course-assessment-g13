package models

type MessageQueueRequestPacket struct {
	RequestBody MatchRequest `json:"request_body"`
}

type MatchRequest struct {
	Username      string `json:"username"`
	MatchCriteria string `json:"match_criteria"`
}

type MessageQueueResponsePacket struct {
}
