package handlers

import (
	"api-gateway/config"
	"api-gateway/models"
	"api-gateway/utils/cookie"
	"api-gateway/utils/expiry"
	"api-gateway/utils/message"
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

	c.Set(USER_CONTEXT_KEY, user)
	c.Set(SUCCESS_MESSAGE_CONTEXT_KEY, SUCCESS_LOGIN)
	c.Set(EXPIRATION_TIME_CONTEXT_KEY, expiry.ExpireIn5Minutes())

	return GenerateTokenAndSetCookie(c)
}

func Logout(c echo.Context) error {
	cookie_, statusCode, responseMessage := cookie.Service.SetCookieExpires(c.Cookie(JWT_COOKIE_NAME))
	if statusCode != http.StatusOK {
		return c.JSON(statusCode, message.CreateErrorMessage(responseMessage))
	}
	c.SetCookie(cookie_)
	return c.JSON(http.StatusOK, message.CreateSuccessMessage(SUCCESS_LOGOUT))
}

func Refresh(c echo.Context) error {
	tokenClaims := c.Get(TOKEN_CLAIMS_CONTEXT_KEY).(*models.Claims)

	user := tokenClaims.User

	c.Set(USER_CONTEXT_KEY, user)
	c.Set(SUCCESS_MESSAGE_CONTEXT_KEY, SUCCESS_TOKEN_REFRESHED)
	c.Set(EXPIRATION_TIME_CONTEXT_KEY, expiry.ExpireIn5Minutes())

	return GenerateTokenAndSetCookie(c)
}
