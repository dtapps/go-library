package golog

import (
	"go.dtapp.net/library/utils/dorm"
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
}

// WithGormClient 数据库驱动
func WithGormClient(client *gorm.DB) *OperationAttr {
	return &OperationAttr{gormClient: client, logType: logTypeGorm}
}

// WithMongoCollectionClient 数据库驱动(温馨提示：需要已选择库)
func WithMongoCollectionClient(client *dorm.MongoClient) *OperationAttr {
	return &OperationAttr{mongoCollectionClient: client, logType: logTypeMongo}
}

// WithTableName 表名
func WithTableName(tableName string) *OperationAttr {
	return &OperationAttr{tableName: tableName}
}
