package golog

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"go.dtapp.net/library/utils/dorm"
	"go.dtapp.net/library/utils/gorequest"
	"go.dtapp.net/library/utils/gotime"
	"go.dtapp.net/library/utils/goxml"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"io/ioutil"
	"net"
)

// 模型结构体
type ginMongoLog struct {
	LogId             primitive.ObjectID `json:"log_id,omitempty" bson:"_id,omitempty"`                              //【记录】编号
	TraceId           string             `json:"trace_id,omitempty" bson:"trace_id,omitempty"`                       //【系统】链编号
	RequestTime       dorm.BsonTime      `json:"request_time,omitempty" bson:"request_time,omitempty"`               //【请求】时间
	RequestUri        string             `json:"request_uri,omitempty" bson:"request_uri,omitempty"`                 //【请求】请求链接 域名+路径+参数
	RequestUrl        string             `json:"request_url,omitempty" bson:"request_url,omitempty"`                 //【请求】请求链接 域名+路径
	RequestApi        string             `json:"request_api,omitempty" bson:"request_api,omitempty"`                 //【请求】请求接口 路径
	RequestMethod     string             `json:"request_method,omitempty" bson:"request_method,omitempty"`           //【请求】请求方式
	RequestProto      string             `json:"request_proto,omitempty" bson:"request_proto,omitempty"`             //【请求】请求协议
	RequestUa         string             `json:"request_ua,omitempty" bson:"request_ua,omitempty"`                   //【请求】请求UA
	RequestReferer    string             `json:"request_referer,omitempty" bson:"request_referer,omitempty"`         //【请求】请求referer
	RequestBody       interface{}        `json:"request_body,omitempty" bson:"request_body,omitempty"`               //【请求】请求主体
	RequestUrlQuery   interface{}        `json:"request_url_query,omitempty" bson:"request_url_query,omitempty"`     //【请求】请求URL参数
	RequestIp         string             `json:"request_ip,omitempty" bson:"request_ip,omitempty"`                   //【请求】请求客户端Ip
	RequestIpCountry  string             `json:"request_ip_country,omitempty" bson:"request_ip_country,omitempty"`   //【请求】请求客户端城市
	RequestIpRegion   string             `json:"request_ip_region,omitempty" bson:"request_ip_region,omitempty"`     //【请求】请求客户端区域
	RequestIpProvince string             `json:"request_ip_province,omitempty" bson:"request_ip_province,omitempty"` //【请求】请求客户端省份
	RequestIpCity     string             `json:"request_ip_city,omitempty" bson:"request_ip_city,omitempty"`         //【请求】请求客户端城市
	RequestIpIsp      string             `json:"request_ip_isp,omitempty" bson:"request_ip_isp,omitempty"`           //【请求】请求客户端运营商
	RequestHeader     interface{}        `json:"request_header,omitempty" bson:"request_header,omitempty"`           //【请求】请求头
	ResponseTime      dorm.BsonTime      `json:"response_time,omitempty" bson:"response_time,omitempty"`             //【返回】时间
	ResponseCode      int                `json:"response_code,omitempty" bson:"response_code,omitempty"`             //【返回】状态码
	ResponseMsg       string             `json:"response_msg,omitempty" bson:"response_msg,omitempty"`               //【返回】描述
	ResponseData      interface{}        `json:"response_data,omitempty" bson:"response_data,omitempty"`             //【返回】数据
	CostTime          int64              `json:"cost_time,omitempty" bson:"cost_time,omitempty"`                     //【系统】花费时间
	SystemHostName    string             `json:"system_host_name,omitempty" bson:"system_host_name,omitempty"`       //【系统】主机名
	SystemInsideIp    string             `json:"system_inside_ip,omitempty" bson:"system_inside_ip,omitempty"`       //【系统】内网ip
	GoVersion         string             `json:"go_version,omitempty" bson:"go_version,omitempty"`                   //【程序】Go版本
}

// 记录日志
func (c *GinClient) mongoRecord(mongoLog ginMongoLog) error {

	mongoLog.SystemHostName = c.config.hostname
	if mongoLog.SystemInsideIp == "" {
		mongoLog.SystemInsideIp = c.config.insideIp
	}
	mongoLog.GoVersion = c.config.goVersion

	mongoLog.LogId = primitive.NewObjectID()

	_, err := c.mongoClient.Database(c.config.databaseName).Collection(c.config.collectionName).InsertOne(mongoLog)

	return err
}

// MongoQuery 查询
func (c *GinClient) MongoQuery() *dorm.MongoClient {
	return c.mongoClient.Database(c.config.databaseName).Collection(c.config.collectionName)
}

// MongoMiddleware 中间件
func (c *GinClient) MongoMiddleware() gin.HandlerFunc {
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
			if c.mongoClient != nil {
				host := ""
				if ginCtx.Request.TLS == nil {
					host = "http://" + ginCtx.Request.Host
				} else {
					host = "https://" + ginCtx.Request.Host
				}
				if len(jsonBody) > 0 {
					c.mongoRecord(ginMongoLog{
						TraceId:           ginCtx.MustGet("trace_id").(string),                              //【系统】链编号
						RequestTime:       dorm.BsonTime(requestTime),                                       //【请求】时间
						RequestUri:        host + ginCtx.Request.RequestURI,                                 //【请求】请求链接
						RequestUrl:        ginCtx.Request.RequestURI,                                        //【请求】请求链接
						RequestApi:        gorequest.UriFilterExcludeQueryString(ginCtx.Request.RequestURI), //【请求】请求接口
						RequestMethod:     ginCtx.Request.Method,                                            //【请求】请求方式
						RequestProto:      ginCtx.Request.Proto,                                             //【请求】请求协议
						RequestUa:         ginCtx.Request.UserAgent(),                                       //【请求】请求UA
						RequestReferer:    ginCtx.Request.Referer(),                                         //【请求】请求referer
						RequestBody:       jsonBody,                                                         //【请求】请求主体
						RequestUrlQuery:   ginCtx.Request.URL.Query(),                                       //【请求】请求URL参数
						RequestIp:         ginCtx.ClientIP(),                                                //【请求】请求客户端Ip
						RequestIpCountry:  requestClientIpCountry,                                           //【请求】请求客户端城市
						RequestIpRegion:   requestClientIpRegion,                                            //【请求】请求客户端区域
						RequestIpProvince: requestClientIpProvince,                                          //【请求】请求客户端省份
						RequestIpCity:     requestClientIpCity,                                              //【请求】请求客户端城市
						RequestIpIsp:      requestClientIpIsp,                                               //【请求】请求客户端运营商
						RequestHeader:     ginCtx.Request.Header,                                            //【请求】请求头
						ResponseTime:      dorm.BsonTime(gotime.Current().Time),                             //【返回】时间
						ResponseCode:      responseCode,                                                     //【返回】状态码
						ResponseData:      c.jsonUnmarshal(responseBody),                                    //【返回】数据
						CostTime:          endTime - startTime,                                              //【系统】花费时间
					})
				} else {
					c.mongoRecord(ginMongoLog{
						TraceId:           ginCtx.MustGet("trace_id").(string),                              //【系统】链编号
						RequestTime:       dorm.BsonTime(requestTime),                                       //【请求】时间
						RequestUri:        host + ginCtx.Request.RequestURI,                                 //【请求】请求链接
						RequestUrl:        ginCtx.Request.RequestURI,                                        //【请求】请求链接
						RequestApi:        gorequest.UriFilterExcludeQueryString(ginCtx.Request.RequestURI), //【请求】请求接口
						RequestMethod:     ginCtx.Request.Method,                                            //【请求】请求方式
						RequestProto:      ginCtx.Request.Proto,                                             //【请求】请求协议
						RequestUa:         ginCtx.Request.UserAgent(),                                       //【请求】请求UA
						RequestReferer:    ginCtx.Request.Referer(),                                         //【请求】请求referer
						RequestBody:       xmlBody,                                                          //【请求】请求主体
						RequestUrlQuery:   ginCtx.Request.URL.Query(),                                       //【请求】请求URL参数
						RequestIp:         ginCtx.ClientIP(),                                                //【请求】请求客户端Ip
						RequestIpCountry:  requestClientIpCountry,                                           //【请求】请求客户端城市
						RequestIpRegion:   requestClientIpRegion,                                            //【请求】请求客户端区域
						RequestIpProvince: requestClientIpProvince,                                          //【请求】请求客户端省份
						RequestIpCity:     requestClientIpCity,                                              //【请求】请求客户端城市
						RequestIpIsp:      requestClientIpIsp,                                               //【请求】请求客户端运营商
						RequestHeader:     ginCtx.Request.Header,                                            //【请求】请求头
						ResponseTime:      dorm.BsonTime(gotime.Current().Time),                             //【返回】时间
						ResponseCode:      responseCode,                                                     //【返回】状态码
						ResponseData:      c.jsonUnmarshal(responseBody),                                    //【返回】数据
						CostTime:          endTime - startTime,                                              //【系统】花费时间
					})
				}
			}
		}()
	}
}
