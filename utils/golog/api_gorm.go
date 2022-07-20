package golog

import (
	"github.com/dtapps/go-library/utils/dorm"
	"github.com/dtapps/go-library/utils/gojson"
	"github.com/dtapps/go-library/utils/gorequest"
	"gorm.io/datatypes"
	"gorm.io/gorm"
	"log"
	"time"
	"unicode/utf8"
)

// 模型结构体
type apiPostgresqlLog struct {
	LogId                 uint           `gorm:"primaryKey" json:"log_id,omitempty"`          //【记录】编号
	RequestTime           time.Time      `gorm:"index" json:"request_time,omitempty"`         //【请求】时间
	RequestUri            string         `json:"request_uri,omitempty"`                       //【请求】链接
	RequestUrl            string         `json:"request_url,omitempty"`                       //【请求】链接
	RequestApi            string         `gorm:"index" json:"request_api,omitempty"`          //【请求】接口
	RequestMethod         string         `gorm:"index" json:"request_method,omitempty"`       //【请求】方式
	RequestParams         datatypes.JSON `gorm:"type:jsonb" json:"request_params,omitempty"`  //【请求】参数
	RequestHeader         datatypes.JSON `gorm:"type:jsonb" json:"request_header,omitempty"`  //【请求】头部
	ResponseHeader        datatypes.JSON `gorm:"type:jsonb" json:"response_header,omitempty"` //【返回】头部
	ResponseStatusCode    int            `gorm:"index" json:"response_status_code,omitempty"` //【返回】状态码
	ResponseBody          datatypes.JSON `gorm:"type:jsonb" json:"response_body,omitempty"`   //【返回】内容
	ResponseContentLength int64          `json:"response_content_length,omitempty"`           //【返回】大小
	ResponseTime          time.Time      `gorm:"index" json:"response_time,omitempty"`        //【返回】时间
	SystemHostName        string         `gorm:"index" json:"system_host_name,omitempty"`     //【系统】主机名
	SystemInsideIp        string         `gorm:"index" json:"system_inside_ip,omitempty"`     //【系统】内网ip
	GoVersion             string         `gorm:"index" json:"go_version,omitempty"`           //【程序】Go版本
}

// 记录日志
func (c *ApiClient) gormRecord(postgresqlLog apiPostgresqlLog) error {

	if utf8.ValidString(string(postgresqlLog.ResponseBody)) == false {
		log.Println("内容格式无法记录")
		postgresqlLog.ResponseBody = datatypes.JSON("")
	}

	postgresqlLog.SystemHostName = c.config.hostname
	if postgresqlLog.SystemInsideIp == "" {
		postgresqlLog.SystemInsideIp = c.config.insideIp
	}
	postgresqlLog.GoVersion = c.config.goVersion

	return c.gormClient.Table(c.config.tableName).Create(&postgresqlLog).Error
}

// GormQuery 查询
func (c *ApiClient) GormQuery() *gorm.DB {
	return c.gormClient.Table(c.config.tableName)
}

// GormMiddleware 中间件
func (c *ApiClient) GormMiddleware(request gorequest.Response) {
	if request.ResponseHeader.Get("Content-Type") == "image/jpeg" || request.ResponseHeader.Get("Content-Type") == "image/png" {
		log.Println("内容格式无法记录")
		return
	}
	c.gormRecord(apiPostgresqlLog{
		RequestTime:           request.RequestTime,                                              //【请求】时间
		RequestUri:            request.RequestUri,                                               //【请求】链接
		RequestUrl:            gorequest.UriParse(request.RequestUri).Url,                       //【请求】链接
		RequestApi:            gorequest.UriParse(request.RequestUri).Path,                      //【请求】接口
		RequestMethod:         request.RequestMethod,                                            //【请求】方式
		RequestParams:         datatypes.JSON(gojson.JsonEncodeNoError(request.RequestParams)),  //【请求】参数
		RequestHeader:         datatypes.JSON(gojson.JsonEncodeNoError(request.RequestHeader)),  //【请求】头部
		ResponseHeader:        datatypes.JSON(gojson.JsonEncodeNoError(request.ResponseHeader)), //【返回】头部
		ResponseStatusCode:    request.ResponseStatusCode,                                       //【返回】状态码
		ResponseBody:          request.ResponseBody,                                             //【返回】内容
		ResponseContentLength: request.ResponseContentLength,                                    //【返回】大小
		ResponseTime:          request.ResponseTime,                                             //【返回】时间
	})
}

// GormMiddlewareXml 中间件
func (c *ApiClient) GormMiddlewareXml(request gorequest.Response) {
	c.gormRecord(apiPostgresqlLog{
		RequestTime:           request.RequestTime,                                                                   //【请求】时间
		RequestUri:            request.RequestUri,                                                                    //【请求】链接
		RequestUrl:            gorequest.UriParse(request.RequestUri).Url,                                            //【请求】链接
		RequestApi:            gorequest.UriParse(request.RequestUri).Path,                                           //【请求】接口
		RequestMethod:         request.RequestMethod,                                                                 //【请求】方式
		RequestParams:         datatypes.JSON(gojson.JsonEncodeNoError(request.RequestParams)),                       //【请求】参数
		RequestHeader:         datatypes.JSON(gojson.JsonEncodeNoError(request.RequestHeader)),                       //【请求】头部
		ResponseHeader:        datatypes.JSON(gojson.JsonEncodeNoError(request.ResponseHeader)),                      //【返回】头部
		ResponseStatusCode:    request.ResponseStatusCode,                                                            //【返回】状态码
		ResponseBody:          datatypes.JSON(gojson.JsonEncodeNoError(dorm.XmlDecodeNoError(request.ResponseBody))), //【返回】内容
		ResponseContentLength: request.ResponseContentLength,                                                         //【返回】大小
		ResponseTime:          request.ResponseTime,                                                                  //【返回】时间
	})
}

// GormMiddlewareCustom 中间件
func (c *ApiClient) GormMiddlewareCustom(api string, request gorequest.Response) {
	c.gormRecord(apiPostgresqlLog{
		RequestTime:           request.RequestTime,                                              //【请求】时间
		RequestUri:            request.RequestUri,                                               //【请求】链接
		RequestUrl:            gorequest.UriParse(request.RequestUri).Url,                       //【请求】链接
		RequestApi:            api,                                                              //【请求】接口
		RequestMethod:         request.RequestMethod,                                            //【请求】方式
		RequestParams:         datatypes.JSON(gojson.JsonEncodeNoError(request.RequestParams)),  //【请求】参数
		RequestHeader:         datatypes.JSON(gojson.JsonEncodeNoError(request.RequestHeader)),  //【请求】头部
		ResponseHeader:        datatypes.JSON(gojson.JsonEncodeNoError(request.ResponseHeader)), //【返回】头部
		ResponseStatusCode:    request.ResponseStatusCode,                                       //【返回】状态码
		ResponseBody:          request.ResponseBody,                                             //【返回】内容
		ResponseContentLength: request.ResponseContentLength,                                    //【返回】大小
		ResponseTime:          request.ResponseTime,                                             //【返回】时间
	})
}
