package dorm

import "reflect"

// Database 设置库名
func (c *MongoClient) Database(databaseName string) *MongoClient {
	c.databaseName = databaseName
	return c
}

// Collection 设置表名
func (c *MongoClient) Collection(collectionName string) *MongoClient {
	c.collectionName = collectionName
	return c
}

// Model 传入模型自动获取库名和表名
func (c *MongoClient) Model(value interface{}) *MongoClient {
	// https://studygolang.com/articles/896
	val := reflect.ValueOf(value)
	if methodValue := val.MethodByName("Database"); methodValue.IsValid() {
		c.databaseName = methodValue.Call(nil)[0].String()
	}
	if methodValue := val.MethodByName("TableName"); methodValue.IsValid() {
		c.collectionName = methodValue.Call(nil)[0].String()
	}
	return c
}
