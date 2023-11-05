package handlers

import (
	"api-gateway/config"
	"api-gateway/models"
	"api-gateway/utils/cookie"
	"api-gateway/utils/message"
	"api-gateway/utils/path"
	"api-gateway/utils/token"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

var bypassLoginList = map[string]bool{
	path.SIGNUP:  true,
	path.LOGIN:   true,
	path.REFRESH: true,
	path.LOGOUT:  true,
}

func RequireAuthenticationMiddleWare(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		uri := strings.Split(c.Request().RequestURI, "?")[0]
		_, isInList := bypassLoginList[uri]
		if isInList {
			return next(c)
		}

		tokenString, statusCode, responseMessage := cookie.Service.GetCookieValue(c.Cookie(ACCESS_TOKEN_COOKIE_NAME))
		if statusCode != http.StatusOK {
			return c.JSON(statusCode, message.CreateErrorMessage(responseMessage))
		}

		tokenClaims, statusCode, responseMessage := token.AccessTokenService.Validate(tokenString)
		if statusCode != http.StatusOK {
			return c.JSON(statusCode, message.CreateErrorMessage(responseMessage))
		}

		oauthId := tokenClaims.User.OauthID
		oauthProvider := tokenClaims.User.OauthProvider

		var user models.User
		err := config.DB.Where("oauth_id = ? AND oauth_provider = ?", oauthId, oauthProvider).First(&user).Error
		if err != nil {
			if err == gorm.ErrRecordNotFound {
				return c.JSON(http.StatusBadRequest, message.CreateErrorMessage(INVALID_USER_NOT_FOUND))
			}
			return c.JSON(http.StatusInternalServerError, message.CreateErrorMessage(INVALID_DB_ERROR))
		}

		c.Set(TOKEN_CLAIMS_CONTEXT_KEY, tokenClaims)
		c.Request().Header.Set(USER_ROLE_KEY_REQUEST_HEADER, "")
		c.Request().Header.Set(USER_ROLE_KEY_REQUEST_HEADER, user.Role)
		return next(c)
	}
}
