package dorm

// 获取库名
func (ms *MongoTransaction) getDatabaseName() string {
	return ms.databaseName
}

// 获取表名
func (ms *MongoTransaction) getCollectionName() string {
	return ms.collectionName
}
