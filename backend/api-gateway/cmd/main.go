package main

import (
	"api-gateway/config"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	config.ConnectDb()
	log.Println("Starting development server...")

	API_GATEWAY := echo.New()

	API_GATEWAY.GET("/", hello)
	// API_GATEWAY.POST("/auth/register", createUser)
	// API_GATEWAY.GET("/auth/users", getUsers)

	API_GATEWAY.Start(":1234")
}

func hello(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{"message": "Hello from Echo!"})
}
