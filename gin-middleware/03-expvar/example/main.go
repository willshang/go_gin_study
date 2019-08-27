package main

import (
	"github.com/gin-gonic/gin"
	expvar_middleware "go_gin_study/gin-middleware/03-expvar"
)

func main() {
	r := gin.Default()

	r.GET("/debug/vars", expvar_middleware.Handler())
	r.Run()
}
