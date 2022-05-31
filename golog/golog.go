package golog

import (
	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"time"
)

type ConfigGoLog struct {
	LogPath      string // 日志文件路径
	LogName      string // 日志文件名
	LogLevel     string // 日志级别 debug/info/warn/error，debug输出：debug/info/warn/error日志。 info输出：info/warn/error日志。 warn输出：warn/error日志。 error输出：error日志。
	MaxSize      int    // 单个文件大小,MB
	MaxBackups   int    // 保存的文件个数
	MaxAge       int    // 保存的天数 0=不删除
	Compress     bool   // 压缩
	JsonFormat   bool   // 是否输出为json格式
	ShowLine     bool   // 显示代码行
	LogInConsole bool   // 是否同时输出到控制台
}

type GoLog struct {
	ConfigGoLog
	Logger *zap.Logger
}

func NewGoLog(config *ConfigGoLog) *GoLog {
	g := &GoLog{}
	g.LogPath = config.LogPath
	g.LogName = config.LogName
	g.LogLevel = config.LogLevel
	g.MaxSize = config.MaxSize
	g.MaxBackups = config.MaxBackups
	g.MaxAge = config.MaxAge
	g.Compress = config.Compress
	g.JsonFormat = config.JsonFormat
	g.ShowLine = config.ShowLine
	g.LogInConsole = config.LogInConsole

	// 设置日志级别
	var level zapcore.Level
	switch g.LogLevel {
	case "debug":
		level = zap.DebugLevel
	case "info":
		level = zap.InfoLevel
	case "warn":
		level = zap.WarnLevel
	case "error":
		level = zap.ErrorLevel
	default:
		level = zap.InfoLevel
	}

	var (
		syncer zapcore.WriteSyncer

		// 自定义时间输出格式
		customTimeEncoder = func(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendString(t.Format("2006-01-02 15:04:05.000"))
		}

		// 自定义日志级别显示
		customLevelEncoder = func(level zapcore.Level, enc zapcore.PrimitiveArrayEncoder) {
			enc.AppendString(level.CapitalString())
		}
	)

	// 定义日志切割配置
	hook := lumberjack.Logger{
		Filename:   g.LogPath + g.LogName, // 日志文件的位置
		MaxSize:    g.MaxSize,             // 在进行切割之前，日志文件的最大大小（以MB为单位）
		MaxBackups: g.MaxBackups,          // 保留旧文件的最大个数
		Compress:   g.Compress,            // 是否压缩 disabled by default
	}
	if g.MaxAge > 0 {
		hook.MaxAge = g.MaxAge // days
	}

	// 判断是否控制台输出日志
	if g.LogInConsole {
		syncer = zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(&hook))
	} else {
		syncer = zapcore.AddSync(&hook)
	}

	// 定义zap配置信息
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level",
		NameKey:        "logger",
		CallerKey:      "line",
		MessageKey:     "msg",
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,
		EncodeTime:     customTimeEncoder,          // 自定义时间格式
		EncodeLevel:    customLevelEncoder,         // 小写编码器
		EncodeCaller:   zapcore.ShortCallerEncoder, // 全路径编码器
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeName:     zapcore.FullNameEncoder,
	}

	var encoder zapcore.Encoder
	// 判断是否json格式输出
	if g.JsonFormat {
		encoder = zapcore.NewJSONEncoder(encoderConfig)
	} else {
		encoder = zapcore.NewConsoleEncoder(encoderConfig)
	}

	core := zapcore.NewCore(
		encoder,
		syncer,
		level,
	)

	g.Logger = zap.New(core)

	// 判断是否显示代码行号
	if g.ShowLine {
		g.Logger = g.Logger.WithOptions(zap.AddCaller())
	}

	return g
}
