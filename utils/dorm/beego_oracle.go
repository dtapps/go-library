package dorm

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/beego/beego/v2/client/orm"
	_ "github.com/lib/pq"
)

func NewBeegoOracleClient(config *ConfigBeegoClient) (*BeegoClient, error) {

	var err error
	c := &BeegoClient{config: config}

	err = orm.RegisterDriver("oracle", orm.DROracle)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("加载驱动失败：%v", err))
	}

	var db *sql.DB
	o, err := orm.NewOrmWithDB("oracle", "default", db)
	if err != nil {
		return nil, err
	}
	c.Db = &o

	return c, nil
}
