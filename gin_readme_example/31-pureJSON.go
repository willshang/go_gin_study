package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()

	// Serves unicode entities
	// 提供 unicode 实体
	// http://127.0.0.1:8080/json
	r.GET("/json", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"html": "<b>Hello, world!</b>",
		})
	})

	// Serves literal characters
	// 提供字面字符
	// http://127.0.0.1:8080/purejson
	r.GET("/purejson", func(c *gin.Context) {
		c.PureJSON(200, gin.H{
			"html": "<b>Hello, world!</b>",
		})
	})

	// listen and serve on 0.0.0.0:8080
	r.Run(":8080")
}
