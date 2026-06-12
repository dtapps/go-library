package dorm

import (
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// InsertOne 插入一个文档
func (csc *MongoSessionCollectionOptions) InsertOne(document any, opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error) {
	return csc.dbCollection.InsertOne(csc.sessionContext, document, opts...)
}

// InsertMany 插入多个文档
func (csc *MongoSessionCollectionOptions) InsertMany(document []any, opts ...*options.InsertManyOptions) (*mongo.InsertManyResult, error) {
	return csc.dbCollection.InsertMany(csc.sessionContext, document, opts...)
}

// DeleteOne 删除一个文档
func (csc *MongoSessionCollectionOptions) DeleteOne(filter any, opts ...*options.DeleteOptions) (*mongo.DeleteResult, error) {
	return csc.dbCollection.DeleteOne(csc.sessionContext, filter, opts...)
}

// DeleteMany 删除多个文档
func (csc *MongoSessionCollectionOptions) DeleteMany(filter any, opts ...*options.DeleteOptions) (*mongo.DeleteResult, error) {
	return csc.dbCollection.DeleteMany(csc.sessionContext, filter, opts...)
}

// UpdateByID 按ID更新
func (csc *MongoSessionCollectionOptions) UpdateByID(id any, update any, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
	return csc.dbCollection.UpdateByID(csc.sessionContext, id, update, opts...)
}

// UpdateOne 更新一个文档
func (csc *MongoSessionCollectionOptions) UpdateOne(filter any, update any, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
	return csc.dbCollection.UpdateOne(csc.sessionContext, filter, update, opts...)
}

// UpdateMany 更新多个文档
func (csc *MongoSessionCollectionOptions) UpdateMany(filter any, update any, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
	return csc.dbCollection.UpdateMany(csc.sessionContext, filter, update, opts...)
}

// ReplaceOne 替换一个文档
func (csc *MongoSessionCollectionOptions) ReplaceOne(filter any, update any, opts ...*options.ReplaceOptions) (*mongo.UpdateResult, error) {
	return csc.dbCollection.ReplaceOne(csc.sessionContext, filter, update, opts...)
}

// Aggregate 统计分析
func (csc *MongoSessionCollectionOptions) Aggregate(pipeline any, opts ...*options.AggregateOptions) (*mongo.Cursor, error) {
	return csc.dbCollection.Aggregate(csc.sessionContext, pipeline, opts...)
}

// CountDocuments 计数文档
func (csc *MongoSessionCollectionOptions) CountDocuments(filter any, opts ...*options.CountOptions) (int64, error) {
	return csc.dbCollection.CountDocuments(csc.sessionContext, filter, opts...)
}

// EstimatedDocumentCount 估计文档计数
func (csc *MongoSessionCollectionOptions) EstimatedDocumentCount(opts ...*options.EstimatedDocumentCountOptions) (int64, error) {
	return csc.dbCollection.EstimatedDocumentCount(csc.sessionContext, opts...)
}

func (csc *MongoSessionCollectionOptions) Distinct(fieldName string, filter any, opts ...*options.DistinctOptions) ([]any, error) {
	return csc.dbCollection.Distinct(csc.sessionContext, fieldName, filter, opts...)
}

// Find 查询多个文档
func (csc *MongoSessionCollectionOptions) Find(filter any, opts ...*options.FindOptions) (*mongo.Cursor, error) {
	return csc.dbCollection.Find(csc.sessionContext, filter, opts...)
}

// FindOne 查询一个文档
func (csc *MongoSessionCollectionOptions) FindOne(filter any, opts ...*options.FindOneOptions) *mongo.SingleResult {
	return csc.dbCollection.FindOne(csc.sessionContext, filter, opts...)
}

func (csc *MongoSessionCollectionOptions) FindOneAndDelete(filter any, opts ...*options.FindOneAndDeleteOptions) *mongo.SingleResult {
	return csc.dbCollection.FindOneAndDelete(csc.sessionContext, filter, opts...)
}

func (csc *MongoSessionCollectionOptions) FindOneAndReplace(filter any, replacement any, opts ...*options.FindOneAndReplaceOptions) *mongo.SingleResult {
	return csc.dbCollection.FindOneAndReplace(csc.sessionContext, filter, replacement, opts...)
}

func (csc *MongoSessionCollectionOptions) FindOneAndUpdate(filter any, replacement any, opts ...*options.FindOneAndUpdateOptions) *mongo.SingleResult {
	return csc.dbCollection.FindOneAndUpdate(csc.sessionContext, filter, replacement, opts...)
}

func (csc *MongoSessionCollectionOptions) Watch(pipeline any, opts ...*options.ChangeStreamOptions) (*mongo.ChangeStream, error) {
	return csc.dbCollection.Watch(csc.sessionContext, pipeline, opts...)
}

func (csc *MongoSessionCollectionOptions) Indexes() mongo.IndexView {
	return csc.dbCollection.Indexes()
}

func (csc *MongoSessionCollectionOptions) Drop() error {
	return csc.dbCollection.Drop(csc.sessionContext)
}
