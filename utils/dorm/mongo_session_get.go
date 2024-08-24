package dorm

import "go.mongodb.org/mongo-driver/mongo"

// GetDb 获取驱动
func (cs *MongoSessionOptions) GetDb() *mongo.Client {
	return cs.db
}

// GetSession 获取会话
func (cs *MongoSessionOptions) GetSession() mongo.Session {
	return cs.session
}

// GetSessionContext 获取会话上下文
func (cs *MongoSessionOptions) GetSessionContext() mongo.SessionContext {
	return cs.sessionContext
}
