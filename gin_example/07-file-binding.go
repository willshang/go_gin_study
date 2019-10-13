package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"mime/multipart"
	"net/http"
	"path/filepath"
)

type BindFile struct {
	Name  string                `form:"name" binding:"required"`
	Email string                `form:"email" binding:"required"`
	File  *multipart.FileHeader `form:"file" binding:"required"`
}

func main() {
	router := gin.Default()

	router.MaxMultipartMemory = 8 << 20 // 8M
	router.Static("/", "./public")
	router.POST("/upload", func(c *gin.Context) {
		var bindFile BindFile

		if err := c.ShouldBind(&bindFile); err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("err:%s", err.Error()))
			return
		}

		file := bindFile.File
		dst := filepath.Base(file.Filename)
		if err := c.SaveUploadedFile(file, dst); err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("upload file err: %s", err.Error()))
			return
		}

		c.String(http.StatusOK, fmt.Sprintf("File %s uploaded successfully with fields name=%s and email=%s.",
			file.Filename, bindFile.Name, bindFile.Email))
	})

	router.Run(":8080")
}
