package dorm

import "go.mongodb.org/mongo-driver/mongo"

type FindOneResult struct {
	singleResult *mongo.SingleResult
}

func (f *FindOneResult) One(result interface{}) error {
	return f.singleResult.Decode(result)
}
