package dorm

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// InsertOne 插入一个文档
func (cc *MongoCollectionOptions) InsertOne(ctx context.Context, document interface{}, opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error) {
	return cc.dbCollection.InsertOne(ctx, document, opts...)
}

// InsertMany 插入多个文档
func (cc *MongoCollectionOptions) InsertMany(ctx context.Context, document []interface{}, opts ...*options.InsertManyOptions) (*mongo.InsertManyResult, error) {
	return cc.dbCollection.InsertMany(ctx, document, opts...)
}

// DeleteOne 删除一个文档
func (cc *MongoCollectionOptions) DeleteOne(ctx context.Context, filter interface{}, opts ...*options.DeleteOptions) (*mongo.DeleteResult, error) {
	return cc.dbCollection.DeleteOne(ctx, filter, opts...)
}

// DeleteMany 删除多个文档
func (cc *MongoCollectionOptions) DeleteMany(ctx context.Context, filter interface{}, opts ...*options.DeleteOptions) (*mongo.DeleteResult, error) {
	return cc.dbCollection.DeleteMany(ctx, filter, opts...)
}

// UpdateByID 按ID更新
func (cc *MongoCollectionOptions) UpdateByID(ctx context.Context, id interface{}, update interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
	return cc.dbCollection.UpdateByID(ctx, id, update, opts...)
}

// UpdateOne 更新一个文档
func (cc *MongoCollectionOptions) UpdateOne(ctx context.Context, filter interface{}, update interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
	return cc.dbCollection.UpdateOne(ctx, filter, update, opts...)
}

// UpdateMany 更新多个文档
func (cc *MongoCollectionOptions) UpdateMany(ctx context.Context, filter interface{}, update interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
	return cc.dbCollection.UpdateMany(ctx, filter, update, opts...)
}

// ReplaceOne 替换一个文档
func (cc *MongoCollectionOptions) ReplaceOne(ctx context.Context, filter interface{}, update interface{}, opts ...*options.ReplaceOptions) (*mongo.UpdateResult, error) {
	return cc.dbCollection.ReplaceOne(ctx, filter, update, opts...)
}

// Aggregate 统计分析
func (cc *MongoCollectionOptions) Aggregate(ctx context.Context, pipeline interface{}, opts ...*options.AggregateOptions) (*mongo.Cursor, error) {
	return cc.dbCollection.Aggregate(ctx, pipeline, opts...)
}

// CountDocuments 计数文档
func (cc *MongoCollectionOptions) CountDocuments(ctx context.Context, filter interface{}, opts ...*options.CountOptions) (int64, error) {
	return cc.dbCollection.CountDocuments(ctx, filter, opts...)
}

// EstimatedDocumentCount 估计文档计数
func (cc *MongoCollectionOptions) EstimatedDocumentCount(ctx context.Context, opts ...*options.EstimatedDocumentCountOptions) (int64, error) {
	return cc.dbCollection.EstimatedDocumentCount(ctx, opts...)
}

func (cc *MongoCollectionOptions) Distinct(ctx context.Context, fieldName string, filter interface{}, opts ...*options.DistinctOptions) ([]interface{}, error) {
	return cc.dbCollection.Distinct(ctx, fieldName, filter, opts...)
}

// Find 查询多个文档
func (cc *MongoCollectionOptions) Find(ctx context.Context, filter interface{}, opts ...*options.FindOptions) (*mongo.Cursor, error) {
	return cc.dbCollection.Find(ctx, filter, opts...)
}

// FindOne 查询一个文档
func (cc *MongoCollectionOptions) FindOne(ctx context.Context, filter interface{}, opts ...*options.FindOneOptions) *mongo.SingleResult {
	return cc.dbCollection.FindOne(ctx, filter, opts...)
}

func (cc *MongoCollectionOptions) FindOneAndDelete(ctx context.Context, filter interface{}, opts ...*options.FindOneAndDeleteOptions) *mongo.SingleResult {
	return cc.dbCollection.FindOneAndDelete(ctx, filter, opts...)
}

func (cc *MongoCollectionOptions) FindOneAndReplace(ctx context.Context, filter interface{}, replacement interface{}, opts ...*options.FindOneAndReplaceOptions) *mongo.SingleResult {
	return cc.dbCollection.FindOneAndReplace(ctx, filter, replacement, opts...)
}

func (cc *MongoCollectionOptions) FindOneAndUpdate(ctx context.Context, filter interface{}, replacement interface{}, opts ...*options.FindOneAndUpdateOptions) *mongo.SingleResult {
	return cc.dbCollection.FindOneAndUpdate(ctx, filter, replacement, opts...)
}

func (cc *MongoCollectionOptions) Watch(ctx context.Context, pipeline interface{}, opts ...*options.ChangeStreamOptions) (*mongo.ChangeStream, error) {
	return cc.dbCollection.Watch(ctx, pipeline, opts...)
}

func (cc *MongoCollectionOptions) Indexes(ctx context.Context) mongo.IndexView {
	return cc.dbCollection.Indexes()
}

func (cc *MongoCollectionOptions) Drop(ctx context.Context) error {
	return cc.dbCollection.Drop(ctx)
}
