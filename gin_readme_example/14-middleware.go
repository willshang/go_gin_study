package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	// Creates a router without any middleware by default
	// 创建一个不带中间件的路由
	r := gin.New()

	// Global middleware
	// Logger middleware will write the logs to gin.DefaultWriter even if you set with GIN_MODE=release.
	// By default gin.DefaultWriter = os.Stdout
	// 全局中间件
	// Logger中间件会将日志写入 gin.DefaultWriter，即使您设置了GIN_MODE=release。
	// 默认情况下，gin.DefaultWriter = os.Stdout
	r.Use(gin.Logger())

	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	// Recovery中间件从任何panic中恢复，如果有，则写入500。
	r.Use(gin.Recovery())

	// Per route middleware, you can add as many as you desire.
	// 对于每个路由中间件，您可以根据需要添加任意多个。
	r.GET("/benchmark", MyBenchLogger(), benchEndpoint)

	// Authorization group
	// authorized := r.Group("/", AuthRequired())
	// exactly the same as:
	// 授权组
	// authorized := r.Group("/", AuthRequired())
	// 可以用下面的方式
	authorized := r.Group("/")

	// per group middleware! in this case we use the custom created
	// AuthRequired() middleware just in the "authorized" group.
	// 每组中间件！在本例中，我们使用自定义创建的
	// AuthRequired中间件仅在 authorized 组内
	authorized.Use(AuthRequired())
	{
		authorized.POST("/login", loginEndpoint)
		authorized.POST("/submit", submitEndpoint)
		authorized.POST("/read", readEndpoint)

		// nested group
		testing := authorized.Group("testing")
		testing.GET("/analytics", analyticsEndpoint)
	}

	// Listen and serve on 0.0.0.0:8080
	r.Run(":8080")
}

func loginEndpoint(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "login",
		"method":  c.Request.Method,
		"url":     c.Request.URL.Path,
	})
	return
}

func submitEndpoint(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "submt",
		"method":  c.Request.Method,
		"url":     c.Request.URL.Path,
	})
	return
}

func readEndpoint(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "read",
		"method":  c.Request.Method,
		"url":     c.Request.URL.Path,
	})
	return
}
