package main

import (
	"api-gateway/config"
	"api-gateway/handlers"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	config.ConnectDb()
	log.Println("Starting development server...")

	API_GATEWAY := echo.New()

	API_GATEWAY.GET("/", hello)
	API_GATEWAY.POST("/auth/register", handlers.CreateUser)
	API_GATEWAY.GET("/auth/users/:id", handlers.GetUser)
	API_GATEWAY.DELETE("/auth/users/:id", handlers.DeleteUser)
	API_GATEWAY.POST("/auth/login", handlers.Login)
	API_GATEWAY.GET("/auth/logout", handlers.Logout)

	API_GATEWAY.Start(":1234")
}

func hello(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{"message": "Hello from Echo!"})
}
