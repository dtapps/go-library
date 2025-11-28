package golog

import (
	"context"
	"fmt"
	"log/slog"
)

type ContextHandler struct {
	slog.Handler
}

// Handle 添加上下文属性到 Record 中，然后调用底层的 handler
func (h ContextHandler) Handle(ctx context.Context, r slog.Record) error {

	xRequestID := getRequestIDContext(ctx)
	if xRequestID != "" {
		r.AddAttrs(slog.String("X-Request-ID", xRequestID))
	}

	return h.Handler.Handle(ctx, r)
}

// 获取请求编号
func getRequestIDContext(ctx context.Context) string {
	traceId := fmt.Sprintf("%s", ctx.Value("X-Request-ID"))
	if traceId == "%!s(<nil>)" {
		return ""
	}
	if len(traceId) <= 0 {
		return ""
	}
	return traceId
}
