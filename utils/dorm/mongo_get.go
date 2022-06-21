package dorm

import "go.mongodb.org/mongo-driver/mongo"

// GetDb 获取驱动
func (c *MongoClient) GetDb() *mongo.Client {
	return c.Db
}

// 获取库名
func (c *MongoClient) getDatabaseName() string {
	return c.databaseName
}

// 获取表名
func (c *MongoClient) getCollectionName() string {
	return c.collectionName
}
