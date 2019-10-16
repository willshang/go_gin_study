package middleware

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
)

type ResponseCode int

const (
	SuccessCode             ResponseCode = 0
	UndefErrorCode          ResponseCode = 1
	ValidErrorCode          ResponseCode = 2
	InternalErrorCode       ResponseCode = 3
	InvalidRequestErrorCode ResponseCode = 401
	CustomizeCode           ResponseCode = 1000
	GroupAllSaveFlowError   ResponseCode = 2001
)

type Response struct {
	ErrorCode ResponseCode `json:"error_code"`
	ErrorMsg  string       `json:"error_msg"`
	Data      interface{}  `json:"data"`
	TraceID   interface{}  `json:"trace_id"`
}

func ResponseError(c *gin.Context, code ResponseCode, err error) {
	trace, _ := c.Get("trace")
	traceContext, _ := trace.(*lib.TraceContext)
	traceID := ""
	if traceContext != nil {
		traceID = traceContext.TraceID
	}

	resp := &Response{
		ErrorCode: code,
		ErrorMsg:  err.Error(),
		Data:      "",
		TraceID:   traceID,
	}
	c.JSON(200, resp)
	response, _ := json.Marshal(resp)
	c.Set("response", string(response))
	c.AbortWithError(200, err)
}

func ResponseSuccess(c *gin.Context, data interface{}) {
	trace, _ := c.Get("trace")
	traceContext, _ := trace.(*lib.TraceContext)
	traceID := ""
	if traceContext != nil {
		traceID = traceContext.TraceID
	}

	resp := &Response{
		ErrorCode: SuccessCode,
		ErrorMsg:  "",
		Data:      data,
		TraceID:   traceID,
	}
	c.JSON(200, resp)
	response, _ := json.Marshal(resp)
	c.Set("response", string(response))
}
