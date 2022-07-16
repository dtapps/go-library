package dorm

import (
	"errors"
	"fmt"
	"github.com/upper/db/v4/adapter/postgresql"
)

func NewUpperPostgresqlClient(settings postgresql.ConnectionURL) (*UpperClient, error) {

	var err error
	c := &UpperClient{}

	sess, err := postgresql.Open(settings)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("连接失败：%v", err))
	}
	defer sess.Close()

	c.Db = &sess

	return c, nil
}
