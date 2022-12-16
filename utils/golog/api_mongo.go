package golog

import (
	"context"
	"github.com/dtapps/go-library/utils/dorm"
	"github.com/dtapps/go-library/utils/gorequest"
	"github.com/dtapps/go-library/utils/gotime"
	"github.com/dtapps/go-library/utils/gotrace_id"
	"github.com/dtapps/go-library/utils/gourl"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// 模型结构体
type apiMongolLog struct {
	LogId                 primitive.ObjectID `json:"log_id,omitempty" bson:"_id,omitempty"`                                      //【记录】编号
	LogTime               primitive.DateTime `json:"log_time,omitempty" bson:"log_time"`                                         //【记录】时间
	TraceId               string             `json:"trace_id,omitempty" bson:"trace_id,omitempty"`                               //【记录】跟踪编号
	RequestTime           string             `json:"request_time,omitempty" bson:"request_time,omitempty"`                       //【请求】时间
	RequestUri            string             `json:"request_uri,omitempty" bson:"request_uri,omitempty"`                         //【请求】链接
	RequestUrl            string             `json:"request_url,omitempty" bson:"request_url,omitempty"`                         //【请求】链接
	RequestApi            string             `json:"request_api,omitempty" bson:"request_api,omitempty"`                         //【请求】接口
	RequestMethod         string             `json:"request_method,omitempty" bson:"request_method,omitempty"`                   //【请求】方式
	RequestParams         interface{}        `json:"request_params,omitempty" bson:"request_params,omitempty"`                   //【请求】参数
	RequestHeader         interface{}        `json:"request_header,omitempty" bson:"request_header,omitempty"`                   //【请求】头部
	ResponseHeader        interface{}        `json:"response_header,omitempty" bson:"response_header,omitempty"`                 //【返回】头部
	ResponseStatusCode    int                `json:"response_status_code,omitempty" bson:"response_status_code,omitempty"`       //【返回】状态码
	ResponseBody          interface{}        `json:"response_body,omitempty" bson:"response_body,omitempty"`                     //【返回】内容
	ResponseContentLength int64              `json:"response_content_length,omitempty" bson:"response_content_length,omitempty"` //【返回】大小
	ResponseTime          string             `json:"response_time,omitempty" bson:"response_time,omitempty"`                     //【返回】时间
	System                struct {
		Hostname        string  `json:"hostname" bson:"hostname"`                                       //【系统】主机名
		Os              string  `json:"os" bson:"os"`                                                   //【系统】系统类型
		Version         string  `json:"version" bson:"version"`                                         //【系统】系统版本
		Kernel          string  `json:"kernel" bson:"kernel"`                                           //【系统】系统内核
		KernelVersion   string  `json:"kernel_version" bson:"kernel_version"`                           //【系统】系统内核版本
		BootTime        string  `json:"boot_time" bson:"boot_time"`                                     //【系统】系统开机时间
		CpuCores        int     `json:"cpu_cores,omitempty" bson:"cpu_cores,omitempty"`                 //【系统】CPU核数
		CpuModelName    string  `json:"cpu_model_name,omitempty" bson:"cpu_model_name,omitempty"`       //【系统】CPU型号名称
		CpuMhz          float64 `json:"cpu_mhz,omitempty" bson:"cpu_mhz,omitempty"`                     //【系统】CPU兆赫
		InsideIp        string  `json:"inside_ip" bson:"inside_ip"`                                     //【系统】内网ip
		OutsideIp       string  `json:"outside_ip" bson:"outside_ip"`                                   //【系统】外网ip
		GoVersion       string  `json:"go_version,omitempty" bson:"go_version,omitempty"`               //【系统】go版本
		SdkVersion      string  `json:"sdk_version,omitempty" bson:"sdk_version,omitempty"`             //【系统】sdk版本
		MongoVersion    string  `json:"mongo_version,omitempty" bson:"mongo_version,omitempty"`         //【系统】mongo版本
		MongoSdkVersion string  `json:"mongo_sdk_version,omitempty" bson:"mongo_sdk_version,omitempty"` //【系统】mongo sdk版本
	} `json:"system,omitempty" bson:"system,omitempty"` //【系统】信息
}

// 创建时间序列集合
func (c *ApiClient) mongoCreateCollection(ctx context.Context) {
	err := c.mongoClient.Database(c.mongoConfig.databaseName).CreateCollection(ctx, c.mongoConfig.collectionName, options.CreateCollection().SetTimeSeriesOptions(options.TimeSeries().SetTimeField("log_time")))
	if err != nil {
		c.zapLog.WithTraceId(ctx).Sugar().Errorf("创建时间序列集合：%s", err)
	}
}

// 创建索引
func (c *ApiClient) mongoCreateIndexes(ctx context.Context) {
	_, err := c.mongoClient.Database(c.mongoConfig.databaseName).Collection(c.mongoConfig.collectionName).CreateManyIndexes(ctx, []mongo.IndexModel{
		{
			Keys: bson.D{{
				Key:   "log_time",
				Value: -1,
			}},
		}})
	if err != nil {
		c.zapLog.WithTraceId(ctx).Sugar().Errorf("创建索引：%s", err)
	}
}

// MongoDelete 删除
func (c *ApiClient) MongoDelete(ctx context.Context, hour int64) (*mongo.DeleteResult, error) {
	filter := bson.D{{"log_time", bson.D{{"$lt", primitive.NewDateTimeFromTime(gotime.Current().BeforeHour(hour).Time)}}}}
	return c.mongoClient.Database(c.mongoConfig.databaseName).Collection(c.mongoConfig.collectionName).DeleteMany(ctx, filter)
}

// 记录日志
func (c *ApiClient) mongoRecord(ctx context.Context, data apiMongolLog, sdkVersion string) {

	data.LogId = primitive.NewObjectID()                                                                          //【记录】编号
	data.TraceId = gotrace_id.GetTraceIdContext(ctx)                                                              //【记录】跟踪编号
	data.System.Hostname = c.config.systemHostname                                                                //【系统】主机名
	data.System.Os = c.config.systemOs                                                                            //【系统】系统类型
	data.System.Version = c.config.systemVersion                                                                  //【系统】系统版本
	data.System.Kernel = c.config.systemKernel                                                                    //【系统】系统内核
	data.System.KernelVersion = c.config.systemKernelVersion                                                      //【系统】系统内核版本
	data.System.BootTime = gotime.SetCurrent(gotime.SetCurrentUnix(int64(c.config.systemBootTime)).Time).Format() //【系统】系统开机时间
	data.System.CpuCores = c.config.cpuCores                                                                      //【系统】CPU核数
	data.System.CpuModelName = c.config.cpuModelName                                                              //【系统】CPU型号名称
	data.System.CpuMhz = c.config.cpuMhz                                                                          //【程序】CPU兆赫
	data.System.InsideIp = c.config.systemInsideIp                                                                //【系统】内网ip
	data.System.OutsideIp = c.config.systemOutsideIp                                                              //【系统】外网ip
	data.System.GoVersion = c.config.goVersion                                                                    //【系统】Go版本
	data.System.SdkVersion = sdkVersion                                                                           //【系统】Sdk版本
	data.System.MongoVersion = c.config.mongoVersion                                                              //【系统】mongo版本
	data.System.MongoSdkVersion = c.config.mongoSdkVersion                                                        //【系统】mongo sdk版本

	_, err := c.mongoClient.Database(c.mongoConfig.databaseName).Collection(c.mongoConfig.collectionName).InsertOne(ctx, data)
	if err != nil {
		c.zapLog.WithTraceId(ctx).Sugar().Errorf("保存接口日志错误：%s", err)
		c.zapLog.WithTraceId(ctx).Sugar().Errorf("保存接口日志数据：%+v", data)
	}
}

// 中间件
func (c *ApiClient) mongoMiddleware(ctx context.Context, request gorequest.Response, sdkVersion string) {
	data := apiMongolLog{
		LogTime:               primitive.NewDateTimeFromTime(request.RequestTime), //【记录】时间
		RequestTime:           gotime.SetCurrent(request.RequestTime).Format(),    //【请求】时间
		RequestUri:            request.RequestUri,                                 //【请求】链接
		RequestUrl:            gourl.UriParse(request.RequestUri).Url,             //【请求】链接
		RequestApi:            gourl.UriParse(request.RequestUri).Path,            //【请求】接口
		RequestMethod:         request.RequestMethod,                              //【请求】方式
		RequestParams:         request.RequestParams,                              //【请求】参数
		RequestHeader:         request.RequestHeader,                              //【请求】头部
		ResponseHeader:        request.ResponseHeader,                             //【返回】头部
		ResponseStatusCode:    request.ResponseStatusCode,                         //【返回】状态码
		ResponseContentLength: request.ResponseContentLength,                      //【返回】大小
		ResponseTime:          gotime.SetCurrent(request.ResponseTime).Format(),   //【返回】时间
	}
	if !request.HeaderIsImg() {
		if len(request.ResponseBody) > 0 {
			data.ResponseBody = dorm.JsonDecodeNoError(request.ResponseBody) //【返回】内容
		}
	}

	c.mongoRecord(ctx, data, sdkVersion)
}

// 中间件
func (c *ApiClient) mongoMiddlewareXml(ctx context.Context, request gorequest.Response, sdkVersion string) {
	data := apiMongolLog{
		LogTime:               primitive.NewDateTimeFromTime(request.RequestTime), //【记录】时间
		RequestTime:           gotime.SetCurrent(request.RequestTime).Format(),    //【请求】时间
		RequestUri:            request.RequestUri,                                 //【请求】链接
		RequestUrl:            gourl.UriParse(request.RequestUri).Url,             //【请求】链接
		RequestApi:            gourl.UriParse(request.RequestUri).Path,            //【请求】接口
		RequestMethod:         request.RequestMethod,                              //【请求】方式
		RequestParams:         request.RequestParams,                              //【请求】参数
		RequestHeader:         request.RequestHeader,                              //【请求】头部
		ResponseHeader:        request.ResponseHeader,                             //【返回】头部
		ResponseStatusCode:    request.ResponseStatusCode,                         //【返回】状态码
		ResponseContentLength: request.ResponseContentLength,                      //【返回】大小
		ResponseTime:          gotime.SetCurrent(request.ResponseTime).Format(),   //【返回】时间
	}
	if !request.HeaderIsImg() {
		if len(request.ResponseBody) > 0 {
			data.ResponseBody = dorm.XmlDecodeNoError(request.ResponseBody) //【返回】内容
		}
	}

	c.mongoRecord(ctx, data, sdkVersion)
}

// 中间件
func (c *ApiClient) mongoMiddlewareCustom(ctx context.Context, api string, request gorequest.Response, sdkVersion string) {
	data := apiMongolLog{
		LogTime:               primitive.NewDateTimeFromTime(request.RequestTime), //【记录】时间
		RequestTime:           gotime.SetCurrent(request.RequestTime).Format(),    //【请求】时间
		RequestUri:            request.RequestUri,                                 //【请求】链接
		RequestUrl:            gourl.UriParse(request.RequestUri).Url,             //【请求】链接
		RequestApi:            api,                                                //【请求】接口
		RequestMethod:         request.RequestMethod,                              //【请求】方式
		RequestParams:         request.RequestParams,                              //【请求】参数
		RequestHeader:         request.RequestHeader,                              //【请求】头部
		ResponseHeader:        request.ResponseHeader,                             //【返回】头部
		ResponseStatusCode:    request.ResponseStatusCode,                         //【返回】状态码
		ResponseContentLength: request.ResponseContentLength,                      //【返回】大小
		ResponseTime:          gotime.SetCurrent(request.ResponseTime).Format(),   //【返回】时间
	}
	if !request.HeaderIsImg() {
		if len(request.ResponseBody) > 0 {
			data.ResponseBody = dorm.JsonDecodeNoError(request.ResponseBody) //【返回】内容
		}
	}

	c.mongoRecord(ctx, data, sdkVersion)
}
