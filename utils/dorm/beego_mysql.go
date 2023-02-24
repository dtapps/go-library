package dorm

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	_ "github.com/go-sql-driver/mysql"
)

func NewBeegoMysqlClient(config *ConfigBeegoClient) (*BeegoClient, error) {

	var err error
	c := &BeegoClient{config: config}

	err = orm.RegisterDriver("mysql", orm.DRMySQL)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("加载驱动失败：%v", err))
	}

	var db *sql.DB
	o, err := orm.NewOrmWithDB("mysql", "default", db)
	if err != nil {
		return nil, err
	}
	c.Db = &o

	return c, nil
}
