package golog

import (
	"context"
	"fmt"
	"github.com/dtapps/go-library/utils/dorm"
	"github.com/dtapps/go-library/utils/goip"
	"github.com/dtapps/go-library/utils/gorequest"
	"github.com/dtapps/go-library/utils/gotime"
	"github.com/dtapps/go-library/utils/gotrace_id"
	"github.com/dtapps/go-library/utils/gourl"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"time"
)

// 模型
type ginPostgresqlLogCustom struct {
	LogID              int64     `gorm:"primaryKey;comment:【日志】编号" json:"log_id,omitempty"`                     //【日志】编号
	LogTime            time.Time `gorm:"comment:【日志】时间" json:"log_time,omitempty"`                              //【记日志录】时间
	TraceId            string    `gorm:"index;comment:【系统】跟踪编号" json:"trace_id,omitempty"`                      //【系统】跟踪编号
	RequestUri         string    `gorm:"comment:【请求】请求链接 域名+路径+参数" json:"request_uri,omitempty"`                //【请求】请求链接 域名+路径+参数
	RequestUrl         string    `gorm:"comment:【请求】请求链接 域名+路径" json:"request_url,omitempty"`                   //【请求】请求链接 域名+路径
	RequestApi         string    `gorm:"index;comment:【请求】请求接口 路径" json:"request_api,omitempty"`                //【请求】请求接口 路径
	RequestMethod      string    `gorm:"index;comment:【请求】请求方式" json:"request_method,omitempty"`                //【请求】请求方式
	RequestProto       string    `gorm:"comment:【请求】请求协议" json:"request_proto,omitempty"`                       //【请求】请求协议
	RequestUa          string    `gorm:"comment:【请求】请求UA" json:"request_ua,omitempty"`                          //【请求】请求UA
	RequestReferer     string    `gorm:"comment:【请求】请求referer" json:"request_referer,omitempty"`                //【请求】请求referer
	RequestUrlQuery    string    `gorm:"comment:【请求】请求URL参数" json:"request_url_query,omitempty"`                //【请求】请求URL参数
	RequestHeader      string    `gorm:"comment:【请求】请求头" json:"request_header,omitempty"`                       //【请求】请求头
	RequestIp          string    `gorm:"default:0.0.0.0;index;comment:【请求】请求客户端Ip" json:"request_ip,omitempty"` //【请求】请求客户端Ip
	RequestIpCountry   string    `gorm:"index;comment:【请求】请求客户端城市" json:"request_ip_country,omitempty"`         //【请求】请求客户端城市
	RequestIpProvince  string    `gorm:"index;comment:【请求】请求客户端省份" json:"request_ip_province,omitempty"`        //【请求】请求客户端省份
	RequestIpCity      string    `gorm:"index;comment:【请求】请求客户端城市" json:"request_ip_city,omitempty"`            //【请求】请求客户端城市
	RequestIpIsp       string    `gorm:"index;comment:【请求】请求客户端运营商" json:"request_ip_isp,omitempty"`            //【请求】请求客户端运营商
	RequestIpLongitude float64   `gorm:"index;comment:【请求】请求客户端经度" json:"request_ip_longitude,omitempty"`       //【请求】请求客户端经度
	RequestIpLatitude  float64   `gorm:"index;comment:【请求】请求客户端纬度" json:"request_ip_latitude,omitempty"`        //【请求】请求客户端纬度
	SystemHostName     string    `gorm:"index;comment:【系统】主机名" json:"system_host_name,omitempty"`               //【系统】主机名
	SystemInsideIp     string    `gorm:"default:0.0.0.0;comment:【系统】内网ip" json:"system_inside_ip,omitempty"`    //【系统】内网ip
	SystemOs           string    `gorm:"index;comment:【系统】系统类型" json:"system_os,omitempty"`                     //【系统】系统类型
	SystemArch         string    `gorm:"index;comment:【系统】系统架构" json:"system_arch,omitempty"`                   //【系统】系统架构
	GoVersion          string    `gorm:"comment:【程序】Go版本" json:"go_version,omitempty"`                          //【程序】Go版本
	SdkVersion         string    `gorm:"comment:【程序】Sdk版本" json:"sdk_version,omitempty"`                        //【程序】Sdk版本
	CustomId           string    `gorm:"index;comment:【日志】自定义编号" json:"custom_id,omitempty"`                    //【日志】自定义编号
	CustomType         string    `gorm:"index;comment:【日志】自定义类型" json:"custom_type,omitempty"`                  //【日志】自定义类型
	CustomContent      string    `gorm:"comment:【日志】自定义内容" json:"custom_content,omitempty"`                     //【日志】自定义内容
}

// 创建模型
func (c *GinCustomClient) autoMigrate(ctx context.Context) error {
	return c.gormClient.GetDb().Table(c.gormConfig.tableName).AutoMigrate(&ginPostgresqlLogCustom{})
}

type GinCustomClientGinRecordOperation struct {
	gormClient *dorm.GormClient        // 数据库驱动
	ipService  *goip.Client            // ip服务
	tableName  string                  // 表名
	data       *ginPostgresqlLogCustom // 数据
}

// GinRecord 记录日志
func (c *GinCustomClient) GinRecord(ginCtx *gin.Context) *GinCustomClientGinRecordOperation {
	operation := &GinCustomClientGinRecordOperation{
		gormClient: c.gormClient,
		ipService:  c.ipService,
		tableName:  c.gormConfig.tableName,
	}
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
	operation.data.SystemHostName = c.config.systemHostname                                  //【系统】主机名
	operation.data.SystemInsideIp = c.config.systemInsideIp                                  //【系统】内网ip
	operation.data.SystemOs = c.config.systemOs                                              //【系统】系统类型
	operation.data.SystemArch = c.config.systemKernel                                        //【系统】系统架构
	operation.data.GoVersion = c.config.goVersion                                            //【程序】Go版本
	operation.data.SdkVersion = c.config.sdkVersion                                          //【程序】Sdk版本
	return operation
}

func (o *GinCustomClientGinRecordOperation) CustomInfo(customId any, customType any, customContent any) *GinCustomClientGinRecordOperation {
	o.data.CustomId = fmt.Sprintf("%s", customId)           //【日志】自定义编号
	o.data.CustomType = fmt.Sprintf("%s", customType)       //【日志】自定义类型
	o.data.CustomContent = fmt.Sprintf("%s", customContent) //【日志】自定义内容
	return o
}

func (o *GinCustomClientGinRecordOperation) CreateData() error {
	err := o.gormClient.GetDb().Table(o.tableName).Create(&o.data).Error
	if o.data.LogID != 0 {
		go func() {
			ginCustomClientRecordUpdateIpInfo(o.ipService, o.gormClient.GetDb(), o.tableName, o.data.LogID, o.data.RequestIp)
		}()
	}
	return err
}

func (o *GinCustomClientGinRecordOperation) CreateDataNoError() {
	o.gormClient.GetDb().Table(o.tableName).Create(&o.data)
	if o.data.LogID != 0 {
		go func() {
			ginCustomClientRecordUpdateIpInfo(o.ipService, o.gormClient.GetDb(), o.tableName, o.data.LogID, o.data.RequestIp)
		}()
	}
}

func ginCustomClientRecordUpdateIpInfo(ipService *goip.Client, pdTx *gorm.DB, tableName string, logId int64, requestIp string) {
	ipInfo := ipService.Analyse(requestIp)
	pdTx.Table(tableName).Where("log_id = ?", logId).
		Select(
			"request_ip_country",
			"request_ip_province",
			"request_ip_city",
			"request_ip_isp",
			"request_ip_longitude",
			"request_ip_longitude",
		).
		Updates(ginPostgresqlLogCustom{
			RequestIpCountry:   ipInfo.Country,           //【请求】请求客户端城市
			RequestIpProvince:  ipInfo.Province,          //【请求】请求客户端省份
			RequestIpCity:      ipInfo.City,              //【请求】请求客户端城市
			RequestIpIsp:       ipInfo.Isp,               //【请求】请求客户端运营商
			RequestIpLongitude: ipInfo.LocationLatitude,  //【请求】请求客户端纬度
			RequestIpLatitude:  ipInfo.LocationLongitude, //【请求】请求客户端经度
		})
}
