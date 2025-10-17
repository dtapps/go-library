package gin_ent_log

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"log/slog"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"go.dtapp.net/library/contrib/gin_requestid"
	"go.dtapp.net/library/utils/gorequest"
	"go.dtapp.net/library/utils/gotime"
)

// GinLogFunc Gin框架日志函数
type GinLogFunc func(ctx context.Context, response *GinLogData)

// GinLog 框架日志
type GinLog struct {
	ginLogFunc GinLogFunc // Gin框架日志函数
	routeDebug bool       // 路由是否开启调试
}

// GinLogFun *GinLog 框架日志驱动
type GinLogFun func() *GinLog

// NewGinLog 创建Gin框架实例
func NewGinLog(ctx context.Context) (*GinLog, error) {
	gg := &GinLog{}

	return gg, nil
}

// SetLogFunc 设置 hertzLogFunc
func (gg *GinLog) SetLogFunc(ginLogFunc GinLogFunc) {
	gg.ginLogFunc = ginLogFunc
}

// SetRouteDebug 设置 routeDebug
func (gg *GinLog) SetRouteDebug(routeDebug bool) {
	gg.routeDebug = routeDebug
}

// GinLogBodyWriter 定义一个自定义的 ResponseWriter
type GinLogBodyWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

// 实现 http.ResponseWriter 的 Write 方法
func (w GinLogBodyWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

// WriteString 实现 http.ResponseWriter 的 WriteString 方法
func (w GinLogBodyWriter) WriteString(s string) (int, error) {
	w.body.WriteString(s)
	return w.ResponseWriter.WriteString(s)
}

// WriteHeader 实现 http.ResponseWriter 的 WriteHeader 方法
func (w GinLogBodyWriter) WriteHeader(statusCode int) {
	w.ResponseWriter.WriteHeader(statusCode)
}

// Header 实现 http.ResponseWriter 的 Header 方法
func (w GinLogBodyWriter) Header() http.Header {
	return w.ResponseWriter.Header()
}

// Middleware 中间件
func (gg *GinLog) Middleware() gin.HandlerFunc {
	return func(g *gin.Context) {

		// 开始时间
		start := time.Now().UTC()

		// 模型
		var log = GinLogData{}

		// 请求时间
		log.RequestTime = gotime.Current().Time

		// Read the Body content
		var bodyBytes []byte
		if g.Request.Body != nil {
			bodyBytes, _ = io.ReadAll(g.Request.Body)
		}

		// 将io.ReadCloser恢复到其原始状态
		g.Request.Body = io.NopCloser(bytes.NewBuffer(bodyBytes))

		// 创建自定义的 ResponseWriter 并替换原有的
		blw := &GinLogBodyWriter{
			ResponseWriter: g.Writer,
			body:           bytes.NewBufferString(""),
		}
		g.Writer = blw

		// 可插拔 Tracer
		spanCtx := g.Request.Context()
		if tracer != nil {
			reqInfo := RequestInfo{
				Method: g.Request.Method,
				Path:   gorequest.NewUri(g.Request.RequestURI).UriFilterExcludeQueryString(),
				Host:   g.Request.Host,
				Header: g.Request.Header,
				Start:  start,
			}
			spanCtx = tracer.Start(spanCtx, reqInfo)
			// 将新的上下文注入到 http.Request
			g.Request = g.Request.WithContext(spanCtx)
		}

		// 处理请求
		g.Next()

		// 结束时间
		end := time.Now().UTC()

		// 请求消耗时长
		log.RequestCostTime = end.Sub(start).Milliseconds()

		// 响应时间
		log.ResponseTime = gotime.Current().Time

		// 输出路由日志
		if gg.routeDebug {
			status := g.Writer.Status()
			logFn := slog.InfoContext
			if status >= 500 {
				logFn = slog.ErrorContext
			} else if status >= 400 {
				logFn = slog.WarnContext
			}
			logFn(spanCtx, "hertz route",
				slog.Int("status", status),
				slog.Int64("cost_ms", log.RequestCostTime),
				slog.String("method", g.Request.Method),
				slog.String("full_path", gorequest.NewUri(g.Request.RequestURI).UriFilterExcludeQueryString()),
				slog.String("client_ip", g.ClientIP()),
				slog.String("host", g.Request.Host),
			)
		}

		// 请求编号
		log.RequestID = gin_requestid.Get(g)
		if log.RequestID == "" {
			log.RequestID = gin_requestid.GetX(g)
			if log.RequestID == "" {
				log.RequestID = gorequest.GetRequestIDContext(spanCtx)
			}
		}

		// 请求主机
		log.RequestHost = g.Request.Host

		// 请求地址
		log.RequestPath = gorequest.NewUri(g.Request.RequestURI).UriFilterExcludeQueryString()

		// 请求参数
		log.RequestQuery = gorequest.ParseQueryString(g.Request.RequestURI)

		// 请求方式
		log.RequestMethod = g.Request.Method

		// 请求IP
		log.RequestIP = g.ClientIP()

		// 请求头
		log.RequestHeader = g.Request.Header

		// 响应头
		log.ResponseHeader = blw.Header()

		// 响应状态
		log.ResponseCode = g.Writer.Status()

		// 响应内容
		if gorequest.IsValidJSON(blw.body.String()) {
			_ = json.Unmarshal(blw.body.Bytes(), &log.ResponseBody)
			//log.ResponseBody = gojson.JsonDecodeNoError(blw.body.String())
		} else {
			//log.ResponseBody = blw.body.String()
		}

		// 可插拔 Tracer
		if tracer != nil {
			respInfo := ResponseInfo{
				Status:     g.Writer.Status(),
				Header:     blw.Header(),
				Body:       blw.body.Bytes(),
				End:        end,
				DurationMs: log.RequestCostTime,
			}
			tracer.End(spanCtx, respInfo)
		}

		// 调用Gin框架日志函数
		if gg.ginLogFunc != nil {
			gg.ginLogFunc(spanCtx, &log)
		}

	}
}
