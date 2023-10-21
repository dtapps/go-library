package golog

import (
	"github.com/dtapps/go-library/utils/dorm"
	"github.com/dtapps/go-library/utils/goip"
	"github.com/dtapps/go-library/utils/gorequest"
	"github.com/dtapps/go-library/utils/gotime"
	"github.com/dtapps/go-library/utils/gourl"
	"github.com/gin-gonic/gin"
	"time"
)

// 结构体
type ginSLog struct {
	TraceId            string                 `json:"trace_id,omitempty"`             //【系统】跟踪编号
	RequestTime        time.Time              `json:"request_time,omitempty"`         //【请求】时间
	RequestUri         string                 `json:"request_uri,omitempty"`          //【请求】请求链接 域名+路径+参数
	RequestUrl         string                 `json:"request_url,omitempty"`          //【请求】请求链接 域名+路径
	RequestApi         string                 `json:"request_api,omitempty"`          //【请求】请求接口 路径
	RequestMethod      string                 `json:"request_method,omitempty"`       //【请求】请求方式
	RequestProto       string                 `json:"request_proto,omitempty"`        //【请求】请求协议
	RequestUa          string                 `json:"request_ua,omitempty"`           //【请求】请求UA
	RequestReferer     string                 `json:"request_referer,omitempty"`      //【请求】请求referer
	RequestBody        string                 `json:"request_body,omitempty"`         //【请求】请求主体
	RequestUrlQuery    map[string][]string    `json:"request_url_query,omitempty"`    //【请求】请求URL参数
	RequestIp          string                 `json:"request_ip,omitempty"`           //【请求】请求客户端Ip
	RequestIpCountry   string                 `json:"request_ip_country,omitempty"`   //【请求】请求客户端城市
	RequestIpProvince  string                 `json:"request_ip_province,omitempty"`  //【请求】请求客户端省份
	RequestIpCity      string                 `json:"request_ip_city,omitempty"`      //【请求】请求客户端城市
	RequestIpIsp       string                 `json:"request_ip_isp,omitempty"`       //【请求】请求客户端运营商
	RequestIpLongitude float64                `json:"request_ip_longitude,omitempty"` //【请求】请求客户端经度
	RequestIpLatitude  float64                `json:"request_ip_latitude,omitempty"`  //【请求】请求客户端纬度
	RequestHeader      map[string][]string    `json:"request_header,omitempty"`       //【请求】请求头
	RequestAllContent  map[string]interface{} `json:"request_all_content,omitempty"`  // 【请求】请求全部内容
	ResponseTime       time.Time              `json:"response_time,omitempty"`        //【返回】时间
	ResponseCode       int                    `json:"response_code,omitempty"`        //【返回】状态码
	ResponseMsg        string                 `json:"response_msg,omitempty"`         //【返回】描述
	ResponseData       string                 `json:"response_data,omitempty"`        //【返回】数据
	CostTime           int64                  `json:"cost_time,omitempty"`            //【系统】花费时间
	SystemHostName     string                 `json:"system_host_name,omitempty"`     //【系统】主机名
	SystemInsideIp     string                 `json:"system_inside_ip,omitempty"`     //【系统】内网ip
	SystemOs           string                 `json:"system_os,omitempty"`            //【系统】系统类型
	SystemArch         string                 `json:"system_arch,omitempty"`          //【系统】系统架构
	GoVersion          string                 `json:"go_version,omitempty"`           //【程序】Go版本
	SdkVersion         string                 `json:"sdk_version,omitempty"`          //【程序】Sdk版本
}

// record 记录日志
func (c *GinClient) record(msg string, data ginSLog) {

	data.SystemHostName = c.config.systemHostname //【系统】主机名
	data.SystemInsideIp = c.config.systemInsideIp //【系统】内网ip
	data.GoVersion = c.config.goVersion           //【程序】Go版本
	data.SdkVersion = c.config.sdkVersion         //【程序】Sdk版本
	data.SystemOs = c.config.systemOs             //【系统】系统类型
	data.SystemArch = c.config.systemKernel       //【系统】系统架构

	c.slog.client.WithTraceIdStr(data.TraceId).Info(msg,
		"request_time", data.RequestTime,
		"request_uri", data.RequestUri,
		"request_url", data.RequestUrl,
		"request_api", data.RequestApi,
		"request_method", data.RequestMethod,
		"request_proto", data.RequestProto,
		"request_ua", data.RequestUa,
		"request_referer", data.RequestReferer,
		"request_body", data.RequestBody,
		"request_url_query", data.RequestUrlQuery,
		"request_ip", data.RequestIp,
		"request_ip_country", data.RequestIpCountry,
		"request_ip_province", data.RequestIpProvince,
		"request_ip_city", data.RequestIpCity,
		"request_ip_isp", data.RequestIpIsp,
		"request_ip_longitude", data.RequestIpLongitude,
		"request_ip_latitude", data.RequestIpLatitude,
		"request_header", data.RequestHeader,
		"request_all_content", data.RequestAllContent,
		"response_time", data.ResponseTime,
		"response_code", data.ResponseCode,
		"response_msg", data.ResponseMsg,
		"response_data", data.ResponseData,
		"cost_time", data.CostTime,
		"system_host_name", data.SystemHostName,
		"system_inside_ip", data.SystemInsideIp,
		"system_os", data.SystemOs,
		"system_arch", data.SystemArch,
		"go_version", data.GoVersion,
		"sdk_version", data.SdkVersion,
	)
}

func (c *GinClient) recordJson(ginCtx *gin.Context, traceId string, requestTime time.Time, paramsBody gorequest.Params, startTime, endTime int64, ipInfo goip.AnalyseResult) {

	data := ginSLog{
		TraceId:            traceId,                                                      //【系统】跟踪编号
		RequestTime:        requestTime,                                                  //【请求】时间
		RequestUrl:         ginCtx.Request.RequestURI,                                    //【请求】请求链接
		RequestApi:         gourl.UriFilterExcludeQueryString(ginCtx.Request.RequestURI), //【请求】请求接口
		RequestMethod:      ginCtx.Request.Method,                                        //【请求】请求方式
		RequestProto:       ginCtx.Request.Proto,                                         //【请求】请求协议
		RequestUa:          ginCtx.Request.UserAgent(),                                   //【请求】请求UA
		RequestReferer:     ginCtx.Request.Referer(),                                     //【请求】请求referer
		RequestUrlQuery:    ginCtx.Request.URL.Query(),                                   //【请求】请求URL参数
		RequestIp:          ipInfo.Ip,                                                    //【请求】请求客户端Ip
		RequestIpCountry:   ipInfo.Country,                                               //【请求】请求客户端城市
		RequestIpProvince:  ipInfo.Province,                                              //【请求】请求客户端省份
		RequestIpCity:      ipInfo.City,                                                  //【请求】请求客户端城市
		RequestIpIsp:       ipInfo.Isp,                                                   //【请求】请求客户端运营商
		RequestIpLatitude:  ipInfo.LocationLatitude,                                      //【请求】请求客户端纬度
		RequestIpLongitude: ipInfo.LocationLongitude,                                     //【请求】请求客户端经度
		RequestHeader:      ginCtx.Request.Header,                                        //【请求】请求头
		RequestAllContent:  paramsBody,                                                   //【请求】请求全部内容
		ResponseTime:       gotime.Current().Time,                                        //【返回】时间
		CostTime:           endTime - startTime,                                          //【系统】花费时间
	}
	if ginCtx.Request.TLS == nil {
		data.RequestUri = "http://" + ginCtx.Request.Host + ginCtx.Request.RequestURI //【请求】请求链接
	} else {
		data.RequestUri = "https://" + ginCtx.Request.Host + ginCtx.Request.RequestURI //【请求】请求链接
	}

	c.record("json", data)
}

func (c *GinClient) recordXml(ginCtx *gin.Context, traceId string, requestTime time.Time, requestBody []byte, paramsBody gorequest.Params, responseCode int, responseBody string, startTime, endTime int64, ipInfo goip.AnalyseResult) {

	data := ginSLog{
		TraceId:            traceId,                                                      //【系统】跟踪编号
		RequestTime:        requestTime,                                                  //【请求】时间
		RequestUrl:         ginCtx.Request.RequestURI,                                    //【请求】请求链接
		RequestApi:         gourl.UriFilterExcludeQueryString(ginCtx.Request.RequestURI), //【请求】请求接口
		RequestMethod:      ginCtx.Request.Method,                                        //【请求】请求方式
		RequestProto:       ginCtx.Request.Proto,                                         //【请求】请求协议
		RequestUa:          ginCtx.Request.UserAgent(),                                   //【请求】请求UA
		RequestReferer:     ginCtx.Request.Referer(),                                     //【请求】请求referer
		RequestUrlQuery:    ginCtx.Request.URL.Query(),                                   //【请求】请求URL参数
		RequestIp:          ipInfo.Ip,                                                    //【请求】请求客户端Ip
		RequestIpCountry:   ipInfo.Country,                                               //【请求】请求客户端城市
		RequestIpProvince:  ipInfo.Province,                                              //【请求】请求客户端省份
		RequestIpCity:      ipInfo.City,                                                  //【请求】请求客户端城市
		RequestIpIsp:       ipInfo.Isp,                                                   //【请求】请求客户端运营商
		RequestIpLatitude:  ipInfo.LocationLatitude,                                      //【请求】请求客户端纬度
		RequestIpLongitude: ipInfo.LocationLongitude,                                     //【请求】请求客户端经度
		RequestHeader:      ginCtx.Request.Header,                                        //【请求】请求头
		RequestAllContent:  paramsBody,                                                   //【请求】请求全部内容
		ResponseTime:       gotime.Current().Time,                                        //【返回】时间
		ResponseCode:       responseCode,                                                 //【返回】状态码
		ResponseData:       responseBody,                                                 //【返回】数据
		CostTime:           endTime - startTime,                                          //【系统】花费时间
	}
	if ginCtx.Request.TLS == nil {
		data.RequestUri = "http://" + ginCtx.Request.Host + ginCtx.Request.RequestURI //【请求】请求链接
	} else {
		data.RequestUri = "https://" + ginCtx.Request.Host + ginCtx.Request.RequestURI //【请求】请求链接
	}

	if len(requestBody) > 0 {
		data.RequestBody = dorm.XmlEncodeNoError(dorm.XmlDecodeNoError(requestBody)) //【请求】请求内容
	}

	c.record("xml", data)
}
