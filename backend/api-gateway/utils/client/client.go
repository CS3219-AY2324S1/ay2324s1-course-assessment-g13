package client

import (
	"api-gateway/models"
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"os"
	"strconv"
)

type UserServiceClient interface {
	CreateUser(models.UserServiceCreateUserRequestBody) (statusCode int, message string)
	DeleteUser(authUserId uint) (statusCode int, message string)
}

type userServiceClient struct {
	userServiceUrl string
	httpClient     *http.Client
}

var UserService = CreateUserServiceClient()

const (
	ERROR_MARSHAL_JSON                     = "Error Marshaling JSON"
	ERROR_UNMARSHAL_JSON                   = "Error Unmarshaling JSON"
	ERROR_REQUEST_TO_USER_SERVICE          = "Error Requesting User Service"
	ERROR_CREATING_REQUEST_TO_USER_SERVICE = "Error Making Request to User Service"
	ERROR_READING_USER_SERVICE_RESPONSE    = "Error Reading User Service Response"
)

const endpoint = "/users"

func (client *userServiceClient) CreateUser(user models.UserServiceCreateUserRequestBody) (int, string) {
	userJson, err := json.Marshal(user)
	if err != nil {
		return http.StatusInternalServerError, ERROR_MARSHAL_JSON
	}

	response, err := http.Post(client.userServiceUrl+endpoint, "application/json", bytes.NewBuffer(userJson))
	if err != nil {
		return http.StatusInternalServerError, ERROR_REQUEST_TO_USER_SERVICE
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusCreated {
		responseBody, err := io.ReadAll(response.Body)
		if err != nil {
			return http.StatusInternalServerError, ERROR_READING_USER_SERVICE_RESPONSE
		}
		var errorResponse models.ErrorMessage
		if err := json.Unmarshal(responseBody, &errorResponse); err != nil {
			return http.StatusInternalServerError, ERROR_UNMARSHAL_JSON
		}
		return response.StatusCode, errorResponse.Message
	}
	return response.StatusCode, ""
}

func (client *userServiceClient) DeleteUser(authUserId uint) (int, string) {
	authUserIdString := strconv.FormatUint(uint64(authUserId), 10)
	uri := client.userServiceUrl + endpoint + "/" + authUserIdString
	request, err := http.NewRequest(http.MethodDelete, uri, nil)
	if err != nil {
		return http.StatusInternalServerError, ERROR_CREATING_REQUEST_TO_USER_SERVICE
	}

	response, err := client.httpClient.Do(request)
	if err != nil {
		return http.StatusInternalServerError, ERROR_REQUEST_TO_USER_SERVICE
	}

	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		responseBody, err := io.ReadAll(response.Body)
		if err != nil {
			return http.StatusInternalServerError, ERROR_READING_USER_SERVICE_RESPONSE
		}
		var errorResponse models.ErrorMessage
		if err := json.Unmarshal(responseBody, &errorResponse); err != nil {
			return http.StatusInternalServerError, ERROR_UNMARSHAL_JSON
		}
		return response.StatusCode, errorResponse.Message
	}

	return response.StatusCode, ""
}

func CreateUserServiceClient() UserServiceClient {
	url := os.Getenv("USER_SERVICE_URL")
	return &userServiceClient{
		userServiceUrl: url,
		httpClient:     &http.Client{},
	}
}
