package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

var db = make(map[string]string)

func setupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/ping", func(context *gin.Context) {
		context.String(http.StatusOK, "pong")
	})

	r.GET("/user/:name", func(c *gin.Context) {
		user := c.Params.ByName("name")
		value, ok := db[user]
		if ok {
			c.JSON(http.StatusOK, gin.H{
				"user":  user,
				"value": value,
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"user":   user,
				"status": "no value",
			})
		}
	})

	// gin.BasicAuth() 中间件
	authorized := r.Group("/", gin.BasicAuth(gin.Accounts{
		"foo":  "bar",
		"manu": "123",
	}))

	authorized2 := r.Group("/auth2")
	authorized2.Use(gin.BasicAuth(gin.Accounts{
		"foo":  "bar",
		"manu": "123",
	}))

	// localhost:8080/admin
	authorized.GET("admin", func(c *gin.Context) {
		user := c.MustGet(gin.AuthUserKey).(string)

		var json struct {
			Value string `json:"value" binding:"required"`
		}

		if c.Bind(&json) == nil {
			db[user] = json.Value
			c.JSON(http.StatusOK, gin.H{
				"status": "ok",
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"status": "fail",
			})
		}
	})

	authorized2.GET("admin", func(c *gin.Context) {
		user := c.MustGet(gin.AuthUserKey).(string)

		var json struct {
			Value string `json:"value" binding:"required"`
		}

		if c.Bind(&json) == nil {
			db[user] = json.Value
			c.JSON(http.StatusOK, gin.H{
				"status": "ok",
			})
		}
	})
	return r
}

func main() {
	r := setupRouter()

	r.Run(":8080")
}
