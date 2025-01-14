package golog

import (
	"context"
	"go.dtapp.net/library/utils/gorequest"
	"log/slog"
)

type ContextHandler struct {
	slog.Handler
}

// Handle 添加上下文属性到 Record 中，然后调用底层的 handler
func (h ContextHandler) Handle(ctx context.Context, r slog.Record) error {

	xRequestID := gorequest.GetRequestIDContext(ctx)
	if xRequestID != "" {
		r.AddAttrs(slog.String("X-Request-ID", xRequestID))
	}

	return h.Handler.Handle(ctx, r)
}
