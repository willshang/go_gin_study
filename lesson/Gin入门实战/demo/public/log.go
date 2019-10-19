package public

import (
	"context"
	"github.com/gin-gonic/gin"
	"go_gin_study/lesson/Gin入门实战/common/lib"
)

func ContextWarning(c context.Context, dlTag string, m map[string]interface{}) {
	v := c.Value("trace")
	traceContext, ok := v.(*lib.TraceContext)
	if !ok {
		traceContext = lib.NewTrace()
	}
	lib.Log.TagWarn(traceContext, dlTag, m)
}

func ContextError(c context.Context, dlTag string, m map[string]interface{}) {
	v := c.Value("trace")
	traceContext, ok := v.(*lib.TraceContext)
	if !ok {
		traceContext = lib.NewTrace()
	}
	lib.Log.TagError(traceContext, dlTag, m)
}

func ContextNotice(c context.Context, dlTag string, m map[string]interface{}) {
	v := c.Value("trace")
	traceContext, ok := v.(*lib.TraceContext)
	if !ok {
		traceContext = lib.NewTrace()
	}
	lib.Log.TagInfo(traceContext, dlTag, m)
}

func ComLogWarning(c *gin.Context, dlTag string, m map[string]interface{}) {
	traceContext := GetGinTraceContext(c)
	lib.Log.TagError(traceContext, dlTag, m)
}

func ComLogNotice(c *gin.Context, dlTag string, m map[string]interface{}) {
	traceContext := GetGinTraceContext(c)
	lib.Log.TagInfo(traceContext, dlTag, m)
}

func GetGinTraceContext(c *gin.Context) *lib.TraceContext {
	if c == nil {
		return lib.NewTrace()
	}
	traceContext, exists := c.Get("trace")
	if exists {
		if tc, ok := traceContext.(*lib.TraceContext); ok {
			return tc
		}
	}
	return lib.NewTrace()
}

func GetTraceContext(c *gin.Context) *lib.TraceContext {
	if c == nil {
		return lib.NewTrace()
	}
	traceContext := c.Value("trace")
	if tc, ok := traceContext.(*lib.TraceContext); ok {
		return tc
	}
	return lib.NewTrace()
}
