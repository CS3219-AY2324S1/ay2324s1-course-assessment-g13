package handlers

import (
	"api-gateway/config"
	"api-gateway/models"
	"api-gateway/utils/cookie"
	"api-gateway/utils/message"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func GetUser(c echo.Context) error {
	tokenClaims := c.Get(TOKEN_CLAIMS_CONTEXT_KEY).(*models.Claims)

	user := tokenClaims.User

	return c.JSON(http.StatusOK, message.CreateSuccessUserMessage(SUCCESS_USER_FOUND, user))
}

func DeleteUser(c echo.Context) error {
	tokenClaims := c.Get(TOKEN_CLAIMS_CONTEXT_KEY).(*models.Claims)

	user := tokenClaims.User

	err := config.DB.Delete(&user).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, message.CreateErrorMessage(ERROR_OCCURRED))
	}

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

	return c.JSON(http.StatusOK, message.CreateSuccessUserMessage(SUCCESS_USER_DELETED, user))
}

func CreateUser(c echo.Context) error {
	payload := new(models.CreateUserPayload)
	if err := c.Bind(payload); err != nil {
		return c.JSON(http.StatusBadRequest, message.CreateErrorMessage(INVALID_JSON_REQUEST))
	}

	validator := validator.New()
	if err := validator.Struct(payload); err != nil {
		return c.JSON(http.StatusBadRequest, message.CreateErrorMessage(INVALID_USER_INPUT))
	}

	var existingUser models.User
	err := config.DB.Where("oauth_id = ? AND oauth_provider = ?", payload.OauthID, payload.OauthProvider).First(&existingUser).Error
	if err != nil {
		if err != gorm.ErrRecordNotFound {
			return c.JSON(http.StatusInternalServerError, message.CreateErrorMessage(INVALID_DB_ERROR))
		}
	} else {
		return c.JSON(http.StatusBadRequest, message.CreateErrorMessage(INVALID_USER_EXIST))
	}

	var newUser models.User
	newUser.OauthID = payload.OauthID
	newUser.OauthProvider = payload.OauthProvider
	err = config.DB.Create(&newUser).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, message.CreateErrorMessage(FAILURE_CREATE_USER))
	}

	return c.JSON(http.StatusCreated, message.CreateSuccessMessage(SUCCESS_USER_CREATED))
}
