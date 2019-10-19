package main

import (
	"github.com/gin-gonic/gin"
	"github.com/thinkerou/favicon"
	"net/http"
)

func main() {
	r := gin.Default()
	r.Use(favicon.New("./gin_example/favicon.ico"))
	// localhost:8080/ping
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello favicon.")
	})
	r.Run(":8080")
}
