package golog

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"github.com/dtapps/go-library/utils/dorm"
	"github.com/dtapps/go-library/utils/goip"
	"github.com/dtapps/go-library/utils/gorequest"
	"github.com/dtapps/go-library/utils/gotime"
	"github.com/dtapps/go-library/utils/gotrace_id"
	"github.com/gin-gonic/gin"
	"io/ioutil"
)

// GinClientFun *GinClient 驱动
type GinClientFun func() *GinClient

// GinClient 框架
type GinClient struct {
	gormClient  *dorm.GormClient  // 数据库驱动
	mongoClient *dorm.MongoClient // 数据库驱动
	ipService   *goip.Client      // ip服务
	zapLog      *ZapLog           // 日志服务
	config      struct {
		systemHostname      string  // 主机名
		systemOs            string  // 系统类型
		systemVersion       string  // 系统版本
		systemKernel        string  // 系统内核
		systemKernelVersion string  // 系统内核版本
		systemBootTime      uint64  // 系统开机时间
		cpuCores            int     // CPU核数
		cpuModelName        string  // CPU型号名称
		cpuMhz              float64 // CPU兆赫
		systemInsideIp      string  // 内网ip
		systemOutsideIp     string  // 外网ip
		goVersion           string  // go版本
		sdkVersion          string  // sdk版本
		mongoVersion        string  // mongo版本
		mongoSdkVersion     string  // mongo sdk版本
	}
	gormConfig struct {
		stats     bool   // 状态
		tableName string // 表名
	}
	mongoConfig struct {
		stats          bool   // 状态
		databaseName   string // 库名
		collectionName string // 表名
	}
}

// GinClientConfig 框架实例配置
type GinClientConfig struct {
	IpService      *goip.Client                  // ip服务
	GormClientFun  dorm.GormClientTableFun       // 日志配置
	MongoClientFun dorm.MongoClientCollectionFun // 日志配置
	ZapLog         *ZapLog                       // 日志服务
	CurrentIp      string                        // 当前ip
}

// NewGinClient 创建框架实例化
func NewGinClient(config *GinClientConfig) (*GinClient, error) {

	var ctx = context.Background()

	c := &GinClient{}

	c.zapLog = config.ZapLog

	if config.CurrentIp != "" && config.CurrentIp != "0.0.0.0" {
		c.config.systemOutsideIp = config.CurrentIp
	}
	c.config.systemOutsideIp = goip.IsIp(c.config.systemOutsideIp)
	if c.config.systemOutsideIp == "" {
		return nil, currentIpNoConfig
	}

	c.ipService = config.IpService

	// 配置信息
	c.setConfig(ctx)

	gormClient, gormTableName := config.GormClientFun()
	mongoClient, mongoDatabaseName, mongoCollectionName := config.MongoClientFun()

	if (gormClient == nil || gormClient.GetDb() == nil) || (mongoClient == nil || mongoClient.GetDb() == nil) {
		return nil, dbClientFunNoConfig
	}

	// 配置关系数据库
	if gormClient != nil || gormClient.GetDb() != nil {

		c.gormClient = gormClient

		if gormTableName == "" {
			return nil, errors.New("没有设置表名")
		} else {
			c.gormConfig.tableName = gormTableName
		}

		c.gormAutoMigrate(ctx)

		c.gormConfig.stats = true
	}

	// 配置非关系数据库
	if mongoClient != nil || mongoClient.GetDb() != nil {

		c.mongoClient = mongoClient

		if mongoDatabaseName == "" {
			return nil, errors.New("没有设置库名")
		} else {
			c.mongoConfig.databaseName = mongoDatabaseName
		}

		if mongoCollectionName == "" {
			return nil, errors.New("没有设置表名")
		} else {
			c.mongoConfig.collectionName = mongoCollectionName
		}

		// 创建时间序列集合
		//c.mongoCreateCollection(ctx)

		// 创建索引
		c.mongoCreateIndexes(ctx)

		c.mongoConfig.stats = true
	}

	return c, nil
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

func (c *GinClient) jsonUnmarshal(data string) (result interface{}) {
	_ = json.Unmarshal([]byte(data), &result)
	return
}

// Middleware 中间件
func (c *GinClient) Middleware() gin.HandlerFunc {
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

			var dataJson = true

			// 解析请求内容
			var jsonBody map[string]interface{}

			// 判断是否有内容
			if len(data) > 0 {
				err := json.Unmarshal(data, &jsonBody)
				if err != nil {
					dataJson = false
				}
			}

			clientIp := gorequest.ClientIp(ginCtx.Request)
			var info = goip.AnalyseResult{}

			if c.ipService != nil {
				info = c.ipService.Analyse(clientIp)
			}

			var traceId = gotrace_id.GetGinTraceId(ginCtx)

			// 记录
			if c.gormConfig.stats {
				if dataJson {
					c.gormRecordJson(ginCtx, traceId, requestTime, data, responseCode, responseBody, startTime, endTime, info)
				} else {
					c.gormRecordXml(ginCtx, traceId, requestTime, data, responseCode, responseBody, startTime, endTime, info)
				}
			}
			if c.mongoConfig.stats {
				if dataJson {
					c.mongoRecordJson(ginCtx, traceId, requestTime, data, responseCode, responseBody, startTime, endTime, info)
				} else {
					c.mongoRecordXml(ginCtx, traceId, requestTime, data, responseCode, responseBody, startTime, endTime, info)
				}
			}
		}()
	}
}
