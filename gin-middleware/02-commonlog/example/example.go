package main

import (
	"github.com/gin-gonic/gin"
	commonLog "go_gin_study/gin-middleware/02-commonlog"
	"net/http"
)

func main() {
	router := gin.Default()
	router.Use(commonLog.New())

	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"status": http.StatusOK,
		})
	})

	router.Run()
}
