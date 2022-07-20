package golog

import (
	"github.com/dtapps/go-library/utils/dorm"
	"github.com/dtapps/go-library/utils/goip"
	"gorm.io/gorm"
)

const (
	logTypeGorm  = "gorm"
	logTypeMongo = "mongo"
)

// OperationAttr 操作属性
type OperationAttr struct {
	gormClient     *gorm.DB          // 数据库驱动
	mongoClient    *dorm.MongoClient // 数据库驱动
	ipService      *goip.Client      // ip服务
	logType        string            // 类型
	tableName      string            // 表名
	databaseName   string            // 库名
	collectionName string            // 表名
}

// WithGormClient 设置数据库驱动
func WithGormClient(client *gorm.DB) *OperationAttr {
	return &OperationAttr{gormClient: client, logType: logTypeGorm}
}

// WithMongoClient 设置数据库驱动
func WithMongoClient(client *dorm.MongoClient) *OperationAttr {
	return &OperationAttr{mongoClient: client, logType: logTypeMongo}
}

// WithTableName 设置表名
func WithTableName(tableName string) *OperationAttr {
	return &OperationAttr{tableName: tableName}
}

// WithDatabaseName 设置库名
func WithDatabaseName(databaseName string) *OperationAttr {
	return &OperationAttr{databaseName: databaseName}
}

// WithCollectionName 设置表名
func WithCollectionName(collectionName string) *OperationAttr {
	return &OperationAttr{collectionName: collectionName}
}

// WithIpService 设置ip服务
func WithIpService(ipService *goip.Client) *OperationAttr {
	return &OperationAttr{ipService: ipService}
}
