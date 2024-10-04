package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/", func(c *gin.Context) {
		c.String(200, "Hello, Unes!")
	})

	router.GET("/bye", func(c *gin.Context) {
		c.String(200, "Bye, Unes!")
	})

	router.Run(":8080")
}
