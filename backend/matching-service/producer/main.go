package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
	"net/http"
	"producer/handlers"
	"producer/rmq"
)

func main() {
	rmq.Init()
	defer rmq.Reset()

	e := echo.New()
	e.Use(middleware.CORS())
	e.POST("/match", handlers.MatchHandler)
	e.GET("/ping", func(c echo.Context) error {
		return c.JSON(http.StatusOK, "I am the matching producer microservice")
	})

	err := e.Start(":8080")
	if err != nil {
		msg := fmt.Sprintf("[main] Error starting server | err: %v", err)
		log.Println(msg)
		return
	}
}
