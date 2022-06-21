package dorm

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
)

type FindManyResult struct {
	cursor *mongo.Cursor
	err    error
}

func (f *FindManyResult) Many(result interface{}) error {
	return f.cursor.All(context.TODO(), result)
}
