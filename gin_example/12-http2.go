package main

import (
	"github.com/gin-gonic/gin"
	"html/template"
	"log"
	"net/http"
	"os"
)

var htmlx = template.Must(template.New("https").Parse(`
<html>
<head>
	<title>Https Test</title>
</head>
<body>
	<h1 style="color:red;">Welcome, Ginner!</h1>
</body>
</html>
`))

func main() {
	logger := log.New(os.Stderr, "", 0)
	logger.Println("[WARNING] DON'T USE THE EMBED CERTS FROM THIS EXAMPLE IN PRODUCTION ENVIRONMENT, GENERATE YOUR OWN!")
	r := gin.Default()
	r.SetHTMLTemplate(htmlx)

	r.GET("/welcome", func(c *gin.Context) {
		c.HTML(http.StatusOK, "https", gin.H{
			"status": "success",
		})
	})

	r.RunTLS(":8080", "./testdata/server.pem", "./testdata/server.key")
}
