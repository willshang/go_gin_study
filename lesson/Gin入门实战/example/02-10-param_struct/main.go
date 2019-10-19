package main

import (
	"github.com/gin-gonic/gin"
	"time"
)

type Person struct {
	Name     string    `form:"name"`
	Address  string    `form:"address"`
	Birthday time.Time `form:"birthday" time_format:"2006-01-02"`
}

func main() {
	r := gin.Default()

	r.GET("/testing", testing)
	r.POST("/testing", testing)

	r.Run()
}

func testing(c *gin.Context) {
	var person Person
	// 根据请求content-type来做不同bindind操作
	if err := c.ShouldBind(&person); err == nil {
		c.String(200, "%v", person)
	} else {
		c.String(200, "person bind error:%v", err)
	}
}
