package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"log"
	"producer/handlers"
	"producer/rmq"
)

func main() {
	rmq.Init()
	defer rmq.Reset()

	e := echo.New()
	e.POST("/match", handlers.MatchHandler)
	err := e.Start(":8080")
	if err != nil {
		msg := fmt.Sprintf("[main] Error starting server | err: %v", err)
		log.Println(msg)
		return
	}
}
