package handlers

import (
	"net/http"
	"user-service/common/auth"
	"user-service/common/cookie"
	"user-service/common/errors"
	constants "user-service/common/utils"
	"user-service/config"
	model "user-service/models"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

const (
	INVALID_JSON_REQUEST_MESSAGE         = "Invalid JSON Request"
	INVALID_USERNAME_OR_PASSWORD_MESSAGE = "Invalid Username or Password"
	LOGIN_SUCCESSFUL_MESSAGE             = "Login Successful"
	TOKEN_REFRESH_MESSAGE                = "Token Refreshed"
	LOGOUT_SUCCESSFUL_MESSAGE            = "Logout Successful"
)

func Login(c echo.Context) error {
	req := new(model.LoginRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": INVALID_JSON_REQUEST_MESSAGE})
	}

	var user model.User
	config.DB.Where("username = ?", req.Username).First(&user)

	err := bcrypt.CompareHashAndPassword([]byte(user.HashedPassword), []byte(req.Password))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": INVALID_USERNAME_OR_PASSWORD_MESSAGE})
	}

	expirationTime := auth.GetExpirationTime()

	tokenString, err := auth.TokenService.Generate(&user, expirationTime)

	if err != nil {
		status, message := errors.ParseErrorToServiceError(err)
		return c.JSON(status, map[string]string{"message": message})
	}

	cookie := cookie.CreateCookie(constants.JWT_COOKIE_NAME, tokenString, expirationTime)
	c.SetCookie(cookie)

	return c.JSON(http.StatusOK, map[string]string{"message": LOGIN_SUCCESSFUL_MESSAGE})
}

func Refresh(c echo.Context) error {
	claims := c.Get(constants.CLAIMS_KEY).(*model.Claims)

	user := claims.User
	expirationTime := auth.GetExpirationTime()

	newTokenString, err := auth.TokenService.Generate(&user, expirationTime)
	if err != nil {
		status, message := errors.ParseErrorToServiceError(err)
		return c.JSON(status, map[string]string{"message": message})
	}

	cookie := cookie.CreateCookie(constants.JWT_COOKIE_NAME, newTokenString, expirationTime)
	c.SetCookie(cookie)

	return c.JSON(http.StatusOK, map[string]string{"message": TOKEN_REFRESH_MESSAGE})
}

func Logout(c echo.Context) error {
	cookie, err := cookie.SetCookieExpires(c.Cookie(constants.JWT_COOKIE_NAME))
	if err != nil {
		status, message := errors.ParseErrorToServiceError(err)
		return c.JSON(status, map[string]string{"message": message})
	}
	c.SetCookie(cookie)
	return c.JSON(http.StatusOK, map[string]string{"message": LOGOUT_SUCCESSFUL_MESSAGE})
}
