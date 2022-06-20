package dorm

import (
	"context"
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	"reflect"
)

// Database 设置库名
func (c *MongoClient) Database(databaseName string) *MongoClient {
	c.setDatabaseName(databaseName)
	return c
}

// Collection 设置表名
func (c *MongoClient) Collection(collectionName string) *MongoClient {
	c.setCollectionName(collectionName)
	return c
}

// Model 传入模型自动获取库名和表名
func (c *MongoClient) Model(value interface{}) *MongoClient {
	// https://studygolang.com/articles/896
	val := reflect.ValueOf(value)
	if methodValue := val.MethodByName("Database"); methodValue.IsValid() {
		c.setDatabaseName(methodValue.Call(nil)[0].String())
	}
	if methodValue := val.MethodByName("TableName"); methodValue.IsValid() {
		c.setCollectionName(methodValue.Call(nil)[0].String())
	}
	return c
}

// CreateResult 返回查询结果
type CreateResult struct {
	InsertedID  interface{}   // 创建一条记录的ID
	InsertedIDs []interface{} // 创建多条记录的ID
}

// Create 创建数据
func (c *MongoClient) Create(values ...interface{}) (CreateResult, error) {
	collection := c.Db.Database(c.getDatabaseName()).Collection(c.collectionName)

	const (
		insertTypeOne  = "one"
		insertTypeMany = "many"
	)

	var (
		insertType     string
		insertDataOne  interface{}
		insertDataMany []interface{}
	)

	for _, value := range values {
		switch v := value.(type) {
		case map[string]interface{}:
		case []map[string]interface{}:
		case map[string]string:
		case []map[string]string:
		default:
			sliceValue := reflect.Indirect(reflect.ValueOf(value))
			if sliceValue.Kind() == reflect.Slice {
				insertType = insertTypeMany
				size := sliceValue.Len()
				for i := 0; i < size; i++ {
					sv := sliceValue.Index(i)                          // 取出第i个元素
					elemValue := sv.Interface()                        // 原始数据
					insertDataMany = append(insertDataMany, elemValue) // 加入到数组中
				}
			} else {
				insertType = insertTypeOne
				insertDataOne = v
			}
		}
	}

	if insertType == insertTypeOne {
		result, err := collection.InsertOne(context.TODO(), insertDataOne)
		return CreateResult{InsertedID: result.InsertedID}, err
	} else if insertType == insertTypeMany {
		result, err := collection.InsertMany(context.TODO(), insertDataMany)
		return CreateResult{InsertedIDs: result.InsertedIDs}, err
	} else {
		return CreateResult{}, errors.New("values is empty")
	}
}

// 查询条件
type queryFilter struct {
	Key   string
	Value interface{}
}

// Where 条件
func (c *MongoClient) Where(key string, value interface{}) *MongoClient {
	log.Println("key", key)
	log.Println("value", value)
	c.filterArr = append(c.filterArr, queryFilter{key, value})
	c.filter = bson.D{{key, value}}
	return c
}

// QueryResult 返回查询结果
type QueryResult struct {
	RowsAffected int   // 返回找到的记录数
	Error        error // 错误信息
}

// First 获取第一条记录（主键升序）
func (c *MongoClient) First() *QueryResult {
	return &QueryResult{}
}

// Take 获取一条记录，没有指定排序字段
func (c *MongoClient) Take(v interface{}) *QueryResult {
	collection := c.Db.Database(c.getDatabaseName()).Collection(c.collectionName)
	//log.Printf("c.filterArr：%s\n", c.filterArr)
	//log.Printf("c.filterArr：%v\n", c.filterArr)
	//log.Printf("c.filterArr：%+v\n", c.filterArr)
	//log.Printf("c.filter：%s\n", c.filter)
	//log.Printf("c.filter：%v\n", c.filter)
	//log.Printf("c.filter：%+v\n", c.filter)
	err := collection.FindOne(context.TODO(), c.filter).Decode(v)
	return &QueryResult{1, err}
}

// Last 获取最后一条记录（主键降序）
func (c *MongoClient) Last() *QueryResult {
	return &QueryResult{}
}

// Find 获取多条记录
func (c *MongoClient) Find(v interface{}) *QueryResult {
	collection := c.Db.Database(c.getDatabaseName()).Collection(c.collectionName)
	log.Printf("c.filterArr：%s\n", c.filterArr)
	log.Printf("c.filterArr：%v\n", c.filterArr)
	log.Printf("c.filterArr：%+v\n", c.filterArr)
	log.Printf("c.filter：%s\n", c.filter)
	log.Printf("c.filter：%v\n", c.filter)
	log.Printf("c.filter：%+v\n", c.filter)
	cursor, err := collection.Find(context.TODO(), c.filter)
	if err != nil {
		return &QueryResult{0, err}
	}

	// 结果遍历和赋值
	err = cursor.All(context.TODO(), v)

	return &QueryResult{cursor.RemainingBatchLength(), err}
}
