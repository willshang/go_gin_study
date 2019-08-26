package main

import "github.com/gin-gonic/gin"

func main() {
	router := gin.Default()

	// POST http://localhost:8080/form_post
	// form_data: message hello
	router.POST("/form_post", func(c *gin.Context) {
		message := c.PostForm("message")
		nick := c.DefaultPostForm("nick", "anonymous")

		c.JSON(200, gin.H{
			"status":  "posted",
			"message": message,
			"nick":    nick,
		})
		return
	})

	// POST http://localhost:8080/form_post_origin
	// form_data: message hello
	router.POST("/form_post_origin", func(c *gin.Context) {
		message := c.Request.PostFormValue("message")

		c.JSON(200, gin.H{
			"status":  "posted",
			"message": message,
		})
		return
	})
	router.Run(":8080")
}
