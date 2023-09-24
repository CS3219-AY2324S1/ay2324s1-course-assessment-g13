package handlers

import (
	"api-gateway/config"
	"api-gateway/models"
	"api-gateway/utils/env"
	"api-gateway/utils/expiry"
	"api-gateway/utils/message"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/labstack/echo/v4"
)

const (
	GITHUB                             = "Github"
	GITHUB_OAUTH_ACCESS_TOKEN_URL      = "https://github.com/login/oauth/access_token"
	GITHUB_USER_API_URL                = "https://api.github.com/user"
	GITHUB_CALLBACK_REQUEST_QUERY_CODE = "code"
	GITHUB_REQUEST_CREATION_FAILED     = "Github Request Creation Failed"
	GITHUB_API_REQUEST_CREATION_FAILED = "Github API Request Creation Failed"
	GITHUB_RESPONSE_FAILED             = "Github Response Failed"
	GITHUB_ACCESS_TOKEN_SUCCESS        = "Aquire Github Access Token Successfully"
	GITHUB_USER_DATA_SUCCESS           = "Aquire Github User Data Successfully!"

	HTTP_HEADER_CONTENT_TYPE  = "Content-Type"
	HTTP_HEADER_ACCEPT        = "Accept"
	HTTP_HEADER_AUTHORIZATION = "Authorization"
	HTTP_APPLICATION_JSON     = "application/json"
)

func GithubLogin(c echo.Context) error {
	code := c.Request().URL.Query().Get(GITHUB_CALLBACK_REQUEST_QUERY_CODE)

	githubAccessToken, statusCode, responseMessage := getGithubAccessToken(code)
	if statusCode != http.StatusOK {
		return c.JSON(statusCode, message.CreateErrorMessage(responseMessage))
	}
	githubData, statusCode, responseMessage := getGithubData(githubAccessToken)
	if statusCode != http.StatusOK {
		return c.JSON(statusCode, message.CreateErrorMessage(responseMessage))
	}

	provider := GITHUB
	githubUserID := githubData.GithubID

	var existingUser models.User
	config.DB.Where("o_auth_provider = ? AND o_auth_user_id = ?", provider, githubUserID).First(&existingUser)

	if existingUser.ID == 0 {
		c.Set(GITHUB_DATA_CONTEXT_KEY, githubData)
		return oauthCreateUser(c)
	}
	c.Set(USER_CONTEXT_KEY, existingUser)
	c.Set(SUCCESS_MESSAGE_CONTEXT_KEY, GITHUB_USER_DATA_SUCCESS)
	c.Set(EXPIRATION_TIME_CONTEXT_KEY, expiry.ExpireIn5Minutes())

	return GenerateTokenAndSetCookie(c)
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
		http.MethodPost,
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
		http.MethodGet,
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

func oauthCreateUser(c echo.Context) error {
	githubData := c.Get("GithubData").(*models.GithubDataResponseBody)
	user := new(models.User)
	user.OAuthProvider = GITHUB
	user.OAuthUserID = githubData.GithubID
	user.OAuthUsername = githubData.GithubName
	user.OAuthEmail = githubData.GithubEmail
	user.OAuthProfilePictureURL = githubData.GithubProfilePictureURL
	user.OAuthProfileURL = githubData.GithubProfileURL
	if err := config.DB.Create(user).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, message.CreateErrorMessage(FAILURE_CREATE_USER))
	}
	return c.JSON(http.StatusOK, message.CreateSuccessUserMessage(SUCCESS_USER_CREATED, *user))
}
