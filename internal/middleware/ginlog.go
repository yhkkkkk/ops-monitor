package middleware

import (
	"bytes"
	"io"
	"time"

	"github.com/gin-gonic/gin"

	"ops-monitor/pkg/logger"
	"ops-monitor/pkg/tools"
)

// LogConfig 日志配置
type LogConfig struct {
	SkipPaths   []string // 跳过不需要记录日志的路径
	MaxBodySize int64    // 最大记录的请求/响应体大小
}

var DefaultSkipPaths = []string{
	"/metrics",
	"/debug/pprof/",
	"/api/health",
	"/debug/pprof/cmdline",
	"/debug/pprof/profile",
	"/debug/pprof/symbol",
	"/debug/pprof/trace",
	"/favicon.ico",
}

// GinZapLogger gin日志中间件
func GinZapLogger(config ...LogConfig) gin.HandlerFunc {
	var conf LogConfig
	if len(config) > 0 {
		conf = config[0]
	}
	if conf.SkipPaths == nil {
		conf.SkipPaths = DefaultSkipPaths
	}
	if conf.MaxBodySize == 0 {
		conf.MaxBodySize = 1024 * 1024 * 2 // 2MB
	}

	return func(c *gin.Context) {
		// 跳过不需要记录日志的路径
		path := c.Request.URL.Path
		for _, skip := range conf.SkipPaths {
			if path == skip {
				c.Next()
				return
			}
		}

		start := time.Now()
		method := c.Request.Method
		query := c.Request.URL.RawQuery
		requestBody := readRequestBody(c, conf.MaxBodySize)

		// 创建日志上下文
		lc := logger.NewLogContext().
			WithTraceID(tools.GetTraceID(c.Request.Context())).
			WithAction(path)

		// 创建响应体记录器
		blw := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: c.Writer}
		c.Writer = blw

		// 处理请求
		c.Next()

		// 记录响应时间
		latency := time.Since(start)

		// 构建日志参数
		params := map[string]interface{}{
			"method":     method,
			"path":       path,
			"query":      query,
			"status":     c.Writer.Status(),
			"latency":    latency.String(),
			"client_ip":  c.ClientIP(),
			"user_agent": c.Request.UserAgent(),
		}

		// 添加请求和响应体（如果存在）
		if requestBody != "" {
			params["request_body"] = requestBody
		}
		if blw.body.String() != "" {
			params["response_body"] = blw.body.String()
		}

		// 记录日志
		if len(c.Errors) > 0 {
			// 收集所有错误
			var errMsgs []string
			for _, e := range c.Errors {
				errMsgs = append(errMsgs, e.Error())
			}
			params["errors"] = errMsgs
			logger.Error(c.Request.Context(), lc.WithParams(params), c.Errors.Last().Err)
		} else {
			logger.Info(c.Request.Context(), lc.WithParams(params))
		}
	}
}

// bodyLogWriter 响应体记录器
type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

// readRequestBody 读取请求体
func readRequestBody(c *gin.Context, maxSize int64) string {
	if c.Request.Body == nil {
		return ""
	}

	if c.Request.ContentLength > maxSize {
		return "request body too large"
	}

	bodyBytes, err := io.ReadAll(c.Request.Body)
	if err != nil {
		return ""
	}

	c.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

	return string(bodyBytes)
}
