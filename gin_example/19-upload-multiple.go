package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"path/filepath"
)

func main() {
	router := gin.Default()

	router.MaxMultipartMemory = 8 << 20
	router.Static("/", "./mul_public")

	router.POST("/upload", func(c *gin.Context) {
		name := c.PostForm("name")
		email := c.PostForm("email")

		form, err := c.MultipartForm()
		if err != nil {
			c.String(http.StatusBadRequest,
				fmt.Sprintf("get form err: %s", err.Error()))
		}

		files := form.File["files"]

		for _, file := range files {
			filename := filepath.Base(file.Filename)
			if err := c.SaveUploadedFile(file, filename); err != nil {
				c.String(http.StatusBadRequest,
					fmt.Sprintf("upload file err: %s", err.Error()))
				return
			}
		}
		c.String(http.StatusOK,
			fmt.Sprintf("uploaded successfully %d files with fields name=%s and email=%s.",
				len(files), name, email))
	})
	router.Run()

}
