package http_log

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/http/httptrace"
	"runtime"
	"strings"
	"sync"
	"time"
)

// LogData 表示每次请求/响应的日志数据
type LogData struct {
	GoVersion         string              // Go 版本
	PluginVersion     string              // 插件版本
	Hostname          string              // 主机名
	Method            string              // 请求方法
	URL               string              // 请求 URL
	RequestHeaders    map[string][]string // 请求头
	RequestBody       json.RawMessage     // 请求体
	StatusCode        int                 // 状态码
	ResponseHeaders   map[string][]string // 响应头
	ResponseBody      json.RawMessage     // 响应体
	ElapseTime        int64               // 耗时（毫秒）
	ProcessElapseTime int64               // 处理耗时（毫秒）
	IsError           bool                // 是否出错

	elapseTimeStart        time.Time // 开始时间
	elapseTimeEnd          time.Time // 结束时间
	processElapseTimeStart time.Time // 处理开始时间
	processElapseTimeEnd   time.Time // 处理结束时间
}

// LogHandler 定义处理日志数据的接口
type LogHandler interface {
	HandleLog(ctx context.Context, data *LogData) error
}

// LogCallback 定义回调函数类型
type LogCallback func(ctx context.Context, data *LogData) error

// LoggingRoundTripper 定义拦截器
type LoggingRoundTripper struct {
	Transport           http.RoundTripper // 被拦截的 Transport
	Handler             LogHandler        // 接口方式
	OnLog               LogCallback       // 回调方式
	debug               bool              // 是否开启调试模式
	disableRequestBody  bool              // 是否禁用请求体记录
	disableResponseBody bool              // 是否禁用响应体记录
}

// NewLoggingRoundTripper 标准构造器
func NewLoggingRoundTripper(base http.RoundTripper, handler LogHandler, callback LogCallback) *LoggingRoundTripper {
	if base == nil {
		base = http.DefaultTransport
	}
	return &LoggingRoundTripper{
		Transport: base,
		Handler:   handler,
		OnLog:     callback,
	}
}

// baseTransport 获取基础 Transport
func (t *LoggingRoundTripper) baseTransport() http.RoundTripper {
	if t.Transport != nil {
		return t.Transport
	}
	return http.DefaultTransport
}

// EnableDebug 开启调试模式
func (t *LoggingRoundTripper) EnableDebug() {
	t.debug = true
}

// Clone 复制一个新的 LoggingRoundTripper
func (t *LoggingRoundTripper) Clone() *LoggingRoundTripper {
	return &LoggingRoundTripper{
		Transport:           t.baseTransport(),     // 被拦截的 Transport
		Handler:             t.Handler,             // 接口方式
		OnLog:               t.OnLog,               // 回调方式
		debug:               t.debug,               // 是否开启调试模式
		disableRequestBody:  t.disableRequestBody,  // 是否禁用请求体记录
		disableResponseBody: t.disableResponseBody, // 是否禁用响应体记录
	}
}

// CloneNoBody 复制一个新的 LoggingRoundTripper，禁用请求体和响应体记录
func (t *LoggingRoundTripper) CloneNoBody() *LoggingRoundTripper {
	return &LoggingRoundTripper{
		Transport:           t.baseTransport(), // 被拦截的 Transport
		Handler:             t.Handler,         // 接口方式
		OnLog:               t.OnLog,           // 回调方式
		debug:               t.debug,           // 是否开启调试模式
		disableRequestBody:  true,              // 强制禁用
		disableResponseBody: true,              // 强制禁用
	}
}

// Middleware 返回一个 http.RoundTripper 方法
func (t *LoggingRoundTripper) Middleware() func(http.RoundTripper) http.RoundTripper {
	return func(next http.RoundTripper) http.RoundTripper {
		if next == nil {
			next = t.baseTransport()
		}
		return &LoggingRoundTripper{
			Transport:           next,                  // 被拦截的 Transport
			Handler:             t.Handler,             // 接口方式
			OnLog:               t.OnLog,               // 回调方式
			debug:               t.debug,               // 是否开启调试模式
			disableRequestBody:  t.disableRequestBody,  // 是否禁用请求体记录
			disableResponseBody: t.disableResponseBody, // 是否禁用响应体记录
		}
	}
}

// MiddlewareNoBody 返回一个强制不记录 Body 的 http.RoundTripper 方法
func (t *LoggingRoundTripper) MiddlewareNoBody() func(http.RoundTripper) http.RoundTripper {
	return func(next http.RoundTripper) http.RoundTripper {
		if next == nil {
			next = t.baseTransport()
		}
		return &LoggingRoundTripper{
			Transport:           next,      // 被拦截的 Transport
			Handler:             t.Handler, // 接口方式
			OnLog:               t.OnLog,   // 回调方式
			debug:               t.debug,   // 是否开启调试模式
			disableRequestBody:  true,      // 强制禁用
			disableResponseBody: true,      // 强制禁用
		}
	}
}

// Instance 返回一个 http.RoundTripper 实例
func (t *LoggingRoundTripper) Instance(next http.RoundTripper) http.RoundTripper {
	if next == nil {
		next = t.baseTransport()
	}
	return &LoggingRoundTripper{
		Transport:           next,                  // 被拦截的 Transport
		Handler:             t.Handler,             // 接口方式
		OnLog:               t.OnLog,               // 回调方式
		debug:               t.debug,               // 是否开启调试模式
		disableRequestBody:  t.disableRequestBody,  // 是否禁用请求体记录
		disableResponseBody: t.disableResponseBody, // 是否禁用响应体记录
	}
}

// InstanceNoBody 返回一个强制不记录 Body 的 http.RoundTripper 实例
func (t *LoggingRoundTripper) InstanceNoBody(next http.RoundTripper) http.RoundTripper {
	if next == nil {
		next = t.baseTransport()
	}
	return &LoggingRoundTripper{
		Transport:           next,      // 被拦截的 Transport
		Handler:             t.Handler, // 接口方式
		OnLog:               t.OnLog,   // 回调方式
		debug:               t.debug,   // 是否开启调试模式
		disableRequestBody:  true,      // 强制禁用
		disableResponseBody: true,      // 强制禁用
	}
}

// RoundTrip 实现了 http.RoundTripper 接口
func (t *LoggingRoundTripper) RoundTrip(req *http.Request) (*http.Response, error) {

	// 开启调试模式时
	if t.debug {
		fmt.Printf("[LoggingRoundTripper] RoundTrip Start: %s %s\n", req.Method, req.URL.String())
		defer fmt.Printf("[LoggingRoundTripper] RoundTrip End: %s %s\n", req.Method, req.URL.String())
	}

	// 创建 LogData
	logData := &LogData{
		GoVersion:     runtime.Version(),
		PluginVersion: Version,
		Hostname:      req.Host,
		Method:        req.Method,
		URL:           req.URL.String(),
	}

	// 获取上下文
	ctx := req.Context()

	// 开始时间
	logData.elapseTimeStart = time.Now().UTC()
	logData.processElapseTimeStart = time.Now().UTC()

	// 异步捕获网线上真正跑的 Header
	var mu sync.Mutex

	// 初始值为应用层设置的头
	finalReqHeaders := req.Header.Clone()

	// 注入 httptrace 钩子
	// WroteHeaderField 会在每一行 Header 真正写入二进制流时触发
	trace := &httptrace.ClientTrace{
		WroteHeaderField: func(key string, value []string) {
			mu.Lock()
			defer mu.Unlock()

			// 过滤掉 HTTP/2 的伪头部 (如 :method, :path)
			// 这些是协议层逻辑，记在 Header Map 里显得不专业
			if strings.HasPrefix(key, ":") {
				return
			}

			// 规范化 Key 的名称 (例如把 "user-agent" 转回 "User-Agent")
			canonicalKey := http.CanonicalHeaderKey(key)

			if finalReqHeaders == nil {
				finalReqHeaders = make(http.Header)
			}
			finalReqHeaders[canonicalKey] = value
		},
	}
	req = req.WithContext(httptrace.WithClientTrace(ctx, trace))

	// 主机名
	if logData.Hostname == "" && req.URL != nil {
		logData.Hostname = req.URL.Host
	}

	// 请求头
	logData.RequestHeaders = finalReqHeaders

	// 请求体
	if !t.disableRequestBody && req.Body != nil {
		bodyBytes, _ := io.ReadAll(req.Body)
		logData.RequestBody = t.processResponseBody(req.Header, bodyBytes)
		req.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
	}

	// 执行请求
	// 调用 Transport.RoundTrip 会触发上面的 WroteHeaderField 回调
	resp, err := t.baseTransport().RoundTrip(req)

	// 请求结束后，更新最终捕获到的真实请求头
	mu.Lock()
	logData.RequestHeaders = finalReqHeaders
	mu.Unlock()

	// 结束时间
	logData.elapseTimeEnd = time.Now().UTC()

	// 计算耗时
	logData.ElapseTime = time.Since(logData.elapseTimeEnd).Milliseconds()

	// 记录真实 Host
	// if logData.RequestHeaders["Host"] == nil && logData.Hostname != "" {
	// 	logData.RequestHeaders["Host"] = []string{
	// 		logData.Hostname,
	// 	}
	// }

	// 异常处理
	if err != nil {
		logData.IsError = true
		// 触发保存
		t.emit(context.WithoutCancel(ctx), logData)
		return nil, err
	}

	// 状态码
	logData.StatusCode = resp.StatusCode

	// 响应头
	logData.ResponseHeaders = resp.Header.Clone()

	// 是否错误
	logData.IsError = resp.StatusCode >= 400

	// 响应体
	if !t.disableResponseBody && resp.Body != nil {
		respBytes, _ := io.ReadAll(resp.Body)
		logData.ResponseBody = t.processResponseBody(resp.Header, respBytes)
		resp.Body = io.NopCloser(bytes.NewBuffer(respBytes))
	}

	// 触发保存
	t.emit(context.WithoutCancel(ctx), logData)

	return resp, nil
}

// emit 触发接口或回调
func (t *LoggingRoundTripper) emit(ctx context.Context, logData *LogData) {

	// 开启调试模式时
	if t.debug {
		fmt.Printf("[LoggingRoundTripper] emit Start: %s %s\n", logData.Method, logData.URL)
		defer fmt.Printf("[LoggingRoundTripper] emit End: %s %s\n", logData.Method, logData.URL)
	}

	// 处理结束时间
	logData.processElapseTimeEnd = time.Now().UTC()

	// 计算处理耗时
	logData.ProcessElapseTime = time.Since(logData.processElapseTimeStart).Milliseconds()
	if t.OnLog != nil {
		go func(ctx context.Context, data *LogData) {
			if err := t.OnLog(ctx, data); err != nil {
				fmt.Println("save log failed (OnLog):", err)
			}
		}(context.WithoutCancel(ctx), logData)
	} else if t.Handler != nil {
		go func(ctx context.Context, data *LogData) {
			if err := t.Handler.HandleLog(ctx, data); err != nil {
				fmt.Println("save log failed (HandleLog):", err)
			}
		}(context.WithoutCancel(ctx), logData)
	}
}
