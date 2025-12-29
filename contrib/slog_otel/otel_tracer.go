package slog_otel

import (
	"context"
	"log/slog"

	"go.opentelemetry.io/contrib/bridges/otelslog"
)

// NewOtelMiddleware 返回一个 slog 中间件：将日志同时导出到 OpenTelemetry
func NewOtelMiddleware(name string) slog.Handler {
	if name == "" {
		name = "slog"
	}
	return otelslog.NewHandler(name)
}

// multiHandler 会将日志分发到多个 slog.Handler
type multiHandler struct {
	hs []slog.Handler
}

func (m multiHandler) Enabled(ctx context.Context, level slog.Level) bool {
	for _, h := range m.hs {
		if h.Enabled(ctx, level) {
			return true
		}
	}
	return false
}

func (m multiHandler) Handle(ctx context.Context, r slog.Record) error {
	var firstErr error
	for _, h := range m.hs {
		// 克隆 Record，避免下游修改互相影响
		if err := h.Handle(ctx, r.Clone()); err != nil && firstErr == nil {
			firstErr = err
		}
	}
	return firstErr
}

func (m multiHandler) WithAttrs(attrs []slog.Attr) slog.Handler {
	nhs := make([]slog.Handler, len(m.hs))
	for i, h := range m.hs {
		nhs[i] = h.WithAttrs(attrs)
	}
	return multiHandler{hs: nhs}
}

func (m multiHandler) WithGroup(name string) slog.Handler {
	nhs := make([]slog.Handler, len(m.hs))
	for i, h := range m.hs {
		nhs[i] = h.WithGroup(name)
	}
	return multiHandler{hs: nhs}
}
