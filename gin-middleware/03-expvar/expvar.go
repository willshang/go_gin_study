package _3_expvar

import (
	"expvar"
	"fmt"
	"github.com/gin-gonic/gin"
)

func Handler() gin.HandlerFunc {
	return func(c *gin.Context) {
		w := c.Writer
		c.Header("Content-Type", "application/json; charset=utf-8")
		w.Write([]byte("{\n"))
		first := true
		expvar.Do(func(value expvar.KeyValue) {
			if !first {
				w.Write([]byte(",\n"))
			}
			first = false
			fmt.Fprintf(w, "%q: %s", value.Key, value.Value)
		})
		w.Write([]byte("\n}\n"))
		c.AbortWithStatus(200)
	}
}
