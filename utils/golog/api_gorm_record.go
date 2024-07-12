package golog

import (
	"context"
	"go.dtapp.net/library/utils/gojson"
	"go.dtapp.net/library/utils/gorequest"
	"go.dtapp.net/library/utils/gotime"
	"go.dtapp.net/library/utils/gourl"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
	"unicode/utf8"
)

// 记录日志
func (ag *ApiGorm) gormRecord(ctx context.Context, span trace.Span, data GormApiLogModel) {
	if ag.gormConfig.stats == false {
		return
	}

	if utf8.ValidString(data.ResponseBody) == false {
		data.ResponseBody = ""
	}

	// 跟踪编号
	data.TraceID = gorequest.TraceSpanGetTraceID(span)

	// 请求编号
	data.RequestID = gorequest.GetRequestIDContext(ctx)

	// OpenTelemetry链路追踪
	span.SetAttributes(attribute.String("request.time", data.RequestTime.Format(gotime.DateTimeFormat)))
	span.SetAttributes(attribute.String("request.uri", data.RequestUri))
	span.SetAttributes(attribute.String("request.url", data.RequestUrl))
	span.SetAttributes(attribute.String("request.api", data.RequestApi))
	span.SetAttributes(attribute.String("request.method", data.RequestMethod))
	span.SetAttributes(attribute.String("request.params", data.RequestParams))
	span.SetAttributes(attribute.String("request.header", data.RequestHeader))
	span.SetAttributes(attribute.String("request.ip", data.RequestIP))
	span.SetAttributes(attribute.Int64("request.cost_time", data.RequestCostTime))
	span.SetAttributes(attribute.String("response.header", data.ResponseHeader))
	span.SetAttributes(attribute.Int("response.status_code", data.ResponseStatusCode))
	span.SetAttributes(attribute.String("response.body", data.ResponseBody))
	span.SetAttributes(attribute.String("response.time", data.ResponseTime.Format(gotime.DateTimeFormat)))

	err := ag.gormClient.WithContext(ctx).
		Table(ag.gormConfig.tableName).
		Create(&data).Error
	if err != nil {
		span.RecordError(err, trace.WithStackTrace(true))
		span.SetStatus(codes.Error, err.Error())
	}
}

// 中间件
func (ag *ApiGorm) gormMiddleware(ctx context.Context, span trace.Span, request gorequest.Response) {
	data := GormApiLogModel{
		RequestTime:        request.RequestTime,                              // 请求时间
		RequestUri:         request.RequestUri,                               // 请求链接
		RequestUrl:         gourl.UriParse(request.RequestUri).Url,           // 请求链接
		RequestApi:         gourl.UriParse(request.RequestUri).Path,          // 请求接口
		RequestMethod:      request.RequestMethod,                            // 请求方式
		RequestParams:      gojson.JsonEncodeNoError(request.RequestParams),  // 请求参数
		RequestHeader:      gojson.JsonEncodeNoError(request.RequestHeader),  // 请求头部
		RequestCostTime:    request.RequestCostTime,                          // 请求消耗时长
		ResponseHeader:     gojson.JsonEncodeNoError(request.ResponseHeader), // 响应头部
		ResponseStatusCode: request.ResponseStatusCode,                       // 响应状态码
		ResponseTime:       request.ResponseTime,                             // 响应时间
	}
	if !request.HeaderIsImg() && !request.HeaderHtml() {
		if len(request.ResponseBody) > 0 {
			data.ResponseBody = gojson.JsonEncodeNoError(gojson.JsonDecodeNoError(string(request.ResponseBody))) // 响应数据
		}
	}

	ag.gormRecord(ctx, span, data)
}

// 中间件
func (ag *ApiGorm) gormMiddlewareXml(ctx context.Context, span trace.Span, request gorequest.Response) {
	data := GormApiLogModel{
		RequestTime:        request.RequestTime,                              // 请求时间
		RequestUri:         request.RequestUri,                               // 请求链接
		RequestUrl:         gourl.UriParse(request.RequestUri).Url,           // 请求链接
		RequestApi:         gourl.UriParse(request.RequestUri).Path,          // 请求接口
		RequestMethod:      request.RequestMethod,                            // 请求方式
		RequestParams:      gojson.JsonEncodeNoError(request.RequestParams),  // 请求参数
		RequestHeader:      gojson.JsonEncodeNoError(request.RequestHeader),  // 请求头部
		RequestCostTime:    request.RequestCostTime,                          // 请求消耗时长
		ResponseHeader:     gojson.JsonEncodeNoError(request.ResponseHeader), // 响应头部
		ResponseStatusCode: request.ResponseStatusCode,                       // 响应状态码
		ResponseTime:       request.ResponseTime,                             // 响应时间
	}
	if !request.HeaderIsImg() && !request.HeaderHtml() {
		if len(request.ResponseBody) > 0 {
			data.ResponseBody = gojson.XmlEncodeNoError(gojson.XmlDecodeNoError(request.ResponseBody)) // 响应内容
		}
	}

	ag.gormRecord(ctx, span, data)
}

// 中间件
func (ag *ApiGorm) gormMiddlewareCustom(ctx context.Context, span trace.Span, api string, request gorequest.Response) {
	data := GormApiLogModel{
		RequestTime:        request.RequestTime,                              // 请求时间
		RequestUri:         request.RequestUri,                               // 请求链接
		RequestUrl:         gourl.UriParse(request.RequestUri).Url,           // 请求链接
		RequestApi:         api,                                              // 请求接口
		RequestMethod:      request.RequestMethod,                            // 请求方式
		RequestParams:      gojson.JsonEncodeNoError(request.RequestParams),  // 请求参数
		RequestHeader:      gojson.JsonEncodeNoError(request.RequestHeader),  // 请求头部
		RequestCostTime:    request.RequestCostTime,                          // 请求消耗时长
		ResponseHeader:     gojson.JsonEncodeNoError(request.ResponseHeader), // 响应头部
		ResponseStatusCode: request.ResponseStatusCode,                       // 响应状态码
		ResponseTime:       request.ResponseTime,                             // 响应时间
	}
	if !request.HeaderIsImg() && !request.HeaderHtml() {
		if len(request.ResponseBody) > 0 {
			data.ResponseBody = gojson.JsonEncodeNoError(gojson.JsonDecodeNoError(string(request.ResponseBody))) // 响应数据
		}
	}

	ag.gormRecord(ctx, span, data)
}
