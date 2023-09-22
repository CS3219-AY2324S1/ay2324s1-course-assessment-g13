package handlers

import (
	"api-gateway/config"
	"api-gateway/models"
	"api-gateway/utils/message"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

const (
	INVALID_JSON_REQUEST     = "Invalid JSON Request!"
	INVALID_USER_INPUT       = "Invalid User Input!"
	INVALID_USER_EXIST       = "Username Already Exists!"
	INVALID_USER_NOT_FOUND   = "User Not Found!"
	FAILURE_HASHING_PASSWORD = "An Error Occurred while Hashing Password"
	FAILURE_CREATE_USER      = "Failed to Create User!"
	SUCCESS_USER_FOUND       = "User Found!"
	SUCCESS_USER_CREATED     = "User Created Successfully!"
	SUCCESS_USER_DELETED     = "User Deleted Successfully!"
)

func CreateUser(c echo.Context) error {
	requestBody := new(models.UserCredential)
	if err := c.Bind(requestBody); err != nil {
		return c.JSON(http.StatusBadRequest, message.CreateErrorMessage(INVALID_JSON_REQUEST))
	}

	validator := validator.New()
	if err := validator.Struct(requestBody); err != nil {
		return c.JSON(http.StatusBadRequest, message.CreateErrorMessage(INVALID_USER_INPUT))
	}

	var existingUser models.User
	config.DB.Where("username = ?", requestBody.Username).First(&existingUser)
	if existingUser.ID != 0 {
		return c.JSON(http.StatusConflict, message.CreateErrorMessage(INVALID_USER_EXIST))
	}

	user := new(models.User)
	user.Username = requestBody.Username
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(requestBody.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, message.CreateErrorMessage(FAILURE_HASHING_PASSWORD))
	}
	user.HashedPassword = string(hashedPassword)

	if err := config.DB.Create(user).Error; err != nil {
		return c.JSON(http.StatusInternalServerError, message.CreateErrorMessage(FAILURE_CREATE_USER))
	}

	return c.JSON(http.StatusCreated, message.CreateSuccessMessage(SUCCESS_USER_CREATED, user))
}

func GetUser(c echo.Context) error {
	id := c.Param("id")

	var user models.User
	config.DB.Where("id = ?", id).First(&user)
	if user.ID == 0 {
		return c.JSON(http.StatusNotFound, message.CreateErrorMessage(INVALID_USER_NOT_FOUND))
	}

	return c.JSON(http.StatusOK, message.CreateSuccessMessage(SUCCESS_USER_FOUND, user))
}

func DeleteUser(c echo.Context) error {
	id := c.Param("id")

	var user models.User
	config.DB.Where("id = ?", id).First(&user)
	if user.ID == 0 {
		return c.JSON(http.StatusNotFound, message.CreateErrorMessage(INVALID_USER_NOT_FOUND))
	}

	config.DB.Delete(&user)
	return c.JSON(http.StatusOK, message.CreateSuccessMessage(SUCCESS_USER_DELETED, user))
}
