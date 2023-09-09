package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/match", func(c *gin.Context) {
		c.String(200, "I am the matching microservice!")
	})
	r.Run()
}
