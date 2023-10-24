package golog

import (
	"fmt"
	"github.com/dtapps/go-library/utils/dorm"
	"github.com/dtapps/go-library/utils/goip"
	"github.com/dtapps/go-library/utils/gorequest"
	"github.com/dtapps/go-library/utils/gotime"
	"github.com/dtapps/go-library/utils/gotrace_id"
	"github.com/dtapps/go-library/utils/gourl"
	"github.com/gin-gonic/gin"
	"time"
)

// 结构体
type ginSLogCustom struct {
	LogTime            time.Time `json:"log_time,omitempty"`             //【记日志录】时间
	TraceId            string    `json:"trace_id,omitempty"`             //【系统】跟踪编号
	RequestUri         string    `json:"request_uri,omitempty"`          //【请求】请求链接 域名+路径+参数
	RequestUrl         string    `json:"request_url,omitempty"`          //【请求】请求链接 域名+路径
	RequestApi         string    `json:"request_api,omitempty"`          //【请求】请求接口 路径
	RequestMethod      string    `json:"request_method,omitempty"`       //【请求】请求方式
	RequestProto       string    `json:"request_proto,omitempty"`        //【请求】请求协议
	RequestUa          string    `json:"request_ua,omitempty"`           //【请求】请求UA
	RequestReferer     string    `json:"request_referer,omitempty"`      //【请求】请求referer
	RequestUrlQuery    string    `json:"request_url_query,omitempty"`    //【请求】请求URL参数
	RequestHeader      string    `json:"request_header,omitempty"`       //【请求】请求头
	RequestIp          string    `json:"request_ip,omitempty"`           //【请求】请求客户端Ip
	RequestIpCountry   string    `json:"request_ip_country,omitempty"`   //【请求】请求客户端城市
	RequestIpProvince  string    `json:"request_ip_province,omitempty"`  //【请求】请求客户端省份
	RequestIpCity      string    `json:"request_ip_city,omitempty"`      //【请求】请求客户端城市
	RequestIpIsp       string    `json:"request_ip_isp,omitempty"`       //【请求】请求客户端运营商
	RequestIpLongitude float64   `json:"request_ip_longitude,omitempty"` //【请求】请求客户端经度
	RequestIpLatitude  float64   `json:"request_ip_latitude,omitempty"`  //【请求】请求客户端纬度
	CustomId           string    `json:"custom_id,omitempty"`            //【日志】自定义编号
	CustomType         string    `json:"custom_type,omitempty"`          //【日志】自定义类型
	CustomContent      string    `json:"custom_content,omitempty"`       //【日志】自定义内容
}

type GinCustomClientGinRecordOperation struct {
	slogClient *SLog          // 日志服务
	ipService  *goip.Client   // IP服务
	data       *ginSLogCustom // 数据
}

// GinRecord 记录日志
func (c *GinCustomClient) GinRecord(ginCtx *gin.Context) *GinCustomClientGinRecordOperation {
	operation := &GinCustomClientGinRecordOperation{
		slogClient: c.slog.client,
		ipService:  c.ipService,
	}
	operation.data = new(ginSLogCustom)
	operation.data.LogTime = gotime.Current().Time            //【日志】时间
	operation.data.TraceId = gotrace_id.GetGinTraceId(ginCtx) // 【系统】跟踪编号
	if ginCtx.Request.TLS == nil {
		operation.data.RequestUri = "http://" + ginCtx.Request.Host + ginCtx.Request.RequestURI //【请求】请求链接
	} else {
		operation.data.RequestUri = "https://" + ginCtx.Request.Host + ginCtx.Request.RequestURI //【请求】请求链接
	}
	operation.data.RequestUrl = ginCtx.Request.RequestURI                                    //【请求】请求链接 域名+路径
	operation.data.RequestApi = gourl.UriFilterExcludeQueryString(ginCtx.Request.RequestURI) //【请求】请求接口 路径
	operation.data.RequestMethod = ginCtx.Request.Method                                     //【请求】请求方式
	operation.data.RequestProto = ginCtx.Request.Proto                                       //【请求】请求协议
	operation.data.RequestUa = ginCtx.Request.UserAgent()                                    //【请求】请求UA
	operation.data.RequestReferer = ginCtx.Request.Referer()                                 //【请求】请求referer
	operation.data.RequestUrlQuery = dorm.JsonEncodeNoError(ginCtx.Request.URL.Query())      //【请求】请求URL参数
	operation.data.RequestHeader = dorm.JsonEncodeNoError(ginCtx.Request.Header)             //【请求】请求头
	operation.data.RequestIp = gorequest.ClientIp(ginCtx.Request)                            //【请求】请求客户端Ip
	return operation
}

func (o *GinCustomClientGinRecordOperation) CustomInfo(customId any, customType any, customContent any) *GinCustomClientGinRecordOperation {
	o.data.CustomId = fmt.Sprintf("%s", customId)           //【日志】自定义编号
	o.data.CustomType = fmt.Sprintf("%s", customType)       //【日志】自定义类型
	o.data.CustomContent = fmt.Sprintf("%s", customContent) //【日志】自定义内容
	return o
}

func (o *GinCustomClientGinRecordOperation) CreateData() {
	o.slogClient.WithTraceIdStr(o.data.TraceId).Info("custom",
		"log_time", o.data.LogTime,
		"request_uri", o.data.RequestUri,
		"request_url", o.data.RequestUrl,
		"request_api", o.data.RequestApi,
		"request_method", o.data.RequestMethod,
		"request_proto", o.data.RequestProto,
		"request_ua", o.data.RequestUa,
		"request_referer", o.data.RequestReferer,
		"request_url_query", o.data.RequestUrlQuery,
		"request_header", o.data.RequestHeader,
		"request_ip", o.data.RequestIp,
		"request_ip_country", o.data.RequestIpCountry,
		"request_ip_province", o.data.RequestIpProvince,
		"request_ip_city", o.data.RequestIpCity,
		"request_ip_isp", o.data.RequestIpIsp,
		"request_ip_longitude", o.data.RequestIpLongitude,
		"request_ip_latitude", o.data.RequestIpLatitude,
		"custom_id", o.data.CustomId,
		"custom_type", o.data.CustomType,
		"custom_content", o.data.CustomContent,
	)
}

func (o *GinCustomClientGinRecordOperation) CreateDataNoError() {
	o.slogClient.WithTraceIdStr(o.data.TraceId).Info("custom",
		"log_time", o.data.LogTime,
		"request_uri", o.data.RequestUri,
		"request_url", o.data.RequestUrl,
		"request_api", o.data.RequestApi,
		"request_method", o.data.RequestMethod,
		"request_proto", o.data.RequestProto,
		"request_ua", o.data.RequestUa,
		"request_referer", o.data.RequestReferer,
		"request_url_query", o.data.RequestUrlQuery,
		"request_header", o.data.RequestHeader,
		"request_ip", o.data.RequestIp,
		"request_ip_country", o.data.RequestIpCountry,
		"request_ip_province", o.data.RequestIpProvince,
		"request_ip_city", o.data.RequestIpCity,
		"request_ip_isp", o.data.RequestIpIsp,
		"request_ip_longitude", o.data.RequestIpLongitude,
		"request_ip_latitude", o.data.RequestIpLatitude,
		"custom_id", o.data.CustomId,
		"custom_type", o.data.CustomType,
		"custom_content", o.data.CustomContent,
	)
}
