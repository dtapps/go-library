package dorm

import "go.mongodb.org/mongo-driver/mongo"

func (c *MongoClient) GetDb() *mongo.Client {
	return c.Db
}

func (c *MongoClient) getDatabaseName() string {
	return c.config.DatabaseName
}

func (c *MongoClient) getCollectionName() string {
	return c.collectionName
}
