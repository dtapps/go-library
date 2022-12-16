package golog

import (
	"context"
	"github.com/dtapps/go-library/utils/dorm"
	"github.com/dtapps/go-library/utils/gorequest"
	"github.com/dtapps/go-library/utils/gotime"
	"github.com/dtapps/go-library/utils/gotrace_id"
	"github.com/dtapps/go-library/utils/gourl"
	"time"
	"unicode/utf8"
)

// 模型
type apiPostgresqlLog struct {
	LogId                 uint      `gorm:"primaryKey;comment:【记录】编号" json:"log_id,omitempty"`                  //【记录】编号
	TraceId               string    `gorm:"index;comment:【系统】跟踪编号" json:"trace_id,omitempty"`                   //【系统】跟踪编号
	RequestTime           time.Time `gorm:"index;comment:【请求】时间" json:"request_time,omitempty"`                 //【请求】时间
	RequestUri            string    `gorm:"comment:【请求】链接" json:"request_uri,omitempty"`                        //【请求】链接
	RequestUrl            string    `gorm:"comment:【请求】链接" json:"request_url,omitempty"`                        //【请求】链接
	RequestApi            string    `gorm:"index;comment:【请求】接口" json:"request_api,omitempty"`                  //【请求】接口
	RequestMethod         string    `gorm:"index;comment:【请求】方式" json:"request_method,omitempty"`               //【请求】方式
	RequestParams         string    `gorm:"comment:【请求】参数" json:"request_params,omitempty"`                     //【请求】参数
	RequestHeader         string    `gorm:"comment:【请求】头部" json:"request_header,omitempty"`                     //【请求】头部
	RequestIp             string    `gorm:"default:0.0.0.0;index;comment:【请求】请求Ip" json:"request_ip,omitempty"` //【请求】请求Ip
	ResponseHeader        string    `gorm:"comment:【返回】头部" json:"response_header,omitempty"`                    //【返回】头部
	ResponseStatusCode    int       `gorm:"index;comment:【返回】状态码" json:"response_status_code,omitempty"`        //【返回】状态码
	ResponseBody          string    `gorm:"comment:【返回】数据" json:"response_content,omitempty"`                   //【返回】数据
	ResponseContentLength int64     `gorm:"comment:【返回】大小" json:"response_content_length,omitempty"`            //【返回】大小
	ResponseTime          time.Time `gorm:"index;comment:【返回】时间" json:"response_time,omitempty"`                //【返回】时间
	SystemHostName        string    `gorm:"index;comment:【系统】主机名" json:"system_host_name,omitempty"`            //【系统】主机名
	SystemInsideIp        string    `gorm:"default:0.0.0.0;comment:【系统】内网ip" json:"system_inside_ip,omitempty"` //【系统】内网ip
	SystemOs              string    `gorm:"index;comment:【系统】系统类型" json:"system_os,omitempty"`                  //【系统】系统类型
	SystemArch            string    `gorm:"index;comment:【系统】系统架构" json:"system_arch,omitempty"`                //【系统】系统架构
	GoVersion             string    `gorm:"comment:【程序】Go版本" json:"go_version,omitempty"`                       //【程序】Go版本
	SdkVersion            string    `gorm:"comment:【程序】Sdk版本" json:"sdk_version,omitempty"`                     //【程序】Sdk版本
}

// 创建模型
func (c *ApiClient) gormAutoMigrate(ctx context.Context) {
	err := c.gormClient.GetDb().Table(c.gormConfig.tableName).AutoMigrate(&apiPostgresqlLog{})
	if err != nil {
		c.zapLog.WithTraceId(ctx).Sugar().Errorf("创建模型：%s", err)
	}
}

// 记录日志
func (c *ApiClient) gormRecord(ctx context.Context, data apiPostgresqlLog) {

	if utf8.ValidString(data.ResponseBody) == false {
		data.ResponseBody = ""
	}

	data.SystemHostName = c.config.systemHostname    //【系统】主机名
	data.SystemInsideIp = c.config.systemInsideIp    //【系统】内网ip
	data.GoVersion = c.config.goVersion              //【程序】Go版本
	data.TraceId = gotrace_id.GetTraceIdContext(ctx) //【记录】跟踪编号
	data.RequestIp = c.config.systemOutsideIp        //【请求】请求Ip
	data.SystemOs = c.config.systemOs                //【系统】系统类型
	data.SystemArch = c.config.systemKernel          //【系统】系统架构

	err := c.gormClient.GetDb().Table(c.gormConfig.tableName).Create(&data).Error
	if err != nil {
		c.zapLog.WithTraceId(ctx).Sugar().Errorf("记录接口日志错误：%s", err)
		c.zapLog.WithTraceId(ctx).Sugar().Errorf("记录接口日志数据：%+v", data)
	}
}

// GormDelete 删除
func (c *ApiClient) GormDelete(ctx context.Context, hour int64) error {
	err := c.gormClient.GetDb().Table(c.gormConfig.tableName).Where("request_time < ?", gotime.Current().BeforeHour(hour).Format()).Delete(&apiPostgresqlLog{}).Error
	if err != nil {
		c.zapLog.WithTraceId(ctx).Sugar().Errorf("删除失败：%s", err)
	}
	return err
}

// 中间件
func (c *ApiClient) gormMiddleware(ctx context.Context, request gorequest.Response, sdkVersion string) {
	data := apiPostgresqlLog{
		RequestTime:           request.RequestTime,                            //【请求】时间
		RequestUri:            request.RequestUri,                             //【请求】链接
		RequestUrl:            gourl.UriParse(request.RequestUri).Url,         //【请求】链接
		RequestApi:            gourl.UriParse(request.RequestUri).Path,        //【请求】接口
		RequestMethod:         request.RequestMethod,                          //【请求】方式
		RequestParams:         dorm.JsonEncodeNoError(request.RequestParams),  //【请求】参数
		RequestHeader:         dorm.JsonEncodeNoError(request.RequestHeader),  //【请求】头部
		ResponseHeader:        dorm.JsonEncodeNoError(request.ResponseHeader), //【返回】头部
		ResponseStatusCode:    request.ResponseStatusCode,                     //【返回】状态码
		ResponseContentLength: request.ResponseContentLength,                  //【返回】大小
		ResponseTime:          request.ResponseTime,                           //【返回】时间
		SdkVersion:            sdkVersion,                                     //【程序】Sdk版本
	}
	if !request.HeaderIsImg() {
		if len(request.ResponseBody) > 0 {
			data.ResponseBody = dorm.JsonEncodeNoError(dorm.JsonDecodeNoError(request.ResponseBody)) //【返回】数据
		}
	}

	c.gormRecord(ctx, data)
}

// 中间件
func (c *ApiClient) gormMiddlewareXml(ctx context.Context, request gorequest.Response, sdkVersion string) {
	data := apiPostgresqlLog{
		RequestTime:           request.RequestTime,                            //【请求】时间
		RequestUri:            request.RequestUri,                             //【请求】链接
		RequestUrl:            gourl.UriParse(request.RequestUri).Url,         //【请求】链接
		RequestApi:            gourl.UriParse(request.RequestUri).Path,        //【请求】接口
		RequestMethod:         request.RequestMethod,                          //【请求】方式
		RequestParams:         dorm.JsonEncodeNoError(request.RequestParams),  //【请求】参数
		RequestHeader:         dorm.JsonEncodeNoError(request.RequestHeader),  //【请求】头部
		ResponseHeader:        dorm.JsonEncodeNoError(request.ResponseHeader), //【返回】头部
		ResponseStatusCode:    request.ResponseStatusCode,                     //【返回】状态码
		ResponseContentLength: request.ResponseContentLength,                  //【返回】大小
		ResponseTime:          request.ResponseTime,                           //【返回】时间
		SdkVersion:            sdkVersion,                                     //【程序】Sdk版本
	}
	if !request.HeaderIsImg() {
		if len(request.ResponseBody) > 0 {
			data.ResponseBody = dorm.JsonEncodeNoError(request.ResponseBody) //【返回】内容
		}
	}

	c.gormRecord(ctx, data)
}

// 中间件
func (c *ApiClient) gormMiddlewareCustom(ctx context.Context, api string, request gorequest.Response, sdkVersion string) {
	data := apiPostgresqlLog{
		RequestTime:           request.RequestTime,                            //【请求】时间
		RequestUri:            request.RequestUri,                             //【请求】链接
		RequestUrl:            gourl.UriParse(request.RequestUri).Url,         //【请求】链接
		RequestApi:            api,                                            //【请求】接口
		RequestMethod:         request.RequestMethod,                          //【请求】方式
		RequestParams:         dorm.JsonEncodeNoError(request.RequestParams),  //【请求】参数
		RequestHeader:         dorm.JsonEncodeNoError(request.RequestHeader),  //【请求】头部
		ResponseHeader:        dorm.JsonEncodeNoError(request.ResponseHeader), //【返回】头部
		ResponseStatusCode:    request.ResponseStatusCode,                     //【返回】状态码
		ResponseContentLength: request.ResponseContentLength,                  //【返回】大小
		ResponseTime:          request.ResponseTime,                           //【返回】时间
		SdkVersion:            sdkVersion,                                     //【程序】Sdk版本
	}
	if !request.HeaderIsImg() {
		if len(request.ResponseBody) > 0 {
			data.ResponseBody = dorm.JsonEncodeNoError(dorm.JsonDecodeNoError(request.ResponseBody)) //【返回】数据
		}
	}

	c.gormRecord(ctx, data)
}
