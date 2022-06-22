package golog

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"go.dtapp.net/library/utils/dorm"
	"go.dtapp.net/library/utils/goip"
	"go.dtapp.net/library/utils/gorequest"
	"go.dtapp.net/library/utils/gotime"
	"go.dtapp.net/library/utils/goxml"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"gorm.io/gorm"
	"io/ioutil"
	"net"
	"os"
	"runtime"
	"strings"
)

// GinClient 框架
type GinClient struct {
	gormClient            *gorm.DB          // 驱动
	mongoCollectionClient *dorm.MongoClient // 驱动(温馨提示：需要已选择库和表)
	ipService             *goip.Client      // ip服务
	config                struct {
		logType   string // 日志类型
		tableName string // 表名
		insideIp  string // 内网ip
		hostname  string // 主机名
		goVersion string // go版本
	} // 配置
}

// NewGinClient 创建框架实例化
func NewGinClient(attrs ...*OperationAttr) (*GinClient, error) {

	c := &GinClient{}
	for _, attr := range attrs {
		if attr.gormClient != nil {
			c.gormClient = attr.gormClient
			c.config.logType = attr.logType
		}
		if attr.mongoCollectionClient != nil {
			c.mongoCollectionClient = attr.mongoCollectionClient
			c.config.logType = attr.logType
		}
		if attr.tableName != "" {
			c.config.tableName = attr.tableName
		}
		if attr.ipService != nil {
			c.ipService = attr.ipService
		}
	}

	switch c.config.logType {
	case logTypeGorm:

		if c.gormClient == nil {
			return nil, errors.New("驱动不能为空")
		}

		if c.config.tableName == "" {
			return nil, errors.New("表名不能为空")
		}

		err := c.gormClient.Table(c.config.tableName).AutoMigrate(&ApiPostgresqlLog{})
		if err != nil {
			return nil, errors.New("创建表失败：" + err.Error())
		}

	case logTypeMongo:

		if c.mongoCollectionClient.Db == nil {
			return nil, errors.New("驱动不能为空")
		}

		if c.config.tableName == "" {
			return nil, errors.New("表名不能为空")
		}

		c.mongoCollectionClient = c.mongoCollectionClient.Collection(c.config.tableName)

	default:
		return nil, errors.New("驱动为空")
	}

	hostname, _ := os.Hostname()

	c.config.hostname = hostname
	c.config.insideIp = goip.GetInsideIp()
	c.config.goVersion = strings.TrimPrefix(runtime.Version(), "go")

	return c, nil
}

// GormRecord 记录日志
func (c *GinClient) GormRecord(postgresqlLog GinPostgresqlLog) error {

	postgresqlLog.SystemHostName = c.config.hostname
	if postgresqlLog.SystemInsideIp == "" {
		postgresqlLog.SystemInsideIp = c.config.insideIp
	}
	postgresqlLog.GoVersion = c.config.goVersion

	return c.gormClient.Table(c.config.tableName).Create(&postgresqlLog).Error
}

// GormQuery 查询
func (c *GinClient) GormQuery() *gorm.DB {
	return c.gormClient.Table(c.config.tableName)
}

type bodyLogWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w bodyLogWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func (w bodyLogWriter) WriteString(s string) (int, error) {
	w.body.WriteString(s)
	return w.ResponseWriter.WriteString(s)
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
			if c.mongoCollectionClient != nil {
				host := ""
				if ginCtx.Request.TLS == nil {
					host = "http://" + ginCtx.Request.Host
				} else {
					host = "https://" + ginCtx.Request.Host
				}
				if len(jsonBody) > 0 {
					c.mongoRecord(GinMongoLog{
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
						ResponseData:      responseBody,                                                     //【返回】数据
						CostTime:          endTime - startTime,                                              //【系统】花费时间
					})
				} else {
					c.mongoRecord(GinMongoLog{
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
						ResponseData:      responseBody,                                                     //【返回】数据
						CostTime:          endTime - startTime,                                              //【系统】花费时间
					})
				}
			}
		}()
	}
}

// 记录日志
func (c *GinClient) mongoRecord(mongoLog GinMongoLog) error {

	mongoLog.SystemHostName = c.config.hostname
	if mongoLog.SystemInsideIp == "" {
		mongoLog.SystemInsideIp = c.config.insideIp
	}
	mongoLog.GoVersion = c.config.goVersion

	mongoLog.LogId = primitive.NewObjectID()

	_, err := c.mongoCollectionClient.InsertOne(mongoLog)
	return err
}

// MongoQuery 查询
func (c *GinClient) MongoQuery() *dorm.MongoClient {
	return c.mongoCollectionClient
}
