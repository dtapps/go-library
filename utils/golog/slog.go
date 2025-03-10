package golog

import (
	"bufio"
	"go.dtapp.net/library/utils/gotime"
	"gopkg.in/natefinch/lumberjack.v2"
	"io"
	"log/slog"
	"os"
	"time"
)

type SLogFun func() *SLog

type sLogConfig struct {
	consoleOutput          bool               // 控制台输出开关
	showLine               bool               // 显示代码行
	setDefault             bool               // 设置为默认的实例
	setDefaultCtx          bool               // 设置默认上下文
	setJSONFormat          bool               // 设置为json格式
	lumberjackConfig       *lumberjack.Logger // 配置lumberjack
	lumberjackConfigStatus bool               // 配置lumberjack状态
}

type SLog struct {
	option      sLogConfig
	logger      *slog.Logger
	ctxHandler  *ContextHandler
	jsonHandler *slog.JSONHandler
	textHandler *slog.TextHandler
}

// NewSlog 创建
func NewSlog(opts ...SLogOption) *SLog {
	sl := &SLog{}
	for _, opt := range opts {
		opt(sl)
	}
	sl.start()
	return sl
}

func (sl *SLog) start() {
	opts := slog.HandlerOptions{
		AddSource: sl.option.showLine, // 输出日志语句的位置信息
		Level:     slog.LevelDebug,    // 设置最低日志等级
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			if a.Key == slog.TimeKey { // 格式化 key 为 "time" 的属性值
				a.Value = slog.StringValue(a.Value.Time().Format(gotime.DateTimeFormat))
				//return slog.Attr{}
			}
			return a
		},
	}

	// 输出
	var mw io.Writer
	if sl.option.lumberjackConfigStatus {
		if sl.option.consoleOutput {
			// 同时控制台和文件输出日志
			mw = io.MultiWriter(os.Stdout, sl.option.lumberjackConfig)
		} else {
			// 只在文件输出日志
			mw = sl.option.lumberjackConfig
		}
	} else {
		mw = os.Stdout
	}

	// 控制台输出
	if sl.option.consoleOutput {
		// 使用缓冲
		bufferedWriter := bufio.NewWriter(mw)
		if sl.option.setJSONFormat {
			sl.jsonHandler = slog.NewJSONHandler(bufferedWriter, &opts)
		} else {
			sl.textHandler = slog.NewTextHandler(bufferedWriter, &opts)
		}

		// 定时刷新缓冲区
		if bufferedWriter, ok := mw.(*bufio.Writer); ok {
			go func() {
				for range time.Tick(5 * time.Second) { // 5 秒刷新一次
					bufferedWriter.Flush()
				}
			}()
		}
	}
	if sl.option.setJSONFormat {
		// 设置默认上下文
		if sl.option.setDefaultCtx {
			sl.ctxHandler = &ContextHandler{sl.jsonHandler}
			sl.logger = slog.New(sl.ctxHandler)
		} else {
			sl.logger = slog.New(sl.jsonHandler)
		}
	} else {
		// 设置默认上下文
		if sl.option.setDefaultCtx {
			sl.ctxHandler = &ContextHandler{sl.textHandler}
			sl.logger = slog.New(sl.ctxHandler)
		} else {
			sl.logger = slog.New(sl.textHandler)
		}
	}

	// 将这个 slog 对象设置为默认的实例
	if sl.option.setDefault {
		slog.SetDefault(sl.logger)
	}
}

// WithLogger 跟踪编号
func (sl *SLog) WithLogger() (logger *slog.Logger) {
	if sl.option.setDefaultCtx {
		logger = slog.New(sl.ctxHandler)
	} else {
		if sl.option.setJSONFormat {
			logger = slog.New(sl.jsonHandler)
		} else {
			logger = slog.New(sl.textHandler)
		}
	}
	return logger
}

// WriteLog 异步写入日志
//func (sl *SLog) WriteLog(ctx context.Context, r slog.Record) {
//	go func() {
//		_ = sl.logger.Handler().Handle(ctx, r) // 异步写入
//	}()
//}
