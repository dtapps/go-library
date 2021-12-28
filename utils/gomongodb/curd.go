package gomongodb

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func (app *App) Session() (session mongo.Session, err error) {
	session, err = app.Mgo.StartSession()
	return
}

// InsertOne 插入单个文档
func (app *App) InsertOne(value interface{}) (result *mongo.InsertOneResult, err error) {
	collection := app.Mgo.Database(app.Database).Collection(app.Collection)
	result, err = collection.InsertOne(context.TODO(), value)
	return result, err
}

// InsertMany 插入多个文档
func (app *App) InsertMany(values []interface{}) (result *mongo.InsertManyResult, err error) {
	collection := app.Mgo.Database(app.Database).Collection(app.Collection)
	result, err = collection.InsertMany(context.TODO(), values)
	return result, err
}

// Delete 删除文档
func (app *App) Delete(filter interface{}) (int64, error) {
	collection := app.Mgo.Database(app.Database).Collection(app.Collection)
	count, err := collection.DeleteOne(context.TODO(), filter, nil)
	return count.DeletedCount, err
}

// DeleteMany 删除多个文档
func (app *App) DeleteMany(key string, value interface{}) (int64, error) {
	collection := app.Mgo.Database(app.Database).Collection(app.Collection)
	filter := bson.D{{key, value}}
	count, err := collection.DeleteMany(context.TODO(), filter)
	return count.DeletedCount, err
}

// UpdateOne 更新单个文档
// 修改字段的值($set)
// 字段增加值 inc($inc)
// 从数组中增加一个元素 push($push)
// 从数组中删除一个元素 pull($pull)
func (app *App) UpdateOne(filter, update interface{}) (int64, error) {
	collection := app.Mgo.Database(app.Database).Collection(app.Collection)
	result, err := collection.UpdateOne(context.TODO(), filter, update)
	return result.UpsertedCount, err
}

// UpdateMany 更新多个文档
// 修改字段的值($set)
// 字段增加值 inc($inc)
// 从数组中增加一个元素 push($push)
// 从数组中删除一个元素 pull($pull)
func (app *App) UpdateMany(filter, update interface{}) (int64, error) {
	collection := app.Mgo.Database(app.Database).Collection(app.Collection)
	result, err := collection.UpdateMany(context.TODO(), filter, update)
	return result.UpsertedCount, err
}

// Find 查询
func (app *App) Find(filter interface{}, opts ...*options.FindOptions) (result *mongo.Cursor, err error) {
	collection := app.Mgo.Database(app.Database).Collection(app.Collection)
	result, err = collection.Find(context.TODO(), filter, opts...)
	return result, err
}

// FindOne 查询单个文档
func (app *App) FindOne(filter interface{}) (result *mongo.SingleResult) {
	collection := app.Mgo.Database(app.Database).Collection(app.Collection)
	result = collection.FindOne(context.TODO(), filter)
	return result
}

// FindMany 查询多个文档
func (app *App) FindMany(filter interface{}) (result *mongo.Cursor, err error) {
	collection := app.Mgo.Database(app.Database).Collection(app.Collection)
	result, err = collection.Find(context.TODO(), filter)
	return result, err
}

// FindManyByFilters 多条件查询
func (app *App) FindManyByFilters(filter interface{}) (result *mongo.Cursor, err error) {
	collection, err := app.Mgo.Database(app.Database).Collection(app.Collection).Clone()
	result, err = collection.Find(context.TODO(), bson.M{"$and": filter})
	return result, err
}

// FindManyByFiltersSort 多条件查询支持排序
func (app *App) FindManyByFiltersSort(filter interface{}, Sort interface{}) (result *mongo.Cursor, err error) {
	collection, err := app.Mgo.Database(app.Database).Collection(app.Collection).Clone()
	findOptions := options.Find()
	findOptions.SetSort(Sort)
	result, err = collection.Find(context.TODO(), filter, findOptions)
	return result, err
}

// FindCollection 查询集合文档
func (app *App) FindCollection(Limit int64) (result *mongo.Cursor, err error) {
	collection := app.Mgo.Database(app.Database).Collection(app.Collection)
	findOptions := options.Find()
	findOptions.SetLimit(Limit)
	result, err = collection.Find(context.TODO(), bson.D{{}}, findOptions)
	return result, err
}

// FindCollectionSort 查询集合文档支持排序
func (app *App) FindCollectionSort(Sort interface{}, Limit int64) (result *mongo.Cursor, err error) {
	collection := app.Mgo.Database(app.Database).Collection(app.Collection)
	findOptions := options.Find()
	findOptions.SetSort(Sort)
	findOptions.SetLimit(Limit)
	result, err = collection.Find(context.TODO(), bson.D{{}}, findOptions)
	return result, err
}

// FindManyCollectionSort 查询集合文档支持排序支持条件
func (app *App) FindManyCollectionSort(filter interface{}, Sort interface{}) (result *mongo.Cursor, err error) {
	collection := app.Mgo.Database(app.Database).Collection(app.Collection)
	findOptions := options.Find()
	findOptions.SetSort(Sort)
	result, err = collection.Find(context.TODO(), filter, findOptions)
	return result, err
}

// CollectionCount 查询集合里有多少数据
func (app *App) CollectionCount() (name string, size int64) {
	collection := app.Mgo.Database(app.Database).Collection(app.Collection)
	name = collection.Name()
	size, _ = collection.EstimatedDocumentCount(context.TODO())
	return name, size
}

// CollectionDocuments 按选项查询集合
// Skip 跳过
// Limit 读取数量
// sort 1 ，-1 . 1 为升序 ， -1 为降序
func (app *App) CollectionDocuments(Skip, Limit int64, sort int, key string, value interface{}) (result *mongo.Cursor, err error) {
	collection := app.Mgo.Database(app.Database).Collection(app.Collection)
	SORT := bson.D{{"_id", sort}}
	filter := bson.D{{key, value}}
	findOptions := options.Find().SetSort(SORT).SetLimit(Limit).SetSkip(Skip)
	result, err = collection.Find(context.Background(), filter, findOptions)
	return result, err
}

// AggregateByFiltersSort 统计分析
func (app *App) AggregateByFiltersSort(pipeline interface{}, opts ...*options.AggregateOptions) (result *mongo.Cursor, err error) {
	collection := app.Mgo.Database(app.Database).Collection(app.Collection)
	result, err = collection.Aggregate(context.TODO(), pipeline, opts...)
	return result, err
}

// CountDocumentsByFilters 统计数量
func (app *App) CountDocumentsByFilters(filter interface{}) (count int64, err error) {
	collection := app.Mgo.Database(app.Database).Collection(app.Collection)
	count, err = collection.CountDocuments(context.TODO(), filter)
	return count, err
}
