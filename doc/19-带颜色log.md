# 控制log输出颜色
- 17-log-without-color.go
- 18-log-with-color.go
- 不带颜色的log输出
``` 
package main

import "github.com/gin-gonic/gin"

func main() {
	// Disable log's color
	// 禁用log颜色
	gin.DisableConsoleColor()

	// Creates a gin router with default middleware:
	// logger and recovery (crash-free) middleware
	// 创建一个带默认中间件(logger和recovery)的中间件
	router := gin.Default()
	
	// http://localhost:8080/ping
	router.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	router.Run(":8080")
}
```
- 强制控制台带颜色的log输出
``` 
package main

import "github.com/gin-gonic/gin"

func main() {
	// Force log's color
	// ForceConsoleColor 强制在控制台输出颜色
	gin.ForceConsoleColor()

	// Creates a gin router with default middleware:
	// logger and recovery (crash-free) middleware
	// 创建一个带默认中间件(logger和recovery)的中间件
	router := gin.Default()

	// http://localhost:8080/ping
	router.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})

	router.Run(":8080")
}
```
---
### Controlling Log output coloring 

By default, logs output on console should be colorized depending on the detected TTY.

Never colorize logs: 

```go
func main() {
    // Disable log's color
    gin.DisableConsoleColor()
    
    // Creates a gin router with default middleware:
    // logger and recovery (crash-free) middleware
    router := gin.Default()
    
    router.GET("/ping", func(c *gin.Context) {
        c.String(200, "pong")
    })
    
    router.Run(":8080")
}
```

Always colorize logs: 

```go
func main() {
    // Force log's color
    gin.ForceConsoleColor()
    
    // Creates a gin router with default middleware:
    // logger and recovery (crash-free) middleware
    router := gin.Default()
    
    router.GET("/ping", func(c *gin.Context) {
        c.String(200, "pong")
    })
    
    router.Run(":8080")
}
```