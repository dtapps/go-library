package dorm

import (
	"context"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ConfigMongoClient struct {
	Dns          string // 地址
	DatabaseName string // 库名
}

type MongoClient struct {
	Db             *mongo.Client      // 驱动
	config         *ConfigMongoClient // 配置
	collectionName string             // 表名
	filterArr      []queryFilter      // 查询条件数组
	filter         bson.D             // 查询条件
}

func NewMongoClient(config *ConfigMongoClient) (*MongoClient, error) {

	c := &MongoClient{config: config}

	// 连接到MongoDB
	db, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(c.config.Dns))
	if err != nil {
		return nil, errors.New(fmt.Sprintf("连接失败：%v", err))
	}

	// 检查连接
	err = db.Ping(context.TODO(), nil)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("检查连接失败：%v", err))
	}

	return c, nil
}

// Close 关闭
func (c *MongoClient) Close() error {
	err := c.Db.Disconnect(context.TODO())
	if err != nil {
		return errors.New(fmt.Sprintf("关闭失败：%v", err))
	}
	return nil
}
