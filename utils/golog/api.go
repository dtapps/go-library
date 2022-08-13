package golog

import (
	"context"
	"errors"
	"github.com/dtapps/go-library/utils/dorm"
	"github.com/dtapps/go-library/utils/goip"
	"gorm.io/gorm"
	"os"
	"runtime"
)

// ApiClient 接口
type ApiClient struct {
	gormClient  *gorm.DB          // 数据库驱动
	mongoClient *dorm.MongoClient // 数据库驱动
	config      struct {
		logType        string // 日志类型
		tableName      string // 表名
		databaseName   string // 库名
		collectionName string // 表名
		insideIp       string // 内网ip
		hostname       string // 主机名
		goVersion      string // go版本
	} // 配置
}

// NewApiClient 创建接口实例化
// WithGormClient && WithTableName
// WithMongoCollectionClient && WithDatabaseName && WithCollectionName
func NewApiClient(attrs ...*OperationAttr) (*ApiClient, error) {

	c := &ApiClient{}
	for _, attr := range attrs {
		if attr.gormClient != nil {
			c.gormClient = attr.gormClient
			c.config.logType = attr.logType
		}
		if attr.mongoClient != nil {
			c.mongoClient = attr.mongoClient
			c.config.logType = attr.logType
		}
		if attr.tableName != "" {
			c.config.tableName = attr.tableName
		}
		if attr.databaseName != "" {
			c.config.databaseName = attr.databaseName
		}
		if attr.collectionName != "" {
			c.config.collectionName = attr.collectionName
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

		err := c.gormClient.Table(c.config.tableName).AutoMigrate(&apiPostgresqlLog{})
		if err != nil {
			return nil, errors.New("创建表失败：" + err.Error())
		}

	case logTypeMongo:

		if c.mongoClient.Db == nil {
			return nil, errors.New("没有设置驱动")
		}

		if c.config.databaseName == "" {
			return nil, errors.New("没有设置库名")
		}

		if c.config.collectionName == "" {
			return nil, errors.New("没有设置表名")
		}

	default:
		return nil, errors.New("驱动为空")
	}

	hostname, _ := os.Hostname()

	c.config.hostname = hostname
	c.config.insideIp = goip.GetInsideIp(context.Background())
	c.config.goVersion = runtime.Version()

	return c, nil
}
