package main

import "github.com/gin-gonic/gin"

func main() {
	server := gin.Default()
	count := 0
	server.GET("/", func(c *gin.Context) {
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
	server.Run(":8080")
}
