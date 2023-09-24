package handlers

import (
	"api-gateway/config"
	"api-gateway/models"
	"api-gateway/utils/cookie"
	"api-gateway/utils/message"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
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

	return c.JSON(http.StatusOK, message.CreateSuccessUserMessage(SUCCESS_USER_CREATED, *user))
}

func GetUser(c echo.Context) error {
	tokenClaims := c.Get(TOKEN_CLAIMS_CONTEXT_KEY).(*models.Claims)

	user := tokenClaims.User

	return c.JSON(http.StatusOK, message.CreateSuccessUserMessage(SUCCESS_USER_FOUND, user))
}

func DeleteUser(c echo.Context) error {
	tokenClaims := c.Get(TOKEN_CLAIMS_CONTEXT_KEY).(*models.Claims)

	user := tokenClaims.User

	config.DB.Delete(&user)

	cookie_, statusCode, responseMessage := cookie.Service.SetCookieExpires(c.Cookie(JWT_COOKIE_NAME))
	if statusCode != http.StatusOK {
		return c.JSON(statusCode, message.CreateErrorMessage(responseMessage))
	}
	c.SetCookie(cookie_)

	return c.JSON(http.StatusOK, message.CreateSuccessUserMessage(SUCCESS_USER_DELETED, user))
}
