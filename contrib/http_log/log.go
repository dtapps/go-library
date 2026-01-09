package http_log

import (
	"bytes"
	"context"
	"encoding/json"
	"encoding/xml"
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
	GoVersion        string              // Go 版本
	PluginVersion    string              // 插件版本
	Hostname         string              // 主机名
	Method           string              // HTTP 方法
	URL              string              // 请求 URL
	RequestHeaders   map[string][]string // 请求头
	RequestBodyJSON  json.RawMessage     // JSON 格式的请求体
	RequestBodyXML   json.RawMessage     // XML 格式的请求体
	StatusCode       int                 // HTTP 状态码
	ResponseHeaders  map[string][]string // 响应头
	ResponseBodyJSON json.RawMessage     // JSON 格式的响应体
	ResponseBodyXML  json.RawMessage     // XML 格式的响应体
	ElapseTime       int64               // 处理时间（毫秒）
	IsError          bool                // 是否出错
}

// LogHandler 定义了处理日志数据的接口
type LogHandler interface {
	HandleLog(ctx context.Context, data *LogData) error
}

// LogCallback 定义了回调函数类型（如果你更喜欢函数式编程）
type LogCallback func(ctx context.Context, data *LogData) error

// 定义拦截器
type LoggingTransport struct {
	Proxied             http.RoundTripper // 被拦截的 Transport
	Handler             LogHandler        // 接口方式
	OnLog               LogCallback       // 回调方式
	DisableRequestBody  bool              // 是否禁用请求体记录
	DisableResponseBody bool              // 是否禁用响应体记录
}

// NewTransport 标准构造器
func NewTransport(base http.RoundTripper, handler LogHandler, callback LogCallback) http.RoundTripper {
	if base == nil {
		base = http.DefaultTransport
	}
	return &LoggingTransport{
		Proxied: base,
		Handler: handler,
		OnLog:   callback,
	}
}

// Middleware 返回一个 http.RoundTripper 方法
func (l *LoggingTransport) Middleware() func(http.RoundTripper) http.RoundTripper {
	return func(next http.RoundTripper) http.RoundTripper {
		if next == nil {
			next = l.Proxied
		}
		return &LoggingTransport{
			Proxied:             next,                  // 被拦截的 Transport
			Handler:             l.Handler,             // 接口方式
			OnLog:               l.OnLog,               // 回调方式
			DisableRequestBody:  l.DisableRequestBody,  // 保持配置
			DisableResponseBody: l.DisableResponseBody, // 保持配置
		}
	}
}

// MiddlewareNoBody 返回一个强制不记录 Body 的 http.RoundTripper 方法
func (l *LoggingTransport) MiddlewareNoBody() func(http.RoundTripper) http.RoundTripper {
	return func(next http.RoundTripper) http.RoundTripper {
		if next == nil {
			next = l.Proxied
		}
		return &LoggingTransport{
			Proxied:             next,      // 被拦截的 Transport
			Handler:             l.Handler, // 接口方式
			OnLog:               l.OnLog,   // 回调方式
			DisableRequestBody:  true,      // 强制禁用
			DisableResponseBody: true,      // 强制禁用
		}
	}
}

// Instance 返回一个 http.RoundTripper 实例
func (l *LoggingTransport) Instance(next http.RoundTripper) http.RoundTripper {
	if next == nil {
		next = l.Proxied
	}
	return &LoggingTransport{
		Proxied:             next,
		Handler:             l.Handler,
		OnLog:               l.OnLog,
		DisableRequestBody:  l.DisableRequestBody,
		DisableResponseBody: l.DisableResponseBody,
	}
}

// InstanceNoBody 返回一个强制不记录 Body 的 http.RoundTripper 实例
func (l *LoggingTransport) InstanceNoBody(next http.RoundTripper) http.RoundTripper {
	if next == nil {
		next = l.Proxied
	}
	return &LoggingTransport{
		Proxied:             next,
		Handler:             l.Handler,
		OnLog:               l.OnLog,
		DisableRequestBody:  true, // 强制禁用
		DisableResponseBody: true, // 强制禁用
	}
}

// RoundTrip 实现了 http.RoundTripper 接口
func (l *LoggingTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	startTime := time.Now()

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
	req = req.WithContext(httptrace.WithClientTrace(req.Context(), trace))

	// 获取主机名
	hostname := req.Host
	if hostname == "" && req.URL != nil {
		hostname = req.URL.Host
	}

	// 创建 LogData
	logData := &LogData{
		GoVersion:      runtime.Version(),
		PluginVersion:  Version,
		Hostname:       hostname,
		Method:         req.Method,
		URL:            req.URL.String(),
		RequestHeaders: finalReqHeaders,
	}

	// 捕获请求 Body
	if !l.DisableRequestBody && req.Body != nil {
		bodyBytes, _ := io.ReadAll(req.Body)
		l.processBody(bodyBytes, req.Header.Get("Content-Type"), &logData.RequestBodyJSON, &logData.RequestBodyXML)
		req.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
	}

	// 执行请求
	// 调用 Proxied.RoundTrip 会触发上面的 WroteHeaderField 回调
	resp, err := l.Proxied.RoundTrip(req)

	// 请求结束后，更新最终捕获到的真实请求头
	mu.Lock()
	logData.RequestHeaders = finalReqHeaders
	mu.Unlock()

	// 记录真实 Host (如果 Header 里没写，通常取 req.Host)
	if logData.RequestHeaders["Host"] == nil && hostname != "" {
		logData.RequestHeaders["Host"] = []string{hostname}
	}

	// 计算耗时
	logData.ElapseTime = time.Since(startTime).Milliseconds()

	// 异常处理
	if err != nil {
		logData.IsError = true
		l.emit(context.WithoutCancel(req.Context()), logData)
		return nil, err
	}

	// 捕获响应信息
	logData.StatusCode = resp.StatusCode
	logData.ResponseHeaders = resp.Header.Clone()
	logData.IsError = resp.StatusCode >= 400

	if !l.DisableResponseBody && resp.Body != nil {
		respBytes, _ := io.ReadAll(resp.Body)
		l.processBody(respBytes, resp.Header.Get("Content-Type"), &logData.ResponseBodyJSON, &logData.ResponseBodyXML)
		resp.Body = io.NopCloser(bytes.NewBuffer(respBytes))
	}

	// 触发保存
	l.emit(context.WithoutCancel(req.Context()), logData)

	return resp, nil
}

// 内部辅助方法：处理 body，根据类型存到 JSON 或 XML 字段
func (l *LoggingTransport) processBody(data []byte, contentType string, jsonField, xmlField *json.RawMessage) {
	if len(data) == 0 {
		return
	}

	bodyType := l.detectBodyType(contentType, data)
	switch bodyType {
	case "json":
		if l.isValidJSON(data) {
			*jsonField = data
		}
	case "xml":
		if l.isValidXML(data) {
			*xmlField = data
		}
	default:

	}
}

// 内部辅助方法：判断是否为 JSON 格式
func (l *LoggingTransport) isValidJSON(data []byte) bool {
	var js json.RawMessage
	return json.Unmarshal(data, &js) == nil
}

// 内部辅助方法：判断是否为 XML 格式
func (l *LoggingTransport) isValidXML(data []byte) bool {
	var v any
	return xml.Unmarshal(data, &v) == nil
}

// 内部辅助方法：根据 Content-Type 或内容判断 body 类型
func (l *LoggingTransport) detectBodyType(contentType string, data []byte) string {
	if strings.Contains(contentType, "application/json") || strings.HasPrefix(string(data), "{") || strings.HasPrefix(string(data), "[") {
		return "json"
	}
	if strings.Contains(contentType, "xml") || strings.Contains(contentType, "soap+xml") || strings.HasPrefix(string(data), "<") {
		return "xml"
	}
	if l.isValidJSON(data) {
		return "json"
	}
	if l.isValidXML(data) {
		return "xml"
	}
	return "text"
}

// 内部辅助方法：触发接口或回调
func (l *LoggingTransport) emit(ctx context.Context, data *LogData) {
	if l.OnLog != nil {
		go func(ctx context.Context, data *LogData) {
			if err := l.OnLog(ctx, data); err != nil {
				fmt.Println("save log failed (OnLog):", err)
			}
		}(ctx, data)
	} else if l.Handler != nil {
		go func(ctx context.Context, data *LogData) {
			if err := l.Handler.HandleLog(ctx, data); err != nil {
				fmt.Println("save log failed (HandleLog):", err)
			}
		}(ctx, data)
	}
}
