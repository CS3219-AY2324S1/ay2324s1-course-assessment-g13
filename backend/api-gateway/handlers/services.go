package handlers

import (
	"net/http/httputil"
	"net/url"
	"os"

	"github.com/labstack/echo/v4"
)

func HandleUserService(c echo.Context) error {
	targetURL := os.Getenv("USER_SERVICE_URL")
	target, _ := url.Parse(targetURL)
	proxy := httputil.NewSingleHostReverseProxy(target)
	proxy.ServeHTTP(c.Response(), c.Request())
	return nil
}

func HandleQuestionService(c echo.Context) error {
	targetURL := os.Getenv("QUESTION_SERVICE_URL")
	target, _ := url.Parse(targetURL)
	proxy := httputil.NewSingleHostReverseProxy(target)
	proxy.ServeHTTP(c.Response(), c.Request())
	return nil
}

func HandleMatchingService(c echo.Context) error {
	targetURL := os.Getenv("MATCHING_SERVICE_URL")
	target, _ := url.Parse(targetURL)
	proxy := httputil.NewSingleHostReverseProxy(target)
	proxy.ServeHTTP(c.Response(), c.Request())
	return nil
}
