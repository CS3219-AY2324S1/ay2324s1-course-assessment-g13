package models

type MatchRequest struct {
	Username      string `json:"username"`
	MatchCriteria string `json:"match_criteria"`
}

type MatchResponse struct {
	MatchUser    string `json:"match_user"`
	MatchStatus  int    `json:"match_status"`
	RedirectURL  string `json:"redirect_url"`
	ErrorMessage string `json:"error_message"`
}
