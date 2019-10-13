package main

import (
	"github.com/gin-gonic/gin"
	gin_static "go_gin_study/gin-middleware/05-static"
	"net/http"
)

func main() {
	router := gin.Default()

	router.Use(gin_static.Serve("/", gin_static.LocalFile("/tmp", false)))

	router.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "test")
	})

	router.Run()
}
