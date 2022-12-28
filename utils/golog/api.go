package golog

import (
	"context"
	"errors"
	"github.com/dtapps/go-library/utils/dorm"
	"github.com/dtapps/go-library/utils/goip"
	"github.com/dtapps/go-library/utils/gorequest"
)

// ApiClientFun *ApiClient 驱动
type ApiClientFun func() *ApiClient

// ApiClient 接口
type ApiClient struct {
	gormClient  *dorm.GormClient  // 数据库驱动
	mongoClient *dorm.MongoClient // 数据库驱动
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

// ApiClientConfig 接口实例配置
type ApiClientConfig struct {
	GormClientFun  dorm.GormClientTableFun       // 日志配置
	MongoClientFun dorm.MongoClientCollectionFun // 日志配置
	ZapLog         *ZapLog                       // 日志服务
	CurrentIp      string                        // 当前ip
}

// NewApiClient 创建接口实例化
func NewApiClient(config *ApiClientConfig) (*ApiClient, error) {

	var ctx = context.Background()

	c := &ApiClient{}

	c.zapLog = config.ZapLog

	if config.CurrentIp != "" && config.CurrentIp != "0.0.0.0" {
		c.config.systemOutsideIp = config.CurrentIp
	}
	c.config.systemOutsideIp = goip.IsIp(c.config.systemOutsideIp)
	if c.config.systemOutsideIp == "" {
		return nil, currentIpNoConfig
	}

	// 配置信息
	c.setConfig(ctx)

	gormClient, gormTableName := config.GormClientFun()
	//mongoClient, mongoDatabaseName, mongoCollectionName := config.MongoClientFun()

	//if (gormClient == nil || gormClient.GetDb() == nil) || (mongoClient == nil || mongoClient.GetDb() == nil) {
	//	return nil, dbClientFunNoConfig
	//}
	if gormClient == nil || gormClient.GetDb() == nil {
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

		// 创建模型
		c.gormAutoMigrate(ctx)

		c.gormConfig.stats = true
	}

	// 配置非关系数据库
	//if mongoClient != nil || mongoClient.GetDb() != nil {
	//
	//	c.mongoClient = mongoClient
	//
	//	if mongoDatabaseName == "" {
	//		return nil, errors.New("没有设置库名")
	//	} else {
	//		c.mongoConfig.databaseName = mongoDatabaseName
	//	}
	//
	//	if mongoCollectionName == "" {
	//		return nil, errors.New("没有设置表名")
	//	} else {
	//		c.mongoConfig.collectionName = mongoCollectionName
	//	}
	//
	//	// 创建时间序列集合
	//	c.mongoCreateCollection(ctx)
	//
	//	// 创建索引
	//	c.mongoCreateIndexes(ctx)
	//
	//	c.mongoConfig.stats = true
	//}

	return c, nil
}

// Middleware 中间件
func (c *ApiClient) Middleware(ctx context.Context, request gorequest.Response, sdkVersion string) {
	if c.gormConfig.stats {
		c.gormMiddleware(ctx, request, sdkVersion)
	}
	//if c.mongoConfig.stats {
	//	c.mongoMiddleware(ctx, request, sdkVersion)
	//}
}

// MiddlewareXml 中间件
func (c *ApiClient) MiddlewareXml(ctx context.Context, request gorequest.Response, sdkVersion string) {
	if c.gormConfig.stats {
		c.gormMiddlewareXml(ctx, request, sdkVersion)
	}
	//if c.mongoConfig.stats {
	//	c.mongoMiddlewareXml(ctx, request, sdkVersion)
	//}
}

// MiddlewareCustom 中间件
func (c *ApiClient) MiddlewareCustom(ctx context.Context, api string, request gorequest.Response, sdkVersion string) {
	if c.gormConfig.stats {
		c.gormMiddlewareCustom(ctx, api, request, sdkVersion)
	}
	//if c.mongoConfig.stats {
	//	c.mongoMiddlewareCustom(ctx, api, request, sdkVersion)
	//}
}
