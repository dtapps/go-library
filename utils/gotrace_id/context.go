package gotrace_id

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// CustomTraceIdContext 自定义设置跟踪编号上下文
func CustomTraceIdContext(ctx context.Context) context.Context {
	var traceId = uuid.Must(uuid.NewV7()).String()
	return context.WithValue(ctx, TraceIdKey, traceId)
}

// SetCustomTraceId 自定义设置跟踪编号上下文
func SetCustomTraceId(ctx context.Context, traceId string) context.Context {
	return context.WithValue(ctx, TraceIdKey, traceId)
}

// SetGinTraceIdContext 设置跟踪编号上下文
func SetGinTraceIdContext(ctx context.Context, c *gin.Context) context.Context {
	var traceId = GetGinTraceId(c)
	return context.WithValue(ctx, TraceIdKey, traceId)
}

// GetTraceIdContext 通过上下文获取跟踪编号
func GetTraceIdContext(ctx context.Context) string {
	return CustomGetTraceIdContext(ctx, TraceIdKey)
}

// CustomGetTraceIdContext 通过自定义上下文获取跟踪编号
func CustomGetTraceIdContext(ctx context.Context, key string) string {
	traceId := fmt.Sprintf("%s", ctx.Value(key))
	if traceId == Nil {
		return ""
	}
	if len(traceId) <= 0 {
		return ""
	}
	return traceId
}
