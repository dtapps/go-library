package gomongo

import (
	"context"
	"errors"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Client 实例
type Client struct {
	db             *mongo.Client   // MongoDB驱动
	ctx            context.Context // 上下文
	DatabaseName   string          // 库名
	collectionName string          // 表名
	filterArr      []queryFilter   // 查询条件数组
	filter         bson.D          // 查询条件
}

// NewClient 实例化并链接数据库
func NewClient(dns string) *Client {

	// 连接到MongoDB
	db, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(dns))
	if err != nil {
		panic(fmt.Sprintf("连接失败：%v", err))
	}

	// 检查连接
	err = db.Ping(context.TODO(), nil)
	if err != nil {
		panic(fmt.Sprintf("检查连接失败：%v", err))
	}

	return &Client{db: db, ctx: context.TODO()}
}

// NewClientDb 实例化并传入链接
func NewClientDb(db *mongo.Client) *Client {
	return &Client{db: db, ctx: context.TODO()}
}

// Close 关闭
func (c *Client) Close() {
	err := c.db.Disconnect(context.TODO())
	if err != nil {
		panic(errors.New(fmt.Sprintf("关闭失败：%v", err)))
	}
	return
}

// GetDbDriver 获取驱动
func (c *Client) GetDbDriver() *mongo.Client {
	return c.db
}
