package main

import (
	"github.com/gin-gonic/gin"
)

type Person struct {
	Age     int    `form:"age" binding:"required,gt=10"`
	Address string `form:"address" binding:"required"`
	Name    string `form:"name"  binding:"required"`
}

func main() {
	r := gin.Default()

	r.POST("/testing", func(c *gin.Context) {
		var person Person
		// 根据请求content-type来做不同bindind操作
		if err := c.ShouldBind(&person); err == nil {
			c.String(200, "%v", person)
			c.Abort()
			return
		} else {
			c.String(500, "person bind error:%v", err)
		}
	})

	r.Run()
}
