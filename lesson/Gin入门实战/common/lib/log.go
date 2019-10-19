package lib

import (
	"fmt"
	dlog "go_gin_study/lesson/Gin入门实战/common/log"
	"strings"
)

const (
	DLTagUndefined     = "_com_undef"
	DLTagMySqlFailed   = "_com_mysql_failure"
	DLTagRedisFailed   = "_com_redis_failure"
	DLTagMySqlSuccess  = "_com_mysql_success"
	DLTagRedisSuccess  = "_com_redis_success"
	DLTagThriftFailed  = "_com_thrift_failure"
	DLTagThriftSuccess = "_com_thrift_success"
	DLTagHTTPSuccess   = "_com_http_success"
	DLTagHTTPFailed    = "_com_http_failure"
	DLTagTCPFailed     = "_com_tcp_failure"
	DLTagRequestIn     = "_com_request_in"
	DLTagRequestOut    = "_com_request_out"
)

const (
	_dlTag          = "dlTag"
	_traceId        = "traceID"
	_spanId         = "spanID"
	_childSpanId    = "cspanID"
	_dlTagBizPrefix = "_com_"
	_dlTagBizUndef  = "_com_undef"
)

var Log *Logger

type Trace struct {
	TraceID     string
	SpanID      string
	Caller      string
	SrcMethod   string
	HintCode    int64
	HintContent string
}

type TraceContext struct {
	Trace
	CSpanID string
}

type Logger struct {
}

func (l *Logger) TagInfo(trace *TraceContext, dlTag string, m map[string]interface{}) {
	m[_dlTag] = checkDLTag(dlTag)
	m[_traceId] = trace.TraceID
	m[_childSpanId] = trace.CSpanID
	m[_spanId] = trace.SpanID
	dlog.Info(parseParams(m))
}

func (l *Logger) TagWarn(trace *TraceContext, dlTag string, m map[string]interface{}) {
	m[_dlTag] = checkDLTag(dlTag)
	m[_traceId] = trace.TraceID
	m[_childSpanId] = trace.CSpanID
	m[_spanId] = trace.SpanID
	dlog.Warn(parseParams(m))
}

func (l *Logger) TagError(trace *TraceContext, dlTag string, m map[string]interface{}) {
	m[_dlTag] = checkDLTag(dlTag)
	m[_traceId] = trace.TraceID
	m[_childSpanId] = trace.CSpanID
	m[_spanId] = trace.SpanID
	dlog.Error(parseParams(m))
}

func (l *Logger) TagTrace(trace *TraceContext, dlTag string, m map[string]interface{}) {
	m[_dlTag] = checkDLTag(dlTag)
	m[_traceId] = trace.TraceID
	m[_childSpanId] = trace.CSpanID
	m[_spanId] = trace.SpanID
	dlog.Trace(parseParams(m))
}

func (l *Logger) TagDebug(trace *TraceContext, dlTag string, m map[string]interface{}) {
	m[_dlTag] = checkDLTag(dlTag)
	m[_traceId] = trace.TraceID
	m[_childSpanId] = trace.CSpanID
	m[_spanId] = trace.SpanID
	dlog.Debug(parseParams(m))
}

func (l *Logger) Close() {
	dlog.Close()
}

// 生成业务dlTag
func CreateBizDLTag(tagName string) string {
	if tagName == "" {
		return _dlTagBizUndef
	}

	return _dlTagBizPrefix + tagName
}

// 校验dlTag合法性
func checkDLTag(dlTag string) string {
	if strings.HasPrefix(dlTag, _dlTagBizPrefix) {
		return dlTag
	}

	if strings.HasPrefix(dlTag, "_com_") {
		return dlTag
	}

	if dlTag == DLTagUndefined {
		return dlTag
	}
	return dlTag
}

// map格式化为string
func parseParams(m map[string]interface{}) string {
	var dlTag string = "_undef"
	if _dlTag, _have := m["dlTag"]; _have {
		if __val, __ok := _dlTag.(string); __ok {
			dlTag = __val
		}
	}
	for _key, _val := range m {
		if _key == "dlTag" {
			continue
		}
		dlTag = dlTag + "||" + fmt.Sprintf("%v=%+v", _key, _val)
	}
	dlTag = strings.Trim(fmt.Sprintf("%q", dlTag), "\"")
	return dlTag
}
