package dorm

import (
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoSessionCollectionOptions struct {
	db                 *mongo.Client        // 驱动
	configDatabaseName string               // 库名
	session            mongo.Session        // 会话
	sessionContext     mongo.SessionContext // 会话上下文
	dbCollection       *mongo.Collection    // 集合
}

// Collection 选择集合
func (csd *MongoSessionDatabaseOptions) Collection(name string, opts ...*options.CollectionOptions) *MongoSessionCollectionOptions {
	return &MongoSessionCollectionOptions{
		db:                 csd.db,                                   // 驱动
		configDatabaseName: csd.configDatabaseName,                   // 库名
		session:            csd.session,                              // 会话
		sessionContext:     csd.sessionContext,                       // 会话上下文
		dbCollection:       csd.dbDatabase.Collection(name, opts...), // 集合
	}
}
