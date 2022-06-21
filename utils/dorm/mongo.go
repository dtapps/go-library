package dorm

import (
	"context"
	"errors"
	"fmt"
	"github.com/qiniu/qmgo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"gopkg.in/mgo.v2"
	"log"
)

type ConfigMongoClient struct {
	Dns          string // 地址
	Opts         *options.ClientOptions
	DatabaseName string // 库名
}

type MongoClient struct {
	Db             *mongo.Client      // 驱动
	config         *ConfigMongoClient // 配置
	databaseName   string             // 库名
	collectionName string             // 表名
	//filterArr      []queryFilter      // 查询条件数组
	filter bson.D // 查询条件
}

func NewMongoClient(config *ConfigMongoClient) (*MongoClient, error) {

	var err error
	c := &MongoClient{config: config}

	c.databaseName = c.config.DatabaseName

	// 连接到MongoDB
	if c.config.Dns != "" {
		c.Db, err = mongo.Connect(context.Background(), options.Client().ApplyURI(c.config.Dns))
		if err != nil {
			return nil, errors.New(fmt.Sprintf("连接失败：%v", err))
		}
	} else {
		c.Db, err = mongo.Connect(context.Background(), c.config.Opts)
		if err != nil {
			return nil, errors.New(fmt.Sprintf("连接失败：%v", err))
		}
	}

	// 检查连接
	err = c.Db.Ping(context.TODO(), nil)
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

func NewMongoClientOfQmgo(config *ConfigMongoClient) (*MongoClient, error) {

	var err error
	c := &MongoClient{config: config}

	// 连接到MongoDB
	ctx := context.Background()
	client, err := qmgo.NewClient(ctx, &qmgo.Config{Uri: "mongodb://localhost:27017"})
	if err != nil {
		return nil, errors.New(fmt.Sprintf("连接失败：%v", err))
	}
	log.Println(client)
	db := client.Database("class")
	coll := db.Collection("user")
	log.Println(coll)
	//coll.Find().One()
	// 检查连接
	err = c.Db.Ping(context.TODO(), nil)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("检查连接失败：%v", err))
	}

	return c, nil
}

func NewMongoClientOfMgo(config *ConfigMongoClient) (*MongoClient, error) {

	var err error
	c := &MongoClient{config: config}

	// 连接到MongoDB
	client, err := mgo.Dial("localhost:27017")
	if err != nil {
		return nil, errors.New(fmt.Sprintf("连接失败：%v", err))
	}
	log.Println(client)

	// 检查连接
	err = c.Db.Ping(context.TODO(), nil)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("检查连接失败：%v", err))
	}

	return c, nil
}
