package main

import (
	"fmt"

	"github.com/labstack/echo"
)

func main() {
	fmt.Println("Starting development server")
	e := echo.New()

	e.GET("/", func(c echo.Context) error {
		return c.String(200, "Hello World")
	})

	e.Start(":3000")
}
