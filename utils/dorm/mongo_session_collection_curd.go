package dorm

import (
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// InsertOne 插入一个文档
func (csc *MongoSessionCollectionOptions) InsertOne(document interface{}, opts ...*options.InsertOneOptions) (*mongo.InsertOneResult, error) {
	return csc.dbCollection.InsertOne(csc.sessionContext, document, opts...)
}

// InsertMany 插入多个文档
func (csc *MongoSessionCollectionOptions) InsertMany(document []interface{}, opts ...*options.InsertManyOptions) (*mongo.InsertManyResult, error) {
	return csc.dbCollection.InsertMany(csc.sessionContext, document, opts...)
}

// DeleteOne 删除一个文档
func (csc *MongoSessionCollectionOptions) DeleteOne(filter interface{}, opts ...*options.DeleteOptions) (*mongo.DeleteResult, error) {
	return csc.dbCollection.DeleteOne(csc.sessionContext, filter, opts...)
}

// DeleteMany 删除多个文档
func (csc *MongoSessionCollectionOptions) DeleteMany(filter interface{}, opts ...*options.DeleteOptions) (*mongo.DeleteResult, error) {
	return csc.dbCollection.DeleteMany(csc.sessionContext, filter, opts...)
}

// UpdateByID 按ID更新
func (csc *MongoSessionCollectionOptions) UpdateByID(id interface{}, update interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
	return csc.dbCollection.UpdateByID(csc.sessionContext, id, update, opts...)
}

// UpdateOne 更新一个文档
func (csc *MongoSessionCollectionOptions) UpdateOne(filter interface{}, update interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
	return csc.dbCollection.UpdateOne(csc.sessionContext, filter, update, opts...)
}

// UpdateMany 更新多个文档
func (csc *MongoSessionCollectionOptions) UpdateMany(filter interface{}, update interface{}, opts ...*options.UpdateOptions) (*mongo.UpdateResult, error) {
	return csc.dbCollection.UpdateMany(csc.sessionContext, filter, update, opts...)
}

// ReplaceOne 替换一个文档
func (csc *MongoSessionCollectionOptions) ReplaceOne(filter interface{}, update interface{}, opts ...*options.ReplaceOptions) (*mongo.UpdateResult, error) {
	return csc.dbCollection.ReplaceOne(csc.sessionContext, filter, update, opts...)
}

// Aggregate 统计分析
func (csc *MongoSessionCollectionOptions) Aggregate(pipeline interface{}, opts ...*options.AggregateOptions) (*mongo.Cursor, error) {
	return csc.dbCollection.Aggregate(csc.sessionContext, pipeline, opts...)
}

// CountDocuments 计数文档
func (csc *MongoSessionCollectionOptions) CountDocuments(filter interface{}, opts ...*options.CountOptions) (int64, error) {
	return csc.dbCollection.CountDocuments(csc.sessionContext, filter, opts...)
}

// EstimatedDocumentCount 估计文档计数
func (csc *MongoSessionCollectionOptions) EstimatedDocumentCount(opts ...*options.EstimatedDocumentCountOptions) (int64, error) {
	return csc.dbCollection.EstimatedDocumentCount(csc.sessionContext, opts...)
}

func (csc *MongoSessionCollectionOptions) Distinct(fieldName string, filter interface{}, opts ...*options.DistinctOptions) ([]interface{}, error) {
	return csc.dbCollection.Distinct(csc.sessionContext, fieldName, filter, opts...)
}

// Find 查询多个文档
func (csc *MongoSessionCollectionOptions) Find(filter interface{}, opts ...*options.FindOptions) (*mongo.Cursor, error) {
	return csc.dbCollection.Find(csc.sessionContext, filter, opts...)
}

// FindOne 查询一个文档
func (csc *MongoSessionCollectionOptions) FindOne(filter interface{}, opts ...*options.FindOneOptions) *mongo.SingleResult {
	return csc.dbCollection.FindOne(csc.sessionContext, filter, opts...)
}

func (csc *MongoSessionCollectionOptions) FindOneAndDelete(filter interface{}, opts ...*options.FindOneAndDeleteOptions) *mongo.SingleResult {
	return csc.dbCollection.FindOneAndDelete(csc.sessionContext, filter, opts...)
}

func (csc *MongoSessionCollectionOptions) FindOneAndReplace(filter interface{}, replacement interface{}, opts ...*options.FindOneAndReplaceOptions) *mongo.SingleResult {
	return csc.dbCollection.FindOneAndReplace(csc.sessionContext, filter, replacement, opts...)
}

func (csc *MongoSessionCollectionOptions) FindOneAndUpdate(filter interface{}, replacement interface{}, opts ...*options.FindOneAndUpdateOptions) *mongo.SingleResult {
	return csc.dbCollection.FindOneAndUpdate(csc.sessionContext, filter, replacement, opts...)
}

func (csc *MongoSessionCollectionOptions) Watch(pipeline interface{}, opts ...*options.ChangeStreamOptions) (*mongo.ChangeStream, error) {
	return csc.dbCollection.Watch(csc.sessionContext, pipeline, opts...)
}

func (csc *MongoSessionCollectionOptions) Indexes() mongo.IndexView {
	return csc.dbCollection.Indexes()
}

func (csc *MongoSessionCollectionOptions) Drop() error {
	return csc.dbCollection.Drop(csc.sessionContext)
}
