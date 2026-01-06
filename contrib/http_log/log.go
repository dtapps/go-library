package http_log

import (
	"bytes"
	"context"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"runtime"
	"strings"
	"time"
)

// LogData 表示每次请求/响应的日志数据
type LogData struct {
	GoVersion        string
	PluginVersion    string
	Hostname         string
	Method           string
	URL              string
	RequestHeaders   http.Header
	RequestBodyJSON  []byte // JSON 格式的请求体
	RequestBodyXML   []byte // XML 格式的请求体
	StatusCode       int
	ResponseHeaders  http.Header
	ResponseBodyJSON []byte // JSON 格式的响应体
	ResponseBodyXML  []byte // XML 格式的响应体
	ElapseTime       int64
	IsError          bool
}

// LogHandler 定义了处理日志数据的接口
type LogHandler interface {
	HandleLog(ctx context.Context, data *LogData) error
}

// LogCallback 定义了回调函数类型（如果你更喜欢函数式编程）
type LogCallback func(ctx context.Context, data *LogData) error

// 定义拦截器
type LoggingTransport struct {
	Proxied http.RoundTripper // 被拦截的 Transport
	Handler LogHandler        // 接口方式
	OnLog   LogCallback       // 回调方式
}

// RoundTrip 实现了 http.RoundTripper 接口
func (l *LoggingTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	startTime := time.Now()

	hostname := req.Host
	if hostname == "" {
		hostname = req.URL.Host
	}

	// 创建 LogData
	logData := &LogData{
		GoVersion:      runtime.Version(),
		PluginVersion:  Version,
		Hostname:       hostname,
		Method:         req.Method,
		URL:            req.URL.String(),
		RequestHeaders: req.Header,
	}

	// 捕获请求 Body
	if req.Body != nil {
		bodyBytes, _ := io.ReadAll(req.Body)
		l.processBody(bodyBytes, req.Header.Get("Content-Type"), &logData.RequestBodyJSON, &logData.RequestBodyXML)
		req.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
	}

	// 执行请求
	resp, err := l.Proxied.RoundTrip(req)

	// 计算耗时
	logData.ElapseTime = time.Since(startTime).Milliseconds()

	// 异常处理：即使 err != nil，我们也应该触发日志记录
	if err != nil {
		logData.IsError = true
		l.emit(context.WithoutCancel(req.Context()), logData)
		return nil, err
	}

	// 捕获响应信息
	logData.StatusCode = resp.StatusCode
	logData.ResponseHeaders = resp.Header
	logData.IsError = resp.StatusCode >= 400

	if resp != nil && resp.Body != nil {
		respBytes, _ := io.ReadAll(resp.Body)
		l.processBody(respBytes, resp.Header.Get("Content-Type"), &logData.ResponseBodyJSON, &logData.ResponseBodyXML)
		resp.Body = io.NopCloser(bytes.NewBuffer(respBytes))
	}

	// 触发保存
	l.emit(context.WithoutCancel(req.Context()), logData)

	return resp, nil
}

// 内部辅助方法：处理 body，根据类型存到 JSON 或 XML 字段
func (l *LoggingTransport) processBody(data []byte, contentType string, jsonField, xmlField *[]byte) {
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
