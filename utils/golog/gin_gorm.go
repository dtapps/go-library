package golog

import (
	"bytes"
	"encoding/json"
	"github.com/dtapps/go-library/utils/gojson"
	"github.com/dtapps/go-library/utils/gotime"
	"github.com/dtapps/go-library/utils/gotrace_id"
	"github.com/dtapps/go-library/utils/gourl"
	"github.com/dtapps/go-library/utils/goxml"
	"github.com/gin-gonic/gin"
	"gorm.io/datatypes"
	"gorm.io/gorm"
	"io/ioutil"
	"log"
	"net"
	"time"
)

// 模型结构体
type ginPostgresqlLog struct {
	LogId             uint           `gorm:"primaryKey;comment:【记录】编号" json:"log_id,omitempty"`                 //【记录】编号
	TraceId           string         `gorm:"index;comment:【系统】跟踪编号" json:"trace_id,omitempty"`                  //【系统】跟踪编号
	RequestTime       time.Time      `gorm:"index;comment:【请求】时间" json:"request_time,omitempty"`                //【请求】时间
	RequestUri        string         `gorm:"comment:【请求】请求链接 域名+路径+参数" json:"request_uri,omitempty"`            //【请求】请求链接 域名+路径+参数
	RequestUrl        string         `gorm:"comment:【请求】请求链接 域名+路径" json:"request_url,omitempty"`               //【请求】请求链接 域名+路径
	RequestApi        string         `gorm:"index;comment:【请求】请求接口 路径" json:"request_api,omitempty"`            //【请求】请求接口 路径
	RequestMethod     string         `gorm:"index;comment:【请求】请求方式" json:"request_method,omitempty"`            //【请求】请求方式
	RequestProto      string         `gorm:"comment:【请求】请求协议" json:"request_proto,omitempty"`                   //【请求】请求协议
	RequestUa         string         `gorm:"comment:【请求】请求UA" json:"request_ua,omitempty"`                      //【请求】请求UA
	RequestReferer    string         `gorm:"comment:【请求】请求referer" json:"request_referer,omitempty"`            //【请求】请求referer
	RequestBody       datatypes.JSON `gorm:"type:jsonb;comment:【请求】请求主体" json:"request_body,omitempty"`         //【请求】请求主体
	RequestUrlQuery   datatypes.JSON `gorm:"type:jsonb;comment:【请求】请求URL参数" json:"request_url_query,omitempty"` //【请求】请求URL参数
	RequestIp         string         `gorm:"index;comment:【请求】请求客户端Ip" json:"request_ip,omitempty"`             //【请求】请求客户端Ip
	RequestIpCountry  string         `gorm:"index;comment:【请求】请求客户端城市" json:"request_ip_country,omitempty"`     //【请求】请求客户端城市
	RequestIpRegion   string         `gorm:"index;comment:【请求】请求客户端区域" json:"request_ip_region,omitempty"`      //【请求】请求客户端区域
	RequestIpProvince string         `gorm:"index;comment:【请求】请求客户端省份" json:"request_ip_province,omitempty"`    //【请求】请求客户端省份
	RequestIpCity     string         `gorm:"index;comment:【请求】请求客户端城市" json:"request_ip_city,omitempty"`        //【请求】请求客户端城市
	RequestIpIsp      string         `gorm:"index;comment:【请求】请求客户端运营商" json:"request_ip_isp,omitempty"`        //【请求】请求客户端运营商
	RequestHeader     datatypes.JSON `gorm:"type:jsonb;comment:【请求】请求头" json:"request_header,omitempty"`        //【请求】请求头
	ResponseTime      time.Time      `gorm:"index;comment:【返回】时间" json:"response_time,omitempty"`               //【返回】时间
	ResponseCode      int            `gorm:"index;comment:【返回】状态码" json:"response_code,omitempty"`              //【返回】状态码
	ResponseMsg       string         `gorm:"comment:【返回】描述" json:"response_msg,omitempty"`                      //【返回】描述
	ResponseData      datatypes.JSON `gorm:"type:jsonb;comment:【返回】数据" json:"response_data,omitempty"`          //【返回】数据
	CostTime          int64          `gorm:"comment:【系统】花费时间" json:"cost_time,omitempty"`                       //【系统】花费时间
	SystemHostName    string         `gorm:"index;comment:【系统】主机名" json:"system_host_name,omitempty"`           //【系统】主机名
	SystemInsideIp    string         `gorm:"index;comment:【系统】内网ip" json:"system_inside_ip,omitempty"`          //【系统】内网ip
	GoVersion         string         `gorm:"index;comment:【程序】Go版本" json:"go_version,omitempty"`                //【程序】Go版本
	SdkVersion        string         `gorm:"index;comment:【程序】Sdk版本" json:"sdk_version,omitempty"`              //【程序】Sdk版本
}

// gormRecord 记录日志
func (c *GinClient) gormRecord(postgresqlLog ginPostgresqlLog) error {

	postgresqlLog.SystemHostName = c.config.hostname
	if postgresqlLog.SystemInsideIp == "" {
		postgresqlLog.SystemInsideIp = c.config.insideIp
	}
	postgresqlLog.GoVersion = c.config.goVersion

	postgresqlLog.SdkVersion = Version

	return c.gormClient.Table(c.config.tableName).Create(&postgresqlLog).Error
}

// GormQuery 查询
func (c *GinClient) GormQuery() *gorm.DB {
	return c.gormClient.Table(c.config.tableName)
}

// GormMiddleware 中间件
func (c *GinClient) GormMiddleware() gin.HandlerFunc {
	return func(ginCtx *gin.Context) {

		// 开始时间
		startTime := gotime.Current().TimestampWithMillisecond()
		requestTime := gotime.Current().Time

		// 获取
		data, _ := ioutil.ReadAll(ginCtx.Request.Body)

		// 复用
		ginCtx.Request.Body = ioutil.NopCloser(bytes.NewBuffer(data))

		blw := &bodyLogWriter{body: bytes.NewBufferString(""), ResponseWriter: ginCtx.Writer}
		ginCtx.Writer = blw

		// 处理请求
		ginCtx.Next()

		// 响应
		responseCode := ginCtx.Writer.Status()
		responseBody := blw.body.String()

		//结束时间
		endTime := gotime.Current().TimestampWithMillisecond()

		go func() {

			// 解析请求内容
			var xmlBody map[string]string
			var jsonBody map[string]interface{}
			_ = json.Unmarshal(data, &jsonBody)
			if len(jsonBody) <= 0 {
				xmlBody = goxml.XmlDecode(string(data))
			}

			requestClientIpCountry, requestClientIpRegion, requestClientIpProvince, requestClientIpCity, requestClientIpIsp := "", "", "", "", ""
			if c.ipService != nil {
				if net.ParseIP(ginCtx.ClientIP()).To4() != nil {
					// 判断是不是IPV4
					_, info := c.ipService.Ipv4(ginCtx.ClientIP())
					requestClientIpCountry = info.Country
					requestClientIpRegion = info.Region
					requestClientIpProvince = info.Province
					requestClientIpCity = info.City
					requestClientIpIsp = info.ISP
				} else if net.ParseIP(ginCtx.ClientIP()).To16() != nil {
					// 判断是不是IPV6
					info := c.ipService.Ipv6(ginCtx.ClientIP())
					requestClientIpCountry = info.Country
					requestClientIpProvince = info.Province
					requestClientIpCity = info.City
				}
			}

			// 记录
			if c.gormClient != nil {
				host := ""
				if ginCtx.Request.TLS == nil {
					host = "http://" + ginCtx.Request.Host
				} else {
					host = "https://" + ginCtx.Request.Host
				}
				if len(jsonBody) > 0 {
					err := c.gormRecord(ginPostgresqlLog{
						TraceId:           gotrace_id.GetGinTraceId(ginCtx),                                     //【系统】链编号
						RequestTime:       requestTime,                                                          //【请求】时间
						RequestUri:        host + ginCtx.Request.RequestURI,                                     //【请求】请求链接
						RequestUrl:        ginCtx.Request.RequestURI,                                            //【请求】请求链接
						RequestApi:        gourl.UriFilterExcludeQueryString(ginCtx.Request.RequestURI),         //【请求】请求接口
						RequestMethod:     ginCtx.Request.Method,                                                //【请求】请求方式
						RequestProto:      ginCtx.Request.Proto,                                                 //【请求】请求协议
						RequestUa:         ginCtx.Request.UserAgent(),                                           //【请求】请求UA
						RequestReferer:    ginCtx.Request.Referer(),                                             //【请求】请求referer
						RequestBody:       datatypes.JSON(gojson.JsonEncodeNoError(jsonBody)),                   //【请求】请求主体
						RequestUrlQuery:   datatypes.JSON(gojson.JsonEncodeNoError(ginCtx.Request.URL.Query())), //【请求】请求URL参数
						RequestIp:         ginCtx.ClientIP(),                                                    //【请求】请求客户端Ip
						RequestIpCountry:  requestClientIpCountry,                                               //【请求】请求客户端城市
						RequestIpRegion:   requestClientIpRegion,                                                //【请求】请求客户端区域
						RequestIpProvince: requestClientIpProvince,                                              //【请求】请求客户端省份
						RequestIpCity:     requestClientIpCity,                                                  //【请求】请求客户端城市
						RequestIpIsp:      requestClientIpIsp,                                                   //【请求】请求客户端运营商
						RequestHeader:     datatypes.JSON(gojson.JsonEncodeNoError(ginCtx.Request.Header)),      //【请求】请求头
						ResponseTime:      gotime.Current().Time,                                                //【返回】时间
						ResponseCode:      responseCode,                                                         //【返回】状态码
						ResponseData:      datatypes.JSON(responseBody),                                         //【返回】数据
						CostTime:          endTime - startTime,                                                  //【系统】花费时间
					})
					if err != nil {
						log.Println("log.gormRecord：", err.Error())
					}
				} else {
					err := c.gormRecord(ginPostgresqlLog{
						TraceId:           gotrace_id.GetGinTraceId(ginCtx),                                     //【系统】链编号
						RequestTime:       requestTime,                                                          //【请求】时间
						RequestUri:        host + ginCtx.Request.RequestURI,                                     //【请求】请求链接
						RequestUrl:        ginCtx.Request.RequestURI,                                            //【请求】请求链接
						RequestApi:        gourl.UriFilterExcludeQueryString(ginCtx.Request.RequestURI),         //【请求】请求接口
						RequestMethod:     ginCtx.Request.Method,                                                //【请求】请求方式
						RequestProto:      ginCtx.Request.Proto,                                                 //【请求】请求协议
						RequestUa:         ginCtx.Request.UserAgent(),                                           //【请求】请求UA
						RequestReferer:    ginCtx.Request.Referer(),                                             //【请求】请求referer
						RequestBody:       datatypes.JSON(gojson.JsonEncodeNoError(xmlBody)),                    //【请求】请求主体
						RequestUrlQuery:   datatypes.JSON(gojson.JsonEncodeNoError(ginCtx.Request.URL.Query())), //【请求】请求URL参数
						RequestIp:         ginCtx.ClientIP(),                                                    //【请求】请求客户端Ip
						RequestIpCountry:  requestClientIpCountry,                                               //【请求】请求客户端城市
						RequestIpRegion:   requestClientIpRegion,                                                //【请求】请求客户端区域
						RequestIpProvince: requestClientIpProvince,                                              //【请求】请求客户端省份
						RequestIpCity:     requestClientIpCity,                                                  //【请求】请求客户端城市
						RequestIpIsp:      requestClientIpIsp,                                                   //【请求】请求客户端运营商
						RequestHeader:     datatypes.JSON(gojson.JsonEncodeNoError(ginCtx.Request.Header)),      //【请求】请求头
						ResponseTime:      gotime.Current().Time,                                                //【返回】时间
						ResponseCode:      responseCode,                                                         //【返回】状态码
						ResponseData:      datatypes.JSON(responseBody),                                         //【返回】数据
						CostTime:          endTime - startTime,                                                  //【系统】花费时间
					})
					if err != nil {
						log.Println("log.gormRecord：", err.Error())
					}
				}
			}
		}()
	}
}
