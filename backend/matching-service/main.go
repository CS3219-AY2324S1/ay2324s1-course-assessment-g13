package main

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func main() {
	e := echo.New()
	e.GET("/ping", func(c echo.Context) error {
		return c.String(http.StatusOK, "I am the matching microservice!")
	})
	e.Logger.Fatal(e.Start(":8080"))
}
