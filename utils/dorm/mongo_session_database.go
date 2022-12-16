package dorm

import (
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"reflect"
)

type MongoSessionDatabaseOptions struct {
	db                 *mongo.Client        // 驱动
	configDatabaseName string               // 库名
	session            mongo.Session        // 会话
	sessionContext     mongo.SessionContext // 会话上下文
	dbDatabase         *mongo.Database      // 数据库
}

// Database 选择数据库
func (cs *MongoSessionOptions) Database(name string, opts ...*options.DatabaseOptions) *MongoSessionDatabaseOptions {
	return &MongoSessionDatabaseOptions{
		db:                 cs.db,                         // 驱动
		configDatabaseName: cs.configDatabaseName,         // 库名
		session:            cs.session,                    // 会话
		sessionContext:     cs.sessionContext,             // 会话上下文
		dbDatabase:         cs.db.Database(name, opts...), // 数据库
	}
}

// Model 传入模型自动获取库名和表名
// https://studygolang.com/articles/896
// DatabaseName 库名
// CollectionName 集合名
func (cs *MongoSessionOptions) Model(value interface{}) *MongoSessionCollectionOptions {

	var sessionDatabaseOptions *MongoSessionDatabaseOptions
	var sessionCollectionOptions *MongoSessionCollectionOptions

	val := reflect.ValueOf(value)

	methodDatabaseNameValue := val.MethodByName("DatabaseName")
	if methodDatabaseNameValue.IsValid() {
		databaseName := methodDatabaseNameValue.Call(nil)[0].String()
		sessionDatabaseOptions = cs.Database(databaseName)
	} else {
		sessionDatabaseOptions = cs.Database(cs.configDatabaseName)
	}

	methodCollectionNameValue := val.MethodByName("CollectionName")
	if methodCollectionNameValue.IsValid() {
		collectionName := methodCollectionNameValue.Call(nil)[0].String()
		sessionCollectionOptions = sessionDatabaseOptions.Collection(collectionName)
	} else {
		methodTableNameValue := val.MethodByName("TableName")
		if methodTableNameValue.IsValid() {
			collectionName := methodTableNameValue.Call(nil)[0].String()
			sessionCollectionOptions = sessionDatabaseOptions.Collection(collectionName)
		} else {
			panic(NoConfigCollectionName)
		}
	}

	return sessionCollectionOptions
}
