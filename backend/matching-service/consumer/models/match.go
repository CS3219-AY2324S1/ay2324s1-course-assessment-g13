package models

type MessageQueueRequestPacket struct {
	RequestBody MatchRequest `json:"request_body"`
}

type MatchRequest struct {
	Username      string `json:"username"`
	MatchCriteria string `json:"match_criteria"`
}

type MessageQueueResponsePacket struct {
	ResponseBody MatchResponse `json:"response_body"`
}

type MatchResponse struct {
	MatchUser    string `json:"match_user"`
	MatchStatus  int    `json:"match_status"`
	RedirectURL  string `json:"redirect_url"`
	ErrorMessage string `json:"error_message"`
}
