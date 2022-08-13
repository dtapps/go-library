package golog

import (
	"context"
	"github.com/dtapps/go-library/utils/dorm"
	"github.com/dtapps/go-library/utils/gojson"
	"github.com/dtapps/go-library/utils/gorequest"
	"github.com/dtapps/go-library/utils/gotrace_id"
	"github.com/dtapps/go-library/utils/gourl"
	"gorm.io/datatypes"
	"gorm.io/gorm"
	"log"
	"time"
	"unicode/utf8"
)

// 模型结构体
type apiPostgresqlLog struct {
	LogId                 uint           `gorm:"primaryKey;comment:【记录】编号" json:"log_id,omitempty"`           //【记录】编号
	TraceId               string         `gorm:"index;comment:【系统】跟踪编号" json:"trace_id,omitempty"`            //【系统】跟踪编号
	RequestTime           time.Time      `gorm:"index;comment:【请求】时间" json:"request_time,omitempty"`          //【请求】时间
	RequestUri            string         `gorm:"comment:【请求】链接" json:"request_uri,omitempty"`                 //【请求】链接
	RequestUrl            string         `gorm:"comment:【请求】链接" json:"request_url,omitempty"`                 //【请求】链接
	RequestApi            string         `gorm:"index;comment:【请求】接口" json:"request_api,omitempty"`           //【请求】接口
	RequestMethod         string         `gorm:"index;comment:【请求】方式" json:"request_method,omitempty"`        //【请求】方式
	RequestParams         datatypes.JSON `gorm:"type:jsonb;comment:【请求】参数" json:"request_params,omitempty"`   //【请求】参数
	RequestHeader         datatypes.JSON `gorm:"type:jsonb;comment:【请求】头部" json:"request_header,omitempty"`   //【请求】头部
	ResponseHeader        datatypes.JSON `gorm:"type:jsonb;comment:【返回】头部" json:"response_header,omitempty"`  //【返回】头部
	ResponseStatusCode    int            `gorm:"index;comment:【返回】状态码" json:"response_status_code,omitempty"` //【返回】状态码
	ResponseBody          datatypes.JSON `gorm:"type:jsonb;comment:【返回】内容" json:"response_body,omitempty"`    //【返回】内容
	ResponseContentLength int64          `gorm:"comment:【返回】大小" json:"response_content_length,omitempty"`     //【返回】大小
	ResponseTime          time.Time      `gorm:"index;comment:【返回】时间" json:"response_time,omitempty"`         //【返回】时间
	SystemHostName        string         `gorm:"index;comment:【系统】主机名" json:"system_host_name,omitempty"`     //【系统】主机名
	SystemInsideIp        string         `gorm:"index;comment:【系统】内网ip" json:"system_inside_ip,omitempty"`    //【系统】内网ip
	GoVersion             string         `gorm:"index;comment:【程序】Go版本" json:"go_version,omitempty"`          //【程序】Go版本
	SdkVersion            string         `gorm:"index;comment:【程序】Sdk版本" json:"sdk_version,omitempty"`        //【程序】Sdk版本
}

// 记录日志
func (c *ApiClient) gormRecord(ctx context.Context, postgresqlLog apiPostgresqlLog) error {

	if utf8.ValidString(string(postgresqlLog.ResponseBody)) == false {
		postgresqlLog.ResponseBody = datatypes.JSON("")
	}

	postgresqlLog.SystemHostName = c.config.hostname
	if postgresqlLog.SystemInsideIp == "" {
		postgresqlLog.SystemInsideIp = c.config.insideIp
	}
	postgresqlLog.GoVersion = c.config.goVersion

	postgresqlLog.TraceId = gotrace_id.GetTraceIdContext(ctx)

	return c.gormClient.Table(c.config.tableName).Create(&postgresqlLog).Error
}

// GormQuery 查询
func (c *ApiClient) GormQuery() *gorm.DB {
	return c.gormClient.Table(c.config.tableName)
}

// GormMiddleware 中间件
func (c *ApiClient) GormMiddleware(ctx context.Context, request gorequest.Response, sdkVersion string) {
	if request.ResponseHeader.Get("Content-Type") == "image/jpeg" || request.ResponseHeader.Get("Content-Type") == "image/png" {
		return
	}
	err := c.gormRecord(ctx, apiPostgresqlLog{
		RequestTime:           request.RequestTime,                                              //【请求】时间
		RequestUri:            request.RequestUri,                                               //【请求】链接
		RequestUrl:            gourl.UriParse(request.RequestUri).Url,                           //【请求】链接
		RequestApi:            gourl.UriParse(request.RequestUri).Path,                          //【请求】接口
		RequestMethod:         request.RequestMethod,                                            //【请求】方式
		RequestParams:         datatypes.JSON(gojson.JsonEncodeNoError(request.RequestParams)),  //【请求】参数
		RequestHeader:         datatypes.JSON(gojson.JsonEncodeNoError(request.RequestHeader)),  //【请求】头部
		ResponseHeader:        datatypes.JSON(gojson.JsonEncodeNoError(request.ResponseHeader)), //【返回】头部
		ResponseStatusCode:    request.ResponseStatusCode,                                       //【返回】状态码
		ResponseBody:          request.ResponseBody,                                             //【返回】内容
		ResponseContentLength: request.ResponseContentLength,                                    //【返回】大小
		ResponseTime:          request.ResponseTime,                                             //【返回】时间
		SdkVersion:            sdkVersion,                                                       //【程序】Sdk版本
	})
	if err != nil {
		log.Println("log.GormMiddleware：", err.Error())
	}
}

// GormMiddlewareXml 中间件
func (c *ApiClient) GormMiddlewareXml(ctx context.Context, request gorequest.Response, sdkVersion string) {
	err := c.gormRecord(ctx, apiPostgresqlLog{
		RequestTime:           request.RequestTime,                                                                   //【请求】时间
		RequestUri:            request.RequestUri,                                                                    //【请求】链接
		RequestUrl:            gourl.UriParse(request.RequestUri).Url,                                                //【请求】链接
		RequestApi:            gourl.UriParse(request.RequestUri).Path,                                               //【请求】接口
		RequestMethod:         request.RequestMethod,                                                                 //【请求】方式
		RequestParams:         datatypes.JSON(gojson.JsonEncodeNoError(request.RequestParams)),                       //【请求】参数
		RequestHeader:         datatypes.JSON(gojson.JsonEncodeNoError(request.RequestHeader)),                       //【请求】头部
		ResponseHeader:        datatypes.JSON(gojson.JsonEncodeNoError(request.ResponseHeader)),                      //【返回】头部
		ResponseStatusCode:    request.ResponseStatusCode,                                                            //【返回】状态码
		ResponseBody:          datatypes.JSON(gojson.JsonEncodeNoError(dorm.XmlDecodeNoError(request.ResponseBody))), //【返回】内容
		ResponseContentLength: request.ResponseContentLength,                                                         //【返回】大小
		ResponseTime:          request.ResponseTime,                                                                  //【返回】时间
		SdkVersion:            sdkVersion,                                                                            //【程序】Sdk版本
	})
	if err != nil {
		log.Println("log.GormMiddlewareXml：", err.Error())
	}
}

// GormMiddlewareCustom 中间件
func (c *ApiClient) GormMiddlewareCustom(ctx context.Context, api string, request gorequest.Response, sdkVersion string) {
	err := c.gormRecord(ctx, apiPostgresqlLog{
		RequestTime:           request.RequestTime,                                              //【请求】时间
		RequestUri:            request.RequestUri,                                               //【请求】链接
		RequestUrl:            gourl.UriParse(request.RequestUri).Url,                           //【请求】链接
		RequestApi:            api,                                                              //【请求】接口
		RequestMethod:         request.RequestMethod,                                            //【请求】方式
		RequestParams:         datatypes.JSON(gojson.JsonEncodeNoError(request.RequestParams)),  //【请求】参数
		RequestHeader:         datatypes.JSON(gojson.JsonEncodeNoError(request.RequestHeader)),  //【请求】头部
		ResponseHeader:        datatypes.JSON(gojson.JsonEncodeNoError(request.ResponseHeader)), //【返回】头部
		ResponseStatusCode:    request.ResponseStatusCode,                                       //【返回】状态码
		ResponseBody:          request.ResponseBody,                                             //【返回】内容
		ResponseContentLength: request.ResponseContentLength,                                    //【返回】大小
		ResponseTime:          request.ResponseTime,                                             //【返回】时间
		SdkVersion:            sdkVersion,                                                       //【程序】Sdk版本
	})
	if err != nil {
		log.Println("log.GormMiddlewareCustom：", err.Error())
	}
}
