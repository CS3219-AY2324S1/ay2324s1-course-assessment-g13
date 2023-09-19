package main

import (
	"fmt"

	"user-service/common/auth"
	"user-service/config"
	"user-service/handlers"

	"github.com/labstack/echo/v4"
)

func main() {
	config.ConnectDb()

	fmt.Println("Starting development server")
	e := echo.New()
	e.Use(auth.UserLoginRequired)

	e.GET("/users", handlers.GetUsers)
	e.GET("/users/:id", handlers.GetUser)
	e.PUT("/users/:id", handlers.UpdateUser)
	e.DELETE("/users", handlers.DeleteUser)

	e.POST("/users/upgrade", handlers.UpgradeRole)
	e.POST("/users/downgrade", handlers.DowngradeRole)

	e.POST("/register", handlers.CreateUser, auth.PreventLoggedInUser)
	e.POST("/login", handlers.Login, auth.PreventLoggedInUser)
	e.GET("/logout", handlers.Logout)
	e.GET("/refresh", handlers.Refresh)

	e.Start(":3000")
}
