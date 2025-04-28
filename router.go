package main

import "github.com/gin-gonic/gin"

func addRouter(router *gin.Engine) {
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	router.POST("/login", login)
	router.POST("/register", register)
	router.GET("/message", message) // WebSocket连接应该使用GET方法
	router.PUT("/message", updateMessage)
}
