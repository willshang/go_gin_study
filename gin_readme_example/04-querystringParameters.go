package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()

	// Query string parameters are parsed using the existing underlying request object.
	// The request responds to a url matching:  /welcome?firstname=Jane&lastname=Doe
	// 使用现有的基础请求对象分析查询字符串参数。
	// 请求响应的URL匹配：
	// http://localhost:8080/welcome?firstname=Jane&lastname=Doe&word=!
	router.GET("/welcome", welcome)
	router.Run(":8080")
}

func welcome(c *gin.Context) {
	firstname := c.DefaultQuery("firstname", "Guest")
	lastname := c.Query("lastname")
	word := c.Request.URL.Query().Get("word")
	// shortcut for c.Request.URL.Query().Get("lastname")
	// c.DefaultQuery 是带默认值的c.Query
	// c.Query是c.Request.URL.Query().Get()的简写

	c.String(http.StatusOK, "Hello %s %s %s", firstname, lastname, word)
}
