package resty_log

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/propagation"
	"go.opentelemetry.io/otel/trace"
	"resty.dev/v3"
)

const (
	Version = "1.0.3"
)

// LogData 表示每次请求/响应的日志数据
type LogData struct {
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

// LoggerSaveFunc 是保存日志数据的回调类型
type LoggerSaveFunc func(ctx context.Context, data *LogData) error

// LogSaver 定义保存日志的接口
type LogSaver interface {
	SaveLog(ctx context.Context, data *LogData) error
}

type Logger struct {
	saveFunc LoggerSaveFunc // 回调函数
	saver    LogSaver       // 接口实现
	debug    bool           // 调试模式
}

// NewLogger 创建一个支持回调和接口的 Logger
func NewLogger(saveFunc LoggerSaveFunc, saver LogSaver, debug bool) *Logger {
	return &Logger{
		saveFunc: saveFunc,
		saver:    saver,
		debug:    debug,
	}
}

// BeforeRequest Hook
func (l *Logger) BeforeRequest(c *resty.Client, req *resty.Request) error {
	// 记录开始时间并创建 span + 注入上下文
	start := time.Now()
	parentCtx := req.Context()
	tracer := otel.Tracer("go.dtapp.net/library/contrib/resty_log")
	spanCtx, _ := tracer.Start(parentCtx, "HTTP "+req.Method, trace.WithSpanKind(trace.SpanKindClient))

	// 注入 TraceContext 到请求头
	otel.GetTextMapPropagator().Inject(spanCtx, propagation.HeaderCarrier(req.Header))

	// 保存上下文与开始时间
	req.SetContext(context.WithValue(spanCtx, "start_time", start))
	return nil
}

// CopyResponseBodyMiddleware 将响应体拷贝到 Context
func (l *Logger) CopyResponseBodyMiddleware(c *resty.Client, resp *resty.Response) error {
	bodyBytes := resp.Bytes()                                                                 // 读取 body
	resp.Request.SetContext(context.WithValue(resp.Request.Context(), "raw_body", bodyBytes)) // 保存到 Context，方便外部获取
	return nil
}

// AfterResponse Hook
func (l *Logger) AfterResponse(c *resty.Client, resp *resty.Response) error {
	ctx := resp.Request.Context()

	// 耗时统计
	startTime, _ := ctx.Value("start_time").(time.Time)
	elapse := time.Since(startTime).Milliseconds()

	// 取回 span 并在这里设置所有可用属性，然后结束 span
	if span := trace.SpanFromContext(ctx); span != nil {
		// 请求属性
		span.SetAttributes(
			attribute.String("http.method", resp.Request.Method),
			attribute.String("http.url", resp.Request.URL),
			attribute.String("component", "resty"),
		)
		if raw := resp.Request.RawRequest; raw != nil && raw.URL != nil {
			span.SetAttributes(
				attribute.String("net.peer.name", raw.URL.Hostname()),
				attribute.String("net.peer.port", raw.URL.Port()),
				attribute.String("http.scheme", raw.URL.Scheme),
			)
			if ua := raw.Header.Get("User-Agent"); ua != "" {
				span.SetAttributes(attribute.String("http.user_agent", ua))
			}
		}

		// 响应属性
		span.SetAttributes(
			attribute.Int("http.status_code", resp.StatusCode()),
			attribute.Int("http.response_content_length", len(resp.Bytes())),
			attribute.Int64("http.elapsed_ms", elapse),
		)
		if resp.IsError() {
			span.SetAttributes(attribute.Bool("http.response.error", true))
		}

		span.End()
	}

	// 获取 Hostname
	var hostname string
	if rawReq := resp.Request.RawRequest; rawReq != nil {
		hostname = rawReq.URL.Hostname()
	}

	// 请求和响应详情
	reqHeaders := resp.Request.Header
	respHeaders := resp.Header()
	reqBodyBytes := reqBodyToBytes(resp.Request.Body)
	var respBodyBytes []byte
	if v := ctx.Value("raw_body"); v != nil {
		respBodyBytes = v.([]byte)
	}

	// 创建 LogData
	logData := &LogData{
		Hostname:        hostname,
		Method:          resp.Request.Method,
		URL:             resp.Request.URL,
		RequestHeaders:  reqHeaders,
		RequestBody:     safeJSONRawMessage(reqBodyBytes),
		StatusCode:      resp.StatusCode(),
		ResponseHeaders: respHeaders,
		ResponseBody:    safeJSONRawMessage(respBodyBytes),
		ElapseTime:      elapse,
		IsError:         resp.IsError(),
	}

	// 调试日志
	if l.debug {
		fmt.Println("========== REQUEST ==========")
		fmt.Printf("Method: %s\nURL: %s\nHostname: %s\n", logData.Method, logData.URL, logData.Hostname)
		fmt.Println("Headers:")
		for k, v := range logData.RequestHeaders {
			fmt.Printf("  %s: %v\n", k, v)
		}
		if len(logData.RequestBody) > 0 {
			fmt.Println("Body:")
			fmt.Println(prettyPrintJSON(string(logData.RequestBody)))
		}
		fmt.Println("========== RESPONSE =========")
		fmt.Printf("Status Code: %d\n", logData.StatusCode)
		fmt.Println("Headers:")
		for k, v := range logData.ResponseHeaders {
			fmt.Printf("  %s: %v\n", k, v)
		}
		if len(logData.ResponseBody) > 0 {
			fmt.Println("Body:")
			fmt.Println(prettyPrintJSON(string(logData.ResponseBody)))
		}
		fmt.Printf("Elapsed: %d ms\n", logData.ElapseTime)
		fmt.Println("==============================")
	}

	// 保存日志：优先使用回调函数，如果没有回调则使用接口
	if l.saveFunc != nil {
		go func() {
			if err := l.saveFunc(context.Background(), logData); err != nil {
				fmt.Println("save log failed (callback):", err)
			}
		}()
	} else if l.saver != nil {
		go func() {
			if err := l.saver.SaveLog(context.Background(), logData); err != nil {
				fmt.Println("save log failed (saver):", err)
			}
		}()
	}

	return nil
}

// reqBodyToBytes 将请求体转换为 []byte
func reqBodyToBytes(body any) []byte {
	switch v := body.(type) {
	case []byte:
		return v
	case string:
		return []byte(v)
	default:
		b, err := json.Marshal(v)
		if err != nil {
			return []byte(fmt.Sprintf("%v", v))
		}
		return b
	}
}

// prettyPrintJSON JSON 美化
func prettyPrintJSON(s string) string {
	var out bytes.Buffer
	if err := json.Indent(&out, []byte(s), "", "  "); err != nil {
		return s
	}
	return out.String()
}

// safeJSONRawMessage 将数据转换为 json.RawMessage
func safeJSONRawMessage(data []byte) json.RawMessage {
	if len(data) == 0 {
		return json.RawMessage("null")
	}
	var js json.RawMessage
	if json.Unmarshal(data, &js) == nil {
		return data
	}
	escaped, _ := json.Marshal(string(data))
	return escaped
}
