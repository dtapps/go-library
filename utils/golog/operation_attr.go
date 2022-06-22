package golog

import (
	"go.dtapp.net/library/utils/dorm"
	"go.dtapp.net/library/utils/goip"
	"gorm.io/gorm"
)

const (
	logTypeGorm  = "gorm"
	logTypeMongo = "mongo"
)

// OperationAttr 操作属性
type OperationAttr struct {
	gormClient            *gorm.DB          // 驱动
	mongoCollectionClient *dorm.MongoClient // 驱动
	logType               string            // 类型
	tableName             string            // 表名
	ipService             *goip.Client      // ip服务
}

// WithGormClient 数据库驱动
func WithGormClient(client *gorm.DB) *OperationAttr {
	return &OperationAttr{gormClient: client, logType: logTypeGorm}
}

// WithMongoCollectionClient 数据库驱动(温馨提示：需要已选择库和表)
func WithMongoCollectionClient(client *dorm.MongoClient) *OperationAttr {
	return &OperationAttr{mongoCollectionClient: client, logType: logTypeMongo}
}

// WithTableName 表名
func WithTableName(tableName string) *OperationAttr {
	return &OperationAttr{tableName: tableName}
}

// WithIpService ip服务
func WithIpService(ipService *goip.Client) *OperationAttr {
	return &OperationAttr{ipService: ipService}
}
