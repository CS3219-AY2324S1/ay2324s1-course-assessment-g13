package models

type GithubAccessTokenResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	Scope       string `json:"scope"`
}

type GithubRequestBody struct {
	ClientID     string `json:"client_id"`
	ClientSecret string `json:"client_secret"`
	Code         string `json:"code"`
}

type GithubDataResponseBody struct {
	GithubID                int    `json:"id"`
	GithubName              string `json:"name"`
	GithubEmail             string `json:"email"`
	GithubProfilePictureURL string `json:"avatar_url"`
	GithubProfileURL        string `json:"html_url"`
}
