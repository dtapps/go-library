package dorm

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"reflect"
)

type MongoDatabaseOptions struct {
	db                 *mongo.Client   // 驱动
	configDatabaseName string          // 库名
	dbDatabase         *mongo.Database // 数据库
}

// Database 选择数据库
func (c *MongoClient) Database(name string, opts ...*options.DatabaseOptions) *MongoDatabaseOptions {
	return &MongoDatabaseOptions{
		db:                 c.db,                         // 驱动
		configDatabaseName: c.configDatabaseName,         // 库名
		dbDatabase:         c.db.Database(name, opts...), // 数据库
	}
}

// CreateCollection 创建集合
func (cd *MongoDatabaseOptions) CreateCollection(ctx context.Context, name string, opts ...*options.CreateCollectionOptions) error {
	return cd.dbDatabase.CreateCollection(ctx, name, opts...)
}

// CreateTimeSeriesCollection 创建时间序列集合
func (cd *MongoDatabaseOptions) CreateTimeSeriesCollection(ctx context.Context, name string, timeField string) error {
	return cd.dbDatabase.CreateCollection(ctx, name, options.CreateCollection().SetTimeSeriesOptions(options.TimeSeries().SetTimeField(timeField)))
}

// Model 传入模型自动获取库名和表名
// https://studygolang.com/articles/896
// DatabaseName 库名
// CollectionName 集合名
func (c *MongoClient) Model(value interface{}) *MongoCollectionOptions {

	var databaseOptions *MongoDatabaseOptions
	var collectionOptions *MongoCollectionOptions

	val := reflect.ValueOf(value)

	methodDatabaseNameValue := val.MethodByName("DatabaseName")
	if methodDatabaseNameValue.IsValid() {
		databaseName := methodDatabaseNameValue.Call(nil)[0].String()
		databaseOptions = c.Database(databaseName)
	} else {
		databaseOptions = c.Database(c.configDatabaseName)
	}

	methodCollectionNameValue := val.MethodByName("CollectionName")
	if methodCollectionNameValue.IsValid() {
		collectionName := methodCollectionNameValue.Call(nil)[0].String()
		collectionOptions = databaseOptions.Collection(collectionName)
	} else {
		methodTableNameValue := val.MethodByName("TableName")
		if methodTableNameValue.IsValid() {
			collectionName := methodTableNameValue.Call(nil)[0].String()
			collectionOptions = databaseOptions.Collection(collectionName)
		} else {
			panic(NoConfigCollectionName)
		}
	}

	return collectionOptions
}
