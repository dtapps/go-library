package dorm

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/sqlitedialect"
	"github.com/uptrace/bun/driver/sqliteshim"
)

func NewBunSqliteClient(config *ConfigBunClient) (*BunClient, error) {

	var err error
	c := &BunClient{config: config}

	sqlDb, err := sql.Open(sqliteshim.ShimName, c.config.Dns)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("加载驱动失败：%v", err))
	}

	c.Db = bun.NewDB(sqlDb, sqlitedialect.New())

	return c, nil
}
