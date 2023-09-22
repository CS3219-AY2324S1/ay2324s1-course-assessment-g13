package handlers

import (
	"api-gateway/config"
	"api-gateway/models"
	"api-gateway/utils/cookie"
	"api-gateway/utils/expiry"
	"api-gateway/utils/message"
	"api-gateway/utils/token"
	"net/http"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

func Login(c echo.Context) error {
	requestBody := new(models.UserCredential)
	if err := c.Bind(requestBody); err != nil {
		return c.JSON(http.StatusBadRequest, message.CreateErrorMessage(INVALID_JSON_REQUEST))
	}

	var user models.User
	config.DB.Where("username = ?", requestBody.Username).First(&user)
	if user.ID == 0 {
		return c.JSON(http.StatusNotFound, message.CreateErrorMessage(INVALID_USER_NOT_FOUND))
	}
	err := bcrypt.CompareHashAndPassword([]byte(user.HashedPassword), []byte(requestBody.Password))
	if err != nil {
		return c.JSON(http.StatusInternalServerError, message.CreateErrorMessage(FAILURE_HASHING_PASSWORD))
	}

	expirationTime := expiry.ExpireIn5Minutes()
	token, statusCode, responseMessage := token.Service.Generate(&user, expirationTime)
	if statusCode != http.StatusOK {
		return c.JSON(statusCode, message.CreateErrorMessage(responseMessage))
	}

	cookie_ := cookie.Service.CreateCookie(JWT_COOKIE_NAME, token, expirationTime)
	c.SetCookie(cookie_)

	return c.JSON(http.StatusOK, message.CreateSuccessUserMessage(SUCCESS_LOGIN, user))
}

func Logout(c echo.Context) error {
	cookie_, statusCode, responseMessage := cookie.Service.SetCookieExpires(c.Cookie(JWT_COOKIE_NAME))
	if statusCode != http.StatusOK {
		return c.JSON(statusCode, message.CreateErrorMessage(responseMessage))
	}
	c.SetCookie(cookie_)
	return c.JSON(http.StatusOK, message.CreateSuccessMessage(SUCCESS_LOGOUT))
}
