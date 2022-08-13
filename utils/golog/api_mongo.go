package golog

import (
	"context"
	"github.com/dtapps/go-library/utils/dorm"
	"github.com/dtapps/go-library/utils/gorequest"
	"github.com/dtapps/go-library/utils/gotime"
	"github.com/dtapps/go-library/utils/gotrace_id"
	"github.com/dtapps/go-library/utils/gourl"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// 模型结构体
type apiMongoLog struct {
	LogId                 primitive.ObjectID     `json:"log_id,omitempty" bson:"_id,omitempty"`                                      //【记录】编号
	TraceId               string                 `json:"trace_id,omitempty" bson:"trace_id,omitempty"`                               //【系统】跟踪编号
	RequestTime           int64                  `json:"request_time,omitempty" bson:"request_time,omitempty"`                       //【请求】时间
	RequestUri            string                 `json:"request_uri,omitempty" bson:"request_uri,omitempty"`                         //【请求】链接
	RequestUrl            string                 `json:"request_url,omitempty" bson:"request_url,omitempty"`                         //【请求】链接
	RequestApi            string                 `json:"request_api,omitempty" bson:"request_api,omitempty"`                         //【请求】接口
	RequestMethod         string                 `json:"request_method,omitempty" bson:"request_method,omitempty"`                   //【请求】方式
	RequestParams         map[string]interface{} `json:"request_params,omitempty" bson:"request_params,omitempty"`                   //【请求】参数
	RequestHeader         map[string]string      `json:"request_header,omitempty" bson:"request_header,omitempty"`                   //【请求】头部
	ResponseHeader        map[string][]string    `json:"response_header,omitempty" bson:"response_header,omitempty"`                 //【返回】头部
	ResponseStatusCode    int                    `json:"response_status_code,omitempty" bson:"response_status_code,omitempty"`       //【返回】状态码
	ResponseBody          interface{}            `json:"response_body,omitempty" bson:"response_body,omitempty"`                     //【返回】内容
	ResponseContentLength int64                  `json:"response_content_length,omitempty" bson:"response_content_length,omitempty"` //【返回】大小
	ResponseTime          int64                  `json:"response_time,omitempty" bson:"response_time,omitempty"`                     //【返回】时间
	SystemHostName        string                 `json:"system_host_name,omitempty" bson:"system_host_name,omitempty"`               //【系统】主机名
	SystemInsideIp        string                 `json:"system_inside_ip,omitempty" bson:"system_inside_ip,omitempty"`               //【系统】内网ip
	GoVersion             string                 `json:"go_version,omitempty" bson:"go_version,omitempty"`                           //【程序】Go版本
	SdkVersion            string                 `json:"sdk_version,omitempty" bson:"sdk_version,omitempty"`                         //【程序】Sdk版本
}

// 记录日志
func (c *ApiClient) mongoRecord(ctx context.Context, mongoLog apiMongoLog) error {

	mongoLog.SystemHostName = c.config.hostname
	if mongoLog.SystemInsideIp == "" {
		mongoLog.SystemInsideIp = c.config.insideIp
	}
	mongoLog.GoVersion = c.config.goVersion

	mongoLog.TraceId = gotrace_id.GetTraceIdContext(ctx)

	mongoLog.LogId = primitive.NewObjectID()

	_, err := c.mongoClient.Database(c.config.databaseName).Collection(c.config.collectionName).InsertOne(mongoLog)

	return err
}

// MongoQuery 查询
func (c *ApiClient) MongoQuery() *dorm.MongoClient {
	return c.mongoClient.Database(c.config.databaseName).Collection(c.config.collectionName)
}

// MongoMiddleware 中间件
func (c *ApiClient) MongoMiddleware(ctx context.Context, request gorequest.Response, sdkVersion string) {
	c.mongoRecord(ctx, apiMongoLog{
		RequestTime:           gotime.SetCurrent(request.RequestTime).Timestamp(),  //【请求】时间
		RequestUri:            request.RequestUri,                                  //【请求】链接
		RequestUrl:            gourl.UriParse(request.RequestUri).Url,              //【请求】链接
		RequestApi:            gourl.UriParse(request.RequestUri).Path,             //【请求】接口
		RequestMethod:         request.RequestMethod,                               //【请求】方式
		RequestParams:         request.RequestParams,                               //【请求】参数
		RequestHeader:         request.RequestHeader,                               //【请求】头部
		ResponseHeader:        request.ResponseHeader,                              //【返回】头部
		ResponseStatusCode:    request.ResponseStatusCode,                          //【返回】状态码
		ResponseBody:          request.ResponseBody,                                //【返回】内容
		ResponseContentLength: request.ResponseContentLength,                       //【返回】大小
		ResponseTime:          gotime.SetCurrent(request.ResponseTime).Timestamp(), //【返回】时间
		SdkVersion:            sdkVersion,                                          //【程序】Sdk版本
	})
}

// MongoMiddlewareXml 中间件
func (c *ApiClient) MongoMiddlewareXml(ctx context.Context, request gorequest.Response, sdkVersion string) {
	c.mongoRecord(ctx, apiMongoLog{
		RequestTime:           gotime.SetCurrent(request.RequestTime).Timestamp(),  //【请求】时间
		RequestUri:            request.RequestUri,                                  //【请求】链接
		RequestUrl:            gourl.UriParse(request.RequestUri).Url,              //【请求】链接
		RequestApi:            gourl.UriParse(request.RequestUri).Path,             //【请求】接口
		RequestMethod:         request.RequestMethod,                               //【请求】方式
		RequestParams:         request.RequestParams,                               //【请求】参数
		RequestHeader:         request.RequestHeader,                               //【请求】头部
		ResponseHeader:        request.ResponseHeader,                              //【返回】头部
		ResponseStatusCode:    request.ResponseStatusCode,                          //【返回】状态码
		ResponseBody:          dorm.XmlDecodeNoError(request.ResponseBody),         //【返回】内容
		ResponseContentLength: request.ResponseContentLength,                       //【返回】大小
		ResponseTime:          gotime.SetCurrent(request.ResponseTime).Timestamp(), //【返回】时间
		SdkVersion:            sdkVersion,                                          //【程序】Sdk版本
	})
}

// MongoMiddlewareCustom 中间件
func (c *ApiClient) MongoMiddlewareCustom(ctx context.Context, api string, request gorequest.Response, sdkVersion string) {
	c.mongoRecord(ctx, apiMongoLog{
		RequestTime:           gotime.SetCurrent(request.RequestTime).Timestamp(),  //【请求】时间
		RequestUri:            request.RequestUri,                                  //【请求】链接
		RequestUrl:            gourl.UriParse(request.RequestUri).Url,              //【请求】链接
		RequestApi:            api,                                                 //【请求】接口
		RequestMethod:         request.RequestMethod,                               //【请求】方式
		RequestParams:         request.RequestParams,                               //【请求】参数
		RequestHeader:         request.RequestHeader,                               //【请求】头部
		ResponseHeader:        request.ResponseHeader,                              //【返回】头部
		ResponseStatusCode:    request.ResponseStatusCode,                          //【返回】状态码
		ResponseBody:          request.ResponseBody,                                //【返回】内容
		ResponseContentLength: request.ResponseContentLength,                       //【返回】大小
		ResponseTime:          gotime.SetCurrent(request.ResponseTime).Timestamp(), //【返回】时间
		SdkVersion:            sdkVersion,                                          //【程序】Sdk版本
	})
}
