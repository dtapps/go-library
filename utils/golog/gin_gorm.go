package golog

import (
	"bytes"
	"context"
	"github.com/gin-gonic/gin"
	"go.dtapp.net/library/utils/gojson"
	"go.dtapp.net/library/utils/gorequest"
	"go.dtapp.net/library/utils/gotime"
	"go.opentelemetry.io/otel/attribute"
	"io"
	"net/http"
	"runtime"
	"time"
)

// GinLogFunc Gin框架日志函数
type GinLogFunc func(ctx context.Context, response *GormGinLogModel)

// GinGorm 框架日志
type GinGorm struct {
	ginLogFunc GinLogFunc // Gin框架日志函数
}

// GinGormFun *GinGorm 框架日志驱动
type GinGormFun func() *GinGorm

// NewGinGorm 创建Gin框架实例
func NewGinGorm(ctx context.Context) (*GinGorm, error) {
	gg := &GinGorm{}

	return gg, nil
}

// 定义一个自定义的 ResponseWriter
type ginGormBodyWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

// 实现 http.ResponseWriter 的 Write 方法
func (w ginGormBodyWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

// WriteString 实现 http.ResponseWriter 的 WriteString 方法
func (w ginGormBodyWriter) WriteString(s string) (int, error) {
	w.body.WriteString(s)
	return w.ResponseWriter.WriteString(s)
}

// WriteHeader 实现 http.ResponseWriter 的 WriteHeader 方法
func (w ginGormBodyWriter) WriteHeader(statusCode int) {
	w.ResponseWriter.WriteHeader(statusCode)
}

// Header 实现 http.ResponseWriter 的 Header 方法
func (w ginGormBodyWriter) Header() http.Header {
	return w.ResponseWriter.Header()
}

// Middleware 中间件
func (gg *GinGorm) Middleware() gin.HandlerFunc {
	return func(g *gin.Context) {

		// OpenTelemetry链路追踪
		//g.Request = g.Request.WithContext(gg.TraceStartSpan(g))
		ctx, span := TraceStartSpan(g.Request.Context(), "gin")
		defer span.End()

		// 开始时间
		start := time.Now().UTC()

		// 模型
		var log = GormGinLogModel{}

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
		blw := &ginGormBodyWriter{
			ResponseWriter: g.Writer,
			body:           bytes.NewBufferString(""),
		}
		g.Writer = blw

		// 处理请求
		g.Next()

		// 结束时间
		end := time.Now().UTC()

		// 请求消耗时长
		log.RequestCostTime = end.Sub(start).Milliseconds()

		// 响应时间
		log.ResponseTime = gotime.Current().Time

		// 跟踪编号
		log.TraceID = gorequest.TraceSpanGetTraceID(span)

		// 请求编号
		log.RequestID = gorequest.GetRequestIDContext(ctx)

		// 请求主机
		log.RequestHost = g.Request.Host

		// 请求地址
		log.RequestPath = gorequest.NewUri(g.Request.RequestURI).UriFilterExcludeQueryString()

		// 请求参数
		log.RequestQuery = gojson.JsonEncodeNoError(gojson.ParseQueryString(g.Request.RequestURI))

		// 请求方式
		log.RequestMethod = g.Request.Method

		// 请求协议
		log.RequestScheme = g.Request.Proto

		// 请求类型
		log.RequestContentType = g.ContentType()

		// 请求IP
		log.RequestClientIP = g.ClientIP()

		// 请求UA
		log.RequestUserAgent = g.Request.UserAgent()

		// 请求头
		log.RequestHeader = gojson.JsonEncodeNoError(g.Request.Header)

		// 响应头
		log.ResponseHeader = gojson.JsonEncodeNoError(blw.Header())

		// 响应状态
		log.ResponseStatusCode = g.Writer.Status()

		// 响应内容
		if gojson.IsValidJSON(blw.body.String()) {
			log.ResponseBody = gojson.JsonEncodeNoError(gojson.JsonDecodeNoError(blw.body.String()))
		} else {
			log.ResponseBody = blw.body.String()
		}

		// OpenTelemetry链路追踪
		span.SetAttributes(attribute.String("request.time", log.RequestTime.Format(gotime.DateTimeFormat)))
		span.SetAttributes(attribute.String("request.host", log.RequestHost))
		span.SetAttributes(attribute.String("request.path", log.RequestPath))
		span.SetAttributes(attribute.String("request.query", log.RequestQuery))
		span.SetAttributes(attribute.String("request.method", log.RequestMethod))
		span.SetAttributes(attribute.String("request.scheme", log.RequestScheme))
		span.SetAttributes(attribute.String("request.content_type", log.RequestContentType))
		span.SetAttributes(attribute.String("request.body", log.RequestBody))
		span.SetAttributes(attribute.String("request.header", log.RequestHeader))
		span.SetAttributes(attribute.Int64("request.cost_time", log.RequestCostTime))

		span.SetAttributes(attribute.String("response.time", log.ResponseTime.Format(gotime.DateTimeFormat)))
		span.SetAttributes(attribute.String("response.header", log.ResponseHeader))
		span.SetAttributes(attribute.Int("response.status_code", log.ResponseStatusCode))
		span.SetAttributes(attribute.String("response.body", log.ResponseBody))

		// 调用Gin框架日志函数
		log.GoVersion = runtime.Version()
		log.SdkVersion = Version
		if gg.ginLogFunc != nil {
			gg.ginLogFunc(ctx, &log)
		}

	}
}

// SetLogFunc 设置日志记录方法
func (gg *GinGorm) SetLogFunc(ginLogFunc GinLogFunc) {
	gg.ginLogFunc = ginLogFunc
}
