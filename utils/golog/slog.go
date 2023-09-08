package golog

import (
	"context"
	"fmt"
	"github.com/dtapps/go-library/utils/gotime"
	"github.com/dtapps/go-library/utils/gotrace_id"
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
	"log/slog"
	"os"
)

type SLogFun func() *SLog

type SLogConfig struct {
	LogPath      string // 日志文件路径
	LogName      string // 日志文件名
	MaxSize      int    // 单位为MB,默认为512MB
	MaxBackups   int    // 保留旧文件的最大个数
	MaxAge       int    // 文件最多保存多少天 0=不删除
	LocalTime    bool   // 采用本地时间
	Compress     bool   // 是否压缩日志
	JsonFormat   bool   // 是否输出为json格式
	ShowLine     bool   // 显示代码行
	LogSaveFile  bool   // 是否保存到文件
	LogInConsole bool   // 是否同时输出到控制台
}

type SLog struct {
	config      *SLogConfig
	jsonHandler *slog.JSONHandler
	textHandler *slog.TextHandler
	logger      *slog.Logger
}

func NewSlog(ctx context.Context, config *SLogConfig) *SLog {

	sl := &SLog{config: config}

	opts := slog.HandlerOptions{
		AddSource: sl.config.ShowLine,
		Level:     slog.LevelDebug,
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			if a.Key == slog.TimeKey {
				a.Value = slog.StringValue(a.Value.Time().Format(gotime.DateTimeFormat))
				//return slog.Attr{}
			}
			return a
		},
	}

	// 是否保存到文件
	if sl.config.LogSaveFile {

		lumberjackLogger := lumberjack.Logger{
			Filename:   sl.config.LogPath + sl.config.LogName, // ⽇志⽂件路径
			MaxSize:    sl.config.MaxSize,                     // 单位为MB,默认为512MB
			MaxAge:     sl.config.MaxAge,                      // 文件最多保存多少天
			MaxBackups: sl.config.MaxBackups,                  // 保留旧文件的最大个数
			LocalTime:  sl.config.LocalTime,                   // 采用本地时间
			Compress:   sl.config.Compress,                    // 是否压缩日志
		}

		// 是否json格式输出
		if sl.config.JsonFormat {
			if sl.config.LogInConsole {
				sl.jsonHandler = slog.NewJSONHandler(io.MultiWriter(os.Stdout, &lumberjackLogger), &opts)
			} else {
				sl.jsonHandler = slog.NewJSONHandler(&lumberjackLogger, &opts)
			}
			sl.logger = slog.New(sl.jsonHandler)
		} else {
			if sl.config.LogInConsole {
				sl.textHandler = slog.NewTextHandler(io.MultiWriter(os.Stdout, &lumberjackLogger), &opts)
			} else {
				sl.textHandler = slog.NewTextHandler(&lumberjackLogger, &opts)
			}
			sl.logger = slog.New(sl.textHandler)
		}
	} else {
		// 是否json格式输出
		if sl.config.JsonFormat {
			sl.jsonHandler = slog.NewJSONHandler(os.Stdout, &opts)
			sl.logger = slog.New(sl.jsonHandler)
		} else {
			sl.textHandler = slog.NewTextHandler(os.Stdout, &opts)
			sl.logger = slog.New(sl.textHandler)
		}
	}

	return sl
}

// WithLogger 跟踪编号
func (sl *SLog) WithLogger() *LoggerOperation {
	if sl.config.JsonFormat {
		logger := slog.New(sl.jsonHandler)
		return &LoggerOperation{
			logger: logger,
		}
	} else {
		logger := slog.New(sl.textHandler)
		return &LoggerOperation{
			logger: logger,
		}
	}
}

// WithTraceId 跟踪编号
func (sl *SLog) WithTraceId(ctx context.Context) *LoggerOperation {
	if sl.config.JsonFormat {
		jsonHandler := sl.jsonHandler.WithAttrs([]slog.Attr{
			slog.String("trace_id", gotrace_id.GetTraceIdContext(ctx)),
		})
		logger := slog.New(jsonHandler)
		return &LoggerOperation{
			logger: logger,
		}
	} else {
		textHandler := sl.textHandler.WithAttrs([]slog.Attr{
			slog.String("trace_id", gotrace_id.GetTraceIdContext(ctx)),
		})
		logger := slog.New(textHandler)
		return &LoggerOperation{
			logger: logger,
		}
	}
}

// WithTraceIdStr 跟踪编号
func (sl *SLog) WithTraceIdStr(traceId string) *LoggerOperation {
	if sl.config.JsonFormat {
		jsonHandler := sl.jsonHandler.WithAttrs([]slog.Attr{
			slog.String("trace_id", traceId),
		})
		logger := slog.New(jsonHandler)
		return &LoggerOperation{
			logger: logger,
		}
	} else {
		textHandler := sl.textHandler.WithAttrs([]slog.Attr{
			slog.String("trace_id", traceId),
		})
		logger := slog.New(textHandler)
		return &LoggerOperation{
			logger: logger,
		}
	}
}

type LoggerOperation struct {
	logger *slog.Logger
}

// Debug logs at LevelDebug.
func (l *LoggerOperation) Debug(msg string, args ...any) {
	l.logger.Debug(msg, args...)
}

// Debugf formats the message according to the format specifier
func (l *LoggerOperation) Debugf(template string, args ...any) {
	l.logger.Debug(fmt.Sprintf(template, args...))
}

// DebugContext logs at LevelDebug with the given context.
func (l *LoggerOperation) DebugContext(ctx context.Context, msg string, args ...any) {
	l.logger.DebugContext(ctx, msg, args...)
}

// Info logs at LevelInfo.
func (l *LoggerOperation) Info(msg string, args ...any) {
	l.logger.Info(msg, args...)
}

// Infof formats the message according to the format specifier
func (l *LoggerOperation) Infof(template string, args ...any) {
	l.logger.Info(fmt.Sprintf(template, args...))
}

// InfoContext logs at LevelInfo with the given context.
func (l *LoggerOperation) InfoContext(ctx context.Context, msg string, args ...any) {
	l.logger.InfoContext(ctx, msg, args...)
}

// Warn logs at LevelWarn.
func (l *LoggerOperation) Warn(msg string, args ...any) {
	l.logger.Warn(msg, args...)
}

// Warnf formats the message according to the format specifier
func (l *LoggerOperation) Warnf(template string, args ...any) {
	l.logger.Warn(fmt.Sprintf(template, args...))
}

func (l *LoggerOperation) WarnContext(ctx context.Context, msg string, args ...any) {
	l.logger.WarnContext(ctx, msg, args...)
}

// Error logs at LevelError.
func (l *LoggerOperation) Error(msg string, args ...any) {
	l.logger.Error(msg, args...)
}

// Errorf formats the message according to the format specifier
func (l *LoggerOperation) Errorf(template string, args ...any) {
	l.logger.Error(fmt.Sprintf(template, args...))
}

// ErrorContext logs at LevelError with the given context.
func (l *LoggerOperation) ErrorContext(ctx context.Context, msg string, args ...any) {
	l.logger.ErrorContext(ctx, msg, args...)
}
