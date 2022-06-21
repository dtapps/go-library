package dorm

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
)

type FindResult struct {
	cursor *mongo.Cursor
	err    error
}

func (f *FindResult) Many(result interface{}) error {
	return f.cursor.All(context.TODO(), result)
}
