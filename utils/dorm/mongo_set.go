package dorm

func (c *MongoClient) setDatabaseName(databaseName string) {
	c.config.DatabaseName = databaseName
}

func (c *MongoClient) setCollectionName(collectionName string) {
	c.collectionName = collectionName
}
