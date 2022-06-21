package dorm

import "reflect"

// Database 设置库名
func (ms *MongoTransaction) Database(databaseName string) *MongoTransaction {
	ms.databaseName = databaseName
	return ms
}

// Collection 设置表名
func (ms *MongoTransaction) Collection(collectionName string) *MongoTransaction {
	ms.collectionName = collectionName
	return ms
}

// Model 传入模型自动获取库名和表名
func (ms *MongoTransaction) Model(value interface{}) *MongoTransaction {
	// https://studygolang.com/articles/896
	val := reflect.ValueOf(value)
	if methodValue := val.MethodByName("Database"); methodValue.IsValid() {
		ms.databaseName = methodValue.Call(nil)[0].String()
	}
	if methodValue := val.MethodByName("TableName"); methodValue.IsValid() {
		ms.collectionName = methodValue.Call(nil)[0].String()
	}
	return ms
}
