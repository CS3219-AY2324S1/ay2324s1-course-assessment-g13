package handlers

import (
	"net/http"
	"user-service/config"
	model "user-service/models"

	"strconv"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

func Login(c echo.Context) error {
	req := new(model.LoginRequest)
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, "Invalid JSON request")
	}

	var user model.User
	config.DB.Where("username = ?", req.Username).First(&user)

	err := bcrypt.CompareHashAndPassword([]byte(user.HashedPassword), []byte(req.Password))
	if err != nil {
		// Passwords don't match
		return c.JSON(http.StatusBadRequest, "Invalid request")
	}

	sess, _ := session.Get("session", c)
	sess.Options = &sessions.Options{
		Domain:  "localhost",
		Path:     "/",
		MaxAge:   86400 * 7,
		HttpOnly: false,
		Secure:   false, // have to change to true in production
	}
	sess.Values["userId"] = strconv.FormatUint(uint64(user.ID), 10)
	if err := sess.Save(c.Request(), c.Response()); err != nil {
		return c.JSON(http.StatusInternalServerError, "Internal server error")
	}

	res := &model.LoginResponse{
		Id: user.ID,
		Username: user.Username,
		PhotoUrl: user.PhotoUrl,
	}

	return c.JSON(http.StatusOK, res)

}
