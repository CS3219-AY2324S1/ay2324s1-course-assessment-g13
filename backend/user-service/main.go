package main

import (
	"fmt"
	"net/http"
	"user-service/config"

	"user-service/handlers"

	"github.com/gorilla/sessions"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	config.ConnectDb()

	fmt.Println("Starting development server")
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://localhost:3000"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))
	e.Use(session.Middleware(sessions.NewCookieStore([]byte("Secret"))))

	e.GET("/ping", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "I am the user microservice")
	})

	e.POST("/users", handlers.CreateUser)
	e.GET("/users", handlers.GetUsers)
	e.GET("/users/:id", handlers.GetUser)
	e.PUT("/users/:id", handlers.UpdateUser)
	e.DELETE("/users", handlers.DeleteUser)

	e.POST("/login", handlers.Login)

	e.Start(":8080")
}
