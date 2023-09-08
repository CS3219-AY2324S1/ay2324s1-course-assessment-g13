package main

import (
	"fmt"

	"user-service/handlers"
	"user-service/config"
	"github.com/labstack/echo/v4"
)

func main() {
	config.ConnectDb()
	
	fmt.Println("Starting development server")
	e := echo.New()

	e.POST("/users", handlers.CreateUser)
	e.GET("/users/:id", handlers.GetUser)
	// e.PUT("/users/:id", updateUser)
	// e.DELETE("/users/:id", deleteUser)



	e.Start(":3000")
}
