package handlers

import (
	"api-gateway/models"
	"api-gateway/utils/env"
	"api-gateway/utils/message"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/labstack/echo/v4"
)

const (
	GITHUB_OAUTH_AUTHORIZE_URL_FORMAT  = "https://github.com/login/oauth/authorize?client_id=%s&redirect_uri=%s"
	GITHUB_OAUTH_CALLBACK_URL          = "http://localhost:1234/auth/login/github/callback"
	GITHUB_OAUTH_ACCESS_TOKEN_URL      = "https://github.com/login/oauth/access_token"
	GITHUB_USER_API_URL                = "https://api.github.com/user"
	GITHUB_CALLBACK_REQUEST_QUERY_CODE = "code"
	GITHUB_REQUEST_CREATION_FAILED     = "Github Request Creation Failed"
	GITHUB_API_REQUEST_CREATION_FAILED = "Github API Request Creation Failed"
	GITHUB_RESPONSE_FAILED             = "Github Response Failed"
	GITHUB_ACCESS_TOKEN_SUCCESS        = "Aquire Github Access Token Successfully"
	GITHUB_USER_DATA_SUCCESS           = "Aquire Github User Data Successfully!"

	HTTP_POST                 = "POST"
	HTTP_GET                  = "GET"
	HTTP_HEADER_CONTENT_TYPE  = "Content-Type"
	HTTP_HEADER_ACCEPT        = "Accept"
	HTTP_HEADER_AUTHORIZATION = "Authorization"
	HTTP_APPLICATION_JSON     = "application/json"
)

func GithubLogin(c echo.Context) error {
	githubClientID := env.GetGitHubClientID()
	redirectURL := fmt.Sprintf(
		GITHUB_OAUTH_AUTHORIZE_URL_FORMAT,
		githubClientID,
		GITHUB_OAUTH_CALLBACK_URL,
	)
	return c.Redirect(http.StatusMovedPermanently, redirectURL)
}

func GithubCallback(c echo.Context) error {
	code := c.Request().URL.Query().Get("code")

	githubAccessToken, statusCode, responseMessage := getGithubAccessToken(code)
	if statusCode != http.StatusOK {
		return c.JSON(statusCode, message.CreateErrorMessage(responseMessage))
	}
	githubData, statusCode, responseMessage := getGithubData(githubAccessToken)
	if statusCode != http.StatusOK {
		return c.JSON(statusCode, message.CreateErrorMessage(responseMessage))
	}

	return c.JSON(http.StatusOK, githubData)
}

func getGithubAccessToken(code string) (accessToken string, statusCode int, message string) {
	clientID := env.GetGitHubClientID()
	clientSecret := env.GetGitHubClientSecret()

	requestBody := models.GithubRequestBody{
		ClientID:     clientID,
		ClientSecret: clientSecret,
		Code:         code,
	}

	requestBodyJSON, _ := json.Marshal(requestBody)

	request, requestErr := http.NewRequest(
		HTTP_POST,
		GITHUB_OAUTH_ACCESS_TOKEN_URL,
		bytes.NewBuffer(requestBodyJSON),
	)

	if requestErr != nil {
		return "", http.StatusBadRequest, GITHUB_REQUEST_CREATION_FAILED
	}

	request.Header.Set(HTTP_HEADER_CONTENT_TYPE, HTTP_APPLICATION_JSON)
	request.Header.Set(HTTP_HEADER_ACCEPT, HTTP_APPLICATION_JSON)

	response, responseErr := http.DefaultClient.Do(request)
	if responseErr != nil {
		return "", http.StatusBadRequest, GITHUB_REQUEST_CREATION_FAILED
	}

	responseBody, _ := io.ReadAll(response.Body)

	var githubResponse *models.GithubAccessTokenResponse
	json.Unmarshal(responseBody, &githubResponse)

	return githubResponse.AccessToken, http.StatusOK, GITHUB_ACCESS_TOKEN_SUCCESS
}

func getGithubData(accessToken string) (githubData *models.GithubDataResponseBody, statusCode int, message string) {
	request, requestErr := http.NewRequest(
		HTTP_GET,
		GITHUB_USER_API_URL,
		nil,
	)
	if requestErr != nil {
		return nil, http.StatusBadRequest, GITHUB_API_REQUEST_CREATION_FAILED
	}

	authorizationHeaderValue := fmt.Sprintf("token %s", accessToken)
	request.Header.Set(HTTP_HEADER_AUTHORIZATION, authorizationHeaderValue)

	response, responseErr := http.DefaultClient.Do(request)
	if responseErr != nil {
		return nil, http.StatusBadRequest, GITHUB_RESPONSE_FAILED
	}

	responseBody, _ := io.ReadAll(response.Body)
	var githubDataResponseBody *models.GithubDataResponseBody
	json.Unmarshal(responseBody, &githubDataResponseBody)

	return githubDataResponseBody, http.StatusOK, GITHUB_USER_DATA_SUCCESS
}
