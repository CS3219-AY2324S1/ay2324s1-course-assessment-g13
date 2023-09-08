package main

import (
	"fmt"

	"user-service/handlers"
	"github.com/labstack/echo/v4"
)

func main() {
	fmt.Println("Starting development server")
	e := echo.New()

	// e.POST("/users", saveUser)
	e.GET("/users/:id", handlers.GetUser)
	// e.PUT("/users/:id", updateUser)
	// e.DELETE("/users/:id", deleteUser)



	e.Start(":3000")
}
