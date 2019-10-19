package middleware

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"go_gin_study/lesson/Gin入门实战/common/lib"
	"go_gin_study/lesson/Gin入门实战/demo/public"
	"io/ioutil"
	"time"
)

func requestInLog(c *gin.Context) {
	traceContext := lib.NewTrace()
	if traceID := c.Request.Header.Get("com-header-id"); traceID != "" {
		traceContext.TraceID = traceID
	}
	if spanID := c.Request.Header.Get("com-header-spanid"); spanID != "" {
		traceContext.SpanID = spanID
	}

	c.Set("startExecTime", time.Now())
	c.Set("trace", traceContext)

	bodyBytes, _ := ioutil.ReadAll(c.Request.Body)
	c.Request.Body = ioutil.NopCloser(bytes.NewBuffer(bodyBytes))

	lib.Log.TagInfo(traceContext, "_com_request_in", map[string]interface{}{
		"uri":    c.Request.RequestURI,
		"method": c.Request.Method,
		"args":   c.Request.PostForm,
		"body":   string(bodyBytes),
		"from":   c.ClientIP(),
	})
}

func requestOutLog(c *gin.Context) {
	// after request
	endExecTime := time.Now()
	response, _ := c.Get("response")
	st, _ := c.Get("startExecTime")

	startExecTime, _ := st.(time.Time)
	public.ComLogNotice(c, "_com_request_out", map[string]interface{}{
		"uri":       c.Request.RequestURI,
		"method":    c.Request.Method,
		"args":      c.Request.PostForm,
		"from":      c.ClientIP(),
		"response":  response,
		"proc_time": endExecTime.Sub(startExecTime).Seconds(),
	})
}

func RequestLog() gin.HandlerFunc {
	return func(c *gin.Context) {
		requestInLog(c)
		defer requestOutLog(c)
		c.Next()
	}
}
