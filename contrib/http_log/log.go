package http_log

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"runtime"
	"time"
)

// LogData 表示每次请求/响应的日志数据
type LogData struct {
	GoVersion       string
	PluginVersion   string
	Hostname        string
	Method          string
	URL             string
	RequestHeaders  http.Header
	RequestBody     json.RawMessage
	StatusCode      int
	ResponseHeaders http.Header
	ResponseBody    json.RawMessage
	ElapseTime      int64
	IsError         bool
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
		logData.RequestBody = json.RawMessage(bodyBytes)
		req.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))
	}

	// 执行请求
	resp, err := l.Proxied.RoundTrip(req)

	// 计算耗时
	logData.ElapseTime = time.Since(startTime).Milliseconds()

	// 异常处理：即使 err != nil，我们也应该触发日志记录
	if err != nil {
		logData.IsError = true
		// l.emit(req.Context(), logData)
		l.emit(context.Background(), logData)
		return nil, err
	}

	// 捕获响应信息
	logData.StatusCode = resp.StatusCode
	logData.ResponseHeaders = resp.Header
	logData.IsError = resp.StatusCode >= 400

	if resp != nil && resp.Body != nil {
		respBytes, _ := io.ReadAll(resp.Body)
		logData.ResponseBody = json.RawMessage(respBytes)
		resp.Body = io.NopCloser(bytes.NewBuffer(respBytes))
	}

	// 触发保存
	l.emit(context.Background(), logData)
	// l.emit(req.Context(), logData)

	return resp, nil
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
