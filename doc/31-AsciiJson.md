# AsciiJson
- 30-AsciiJson.go
- 使用AsciiJSON生成带有转义非Ascii字符的纯Ascii JSON。
```go
package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()

	// http://127.0.0.1:8080/someJSON
	r.GET("/someJSON", func(c *gin.Context) {
		data := gin.H{
			"lang": "GO语言",
			"tag":  "<br>",
		}

		// will output : {"lang":"GO\u8bed\u8a00","tag":"\u003cbr\u003e"}
		c.AsciiJSON(http.StatusOK, data)
		return
	})

	// Listen and serve on 0.0.0.0:8080
	r.Run(":8080")
}
```
---
#### AsciiJSON

Using AsciiJSON to Generates ASCII-only JSON with escaped non-ASCII chracters.

```go
func main() {
	r := gin.Default()

	r.GET("/someJSON", func(c *gin.Context) {
		data := gin.H{
			"lang": "GO语言",
			"tag":  "<br>",
		}

		// will output : {"lang":"GO\u8bed\u8a00","tag":"\u003cbr\u003e"}
		c.AsciiJSON(http.StatusOK, data)
	})

	// Listen and serve on 0.0.0.0:8080
	r.Run(":8080")
}
```