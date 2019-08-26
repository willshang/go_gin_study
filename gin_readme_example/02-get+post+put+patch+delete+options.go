package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	// Creates a gin router with default middleware:
	// logger and recovery (crash-free) middleware
	// 创建一个gin 路由，自带默认中间件
	// logger + recovery 中间件
	router := gin.Default()

	router.GET("/someGet", getting)
	router.POST("/somePost", posting)
	router.PUT("/somePut", putting)
	router.DELETE("/someDelete", deleting)
	router.PATCH("/somePatch", patching)
	router.HEAD("/someHead", head)
	router.OPTIONS("/someOptions", options)

	// By default it serves on :8080 unless a
	// PORT environment variable was defined.
	// router.Run(":3000") for a hard coded port
	// 默认运行在8080端口
	// 可以设置环境变量 PORT 设定端口
	// 可以使用router.Run(":3000") 设定端口
	router.Run()
}

func getting(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "someGet",
	})
	return
}

func posting(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "somePost",
	})
	return
}

func putting(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "somePut",
	})
	return
}

func deleting(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "someDelete",
	})
	return
}

func patching(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "somePatch",
	})
	return
}

func head(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "someHead",
	})
	return
}

func options(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "someOptions",
	})
	return
}
