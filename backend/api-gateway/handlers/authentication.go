package handlers

import (
	"api-gateway/config"
	"api-gateway/models"
	"api-gateway/utils/cookie"
	"api-gateway/utils/message"
	"api-gateway/utils/token"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func Login(c echo.Context) error {
	requestBody := new(models.LoginRequest)
	if err := c.Bind(requestBody); err != nil {
		return c.JSON(http.StatusBadRequest, message.CreateErrorMessage(INVALID_JSON_REQUEST))
	}

	validator := validator.New()
	if err := validator.Struct(requestBody); err != nil {
		return c.JSON(http.StatusBadRequest, message.CreateErrorMessage(INVALID_USER_INPUT))
	}

	var user models.User
	err := config.DB.Where("oauth_id = ? AND oauth_provider = ?", requestBody.OauthID, requestBody.OauthProvider).First(&user).Error
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return c.JSON(http.StatusBadRequest, message.CreateErrorMessage(INVALID_USER_NOT_FOUND))
		}
		return c.JSON(http.StatusInternalServerError, message.CreateErrorMessage(INVALID_DB_ERROR))
	}

	c.Set(USER_CONTEXT_KEY, user)
	c.Set(SUCCESS_MESSAGE_CONTEXT_KEY, SUCCESS_LOGIN)

	return GenerateTokenAndSetCookie(c)
}

func Logout(c echo.Context) error {
	cookie_, statusCode, responseMessage := cookie.Service.SetCookieExpires(c.Cookie(ACCESS_TOKEN_COOKIE_NAME))
	if statusCode != http.StatusOK {
		return c.JSON(statusCode, message.CreateErrorMessage(responseMessage))
	}
	c.SetCookie(cookie_)

	cookie_, statusCode, responseMessage = cookie.Service.SetCookieExpires(c.Cookie(REFRESH_TOKEN_COOKIE_NAME))
	if statusCode != http.StatusOK {
		return c.JSON(statusCode, message.CreateErrorMessage(responseMessage))
	}
	c.SetCookie(cookie_)

	return c.JSON(http.StatusOK, message.CreateSuccessMessage(SUCCESS_LOGOUT))
}

func Refresh(c echo.Context) error {
	tokenString, statusCode, responseMessage := cookie.Service.GetCookieValue(c.Cookie(REFRESH_TOKEN_COOKIE_NAME))
	if statusCode != http.StatusOK {
		return c.JSON(statusCode, message.CreateErrorMessage(responseMessage))
	}

	tokenClaims, statusCode, responseMessage := token.RefreshTokenService.Validate(tokenString)
	if statusCode != http.StatusOK {
		return c.JSON(statusCode, message.CreateErrorMessage(responseMessage))
	}

	user := tokenClaims.User

	c.Set(USER_CONTEXT_KEY, user)
	c.Set(SUCCESS_MESSAGE_CONTEXT_KEY, SUCCESS_TOKEN_REFRESHED)

	return GenerateTokenAndSetCookie(c)
}
