package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/autotls"
)

func main()  {
	r := gin.Default()
	r.GET("/test", func(c *gin.Context) {
		c.String(200, "hello test")
	})
	fmt.Println()
	autotls.Run(r,"www.abc.com")
}