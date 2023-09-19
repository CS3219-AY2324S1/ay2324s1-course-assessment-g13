package main

import (
	"fmt"

	"user-service/config"
	"user-service/handlers"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

func main() {
	config.ConnectDb()

	fmt.Println("Starting development server")
	e := echo.New()
	e.Use(session.Middleware(sessions.NewCookieStore([]byte("Secret"))))

	e.GET("/users", handlers.GetUsers)
	e.GET("/users/:id", handlers.GetUser)
	e.PUT("/users/:id", handlers.UpdateUser)
	e.DELETE("/users", handlers.DeleteUser)

	e.POST("/register", handlers.CreateUser)
	e.POST("/login", handlers.Login)
	e.POST("/logout", handlers.Logout)
	e.POST("/refresh", handlers.Refresh)

	e.Start(":3000")
}
