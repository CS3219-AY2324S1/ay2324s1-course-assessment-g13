package main

import (
	"github.com/labstack/echo/v4"
	"producer/handlers"
	"producer/rmq"
)

func main() {
	rmq.Init()

	e := echo.New()
	e.POST("/match", handlers.MatchHandler)
	defer rmq.Reset()
	e.Start(":8080")
}
