package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func main() {
	router := gin.Default()
	// Set a lower memory limit for multipart forms (default is 32 MiB)
	// router.MaxMultipartMemory = 8 << 20  // 8 MiB
	// 为 multipart forms 设置较低的内存限制 (默认是 32 MiB)
	// router.MaxMultipartMemory = 8 << 20  // 8 MiB
	router.POST("/upload", func(c *gin.Context) {
		// Multipart form
		form, _ := c.MultipartForm()
		files := form.File["upload[]"]
		
		for _, file := range files {
			dst := "/Users/xx/"
			filename := dst + time.Now().Format("20060102-150405") + "-" + file.Filename
			c.SaveUploadedFile(file, filename)
			
			c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded save!\n", filename))
		}
		c.String(http.StatusOK, fmt.Sprintf("%d files uploaded!\n", len(files)))
		return
	})
	router.Run(":8080")
}