package main

import (
	"fmt"
	"net/http"
	"os"

	"user-service/config"
	"user-service/handlers"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	config.ConnectDb()

	fmt.Println("Starting development server")
	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{os.Getenv("AGW_URL")},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
		AllowMethods: []string{http.MethodGet, http.MethodHead, http.MethodPut, http.MethodPatch, http.MethodPost, http.MethodDelete},
	}))

	e.POST("/users", handlers.CreateUser)
	e.GET("/users", handlers.GetUsers)
	e.GET("/users/:id", handlers.GetUser)

	e.PUT("/users/:id", handlers.UpdateUser)
	e.DELETE("/users/:id", handlers.DeleteUser)

	e.POST("/history", handlers.CreateHistory)

	e.Logger.Fatal(e.Start(":8080"))
}
