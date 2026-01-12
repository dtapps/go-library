package req_log

import (
	"context"
	"encoding/json"
	"fmt"
	"runtime"
	"time"

	"github.com/imroc/req/v3"
)

// LogData 表示每次请求/响应的日志数据
type LogData struct {
	GoVersion          string              // Go 版本
	PluginVersion      string              // 插件版本
	Hostname           string              // 主机名
	Method             string              // 请求方法
	URL                string              // 请求 URL
	RequestHeaders     map[string][]string // 请求头
	RequestBody        json.RawMessage     // 请求体
	RequestQueryParams map[string][]string // 请求参数
	RequestFormData    map[string][]string // 请求表单数据
	RequestPathParams  map[string]string   // 请求路径参数
	StatusCode         int                 // 状态码
	ResponseHeaders    map[string][]string // 响应头
	ResponseBody       json.RawMessage     // 响应体
	ElapseTime         int64               // 耗时（毫秒）
	ProcessElapseTime  int64               // 处理耗时（毫秒）
	IsError            bool                // 是否错误

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

// LoggerMiddleware 定义拦截器
type LoggerMiddleware struct {
	Handler             LogHandler  // 接口实现
	OnLog               LogCallback // 回调方式
	debug               bool        // 是否开启调试模式
	disableRequestBody  bool        // 是否禁用请求体记录
	disableResponseBody bool        // 是否禁用响应体记录
}

// NewLoggerMiddleware 标准构造器
func NewLoggerMiddleware(handler LogHandler, callback LogCallback) *LoggerMiddleware {
	return &LoggerMiddleware{
		Handler: handler,
		OnLog:   callback,
	}
}

// EnableDebug 开启调试模式
func (l *LoggerMiddleware) EnableDebug() {
	l.debug = true
}

// OnBeforeRequest 注入开始时间
func (l *LoggerMiddleware) OnBeforeRequest(c *req.Client, req *req.Request) error {

	// 开启调试模式时
	if l.debug {
		fmt.Printf("[LoggerMiddleware] OnBeforeRequest Start: %s %s\n", req.Method, req.RawURL)
		defer fmt.Printf("[LoggerMiddleware] OnBeforeRequest End: %s %s\n", req.Method, req.RawURL)
	}

	// 获取上下文
	ctx := req.Context()

	// 记录开始时间
	startTime := time.Now().UTC()

	// 保存上下文与开始时间（供 OnAfterResponse 计算耗时）
	req.SetContext(WithStartTimeKey(ctx, startTime))

	return nil
}

// OnAfterResponse 打印/保存
func (l *LoggerMiddleware) OnAfterResponse(c *req.Client, resp *req.Response) error {

	// 开启调试模式时
	if l.debug {
		fmt.Printf("[LoggerMiddleware] OnAfterResponse Start: %d %s %s\n", resp.StatusCode, resp.Request.Method, resp.Request.RawURL)
		defer fmt.Printf("[LoggerMiddleware] OnAfterResponse End: %d %s %s\n", resp.StatusCode, resp.Request.Method, resp.Request.RawURL)
	}

	// 创建 LogData
	logData := &LogData{
		GoVersion:       runtime.Version(),
		PluginVersion:   Version,
		Method:          resp.Request.Method,
		URL:             resp.Request.RawURL,
		RequestHeaders:  resp.Request.Headers.Clone(),
		StatusCode:      resp.StatusCode,
		ResponseHeaders: resp.Header.Clone(),
		IsError:         resp.IsErrorState(),
	}

	// 获取上下文
	ctx := resp.Request.Context()

	// 开始时间
	logData.elapseTimeStart = GetStartTimeKey(ctx)
	logData.processElapseTimeStart = time.Now()

	// 结束时间
	logData.elapseTimeEnd = time.Now().UTC()

	// 计算耗时
	logData.ElapseTime = time.Since(logData.elapseTimeEnd).Milliseconds()

	// 主机名
	if rawReq := resp.Request.RawRequest; rawReq != nil {
		logData.Hostname = rawReq.URL.Hostname()
	}

	// 记录真实 Host
	if logData.RequestHeaders["Host"] == nil && logData.Hostname != "" {
		logData.RequestHeaders["Host"] = []string{
			logData.Hostname,
		}
	}

	// 请求体
	if !l.disableRequestBody && resp.Request.Body != nil {
		contentType := resp.Request.Headers.Get("Content-Type")
		logData.RequestBody = l.processBodyAny(contentType, resp.Request.Body)
	}

	// 请求参数
	if !l.disableRequestBody && resp.Request.QueryParams != nil {
		logData.RequestQueryParams = resp.Request.QueryParams
	}

	// 请求表单数据
	if !l.disableRequestBody && resp.Request.FormData != nil {
		logData.RequestFormData = resp.Request.FormData
	}

	// 请求路径参数
	if !l.disableRequestBody {
		logData.RequestPathParams = resp.Request.PathParams
	}

	// 响应体
	if !l.disableResponseBody && ctx != nil {
		contentType := resp.Header.Get("Content-Type")
		logData.ResponseBody = l.processBodyByte(contentType, resp.Bytes())
	}

	if l.debug {
		fmt.Printf("[LoggerMiddleware] OnAfterResponse TraceInfo:\n")
		fmt.Printf("%+v\n", resp.Request.TraceInfo())
	}

	// 触发保存
	l.emit(context.WithoutCancel(ctx), logData)

	return nil
}

// emit 触发接口或回调
func (l *LoggerMiddleware) emit(ctx context.Context, logData *LogData) {

	// 开启调试模式时
	if l.debug {
		fmt.Printf("[LoggerMiddleware] emit Start: %s %s\n", logData.Method, logData.URL)
		defer fmt.Printf("[LoggerMiddleware] emit End: %s %s\n", logData.Method, logData.URL)
	}

	// 处理结束时间
	logData.processElapseTimeEnd = time.Now().UTC()

	// 计算处理耗时
	logData.ProcessElapseTime = time.Since(logData.processElapseTimeStart).Milliseconds()
	if l.OnLog != nil {
		go func(ctx context.Context, data *LogData) {
			if err := l.OnLog(ctx, data); err != nil {
				fmt.Println("save log failed (OnLog):", err)
			}
		}(context.WithoutCancel(ctx), logData)
	} else if l.Handler != nil {
		go func(ctx context.Context, data *LogData) {
			if err := l.Handler.HandleLog(ctx, data); err != nil {
				fmt.Println("save log failed (HandleLog):", err)
			}
		}(context.WithoutCancel(ctx), logData)
	}
}
