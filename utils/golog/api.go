package golog

import (
	"errors"
	"go.dtapp.net/library/utils/dorm"
	"go.dtapp.net/library/utils/goip"
	"gorm.io/gorm"
	"os"
	"runtime"
	"strings"
)

// ApiClient 接口
type ApiClient struct {
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

// NewApiClient 创建接口实例化
// WithGormClient && WithTableName
// WithMongoCollectionClient && WithTableName
func NewApiClient(attrs ...*OperationAttr) (*ApiClient, error) {

	c := &ApiClient{}
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

		if c.mongoCollectionClient.Db == nil {
			return nil, errors.New("驱动不能为空")
		}

		if c.config.tableName == "" {
			return nil, errors.New("表名不能为空")
		}

	default:
		return nil, errors.New("驱动为空")
	}

	hostname, _ := os.Hostname()

	c.config.hostname = hostname
	c.config.insideIp = goip.GetInsideIp()
	c.config.goVersion = strings.TrimPrefix(runtime.Version(), "go")

	return c, nil
}
