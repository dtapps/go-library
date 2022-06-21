package dorm

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
)

type MongoTransaction struct {
	startSession   mongo.Session
	Session        mongo.SessionContext
	db             *mongo.Client // 驱动
	databaseName   string        // 库名
	collectionName string        // 表名
}

// Begin 开始事务，会同时创建开始会话需要在退出时关闭会话
func (c *MongoClient) Begin() (ms MongoTransaction, err error) {

	ms.db = c.Db

	// 开始会话
	ms.startSession, err = ms.db.StartSession()
	if err != nil {
		panic(err)
	}

	// 会话上下文
	ms.Session = mongo.NewSessionContext(context.Background(), ms.startSession)

	// 会话开启事务
	err = ms.startSession.StartTransaction()
	return ms, err
}

// Close 关闭会话
func (ms *MongoTransaction) Close() {
	ms.startSession.EndSession(context.TODO())
}

// Rollback 回滚事务
func (ms *MongoTransaction) Rollback() error {
	return ms.startSession.AbortTransaction(context.Background())
}

// Commit 提交事务
func (ms *MongoTransaction) Commit() error {
	return ms.startSession.CommitTransaction(context.Background())
}
