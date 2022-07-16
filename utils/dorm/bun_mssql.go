package dorm

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/denisenkom/go-mssqldb"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/mssqldialect"
)

func NewBunMssqlClient(config *ConfigBunClient) (*BunClient, error) {

	var err error
	c := &BunClient{config: config}

	sqlDb, err := sql.Open("sqlserver", c.config.Dns)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("加载驱动失败：%v", err))
	}

	c.Db = bun.NewDB(sqlDb, mssqldialect.New())

	return c, nil
}
