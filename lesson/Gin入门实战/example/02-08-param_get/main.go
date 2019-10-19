package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	r.GET("/test", func(c *gin.Context) {
		firstName := c.Query("first_name")
		lastName := c.DefaultQuery("last_name", "last_default_name")

		c.String(http.StatusOK, "%s,%s", firstName, lastName)
	})

	r.Run()
}
