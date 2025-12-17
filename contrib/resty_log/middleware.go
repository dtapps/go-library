package resty_log

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"runtime"
	"time"

	"resty.dev/v3"
)

// LogData 表示每次请求/响应的日志数据
type LogData struct {
	GoVersion       string
	RestyVersion    string
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

// IntrusionRequest Hook 注入开始时间
// 可以放 resty.PrepareRequestMiddleware 前面
func (l *Logger) IntrusionRequest(c *resty.Client, req *resty.Request) error {

	// 记录开始时间
	start := time.Now().UTC()

	// 获取上下文
	ctx := req.Context()

	// 保存上下文与开始时间（供 AfterResponse 计算耗时）
	// 这里的 start_time 是最准确的计时起点
	req.SetContext(context.WithValue(ctx, "start_time", start))

	return nil
}

// BeforeRequest Hook 记录开始时间和OTel
// 必须放 resty.PrepareRequestMiddleware 后面，否则无法获取到请求体
func (l *Logger) BeforeRequest(c *resty.Client, req *resty.Request) error {

	// 获取上下文
	spanCtx := req.Context()

	// 获取准确的计时起点（由 IntrusionRequest 存入）
	start, ok := spanCtx.Value("start_time").(time.Time)
	if !ok {
		start = time.Now().UTC()
	}

	// 可插拔 tracer
	if tracer != nil && req.RawRequest != nil {
		// 组装 RequestInfo
		reqInfo := RequestInfo{
			Version: resty.Version,

			Header: req.Header, // resty 的 Header 就是 http.Header，便于 tracer 注入
			Start:  start,

			Method:    req.Method,
			URL:       req.URL,
			UserAgent: req.Header.Get("User-Agent"),

			Host: req.RawRequest.URL.Hostname(),
		}
		spanCtx = tracer.Start(spanCtx, reqInfo)
	}

	// 保存上下文与开始时间（供 AfterResponse 计算耗时）
	req.SetContext(context.WithValue(spanCtx, "start_time", start))

	return nil
}

// CopyResponseBodyMiddleware 将响应体拷贝到 Context
// 必须放 resty.AutoParseResponseMiddleware 前面，否则无法获取到响应体
func (l *Logger) CopyResponseBodyMiddleware(c *resty.Client, resp *resty.Response) error {

	// 获取上下文
	ctx := resp.Request.Context()

	// 读取 body
	bodyBytes := resp.Bytes()

	// 保存到 Context，方便外部获取
	resp.Request.SetContext(context.WithValue(ctx, "raw_body", bodyBytes))

	return nil
}

// AfterResponse Hook 打印/保存
// 必须放 resty.AutoParseResponseMiddleware 后面
func (l *Logger) AfterResponse(c *resty.Client, resp *resty.Response) error {

	// 获取上下文
	ctx := resp.Request.Context()

	// 开始时间
	startTime, _ := ctx.Value("start_time").(time.Time)

	// 结束时间
	endTime := time.Now().UTC()

	// 耗时统计
	elapse := endTime.Sub(startTime).Milliseconds()

	// 可插拔 tracer
	if tracer != nil {
		// 组装 ResponseInfo
		var respBodyBytes []byte
		if v := ctx.Value("raw_body"); v != nil {
			respBodyBytes = v.([]byte)
		}

		respInfo := ResponseInfo{
			Status:     resp.StatusCode(),
			Header:     resp.Header(),
			Body:       respBodyBytes,
			End:        endTime, // 记录结束时间
			DurationMs: elapse,  // 使用计算出的总耗时
		}

		// 调用 End 方法，传递包含 Span 的 context
		tracer.End(ctx, respInfo)
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
		GoVersion:       runtime.Version(),
		RestyVersion:    resty.Version,
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
