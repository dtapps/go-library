package dorm

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// InsertOne 插入单个文档
func (c *MongoClient) InsertOne(document interface{}) (result *mongo.InsertOneResult, err error) {
	collection := c.Db.Database(c.getDatabaseName()).Collection(c.collectionName)
	result, err = collection.InsertOne(context.TODO(), document)
	return result, err
}

// InsertMany 插入多个文档
func (c *MongoClient) InsertMany(documents []interface{}) (result *mongo.InsertManyResult, err error) {
	collection := c.Db.Database(c.getDatabaseName()).Collection(c.collectionName)
	result, err = collection.InsertMany(context.TODO(), documents)
	return result, err
}

// Delete 删除文档
func (c *MongoClient) Delete(filter interface{}) (int64, error) {
	collection := c.Db.Database(c.getDatabaseName()).Collection(c.collectionName)
	count, err := collection.DeleteOne(context.TODO(), filter)
	return count.DeletedCount, err
}

// DeleteMany 删除多个文档
func (c *MongoClient) DeleteMany(key string, value interface{}) (int64, error) {
	collection := c.Db.Database(c.getDatabaseName()).Collection(c.collectionName)
	filter := bson.D{{key, value}}
	count, err := collection.DeleteMany(context.TODO(), filter)
	return count.DeletedCount, err
}

// UpdateOne 更新单个文档
// 修改字段的值($set)
// 字段增加值 inc($inc)
// 从数组中增加一个元素 push($push)
// 从数组中删除一个元素 pull($pull)
func (c *MongoClient) UpdateOne(filter interface{}, update interface{}) (int64, error) {
	collection := c.Db.Database(c.getDatabaseName()).Collection(c.collectionName)
	result, err := collection.UpdateOne(context.TODO(), filter, update)
	return result.UpsertedCount, err
}

// UpdateMany 更新多个文档
// 修改字段的值($set)
// 字段增加值 inc($inc)
// 从数组中增加一个元素 push($push)
// 从数组中删除一个元素 pull($pull)
func (c *MongoClient) UpdateMany(filter interface{}, update interface{}) (int64, error) {
	collection := c.Db.Database(c.getDatabaseName()).Collection(c.collectionName)
	result, err := collection.UpdateMany(context.TODO(), filter, update)
	return result.UpsertedCount, err
}

// Find 查询
//func (c *MongoClient) Find(filter interface{}) (result *mongo.Cursor, err error) {
//	collection := c.Db.Database(c.DatabaseName).Collection(c.collectionName)
//	result, err = collection.Find(context.TODO(), filter)
//	return result, err
//}

// FindOne 查询单个文档
func (c *MongoClient) FindOne(filter interface{}) (result *mongo.SingleResult) {
	collection := c.Db.Database(c.getDatabaseName()).Collection(c.collectionName)
	result = collection.FindOne(context.TODO(), filter)
	return result
}

// FindMany 查询多个文档
func (c *MongoClient) FindMany(filter interface{}) (result *mongo.Cursor, err error) {
	collection := c.Db.Database(c.getDatabaseName()).Collection(c.collectionName)
	result, err = collection.Find(context.TODO(), filter)
	return result, err
}

// FindManyByFilters 多条件查询
func (c *MongoClient) FindManyByFilters(filter interface{}) (result *mongo.Cursor, err error) {
	collection, err := c.Db.Database(c.getDatabaseName()).Collection(c.collectionName).Clone()
	result, err = collection.Find(context.TODO(), bson.M{"$and": filter})
	return result, err
}

// FindManyByFiltersSort 多条件查询支持排序
func (c *MongoClient) FindManyByFiltersSort(filter interface{}, Sort interface{}) (result *mongo.Cursor, err error) {
	collection, err := c.Db.Database(c.getDatabaseName()).Collection(c.collectionName).Clone()
	findOptions := options.Find()
	findOptions.SetSort(Sort)
	result, err = collection.Find(context.TODO(), filter, findOptions)
	return result, err
}

// FindCollection 查询集合文档
func (c *MongoClient) FindCollection(Limit int64) (result *mongo.Cursor, err error) {
	collection := c.Db.Database(c.getDatabaseName()).Collection(c.collectionName)
	findOptions := options.Find()
	findOptions.SetLimit(Limit)
	result, err = collection.Find(context.TODO(), bson.D{{}}, findOptions)
	return result, err
}

// FindCollectionSort 查询集合文档支持排序
func (c *MongoClient) FindCollectionSort(Sort interface{}, Limit int64) (result *mongo.Cursor, err error) {
	collection := c.Db.Database(c.getDatabaseName()).Collection(c.collectionName)
	findOptions := options.Find()
	findOptions.SetSort(Sort)
	findOptions.SetLimit(Limit)
	result, err = collection.Find(context.TODO(), bson.D{{}}, findOptions)
	return result, err
}

// FindManyCollectionSort 查询集合文档支持排序支持条件
func (c *MongoClient) FindManyCollectionSort(filter interface{}, Sort interface{}) (result *mongo.Cursor, err error) {
	collection := c.Db.Database(c.getDatabaseName()).Collection(c.collectionName)
	findOptions := options.Find()
	findOptions.SetSort(Sort)
	result, err = collection.Find(context.TODO(), filter, findOptions)
	return result, err
}

// CollectionCount 查询集合里有多少数据
func (c *MongoClient) CollectionCount() (name string, size int64) {
	collection := c.Db.Database(c.getDatabaseName()).Collection(c.collectionName)
	name = collection.Name()
	size, _ = collection.EstimatedDocumentCount(context.TODO())
	return name, size
}

// CollectionDocuments 按选项查询集合
// Skip 跳过
// Limit 读取数量
// sort 1 ，-1 . 1 为升序 ， -1 为降序
func (c *MongoClient) CollectionDocuments(Skip, Limit int64, sort int, key string, value interface{}) (result *mongo.Cursor, err error) {
	collection := c.Db.Database(c.getDatabaseName()).Collection(c.collectionName)
	SORT := bson.D{{"_id", sort}}
	filter := bson.D{{key, value}}
	findOptions := options.Find().SetSort(SORT).SetLimit(Limit).SetSkip(Skip)
	result, err = collection.Find(context.TODO(), filter, findOptions)
	return result, err
}

// AggregateByFiltersSort 统计分析
func (c *MongoClient) AggregateByFiltersSort(pipeline interface{}) (result *mongo.Cursor, err error) {
	collection := c.Db.Database(c.getDatabaseName()).Collection(c.collectionName)
	result, err = collection.Aggregate(context.TODO(), pipeline)
	return result, err
}

// CountDocumentsByFilters 统计数量
func (c *MongoClient) CountDocumentsByFilters(filter interface{}) (count int64, err error) {
	collection := c.Db.Database(c.getDatabaseName()).Collection(c.collectionName)
	count, err = collection.CountDocuments(context.TODO(), filter)
	return count, err
}
