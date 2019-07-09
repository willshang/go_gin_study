package main

import "github.com/gin-gonic/gin"

func main() {
	// Disable log's color
	// 禁用log颜色
	gin.DisableConsoleColor()

	// Creates a gin router with default middleware:
	// logger and recovery (crash-free) middleware
	// 创建一个带默认中间件(logger和recovery)的中间件
	router := gin.Default()
	
	// http://localhost:8080/ping
	router.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	router.Run(":8080")
}
