package dorm

import (
	"context"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// MongoClientFun *MongoClient 驱动
// string 库名
type MongoClientFun func() (*MongoClient, string)

// MongoClientCollectionFun *MongoClient 驱动
// string 库名
// string 集合
type MongoClientCollectionFun func() (*MongoClient, string, string)

// MongoClientConfig 实例配置
type MongoClientConfig struct {
	Dns          string // 地址
	Opts         *options.ClientOptions
	DatabaseName string // 库名
}

// MongoClient 实例
type MongoClient struct {
	db                 *mongo.Client // 驱动
	configDatabaseName string        // 库名
}

// NewMongoClient 创建实例
func NewMongoClient(config *MongoClientConfig) (*MongoClient, error) {

	var ctx = context.Background()
	var err error
	c := &MongoClient{}

	c.configDatabaseName = config.DatabaseName

	// 连接到MongoDB
	if config.Dns != "" {
		c.db, err = mongo.Connect(ctx, options.Client().ApplyURI(config.Dns))
		if err != nil {
			return nil, errors.New(fmt.Sprintf("连接失败：%v", err))
		}
	} else {
		c.db, err = mongo.Connect(ctx, config.Opts)
		if err != nil {
			return nil, errors.New(fmt.Sprintf("连接失败：%v", err))
		}
	}

	// 检查连接
	err = c.db.Ping(ctx, nil)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("检查连接失败：%v", err))
	}

	return c, nil
}

// Close 关闭
func (c *MongoClient) Close(ctx context.Context) error {
	return c.db.Disconnect(ctx)
}
