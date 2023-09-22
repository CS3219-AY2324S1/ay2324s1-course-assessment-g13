package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	API_GATEWAY := echo.New()

	API_GATEWAY.GET("/", hello)

	API_GATEWAY.Start(":1234")
}

func hello(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{"message": "Hello from Echo!"})
}
