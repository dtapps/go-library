package resty_extend

import (
	"fmt"
	"log"

	"gopkg.in/natefinch/lumberjack.v2"
)

// Logger 包装了标准的 log.Logger 并实现了 resty.Logger 接口
type Logger struct {
	*log.Logger
}

// Errorf 实现了 resty.Logger 接口的 Errorf 方法
func (l *Logger) Errorf(format string, v ...any) {
	l.Printf("[ERROR] "+format, v...)
}

// Warnf 实现了 resty.Logger 接口的 Warnf 方法
func (l *Logger) Warnf(format string, v ...any) {
	l.Printf("[WARN] "+format, v...)
}

// Debugf 实现了 resty.Logger 接口的 Debugf 方法
func (l *Logger) Debugf(format string, v ...any) {
	l.Printf("[DEBUG] "+format, v...)
}

func NewLog(filename string, serviceName string) *Logger {

	// 配置 lumberjack 日志轮转
	logFile := &lumberjack.Logger{
		Filename:   filename, // 日志文件路径
		MaxSize:    3,        // 每个日志文件最大尺寸（MB）
		MaxAge:     7,        // 留旧文件的最大天数
		MaxBackups: 30,       // 保留的最大旧文件数量
		Compress:   true,     // 是否压缩/归档旧文件
	}

	prefix := "resty: "
	if serviceName != "" {
		prefix = fmt.Sprintf("resty(%s): ", serviceName)
	}

	// 创建一个新的标准库 logger 实例，并将 lumberjack.Logger 作为 Writer 传入
	stdLogger := log.New(logFile, prefix, log.LstdFlags)

	// 返回包装了标准库 logger 的 Logger
	return &Logger{stdLogger}
}
