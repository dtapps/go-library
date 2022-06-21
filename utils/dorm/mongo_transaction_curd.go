package dorm

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// InsertOne 插入一个文档
func (ms *MongoTransaction) InsertOne(document interface{}) (result *InsertOneResult, err error) {
	collection := ms.db.Database(ms.getDatabaseName()).Collection(ms.collectionName)
	res, err := collection.InsertOne(ms.Session, document)
	return &InsertOneResult{InsertedID: res.InsertedID}, err
}

// InsertMany 插入多个文档
func (ms *MongoTransaction) InsertMany(documents []interface{}) (result *InsertManyResult, err error) {
	collection := ms.db.Database(ms.getDatabaseName()).Collection(ms.collectionName)
	res, err := collection.InsertMany(ms.Session, documents)
	return &InsertManyResult{InsertedIDs: res.InsertedIDs}, err
}

// Delete 删除文档
func (ms *MongoTransaction) Delete(filter interface{}) error {
	collection := ms.db.Database(ms.getDatabaseName()).Collection(ms.collectionName)
	_, err := collection.DeleteOne(ms.Session, filter)
	return err
}

// DeleteId 删除文档
func (ms *MongoTransaction) DeleteId(id interface{}) error {
	collection := ms.db.Database(ms.getDatabaseName()).Collection(ms.collectionName)
	_, err := collection.DeleteOne(ms.Session, bson.M{"_id": id})
	return err
}

// DeleteMany 删除多个文档
func (ms *MongoTransaction) DeleteMany(key string, value interface{}) (result *DeleteResult, err error) {
	collection := ms.db.Database(ms.getDatabaseName()).Collection(ms.collectionName)
	filter := bson.D{{key, value}}
	res, err := collection.DeleteMany(ms.Session, filter)
	return &DeleteResult{DeletedCount: res.DeletedCount}, err
}

// UpdateOne 更新单个文档
// 修改字段的值($set)
// 字段增加值 inc($inc)
// 从数组中增加一个元素 push($push)
// 从数组中删除一个元素 pull($pull)
func (ms *MongoTransaction) UpdateOne(filter interface{}, update interface{}) error {
	collection := ms.db.Database(ms.getDatabaseName()).Collection(ms.collectionName)
	_, err := collection.UpdateOne(ms.Session, filter, update)
	return err
}

// UpdateId 更新单个文档
// 修改字段的值($set)
// 字段增加值 inc($inc)
// 从数组中增加一个元素 push($push)
// 从数组中删除一个元素 pull($pull)
func (ms *MongoTransaction) UpdateId(id interface{}, update interface{}) error {
	collection := ms.db.Database(ms.getDatabaseName()).Collection(ms.collectionName)
	_, err := collection.UpdateOne(context.TODO(), bson.M{"_id": id}, update)
	return err
}

// UpdateMany 更新多个文档
// 修改字段的值($set)
// 字段增加值 inc($inc)
// 从数组中增加一个元素 push($push)
// 从数组中删除一个元素 pull($pull)
func (ms *MongoTransaction) UpdateMany(filter interface{}, update interface{}) (result *UpdateResult, err error) {
	collection := ms.db.Database(ms.getDatabaseName()).Collection(ms.collectionName)
	res, err := collection.UpdateMany(ms.Session, filter, update)
	return &UpdateResult{
		MatchedCount:  res.MatchedCount,
		ModifiedCount: res.ModifiedCount,
		UpsertedCount: res.UpsertedCount,
		UpsertedID:    res.UpsertedID,
	}, err
}

// Find 查询
func (ms *MongoTransaction) Find(filter interface{}) FindResultI {
	collection := ms.db.Database(ms.getDatabaseName()).Collection(ms.collectionName)
	res, err := collection.Find(ms.Session, filter)
	return &FindResult{
		cursor: res,
		err:    err,
	}
}

// FindOne 查询单个文档
func (ms *MongoTransaction) FindOne(filter interface{}) FindOneResultI {
	collection := ms.db.Database(ms.getDatabaseName()).Collection(ms.collectionName)
	res := collection.FindOne(ms.Session, filter)
	return &FindOneResult{
		singleResult: res,
	}
}

// FindMany 查询多个文档
func (ms *MongoTransaction) FindMany(filter interface{}) FindManyResultI {
	collection := ms.db.Database(ms.getDatabaseName()).Collection(ms.collectionName)
	res, err := collection.Find(ms.Session, filter)
	return &FindManyResult{
		cursor: res,
		err:    err,
	}
}

// FindManyByFilters 多条件查询
func (ms *MongoTransaction) FindManyByFilters(filter interface{}) (result *mongo.Cursor, err error) {
	collection, err := ms.db.Database(ms.getDatabaseName()).Collection(ms.collectionName).Clone()
	result, err = collection.Find(ms.Session, bson.M{"$and": filter})
	return result, err
}

// FindManyByFiltersSort 多条件查询支持排序
func (ms *MongoTransaction) FindManyByFiltersSort(filter interface{}, Sort interface{}) (result *mongo.Cursor, err error) {
	collection, err := ms.db.Database(ms.getDatabaseName()).Collection(ms.collectionName).Clone()
	findOptions := options.Find()
	findOptions.SetSort(Sort)
	result, err = collection.Find(ms.Session, filter, findOptions)
	return result, err
}

// FindCollection 查询集合文档
func (ms *MongoTransaction) FindCollection(Limit int64) (result *mongo.Cursor, err error) {
	collection := ms.db.Database(ms.getDatabaseName()).Collection(ms.collectionName)
	findOptions := options.Find()
	findOptions.SetLimit(Limit)
	result, err = collection.Find(ms.Session, bson.D{{}}, findOptions)
	return result, err
}

// FindCollectionSort 查询集合文档支持排序
func (ms *MongoTransaction) FindCollectionSort(Sort interface{}, Limit int64) (result *mongo.Cursor, err error) {
	collection := ms.db.Database(ms.getDatabaseName()).Collection(ms.collectionName)
	findOptions := options.Find()
	findOptions.SetSort(Sort)
	findOptions.SetLimit(Limit)
	result, err = collection.Find(ms.Session, bson.D{{}}, findOptions)
	return result, err
}

// FindManyCollectionSort 查询集合文档支持排序支持条件
func (ms *MongoTransaction) FindManyCollectionSort(filter interface{}, Sort interface{}) (result *mongo.Cursor, err error) {
	collection := ms.db.Database(ms.getDatabaseName()).Collection(ms.collectionName)
	findOptions := options.Find()
	findOptions.SetSort(Sort)
	result, err = collection.Find(ms.Session, filter, findOptions)
	return result, err
}

// CollectionCount 查询集合里有多少数据
func (ms *MongoTransaction) CollectionCount(ctx context.Context) (name string, size int64) {
	collection := ms.db.Database(ms.getDatabaseName()).Collection(ms.collectionName)
	name = collection.Name()
	size, _ = collection.EstimatedDocumentCount(ctx)
	return name, size
}

// CollectionDocuments 按选项查询集合
// Skip 跳过
// Limit 读取数量
// sort 1 ，-1 . 1 为升序 ， -1 为降序
func (ms *MongoTransaction) CollectionDocuments(Skip, Limit int64, sort int, key string, value interface{}) (result *mongo.Cursor, err error) {
	collection := ms.db.Database(ms.getDatabaseName()).Collection(ms.collectionName)
	SORT := bson.D{{"_id", sort}}
	filter := bson.D{{key, value}}
	findOptions := options.Find().SetSort(SORT).SetLimit(Limit).SetSkip(Skip)
	result, err = collection.Find(ms.Session, filter, findOptions)
	return result, err
}

// AggregateByFiltersSort 统计分析
func (ms *MongoTransaction) AggregateByFiltersSort(pipeline interface{}) (result *mongo.Cursor, err error) {
	collection := ms.db.Database(ms.getDatabaseName()).Collection(ms.collectionName)
	result, err = collection.Aggregate(ms.Session, pipeline)
	return result, err
}

// CountDocumentsByFilters 统计数量
func (ms *MongoTransaction) CountDocumentsByFilters(filter interface{}) (count int64, err error) {
	collection := ms.db.Database(ms.getDatabaseName()).Collection(ms.collectionName)
	count, err = collection.CountDocuments(ms.Session, filter)
	return count, err
}
