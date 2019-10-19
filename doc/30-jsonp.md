# jsonp
- 29-jsonp.go
- 使用jsonp从不同域中的服务器请求数据。如果查询参数回调存在，则向响应主体添加回调。
```go
package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	
	// http://127.0.0.1:8080/JSONP?callback=xxx
	r.GET("/JSONP", func(c *gin.Context) {
		data := gin.H{
			"foo": "bar",
		}
		
		// callback is x
		// Will output  :   x({\"foo\":\"bar\"})
		c.JSONP(http.StatusOK, data)
	})
	
	// Listen and serve on 0.0.0.0:8080
	r.Run(":8080")
	
	// client
	// curl http://127.0.0.1:8080/JSONP?callback=x
}
```
---
#### JSONP

Using JSONP to request data from a server  in a different domain. Add callback to response body if the query parameter callback exists.

```go
func main() {
	r := gin.Default()

	r.GET("/JSONP", func(c *gin.Context) {
		data := gin.H{
			"foo": "bar",
		}
		
		// callback is x
		// Will output  :   x({\"foo\":\"bar\"})
		c.JSONP(http.StatusOK, data)
	})

	// Listen and serve on 0.0.0.0:8080
	r.Run(":8080")

        // client
        // curl http://127.0.0.1:8080/JSONP?callback=x
}
```