package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"log"
	"producer/utils"

	"net/http"
	"os"
	"producer/handlers"
	"producer/rmq"
)

func main() {
	rmq.Init()
	defer rmq.Reset()

	handlers.Init()
	utils.InitCancelTracker()
	handlers.InitQueueTracker()

	e := echo.New()
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{os.Getenv("AGW_URL")},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
		AllowMethods: []string{http.MethodGet, http.MethodPost},
	}))

	e.POST("/match/find", handlers.MatchHandler)
	e.POST("/match/cancel", handlers.UserCancelHandler)

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
