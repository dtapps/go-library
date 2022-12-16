package dorm

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoCollectionOptions struct {
	db                 *mongo.Client     // 驱动
	configDatabaseName string            // 库名
	dbCollection       *mongo.Collection // 集合
}

// Collection 选择集合
func (cd *MongoDatabaseOptions) Collection(name string, opts ...*options.CollectionOptions) *MongoCollectionOptions {
	return &MongoCollectionOptions{
		db:                 cd.db,                                   // 驱动
		configDatabaseName: cd.configDatabaseName,                   // 库名
		dbCollection:       cd.dbDatabase.Collection(name, opts...), // 集合
	}
}

// CreateOneIndexes 创建一个索引
func (cc *MongoCollectionOptions) CreateOneIndexes(ctx context.Context, key string, value string) (string, error) {
	return cc.dbCollection.Indexes().CreateOne(ctx, mongo.IndexModel{
		Keys: bson.D{{
			Key:   key,
			Value: value,
		}},
	})
}

// CreateOneUniqueIndexes 创建一个唯一索引
func (cc *MongoCollectionOptions) CreateOneUniqueIndexes(ctx context.Context, key string, value string) (string, error) {
	return cc.dbCollection.Indexes().CreateOne(ctx, mongo.IndexModel{
		Keys: bson.D{{
			Key:   key,
			Value: value,
		}},
		Options: options.Index().SetUnique(true),
	})
}

// CreateOneUniqueIndexesOpts 创建一个索引
func (cc *MongoCollectionOptions) CreateOneUniqueIndexesOpts(ctx context.Context, key string, value string, opts *options.IndexOptions) (string, error) {
	return cc.dbCollection.Indexes().CreateOne(ctx, mongo.IndexModel{
		Keys: bson.D{{
			Key:   key,
			Value: value,
		}},
		Options: opts,
	})
}

// CreateManyIndexes 创建多个索引
func (cc *MongoCollectionOptions) CreateManyIndexes(ctx context.Context, models []mongo.IndexModel) ([]string, error) {
	return cc.dbCollection.Indexes().CreateMany(ctx, models)
}
