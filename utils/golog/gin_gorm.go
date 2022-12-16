package golog

import (
	"context"
	"github.com/dtapps/go-library/utils/dorm"
	"github.com/dtapps/go-library/utils/goip"
	"github.com/dtapps/go-library/utils/gotime"
	"github.com/dtapps/go-library/utils/gourl"
	"github.com/gin-gonic/gin"
	"time"
)

// 模型
type ginPostgresqlLog struct {
	LogId              uint      `gorm:"primaryKey;comment:【记录】编号" json:"log_id,omitempty"`                     //【记录】编号
	TraceId            string    `gorm:"index;comment:【系统】跟踪编号" json:"trace_id,omitempty"`                      //【系统】跟踪编号
	RequestTime        time.Time `gorm:"index;comment:【请求】时间" json:"request_time,omitempty"`                    //【请求】时间
	RequestUri         string    `gorm:"comment:【请求】请求链接 域名+路径+参数" json:"request_uri,omitempty"`                //【请求】请求链接 域名+路径+参数
	RequestUrl         string    `gorm:"comment:【请求】请求链接 域名+路径" json:"request_url,omitempty"`                   //【请求】请求链接 域名+路径
	RequestApi         string    `gorm:"index;comment:【请求】请求接口 路径" json:"request_api,omitempty"`                //【请求】请求接口 路径
	RequestMethod      string    `gorm:"index;comment:【请求】请求方式" json:"request_method,omitempty"`                //【请求】请求方式
	RequestProto       string    `gorm:"comment:【请求】请求协议" json:"request_proto,omitempty"`                       //【请求】请求协议
	RequestUa          string    `gorm:"comment:【请求】请求UA" json:"request_ua,omitempty"`                          //【请求】请求UA
	RequestReferer     string    `gorm:"comment:【请求】请求referer" json:"request_referer,omitempty"`                //【请求】请求referer
	RequestBody        string    `gorm:"comment:【请求】请求主体" json:"request_body,omitempty"`                        //【请求】请求主体
	RequestUrlQuery    string    `gorm:"comment:【请求】请求URL参数" json:"request_url_query,omitempty"`                //【请求】请求URL参数
	RequestIp          string    `gorm:"default:0.0.0.0;index;comment:【请求】请求客户端Ip" json:"request_ip,omitempty"` //【请求】请求客户端Ip
	RequestIpCountry   string    `gorm:"index;comment:【请求】请求客户端城市" json:"request_ip_country,omitempty"`         //【请求】请求客户端城市
	RequestIpProvince  string    `gorm:"index;comment:【请求】请求客户端省份" json:"request_ip_province,omitempty"`        //【请求】请求客户端省份
	RequestIpCity      string    `gorm:"index;comment:【请求】请求客户端城市" json:"request_ip_city,omitempty"`            //【请求】请求客户端城市
	RequestIpIsp       string    `gorm:"index;comment:【请求】请求客户端运营商" json:"request_ip_isp,omitempty"`            //【请求】请求客户端运营商
	RequestIpLongitude float64   `gorm:"index;comment:【请求】请求客户端经度" json:"request_ip_longitude,omitempty"`       //【请求】请求客户端经度
	RequestIpLatitude  float64   `gorm:"index;comment:【请求】请求客户端纬度" json:"request_ip_latitude,omitempty"`        //【请求】请求客户端纬度
	RequestHeader      string    `gorm:"comment:【请求】请求头" json:"request_header,omitempty"`                       //【请求】请求头
	ResponseTime       time.Time `gorm:"index;comment:【返回】时间" json:"response_time,omitempty"`                   //【返回】时间
	ResponseCode       int       `gorm:"index;comment:【返回】状态码" json:"response_code,omitempty"`                  //【返回】状态码
	ResponseMsg        string    `gorm:"comment:【返回】描述" json:"response_msg,omitempty"`                          //【返回】描述
	ResponseData       string    `gorm:"comment:【返回】数据" json:"response_data,omitempty"`                         //【返回】数据
	CostTime           int64     `gorm:"comment:【系统】花费时间" json:"cost_time,omitempty"`                           //【系统】花费时间
	SystemHostName     string    `gorm:"index;comment:【系统】主机名" json:"system_host_name,omitempty"`               //【系统】主机名
	SystemInsideIp     string    `gorm:"default:0.0.0.0;comment:【系统】内网ip" json:"system_inside_ip,omitempty"`    //【系统】内网ip
	SystemOs           string    `gorm:"index;comment:【系统】系统类型" json:"system_os,omitempty"`                     //【系统】系统类型
	SystemArch         string    `gorm:"index;comment:【系统】系统架构" json:"system_arch,omitempty"`                   //【系统】系统架构
	GoVersion          string    `gorm:"comment:【程序】Go版本" json:"go_version,omitempty"`                          //【程序】Go版本
	SdkVersion         string    `gorm:"comment:【程序】Sdk版本" json:"sdk_version,omitempty"`                        //【程序】Sdk版本
}

// 创建模型
func (c *GinClient) gormAutoMigrate(ctx context.Context) {
	err := c.gormClient.GetDb().Table(c.gormConfig.tableName).AutoMigrate(&ginPostgresqlLog{})
	if err != nil {
		c.zapLog.WithTraceId(ctx).Sugar().Errorf("创建模型：%s", err)
	}
}

// gormRecord 记录日志
func (c *GinClient) gormRecord(data ginPostgresqlLog) {

	data.SystemHostName = c.config.systemHostname //【系统】主机名
	data.SystemInsideIp = c.config.systemInsideIp //【系统】内网ip
	data.GoVersion = c.config.goVersion           //【程序】Go版本
	data.SdkVersion = c.config.sdkVersion         //【程序】Sdk版本
	data.SystemOs = c.config.systemOs             //【系统】系统类型
	data.SystemArch = c.config.systemKernel       //【系统】系统架构

	err := c.gormClient.GetDb().Table(c.gormConfig.tableName).Create(&data).Error
	if err != nil {
		c.zapLog.WithTraceIdStr(data.TraceId).Sugar().Errorf("记录框架日志错误：%s", err)
		c.zapLog.WithTraceIdStr(data.TraceId).Sugar().Errorf("记录框架日志数据：%+v", data)
	}
}

func (c *GinClient) gormRecordJson(ginCtx *gin.Context, traceId string, requestTime time.Time, requestBody []byte, responseCode int, responseBody string, startTime, endTime int64, ipInfo goip.AnalyseResult) {

	data := ginPostgresqlLog{
		TraceId:            traceId,                                                      //【系统】跟踪编号
		RequestTime:        requestTime,                                                  //【请求】时间
		RequestUrl:         ginCtx.Request.RequestURI,                                    //【请求】请求链接
		RequestApi:         gourl.UriFilterExcludeQueryString(ginCtx.Request.RequestURI), //【请求】请求接口
		RequestMethod:      ginCtx.Request.Method,                                        //【请求】请求方式
		RequestProto:       ginCtx.Request.Proto,                                         //【请求】请求协议
		RequestUa:          ginCtx.Request.UserAgent(),                                   //【请求】请求UA
		RequestReferer:     ginCtx.Request.Referer(),                                     //【请求】请求referer
		RequestUrlQuery:    dorm.JsonEncodeNoError(ginCtx.Request.URL.Query()),           //【请求】请求URL参数
		RequestIp:          ipInfo.Ip,                                                    //【请求】请求客户端Ip
		RequestIpCountry:   ipInfo.Country,                                               //【请求】请求客户端城市
		RequestIpProvince:  ipInfo.Province,                                              //【请求】请求客户端省份
		RequestIpCity:      ipInfo.City,                                                  //【请求】请求客户端城市
		RequestIpIsp:       ipInfo.Isp,                                                   //【请求】请求客户端运营商
		RequestIpLatitude:  ipInfo.LocationLatitude,                                      //【请求】请求客户端纬度
		RequestIpLongitude: ipInfo.LocationLongitude,                                     //【请求】请求客户端经度
		RequestHeader:      dorm.JsonEncodeNoError(ginCtx.Request.Header),                //【请求】请求头
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
		data.RequestBody = dorm.JsonEncodeNoError(requestBody) //【请求】请求主体
	}

	c.gormRecord(data)
}

func (c *GinClient) gormRecordXml(ginCtx *gin.Context, traceId string, requestTime time.Time, requestBody []byte, responseCode int, responseBody string, startTime, endTime int64, ipInfo goip.AnalyseResult) {

	data := ginPostgresqlLog{
		TraceId:            traceId,                                                      //【系统】跟踪编号
		RequestTime:        requestTime,                                                  //【请求】时间
		RequestUrl:         ginCtx.Request.RequestURI,                                    //【请求】请求链接
		RequestApi:         gourl.UriFilterExcludeQueryString(ginCtx.Request.RequestURI), //【请求】请求接口
		RequestMethod:      ginCtx.Request.Method,                                        //【请求】请求方式
		RequestProto:       ginCtx.Request.Proto,                                         //【请求】请求协议
		RequestUa:          ginCtx.Request.UserAgent(),                                   //【请求】请求UA
		RequestReferer:     ginCtx.Request.Referer(),                                     //【请求】请求referer
		RequestUrlQuery:    dorm.JsonEncodeNoError(ginCtx.Request.URL.Query()),           //【请求】请求URL参数
		RequestIp:          ipInfo.Ip,                                                    //【请求】请求客户端Ip
		RequestIpCountry:   ipInfo.Country,                                               //【请求】请求客户端城市
		RequestIpProvince:  ipInfo.Province,                                              //【请求】请求客户端省份
		RequestIpCity:      ipInfo.City,                                                  //【请求】请求客户端城市
		RequestIpIsp:       ipInfo.Isp,                                                   //【请求】请求客户端运营商
		RequestIpLatitude:  ipInfo.LocationLatitude,                                      //【请求】请求客户端纬度
		RequestIpLongitude: ipInfo.LocationLongitude,                                     //【请求】请求客户端经度
		RequestHeader:      dorm.JsonEncodeNoError(ginCtx.Request.Header),                //【请求】请求头
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

	c.gormRecord(data)
}

// GormDelete 删除
func (c *GinClient) GormDelete(ctx context.Context, hour int64) error {
	err := c.gormClient.GetDb().Table(c.gormConfig.tableName).Where("request_time < ?", gotime.Current().BeforeHour(hour).Format()).Delete(&ginPostgresqlLog{}).Error
	if err != nil {
		c.zapLog.WithTraceId(ctx).Sugar().Errorf("删除失败：%s", err)
	}
	return err
}
