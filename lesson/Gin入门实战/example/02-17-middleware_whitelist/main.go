package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()
	r.Use(IPAuthMiidelware())
	r.GET("/test", func(c *gin.Context) {
		c.String(200, "hello test")
	})

	r.Run()
}

func IPAuthMiidelware() gin.HandlerFunc {
	return func(c *gin.Context) {
		ipList := []string{
			"127.0.0.1",
		}
		flag := false
		clientIP := c.ClientIP()
		for _, host := range ipList {
			if host == clientIP {
				flag = true
				break
			}
		}
		if !flag {
			c.String(401, "%s, not in iplist, ", clientIP)
			c.Abort()
		}
	}
}
