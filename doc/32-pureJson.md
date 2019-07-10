# PureJSON
- 31-pureJSON.go
- JSON 使用 unicode 替换特殊 HTML 字符，例如 < 变为 \ u003c。
- 如果要按字面对这些字符进行编码，则可以使用 PureJSON。Go 1.6 及更低版本无法使用此功能。
```go
package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()

	// Serves unicode entities
	// 提供 unicode 实体
	// http://127.0.0.1:8080/json
	r.GET("/json", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"html": "<b>Hello, world!</b>",
		})
	})

	// Serves literal characters
	// 提供字面字符
	// http://127.0.0.1:8080/purejson
	r.GET("/purejson", func(c *gin.Context) {
		c.PureJSON(200, gin.H{
			"html": "<b>Hello, world!</b>",
		})
	})

	// listen and serve on 0.0.0.0:8080
	r.Run(":8080")
}
```
---
#### PureJSON

Normally, JSON replaces special HTML characters with their unicode entities, e.g. `<` becomes  `\u003c`. If you want to encode such characters literally, you can use PureJSON instead.
This feature is unavailable in Go 1.6 and lower.

```go
func main() {
	r := gin.Default()
	
	// Serves unicode entities
	r.GET("/json", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"html": "<b>Hello, world!</b>",
		})
	})
	
	// Serves literal characters
	r.GET("/purejson", func(c *gin.Context) {
		c.PureJSON(200, gin.H{
			"html": "<b>Hello, world!</b>",
		})
	})
	
	// listen and serve on 0.0.0.0:8080
	r.Run(":8080")
}
```