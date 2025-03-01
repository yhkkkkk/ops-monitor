package middleware

import (
	"ops-monitor/pkg/ctx"
	"ops-monitor/pkg/tools"

	"github.com/gin-gonic/gin"
)

const (
	TraceIDKey    = "X-Trace-ID"
	TraceIDCtxKey = "trace_id"
	LogContent    = "log_content"
)

// TraceMiddleware 链路追踪中间件
func TraceMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		traceID := c.GetHeader(TraceIDKey)

		// 创建带有 TraceID 的新 context
		newCtx := tools.NewTraceContext(c.Request.Context(), traceID)
		c.Request = c.Request.WithContext(newCtx)

		// 更新 Context 结构体中的 Ctx
		ctx.SetRequestContext(newCtx)

		if traceID == "" {
			traceID = tools.GetTraceID(newCtx)
		}
		c.Header(TraceIDKey, traceID)

		c.Next()
	}
}
