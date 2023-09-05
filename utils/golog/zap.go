package golog

import (
	"context"
	"github.com/dtapps/go-library/utils/gotime"
	"github.com/dtapps/go-library/utils/gotrace_id"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
	"time"
)

// ZapLogFun *ApiClient 驱动
type ZapLogFun func() *ZapLog

type ZapLogConfig struct {
	LogPath      string // 日志文件路径
	LogName      string // 日志文件名
	MaxSize      int    // 单位为MB,默认为512MB
	MaxBackups   int    // 保留旧文件的最大个数
	MaxAge       int    // 文件最多保存多少天 0=不删除
	LocalTime    bool   // 采用本地时间
	Compress     bool   // 是否压缩日志
	JsonFormat   bool   // 是否输出为json格式
	ShowLine     bool   // 显示代码行
	LogInConsole bool   // 是否同时输出到控制台
}

type ZapLog struct {
	config  *ZapLogConfig
	logger  *zap.Logger
	zapCore zapcore.Core
}

func NewZapLog(config *ZapLogConfig) *ZapLog {

	zl := &ZapLog{config: config}

	var syncer zapcore.WriteSyncer

	if zl.config.LogPath != "" && zl.config.LogName != "" {
		// 定义日志切割配置
		hook := lumberjack.Logger{
			Filename:   zl.config.LogPath + zl.config.LogName, // ⽇志⽂件路径
			MaxSize:    zl.config.MaxSize,                     // 单位为MB,默认为512MB
			MaxAge:     zl.config.MaxAge,                      // 文件最多保存多少天
			MaxBackups: zl.config.MaxBackups,                  // 保留旧文件的最大个数
			LocalTime:  zl.config.LocalTime,                   // 采用本地时间
			Compress:   zl.config.Compress,                    // 是否压缩日志
		}
		if zl.config.LogInConsole {
			// 在控制台和文件输出日志
			syncer = zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(&hook))
		} else {
			// 在文件输出日志
			syncer = zapcore.AddSync(&hook)
		}
	} else {
		// 在控制台输出日志
		syncer = zapcore.NewMultiWriteSyncer()
	}

	// 自定义时间输出格式
	customTimeEncoder := func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(t.Format(gotime.DateTimeFormat))
	}

	// 自定义日志级别显示
	customLevelEncoder := func(level zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
		enc.AppendString(level.CapitalString())
	}

	// 定义zap配置信息
	encoderConf := zapcore.EncoderConfig{
		CallerKey:      "caller_line", // 打印文件名和行数
		LevelKey:       "level_name",
		MessageKey:     "msg",
		TimeKey:        "time",
		NameKey:        "logger",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeTime:     customTimeEncoder,          // 自定义时间格式
		EncodeLevel:    customLevelEncoder,         // 小写编码器
		EncodeCaller:   zapcore.ShortCallerEncoder, // 全路径编码器
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeName:     zapcore.FullNameEncoder,
	}

	// 判断是否json格式输出
	if zl.config.JsonFormat {
		zl.zapCore = zapcore.NewCore(zapcore.NewJSONEncoder(encoderConf),
			syncer, zap.NewAtomicLevelAt(zapcore.InfoLevel))
	} else {
		zl.zapCore = zapcore.NewCore(zapcore.NewConsoleEncoder(encoderConf),
			syncer, zap.NewAtomicLevelAt(zapcore.InfoLevel))
	}

	zl.logger = zl.withShowLine(zap.New(zl.zapCore))

	return zl
}

// 判断是否显示代码行号
func (zl *ZapLog) withShowLine(logger *zap.Logger) *zap.Logger {
	if zl.config.ShowLine {
		logger = logger.WithOptions(zap.AddCaller())
	}
	return logger
}

// WithLogger 跟踪编号
func (zl *ZapLog) WithLogger() *zap.Logger {
	return zl.logger
}

// WithTraceId 跟踪编号
func (zl *ZapLog) WithTraceId(ctx context.Context) *zap.Logger {
	logger := zl.logger
	logger = logger.With(zapcore.Field{
		Key:    "trace_id",
		Type:   zapcore.StringType,
		String: gotrace_id.GetTraceIdContext(ctx),
	})
	return logger
}

// WithTraceIdStr 跟踪编号
func (zl *ZapLog) WithTraceIdStr(traceId string) *zap.Logger {
	logger := zl.logger
	logger = logger.With(zapcore.Field{
		Key:    "trace_id",
		Type:   zapcore.StringType,
		String: traceId,
	})
	return logger
}
