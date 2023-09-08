package golog

import (
	"context"
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
func (sl *SLog) WithLogger() *slog.Logger {
	if sl.config.JsonFormat {
		logger := slog.New(sl.jsonHandler)
		return logger
	} else {
		logger := slog.New(sl.textHandler)
		return logger
	}
}

// WithTraceId 跟踪编号
func (sl *SLog) WithTraceId(ctx context.Context) *slog.Logger {
	if sl.config.JsonFormat {
		jsonHandler := sl.jsonHandler.WithAttrs([]slog.Attr{
			slog.String("trace_id", gotrace_id.GetTraceIdContext(ctx)),
		})
		logger := slog.New(jsonHandler)
		return logger
	} else {
		textHandler := sl.textHandler.WithAttrs([]slog.Attr{
			slog.String("trace_id", gotrace_id.GetTraceIdContext(ctx)),
		})
		logger := slog.New(textHandler)
		return logger
	}
}

// WithTraceIdStr 跟踪编号
func (sl *SLog) WithTraceIdStr(traceId string) *slog.Logger {
	if sl.config.JsonFormat {
		jsonHandler := sl.jsonHandler.WithAttrs([]slog.Attr{
			slog.String("trace_id", traceId),
		})
		logger := slog.New(jsonHandler)
		return logger
	} else {
		textHandler := sl.textHandler.WithAttrs([]slog.Attr{
			slog.String("trace_id", traceId),
		})
		logger := slog.New(textHandler)
		return logger
	}
}
