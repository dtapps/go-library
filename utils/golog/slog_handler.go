package golog

import (
	"context"
	"go.dtapp.net/library/utils/gorequest"
	"go.opentelemetry.io/otel/trace"
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

	// OpenTelemetry追踪
	span := trace.SpanFromContext(ctx)
	if span.SpanContext().IsValid() {
		r.AddAttrs(slog.String("trace_id", span.SpanContext().TraceID().String()))
		r.AddAttrs(slog.String("span_id", span.SpanContext().SpanID().String()))
	}

	return h.Handler.Handle(ctx, r)
}
