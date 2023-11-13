package models

type MessageQueueRequestPacket struct {
	RequestBody MatchRequest `json:"request_body"`
}

type MessageQueueResponsePacket struct {
	ResponseBody MatchResponse `json:"response_body"`
}

type MatchRequest struct {
	Username      string `json:"username"`
	MatchCriteria string `json:"match_criteria"`
}

type MatchResponse struct {
	MatchUser    string `json:"match_user"`
	MatchStatus  int    `json:"match_status"`
	RoomId       string `json:"room_id"`
	ErrorMessage string `json:"error_message"`
}

type MessageQueueLengthResponse struct {
	MessageStats MessageStatsStruct `json:"message_stats"`
}

type MessageStatsStruct struct {
	Ack     int64 `json:"ack"`
	Publish int64 `json:"publish"`
}

type MessageQueueLengthRequest struct {
	Increment     int    `json:"increment"`
	MatchCriteria string `json:"match_criteria"`
}
