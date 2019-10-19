package _2_commonlog

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"io"
	"strconv"
	"sync"
	"time"
)

// 实例一个日志中间件，将日志写入gin.DefaultWriter。gin.DefaultWriter = os.Stdout
func New() gin.HandlerFunc {
	return NewWithWriter(gin.DefaultWriter)
}

// 指定一个writter buffer, 实例一个日志中间件
// writter buffer 可以是 os.Stdout, 写文件， socket
func NewWithWriter(out io.Writer) gin.HandlerFunc {
	pool := &sync.Pool{
		New: func() interface{} {
			buf := new(bytes.Buffer)
			return buf
		},
	}

	return func(c *gin.Context) {
		c.Next()

		w := pool.Get().(*bytes.Buffer)
		w.Reset()
		w.WriteString("[common_log]")
		w.WriteString(c.ClientIP())
		w.WriteString(" - - ")
		w.WriteString(time.Now().Format("[02/Jan/2006:15:04:05 -0700] "))
		w.WriteString("\"")
		w.WriteString(c.Request.Method)
		w.WriteString(" ")
		w.WriteString(c.Request.URL.Path)
		w.WriteString(" ")
		w.WriteString(c.Request.Proto)
		w.WriteString("\" ")
		w.WriteString(strconv.Itoa(c.Writer.Status()))
		w.WriteString(" ")
		w.WriteString(strconv.Itoa(c.Writer.Size()))
		w.WriteString("\n")
		w.WriteTo(out)
		pool.Put(w)
	}
}
