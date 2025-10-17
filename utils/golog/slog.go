package golog

import (
	"context"
	"io"
	"log/slog"
	"os"
	"time"

	"go.opentelemetry.io/contrib/bridges/otelslog"
	"gopkg.in/natefinch/lumberjack.v2"
)

const (
	Version = "1.0.177"
)

type SLogFun func() *SLog

type sLogConfig struct {
	showLine               bool               // 显示代码行
	setDefault             bool               // 设置为默认的实例
	setDefaultCtx          bool               // 设置默认上下文
	setJSONFormat          bool               // 设置为json格式
	lumberjackConfig       *lumberjack.Logger // 配置lumberjack
	lumberjackConfigStatus bool               // 配置lumberjack状态
	disableLogging         bool               // 完全禁用日志输出（静默模式，使用 io.Discard）
	enableOTel             bool               // 启用 OpenTelemetry slog 桥接
	oTelLoggerName         string             // otelslog 的 logger 名称
}

type SLog struct {
	option       sLogConfig
	logger       *slog.Logger
	ctxHandler   *ContextHandler
	jsonHandler  *slog.JSONHandler
	textHandler  *slog.TextHandler
	finalHandler slog.Handler // 最终用于构建 logger 的 handler（可能是 Tee 后的）
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

	// 配置 slog 的 Handler 选项
	opts := slog.HandlerOptions{
		AddSource: sl.option.showLine, // 输出日志语句的位置信息
		Level:     slog.LevelDebug,    // 设置最低日志等级
		ReplaceAttr: func(groups []string, a slog.Attr) slog.Attr {
			if a.Key == slog.TimeKey { // 格式化 key 为 "time" 的属性值
				a.Value = slog.StringValue(a.Value.Time().Format(time.DateTime))
				//return slog.Attr{}
			}
			return a
		},
	}

	// 核心：决定日志输出目的地
	var mw io.Writer

	// 使用 switch 语句清晰地处理三种情况
	switch {
	case sl.option.disableLogging:
		// 场景一：完全静默
		// 在生产环境或性能敏感场景下，使用 io.Discard 优雅丢弃所有日志。
		// 优势：零内存开销，避免无用 I/O，防止敏感信息泄露。
		mw = io.Discard

	case sl.option.lumberjackConfigStatus:
		// 场景二：同时输出
		// 开发或调试环境，同时输出到控制台和文件，便于实时查看。
		mw = io.MultiWriter(os.Stdout, sl.option.lumberjackConfig)

	default:
		// 场景三：仅文件输出（修正了原逻辑错误）
		// 原代码错误地将“仅文件输出”写成了 os.Stdout。
		// 现在修正为：如果 lumberjack 已配置，则输出到文件；否则，作为兜底，输出到控制台。
		if sl.option.lumberjackConfig != nil {
			mw = sl.option.lumberjackConfig
		} else {
			mw = os.Stdout // 兜底方案，避免 nil Writer 导致 panic
		}
	}

	// 根据用户选择的格式（JSON/Text）创建基础 Handler
	var baseHandler slog.Handler
	if sl.option.setJSONFormat {
		sl.jsonHandler = slog.NewJSONHandler(mw, &opts)
		baseHandler = sl.jsonHandler
	} else {
		sl.textHandler = slog.NewTextHandler(mw, &opts)
		baseHandler = sl.textHandler
	}

	// 如果启用 OTel，则合并一个 OTel Handler，实现同时输出
	if sl.option.enableOTel {
		name := sl.option.oTelLoggerName
		if name == "" {
			name = "slog"
		}
		sl.finalHandler = multiHandler{hs: []slog.Handler{baseHandler, otelslog.NewHandler(name)}}
	} else {
		sl.finalHandler = baseHandler
	}

	// 是否需要默认上下文包装
	if sl.option.setDefaultCtx {
		sl.ctxHandler = &ContextHandler{sl.finalHandler}
		sl.logger = slog.New(sl.ctxHandler)
	} else {
		sl.logger = slog.New(sl.finalHandler)
	}

	// 如果用户要求，将此 logger 设置为全局默认 logger
	if sl.option.setDefault {
		slog.SetDefault(sl.logger)
	}

}

// WithLogger 跟踪编号
func (sl *SLog) WithLogger() (logger *slog.Logger) {
	if sl.option.setDefaultCtx {
		return slog.New(sl.ctxHandler)
	}
	return slog.New(sl.finalHandler)
}
