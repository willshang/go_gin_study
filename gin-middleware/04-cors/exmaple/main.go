package main

import (
	"github.com/gin-gonic/gin"
	cors_new "go_gin_study/gin-middleware/04-cors"
	"time"
)

func main() {
	router := gin.Default()

	router.Use(cors_new.New(cors_new.Config{
		AllowOrigins:     []string{"https://foo.com"},
		AllowMethods:     []string{"PUT", "PATCH"},
		AllowHeaders:     []string{"Origin"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			return origin == "https://github.com"
		},
		MaxAge: 12 * time.Hour,
	}))
	router.Run()
}
