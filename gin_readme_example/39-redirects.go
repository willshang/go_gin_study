package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	// http://127.0.0.1/8080/move
	r.GET("/move", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "http://www.baidu.com/")
	})

	// http://127.0.0.1/8080/test
	r.GET("/test", func(c *gin.Context) {
		c.Request.URL.Path = "/test2"
		r.HandleContext(c)
	})

	// http://127.0.0.1/8080/test2
	r.GET("/test2", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"hello": "world",
		})
		return
	})

	r.Run(":8080")
}
