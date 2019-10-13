package main

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/http"
)

func main() {
	r := gin.Default()

	r.POST("/test", func(c *gin.Context) {
		bodyBytes, err := ioutil.ReadAll(c.Request.Body)
		if err != nil {
			c.String(http.StatusBadRequest, err.Error())
			c.Abort()
		}
		c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))
		firstName := c.PostForm("first_name")
		lastName := c.DefaultPostForm("last_name", "default_last_name")

		c.String(http.StatusOK, "%s,%s,%s\n", firstName, lastName, string(bodyBytes))
	})

	r.Run()
}
