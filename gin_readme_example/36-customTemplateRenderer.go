package main

import (
	"github.com/gin-gonic/gin"
	"html/template"
)

func main() {
	router := gin.Default()
	html := template.Must(template.ParseFiles("templates/file1", "templates/file2"))
	router.SetHTMLTemplate(html)
	router.Run(":8080")
}
