package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	router := gin.Default()

	// POST http://localhost:8080/post?id=5&page=4
	// form_data: message:hello name:jack
	router.POST("/post", func(c *gin.Context) {
		id := c.Query("id")
		page := c.DefaultQuery("page", "0")
		name := c.PostForm("name")
		message := c.PostForm("message")

		fmt.Printf("id: %s; page: %s; name: %s; message: %s", id, page, name, message)
		c.JSON(http.StatusOK, gin.H{
			"data": fmt.Sprintf("id: %s; page: %s; name: %s; message: %s", id, page, name, message),
		})
		return
	})
	router.Run(":8080")
}
