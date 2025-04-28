package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()

	db := initDB()

	// 将 db 实例注入中间件
	router.Use(func(c *gin.Context) {
		c.Set("db", db)
		c.Next()
	})

	addRouter(router)

	router.Run(":8080")
}
