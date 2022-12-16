package dorm

import "go.mongodb.org/mongo-driver/mongo"

// GetDb 获取驱动
func (c *MongoClient) GetDb() *mongo.Client {
	return c.db
}
