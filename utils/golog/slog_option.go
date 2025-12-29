package golog

import (
	"log/slog"

	"gopkg.in/natefinch/lumberjack.v2"
)

// SLogOption 定义了配置 SLog 的函数类型
type SLogOption func(*SLog)

// WithSLogLumberjack Lumberjack配置
// Filename 日志文件的位置
// MaxSize 文件最大尺寸（以MB为单位）
// MaxAge 留旧文件的最大天数
// MaxBackups 保留的最大旧文件数量
// Compress 是否压缩/归档旧文件
// LocalTime 使用本地时间创建时间戳
func WithSLogLumberjack(config *lumberjack.Logger) SLogOption {
	return func(sl *SLog) {
		sl.option.lumberjackConfig = config
		sl.option.lumberjackConfigStatus = true
	}
}

// WithSLogSetLevel 设置日志级别
func WithSLogSetLevel(level slog.Level) SLogOption {
	return func(sl *SLog) {
		sl.option.setLevel = level
	}
}

// WithSLogShowLine 显示代码行
func WithSLogShowLine(status bool) SLogOption {
	return func(sl *SLog) {
		sl.option.showLine = status
	}
}

// WithSLogSetDefault 设置为默认的实例
func WithSLogSetDefault(status bool) SLogOption {
	return func(sl *SLog) {
		sl.option.setDefault = status
	}
}

// WithSLogSetDefaultCtx 设置默认上下文
func WithSLogSetDefaultCtx(status bool) SLogOption {
	return func(sl *SLog) {
		sl.option.setDefaultCtx = status
	}
}

// WithSLogSetJSONFormat 设置JSON格式
func WithSLogSetJSONFormat(status bool) SLogOption {
	return func(sl *SLog) {
		sl.option.setJSONFormat = status
	}
}

// SetDisableLogging 设置是否完全禁用日志输出。
// 当设置为 true 时，所有日志将被 io.Discard 丢弃，适用于生产环境以提升性能。
func SetDisableLogging(disable bool) SLogOption {
	return func(sl *SLog) {
		sl.option.disableLogging = disable
	}
}

// WithSLogEnableOTel 启用/关闭 OpenTelemetry slog 桥接
func WithSLogEnableOTel(status bool) SLogOption {
	return func(sl *SLog) {
		sl.option.enableOTel = status
	}
}

// WithSLogOTelName 配置 otelslog 的 logger 名称（留空则默认 "slog"）
func WithSLogOTelName(name string) SLogOption {
	return func(sl *SLog) {
		sl.option.oTelLoggerName = name
	}
}
