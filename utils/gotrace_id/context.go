package gotrace_id

import (
	"context"
	"fmt"
	"github.com/dtapps/go-library/utils/gostring"
	"github.com/gin-gonic/gin"
)

// CustomTraceIdContext 自定义设置跟踪编号上下文
func CustomTraceIdContext(ctx context.Context) context.Context {
	return context.WithValue(ctx, "trace_id", gostring.GetUuId())
}

// SetCustomTraceId 自定义设置跟踪编号上下文
func SetCustomTraceId(ctx context.Context, traceId string) context.Context {
	return context.WithValue(ctx, "trace_id", traceId)
}

// SetGinTraceIdContext 设置跟踪编号上下文
func SetGinTraceIdContext(ctx context.Context, c *gin.Context) context.Context {
	return context.WithValue(ctx, "trace_id", GetGinTraceId(c))
}

// GetTraceIdContext 通过上下文获取跟踪编号
func GetTraceIdContext(ctx context.Context) string {
	traceId := fmt.Sprintf("%s", ctx.Value("trace_id"))
	if traceId == Nil {
		return ""
	}
	if len(traceId) <= 0 {
		return ""
	}
	return traceId
}
