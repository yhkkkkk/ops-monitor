package tools

import (
	"context"
	"strings"

	"github.com/google/uuid"
)

type traceKey struct{}

// NewTraceContext 创建带有 TraceID 的上下文
func NewTraceContext(ctx context.Context, traceID string) context.Context {
	if traceID == "" {
		traceID = generateTraceID()
	}
	return context.WithValue(ctx, traceKey{}, traceID)
}

// GetTraceID 从上下文获取 TraceID
func GetTraceID(ctx context.Context) string {
	if id, ok := ctx.Value(traceKey{}).(string); ok {
		return id
	}
	return ""
}

// generateTraceID 生成追踪ID
func generateTraceID() string {
	return strings.ReplaceAll(uuid.New().String(), "-", "")
}
