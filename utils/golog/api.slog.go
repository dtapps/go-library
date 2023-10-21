package golog

import (
	"context"
	"github.com/dtapps/go-library"
	"github.com/dtapps/go-library/utils/dorm"
	"github.com/dtapps/go-library/utils/goip"
	"github.com/dtapps/go-library/utils/gorequest"
	"github.com/dtapps/go-library/utils/gotrace_id"
	"github.com/dtapps/go-library/utils/gourl"
	"runtime"
	"time"
)

type ApiSLog struct {
	systemConfig struct {
		systemHostname  string // 主机名
		systemOs        string // 系统类型
		systemKernel    string // 系统内核
		systemInsideIp  string // 内网ip
		systemOutsideIp string // 外网ip
		goVersion       string // go版本
		sdkVersion      string // sdk版本
	}
	slog struct {
		status bool  // 状态
		client *SLog // 日志服务
	}
}

func NewApiSlog(ctx context.Context) *ApiSLog {

	sl := &ApiSLog{}

	sl.setConfig(ctx)

	return sl
}

func (sl *ApiSLog) setConfig(ctx context.Context) {

	info := getSystem()

	sl.systemConfig.systemHostname = info.SystemHostname
	sl.systemConfig.systemOs = info.SystemOs
	sl.systemConfig.systemKernel = info.SystemKernel

	sl.systemConfig.systemInsideIp = goip.GetInsideIp(ctx)

	sl.systemConfig.sdkVersion = go_library.Version()
	sl.systemConfig.goVersion = runtime.Version()

}

// ConfigSLogClientFun 日志配置
func (sl *ApiSLog) ConfigSLogClientFun(sLogFun SLogFun) {
	sLog := sLogFun()
	if sLog != nil {
		sl.slog.client = sLog
		sl.slog.status = true
	}
}

// 结构体
type apiSLog struct {
	TraceId               string                 `json:"trace_id,omitempty"`
	RequestTime           time.Time              `json:"request_time,omitempty"`
	RequestUri            string                 `json:"request_uri,omitempty"`
	RequestUrl            string                 `json:"request_url,omitempty"`
	RequestApi            string                 `json:"request_api,omitempty"`
	RequestMethod         string                 `json:"request_method,omitempty"`
	RequestParams         map[string]interface{} `json:"request_params,omitempty"`
	RequestHeader         map[string]string      `json:"request_header,omitempty"`
	RequestIp             string                 `json:"request_ip,omitempty"`
	ResponseHeader        map[string][]string    `json:"response_header,omitempty"`
	ResponseStatusCode    int                    `json:"response_status_code,omitempty"`
	ResponseBody          map[string]interface{} `json:"response_body,omitempty"`
	ResponseContentLength int64                  `json:"response_content_length,omitempty"`
	ResponseTime          time.Time              `json:"response_time,omitempty,omitempty"`
	SystemHostName        string                 `json:"system_host_name,omitempty"`
	SystemInsideIp        string                 `json:"system_inside_ip,omitempty"`
	SystemOs              string                 `json:"system_os,omitempty"`
	SystemArch            string                 `json:"system_arch,omitempty"`
	GoVersion             string                 `json:"go_version,omitempty"`
	SdkVersion            string                 `json:"sdk_version,omitempty"`
}

// Middleware 中间件
func (sl *ApiSLog) Middleware(ctx context.Context, request gorequest.Response) {

	data := apiSLog{
		TraceId:               gotrace_id.GetTraceIdContext(ctx),
		RequestTime:           request.RequestTime,
		RequestUri:            request.RequestUri,
		RequestUrl:            gourl.UriParse(request.RequestUri).Url,
		RequestApi:            gourl.UriParse(request.RequestUri).Path,
		RequestMethod:         request.RequestMethod,
		RequestParams:         request.RequestParams,
		RequestHeader:         request.RequestHeader,
		RequestIp:             sl.systemConfig.systemOutsideIp,
		ResponseHeader:        request.ResponseHeader,
		ResponseStatusCode:    request.ResponseStatusCode,
		ResponseBody:          dorm.JsonDecodeNoError(request.ResponseBody),
		ResponseContentLength: request.ResponseContentLength,
		ResponseTime:          request.ResponseTime,
		SystemHostName:        sl.systemConfig.systemHostname,
		SystemInsideIp:        sl.systemConfig.systemInsideIp,
		SystemOs:              sl.systemConfig.systemOs,
		SystemArch:            sl.systemConfig.systemKernel,
		GoVersion:             sl.systemConfig.goVersion,
		SdkVersion:            sl.systemConfig.sdkVersion,
	}

	if sl.slog.status {
		sl.slog.client.WithTraceId(ctx).Info("Middleware",
			"request_time", data.RequestTime,
			"request_uri", data.RequestUri,
			"request_url", data.RequestUrl,
			"request_api", data.RequestApi,
			"request_method", data.RequestMethod,
			"request_params", data.RequestParams,
			"request_header", data.RequestHeader,
			"request_ip", data.RequestIp,
			"response_header", data.ResponseHeader,
			"response_status_code", data.ResponseStatusCode,
			"response_body", data.ResponseBody,
			"response_content_length", data.ResponseContentLength,
			"response_time", data.ResponseTime,
			"system_host_name", data.SystemHostName,
			"system_inside_ip", data.SystemInsideIp,
			"system_os", data.SystemOs,
			"system_arch", data.SystemArch,
			"go_version", data.GoVersion,
			"sdk_version", data.SdkVersion,
		)
	}

}

// MiddlewareXml 中间件
func (sl *ApiSLog) MiddlewareXml(ctx context.Context, request gorequest.Response) {

	data := apiSLog{
		TraceId:               gotrace_id.GetTraceIdContext(ctx),
		RequestTime:           request.RequestTime,
		RequestUri:            request.RequestUri,
		RequestUrl:            gourl.UriParse(request.RequestUri).Url,
		RequestApi:            gourl.UriParse(request.RequestUri).Path,
		RequestMethod:         request.RequestMethod,
		RequestParams:         request.RequestParams,
		RequestHeader:         request.RequestHeader,
		RequestIp:             sl.systemConfig.systemOutsideIp,
		ResponseHeader:        request.ResponseHeader,
		ResponseStatusCode:    request.ResponseStatusCode,
		ResponseBody:          dorm.XmlDecodeNoError(request.ResponseBody),
		ResponseContentLength: request.ResponseContentLength,
		ResponseTime:          request.ResponseTime,
		SystemHostName:        sl.systemConfig.systemHostname,
		SystemInsideIp:        sl.systemConfig.systemInsideIp,
		SystemOs:              sl.systemConfig.systemOs,
		SystemArch:            sl.systemConfig.systemKernel,
		GoVersion:             sl.systemConfig.goVersion,
		SdkVersion:            sl.systemConfig.sdkVersion,
	}

	if sl.slog.status {
		sl.slog.client.WithTraceId(ctx).Info("MiddlewareXml",
			"request_time", data.RequestTime,
			"request_uri", data.RequestUri,
			"request_url", data.RequestUrl,
			"request_api", data.RequestApi,
			"request_method", data.RequestMethod,
			"request_params", data.RequestParams,
			"request_header", data.RequestHeader,
			"request_ip", data.RequestIp,
			"response_header", data.ResponseHeader,
			"response_status_code", data.ResponseStatusCode,
			"response_body", data.ResponseBody,
			"response_content_length", data.ResponseContentLength,
			"response_time", data.ResponseTime,
			"system_host_name", data.SystemHostName,
			"system_inside_ip", data.SystemInsideIp,
			"system_os", data.SystemOs,
			"system_arch", data.SystemArch,
			"go_version", data.GoVersion,
			"sdk_version", data.SdkVersion,
		)
	}

}

// MiddlewareCustom 中间件
func (sl *ApiSLog) MiddlewareCustom(ctx context.Context, api string, request gorequest.Response) {

	data := apiSLog{
		TraceId:               gotrace_id.GetTraceIdContext(ctx),
		RequestTime:           request.RequestTime,
		RequestUri:            request.RequestUri,
		RequestUrl:            gourl.UriParse(request.RequestUri).Url,
		RequestApi:            api,
		RequestMethod:         request.RequestMethod,
		RequestParams:         request.RequestParams,
		RequestHeader:         request.RequestHeader,
		RequestIp:             sl.systemConfig.systemOutsideIp,
		ResponseHeader:        request.ResponseHeader,
		ResponseStatusCode:    request.ResponseStatusCode,
		ResponseBody:          dorm.JsonDecodeNoError(request.ResponseBody),
		ResponseContentLength: request.ResponseContentLength,
		ResponseTime:          request.ResponseTime,
		SystemHostName:        sl.systemConfig.systemHostname,
		SystemInsideIp:        sl.systemConfig.systemInsideIp,
		SystemOs:              sl.systemConfig.systemOs,
		SystemArch:            sl.systemConfig.systemKernel,
		GoVersion:             sl.systemConfig.goVersion,
		SdkVersion:            sl.systemConfig.sdkVersion,
	}

	if sl.slog.status {
		sl.slog.client.WithTraceId(ctx).Info("MiddlewareCustom",
			"request_time", data.RequestTime,
			"request_uri", data.RequestUri,
			"request_url", data.RequestUrl,
			"request_api", data.RequestApi,
			"request_method", data.RequestMethod,
			"request_params", data.RequestParams,
			"request_header", data.RequestHeader,
			"request_ip", data.RequestIp,
			"response_header", data.ResponseHeader,
			"response_status_code", data.ResponseStatusCode,
			"response_body", data.ResponseBody,
			"response_content_length", data.ResponseContentLength,
			"response_time", data.ResponseTime,
			"system_host_name", data.SystemHostName,
			"system_inside_ip", data.SystemInsideIp,
			"system_os", data.SystemOs,
			"system_arch", data.SystemArch,
			"go_version", data.GoVersion,
			"sdk_version", data.SdkVersion,
		)
	}

}
