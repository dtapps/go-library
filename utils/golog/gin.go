package golog

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"go.dtapp.net/library/utils/dorm"
	"go.dtapp.net/library/utils/goip"
	"gorm.io/gorm"
	"os"
	"runtime"
)

// GinClient 框架
type GinClient struct {
	gormClient            *gorm.DB          // 驱动
	mongoCollectionClient *dorm.MongoClient // 驱动(温馨提示：需要已选择库)
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
	c.config.goVersion = runtime.Version()

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
