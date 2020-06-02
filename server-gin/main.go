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

	server.POST("values", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"index": 2233,
		})
	})

	server.GET("values/current", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"values": []int{9990, 123, 121, 144},
		})
	})

	server.GET("values/all", func(c *gin.Context) {
		c.JSON(200, []string{"1000", "222"})
	})

	server.Run(":8080")
}
