# 测试
- 55-Tests.go
- 55-Tests_test.go
- HTTP 测试首选 net/http/httptest 包。
```go
package main

import "github.com/gin-gonic/gin"

func setupRouter() *gin.Engine {
	r := gin.Default()
	// http://localhost:8081/ping
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
	return r
}

func main() {
	r := setupRouter()
	r.Run(":8081")
}
```
- 
```go
package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	
	"github.com/stretchr/testify/assert"
)

// go test 55-Tests.go 55-Tests_test.go
func TestPingRoute(t *testing.T) {
	router := setupRouter()
	
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/ping", nil)
	router.ServeHTTP(w, req)
	
	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "pong", w.Body.String())
}
```
---
## Testing

The `net/http/httptest` package is preferable way for HTTP testing.

```go
package main

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
	return r
}

func main() {
	r := setupRouter()
	r.Run(":8080")
}
```

Test for code example above:

```go
package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPingRoute(t *testing.T) {
	router := setupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/ping", nil)
	router.ServeHTTP(w, req)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, "pong", w.Body.String())
}
```