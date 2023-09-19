package handlers

import (
	"net/http"
	"net/http/httputil"
	"net/url"
	"user-service/common/utils"
	model "user-service/models"

	"github.com/labstack/echo/v4"
)

func HandleQuestions(c echo.Context) error {
	questionURL := "http://question-service:8080"
	claims := c.Get(utils.CLAIMS_KEY).(*model.Claims)
	userRole := claims.User.Role

	req := c.Request()

	req.Header.Set("Authorization", userRole)

	target, err := url.Parse(questionURL)
	if err != nil {
		return c.JSON(http.StatusBadRequest, "Bad Request")
	}

	proxy := httputil.NewSingleHostReverseProxy(target)

	proxy.ServeHTTP(c.Response(), req)

	return nil
}
