package handlers

import (
	"net/http"
	"user-service/common/auth"
	constants "user-service/common/constants"
	"user-service/common/cookie"
	"user-service/common/errors"
	"user-service/config"
	model "user-service/models"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

func Login(c echo.Context) error {
	req := new(model.LoginRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid JSON Request"})
	}

	var user model.User
	config.DB.Where("username = ?", req.Username).First(&user)

	err := bcrypt.CompareHashAndPassword([]byte(user.HashedPassword), []byte(req.Password))
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"message": "Invalid Username or Password"})
	}

	expirationTime := auth.GetExpirationTime()

	tokenString, err := auth.TokenService.Generate(&user, expirationTime)

	if err != nil {
		status, message := errors.ParseErrorToServiceError(err)
		return c.JSON(status, map[string]string{"message": message})
	}

	cookie := cookie.CreateCookie(constants.JWT_COOKIE_NAME, tokenString, expirationTime)
	c.SetCookie(cookie)

	return c.JSON(http.StatusOK, map[string]string{"message": "Login Successful"})
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

	return c.JSON(http.StatusOK, map[string]string{"message": "Token Refresh"})
}

func Logout(c echo.Context) error {
	cookie, err := cookie.SetCookieExpires(c.Cookie(constants.JWT_COOKIE_NAME))
	if err != nil {
		status, message := errors.ParseErrorToServiceError(err)
		return c.JSON(status, map[string]string{"message": message})
	}
	c.SetCookie(cookie)
	return c.JSON(http.StatusOK, map[string]string{"message": "Logout Successful"})
}
