package models

type MatchRequest struct {
	Username      string `json:"username"`
	MatchCriteria string `json:"match_criteria"`
}

type MatchResponse struct {
	MatchUser    string `json:"match_user"`
	MatchStatus  int    `json:"match_status"` // 0 is failure, 1 is success
	RoomId  string `json:"room_id"`
	ErrorMessage string `json:"error_message"`
}

type CancelRequest struct {
	Username string `json:"username"`
}

type CancelResponse struct {
	CancelStatus bool `json:"cancel_status"` // True means cancel success, false means failure
}
