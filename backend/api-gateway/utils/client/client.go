package client

import (
	"api-gateway/models"
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"os"
)

type UserServiceClient interface {
	CreateUser(models.UserServiceCreateUserRequestBody) (statusCode int, message string)
	DeleteUser()
}

type userServiceClient struct {
	userServiceUrl string
}

var UserService = CreateUserServiceClient()

func (client *userServiceClient) CreateUser(user models.UserServiceCreateUserRequestBody) (int, string) {
	const endpoint = "/users"
	userJson, err := json.Marshal(user)
	if err != nil {
		return http.StatusInternalServerError, "Error Marshaling JSON"
	}

	response, err := http.Post(client.userServiceUrl+endpoint, "application/json", bytes.NewBuffer(userJson))
	if err != nil {
		return http.StatusInternalServerError, "Error Making Request to User Service"
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusCreated {
		responseBody, err := io.ReadAll(response.Body)
		if err != nil {
			return http.StatusInternalServerError, "Error Reading User Service Response"
		}
		var errorResponse models.ErrorMessage
		if err := json.Unmarshal(responseBody, &errorResponse); err != nil {
			return http.StatusInternalServerError, "Error Unmarshaling JSON"
		}
		return response.StatusCode, errorResponse.Message
	}
	return response.StatusCode, ""
}

func (*userServiceClient) DeleteUser() {
	panic("unimplemented")
}

func CreateUserServiceClient() UserServiceClient {
	url := os.Getenv("USER_SERVICE_URL")
	return &userServiceClient{
		userServiceUrl: url,
	}
}
