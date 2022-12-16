package dorm

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

type MongoSessionOptions struct {
	db                 *mongo.Client        // 驱动
	configDatabaseName string               // 库名
	session            mongo.Session        // 会话
	sessionContext     mongo.SessionContext // 会话上下文
}

// Begin 开始事务，会同时创建开始会话需要在退出时关闭会话
func (c *MongoClient) Begin() *MongoSessionOptions {

	var ctx = context.TODO()
	var err error
	ms := &MongoSessionOptions{}

	ms.db = c.GetDb()
	ms.configDatabaseName = c.configDatabaseName

	// 开始会话
	ms.session, err = ms.db.StartSession()
	if err != nil {
		log.Println("开始会话异常：", err)
	}

	// 会话上下文
	ms.sessionContext = mongo.NewSessionContext(ctx, ms.session)

	// 会话开启事务
	err = ms.session.StartTransaction()

	return ms
}

// Rollback 回滚事务
func (cs *MongoSessionOptions) Rollback() {
	var ctx = context.TODO()
	err := cs.session.AbortTransaction(ctx)
	if err != nil {
		log.Println("回滚事务异常：", err)
	}
	cs.session.EndSession(ctx)
}

// Commit 提交事务
func (cs *MongoSessionOptions) Commit() {
	var ctx = context.TODO()
	err := cs.session.CommitTransaction(ctx)
	if err != nil {
		log.Println("提交事务异常：", err)
	}
	cs.session.EndSession(ctx)
}
