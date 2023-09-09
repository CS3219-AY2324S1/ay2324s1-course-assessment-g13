package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.GET("/collaborate", func(c *gin.Context) {
		c.String(200, "I am the collaboration microservice!")
	})
	r.Run()
}
