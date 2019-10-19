# SecureJSON
- 使用SecureJSON防止json劫持。
- 如果给定的结构是数组值，则默认值会将`"while(1),"`前置到响应主体。 

```go
package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	r := gin.Default()
	
	// You can also use your own secure json prefix
	// r.SecureJsonPrefix(")]}',\n")
	// http://127.0.0.1:8080/someJSON
	r.GET("/someJSON", func(c *gin.Context) {
		names := []string{"lena", "austin", "foo"}
		
		// Will output  :
		// while(1);["lena","austin","foo"]
		c.SecureJSON(http.StatusOK, names)
	})
	
	// Listen and serve on 0.0.0.0:8080
	r.Run(":8080")
}
```
---

#### SecureJSON

Using SecureJSON to prevent json hijacking. Default prepends `"while(1),"` to response body if the given struct is array values.

```go
func main() {
	r := gin.Default()

	// You can also use your own secure json prefix
	// r.SecureJsonPrefix(")]}',\n")

	r.GET("/someJSON", func(c *gin.Context) {
		names := []string{"lena", "austin", "foo"}

		// Will output  :   while(1);["lena","austin","foo"]
		c.SecureJSON(http.StatusOK, names)
	})

	// Listen and serve on 0.0.0.0:8080
	r.Run(":8080")
}
```