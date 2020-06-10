package main

import "github.com/gin-gonic/gin"

func main() {
	server := gin.Default()
	count := 0
	server.GET("/count", func(c *gin.Context) {
		count++
		c.JSON(200, gin.H{
			"message": "hello world",
			"count":   count,
		})
	})
	server.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	server.GET("/jokes/random", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "You don't want any jokes.",
		})
	})

	server.Run(":8080")
}
