package resty_log

import (
	"context"
	"log/slog"
)

// 实现 LogSaver 接口
type TestEntLogSaver struct{}

func (s *TestEntLogSaver) SaveLog(ctx context.Context, data *LogData) error {
	slog.WarnContext(ctx, "[接口方式]准备保存日志",
		slog.String("Hostname", data.Hostname),
		slog.String("Method", data.Method),
		slog.String("URL", data.URL),
	)
	if data == nil {
		return nil
	}
	slog.InfoContext(ctx, "[接口方式]执行保存日志",
		slog.String("Hostname", data.Hostname),
		slog.String("Method", data.Method),
		slog.String("URL", data.URL),
		slog.Int("StatusCode", data.StatusCode),
		slog.Int64("ElapseTime", data.ElapseTime),
		slog.Bool("IsError", data.IsError),
	)
	return nil
}

// 返回一个测试的 Logger
func NewTestLogger(debug bool) *Logger {

	// 定义回调函数
	saveFunc := func(ctx context.Context, data *LogData) error {
		slog.WarnContext(ctx, "[回调方式]准备保存日志",
			slog.String("Hostname", data.Hostname),
			slog.String("Method", data.Method),
			slog.String("URL", data.URL),
		)
		if data == nil {
			return nil
		}
		slog.InfoContext(ctx, "[回调方式]执行保存日志",
			slog.String("Hostname", data.Hostname),
			slog.String("Method", data.Method),
			slog.String("URL", data.URL),
			slog.Int("StatusCode", data.StatusCode),
			slog.Int64("ElapseTime", data.ElapseTime),
			slog.Bool("IsError", data.IsError),
		)
		return nil
	}

	// 创建 LogSaver 实现
	saver := &TestEntLogSaver{}

	// 初始化 Logger，传入回调和接口（回调优先）
	return NewLogger(saveFunc, saver, debug)
}
