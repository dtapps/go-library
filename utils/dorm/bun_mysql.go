package dorm

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/mysqldialect"
)

func NewBunMysqlClient(config *ConfigBunClient) (*BunClient, error) {

	var err error
	c := &BunClient{config: config}

	sqlDb, err := sql.Open("mysql", c.config.Dns)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("加载驱动失败：%v", err))
	}

	c.Db = bun.NewDB(sqlDb, mysqldialect.New())

	return c, nil
}
