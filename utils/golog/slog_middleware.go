package golog

import "log/slog"

// SlogMiddleware 允许在 slog.Handler 链中插入可插拔的处理环节（装饰器）
type SlogMiddleware func(slog.Handler) slog.Handler
