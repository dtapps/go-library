package dorm

//import (
//	"context"
//	"errors"
//	"fmt"
//	_ "github.com/lib/pq"
//	"xorm.io/xorm"
//)
//
//func NewXormPostgresqlClient(ctx context.Context, config *XormClientConfig) (*XormClient, error) {
//
//	var err error
//	c := &XormClient{config: config}
//
//	c.db, err = xorm.NewEngine(XormPostgresDriver, c.config.Dns)
//	if err != nil {
//		return nil, errors.New(fmt.Sprintf("连接失败：%v", err))
//	}
//
//	return c, nil
//}
