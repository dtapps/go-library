package golog

import (
	"context"
	"errors"
	"go.dtapp.net/library/utils/dorm"
	"go.dtapp.net/library/utils/goip"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"gorm.io/gorm"
	"os"
	"runtime"
	"strings"
)

// GinClient 框架
type GinClient struct {
	gormClient            *gorm.DB          // 驱动
	mongoCollectionClient *dorm.MongoClient // 驱动(温馨提示：需要已选择库)
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
		c.gormClient = attr.gormClient
		c.mongoCollectionClient = attr.mongoCollectionClient
		c.config.logType = attr.logType
		c.config.tableName = attr.tableName
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

// MongoRecord 记录日志
func (c *GinClient) MongoRecord(mongoLog GinMongoLog) error {

	mongoLog.SystemHostName = c.config.hostname
	if mongoLog.SystemInsideIp == "" {
		mongoLog.SystemInsideIp = c.config.insideIp
	}
	mongoLog.GoVersion = c.config.goVersion

	mongoLog.LogId = primitive.NewObjectID()

	_, err := c.mongoCollectionClient.InsertOne(context.Background(), mongoLog)
	return err
}

// MongoQuery 查询
func (c *GinClient) MongoQuery() *dorm.MongoClient {
	return c.mongoCollectionClient
}
