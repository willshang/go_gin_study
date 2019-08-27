package _4_cors

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// 中间件所有可选项
type Config struct {
	// 1.
	AllowAllOrigins bool

	// 2.
	// AllowOrigins是一个可以执行跨域请求的源列表。
	// 如果列表中存在特殊的“*”值，则允许使用所有源。
	// 默认值为[]
	AllowOrigins []string

	// 3.
	// AllowOriginFunc是一个用于验证origin的自定义函数。
	// 它将origin作为参数，如果允许返回true，否则返回false。
	// 如果设置此选项，则会忽略AllowOrigins的内容。
	AllowOriginFunc func(origin string) bool

	// 4.
	// AllowMethods是允许客户端对跨域请求使用的方法列表。
	// 默认值是简单方法(GET和POST)
	AllowMethods []string

	// 5.
	// AllowHeaders是允许客户端在跨域请求中使用的非简单头文件列表。
	AllowHeaders []string

	// 6.
	// AllowCredentials表示请求是否可以包含用户凭证(如cookie、HTTP身份验证或客户端SSL证书)。
	AllowCredentials bool

	// 7.
	// ExposeHeaders说明哪些头可以安全地暴露给CORS API规范的API
	ExposeHeaders []string

	// 8
	// MaxAge表示可以缓存预检请求的结果的时间(以秒为单位)
	MaxAge time.Duration

	// 9
	// 允许添加 origins 例如:
	// http://some-domain/*, https://api.* or http://some.*.subdomain.com
	AllowWildcard bool

	// 10
	// 允许使用流行的浏览器扩展模式
	AllowBrowserExtensions bool

	// 11
	// 允许使用WebSocket协议
	AllowWebSockets bool

	// 12
	// 允许使用file://schema(危险!)，只有在100%确定需要时才使用它
	AllowFiles bool
}

func (c *Config) AddAllowMethods(methods ...string) {
	c.AllowMethods = append(c.AllowMethods, methods...)
}

func (c *Config) AddAllowHeaders(headers ...string) {
	c.AllowHeaders = append(c.AllowHeaders, headers...)
}

func (c *Config) AddExposeHeaders(headers ...string) {
	c.ExposeHeaders = append(c.ExposeHeaders, headers...)
}

func (c Config) getAllowedSchemas() []string {
	allowedSchemas := DefaultSchemas
	if c.AllowBrowserExtensions {
		allowedSchemas = append(allowedSchemas, ExtensionSchemas...)
	}
	if c.AllowWebSockets {
		allowedSchemas = append(allowedSchemas, WebSocketSchemas...)
	}
	if c.AllowFiles {
		allowedSchemas = append(allowedSchemas, FileSchemas...)
	}
	return allowedSchemas
}

func (c Config) validateAllowedSchemas(origin string) bool {
	allowedSchemas := c.getAllowedSchemas()
	for _, schema := range allowedSchemas {
		if strings.HasPrefix(origin, schema) {
			return true
		}
	}
	return false
}

func (c *Config) Validate() error {
	if c.AllowAllOrigins && (c.AllowOriginFunc != nil || len(c.AllowOrigins) > 0) {
		return errors.New("conflict settings: all origins are allowed. " +
			"AllowOriginFunc or  AllowOrigins is not needed")
	}
	if !c.AllowAllOrigins && c.AllowOriginFunc == nil && len(c.AllowOrigins) == 0 {
		return errors.New("conflict settings: all origins disabled")
	}

	for _, origin := range c.AllowOrigins {
		if origin == "*" {
			c.AllowAllOrigins = true
			return nil
		} else if !strings.Contains(origin, "*") && !c.validateAllowedSchemas(origin) {
			return errors.New("bad origin: origins must contain '*' or include " +
				strings.Join(c.getAllowedSchemas(), ","))
		}
	}
	return nil
}

func (c Config) parseWildcardRules() [][]string {
	var wRules [][]string

	if !c.AllowWildcard {
		return wRules
	}

	for _, o := range c.AllowOrigins {
		if !strings.Contains(o, "*") {
			continue
		}

		if c := strings.Count(o, "*"); c > 1 {
			panic(errors.New("only one * is allowed").Error())
		}

		i := strings.Index(o, "*")
		if i == 0 {
			wRules = append(wRules, []string{"*", o[1:]})
			continue
		}
		if i == (len(o) - 1) {
			wRules = append(wRules, []string{o[:i-1], "*"})
			continue
		}
		wRules = append(wRules, []string{o[:i], o[i+1:]})
	}
	return wRules
}

func DefaultConfig() Config {
	return Config{
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "HEAD"},
		AllowHeaders:     []string{"Origin", "Content-Length", "Content-Type"},
		AllowCredentials: false,
		MaxAge:           12 * time.Hour,
	}
}

func Default() gin.HandlerFunc {
	config := DefaultConfig()
	config.AllowAllOrigins = true
	return New(config)
}

func New(config Config) gin.HandlerFunc {
	cors := newCors(config)
	return func(c *gin.Context) {
		cors.applyCors(c)
	}
}

// config.go

type cors struct {
	allowAllOrigins  bool
	allowCredentials bool
	allowOriginFunc  func(string) bool
	allowOrigins     []string
	exposeHeaders    []string
	normalHeaders    http.Header
	preflightHeaders http.Header
	wildcardOrigins  [][]string
}

var (
	DefaultSchemas = []string{
		"http://",
		"https://",
	}
	ExtensionSchemas = []string{
		"chrome-extension://",
		"safari-extension://",
		"moz-extension://",
		"ms-browser-extension://",
	}
	FileSchemas = []string{
		"file://",
	}
	WebSocketSchemas = []string{
		"ws://",
		"wss://",
	}
)

func newCors(config Config) *cors {
	if err := config.Validate(); err != nil {
		panic(err.Error())
	}

	return &cors{
		allowOriginFunc:  config.AllowOriginFunc,
		allowAllOrigins:  config.AllowAllOrigins,
		allowCredentials: config.AllowCredentials,
		allowOrigins:     normalize(config.AllowOrigins),
		normalHeaders:    generateNormalHeaders(config),
		preflightHeaders: generatePreflightHeaders(config),
		wildcardOrigins:  config.parseWildcardRules(),
	}
}

func (cors *cors) applyCors(c *gin.Context) {
	origin := c.Request.Header.Get("Origin")
	if len(origin) == 0 {
		return
	}
	host := c.Request.Header.Get("Host")

	if origin == "http://"+host || origin == "https://"+host {
		return
	}

	if !cors.validateOrigin(origin) {
		c.AbortWithStatus(http.StatusForbidden)
		return
	}

	if c.Request.Method == "OPTIONS" {
		cors.handlePreflight(c)
		defer c.AbortWithStatus(http.StatusNoContent)
	} else {
		cors.handleNormal(c)
	}

	if !cors.allowAllOrigins {
		c.Header("Access-Control-Allow-Origin", origin)
	}
}

func (cors *cors) validateWildcardOrigin(origin string) bool {
	for _, w := range cors.wildcardOrigins {
		if w[0] == "*" && strings.HasSuffix(origin, w[1]) {
			return true
		}
		if w[1] == "*" && strings.HasPrefix(origin, w[0]) {
			return true
		}
		if strings.HasPrefix(origin, w[0]) && strings.HasSuffix(origin, w[1]) {
			return true
		}
	}
	return false
}

func (cors *cors) validateOrigin(origin string) bool {
	if cors.allowAllOrigins {
		return true
	}

	for _, value := range cors.allowOrigins {
		if value == origin {
			return true
		}
	}

	if len(cors.wildcardOrigins) > 0 && cors.validateWildcardOrigin(origin) {
		return true
	}

	if cors.allowOriginFunc != nil {
		return cors.allowOriginFunc(origin)
	}

	return false
}

func (cors *cors) handlePreflight(c *gin.Context) {
	header := c.Writer.Header()
	for key, value := range cors.preflightHeaders {
		header[key] = value
	}
}

func (cors *cors) handleNormal(c *gin.Context) {
	header := c.Writer.Header()
	for key, value := range cors.normalHeaders {
		header[key] = value
	}
}

// utils.go
type converter func(string) string

func generateNormalHeaders(c Config) http.Header {
	headers := make(http.Header)
	if c.AllowCredentials {
		headers.Set("Access-Control-Allow-Credentials", "true")
	}
	if len(c.ExposeHeaders) > 0 {
		exposeHeaders := convert(normalize(c.ExposeHeaders), http.CanonicalHeaderKey)
		headers.Set("Access-Control-Expose-Headers", strings.Join(exposeHeaders, ","))
	}
	if c.AllowAllOrigins {
		headers.Set("Access-Control-Allow-Origin", "*")
	} else {
		headers.Set("Vary", "Origin")
	}
	return headers
}

func generatePreflightHeaders(c Config) http.Header {
	headers := make(http.Header)
	if c.AllowCredentials {
		headers.Set("Access-Control-Allow-Credentials", "true")
	}
	if len(c.AllowMethods) > 0 {
		allowMethods := convert(normalize(c.AllowMethods), strings.ToUpper)
		value := strings.Join(allowMethods, ",")
		headers.Set("Access-Control-Allow-Methods", value)
	}
	if len(c.AllowHeaders) > 0 {
		allowHeaders := convert(normalize(c.AllowHeaders), http.CanonicalHeaderKey)
		value := strings.Join(allowHeaders, ",")
		headers.Set("Access-Control-Allow-Headers", value)
	}
	if c.MaxAge > time.Duration(0) {
		value := strconv.FormatInt(int64(c.MaxAge/time.Second), 10)
		headers.Set("Access-Control-Max-Age", value)
	}
	if c.AllowAllOrigins {
		headers.Set("Access-Control-Allow-Origin", "*")
	} else {
		headers.Add("Vary", "Origin")
		headers.Add("Vary", "Access-Control-Request-Method")
		headers.Add("Vary", "Access-Control-Request-Headers")
	}
	return headers
}

func normalize(values []string) []string {
	if values == nil {
		return nil
	}

	distinctMap := make(map[string]bool, len(values))
	normalized := make([]string, 0, len(values))

	for _, value := range values {
		value = strings.TrimSpace(value)
		value = strings.ToLower(value)
		if _, seen := distinctMap[value]; !seen {
			normalized = append(normalized, value)
			distinctMap[value] = true
		}
	}
	return normalized
}

func convert(s []string, c converter) []string {
	var out []string
	for _, i := range s {
		out = append(out, c(i))
	}
	return out
}
