package dorm

import (
	"database/sql"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/pgdialect"
	"github.com/uptrace/bun/driver/pgdriver"
)

func NewBunPgsqlClient(config *ConfigBunClient) (*BunClient, error) {

	c := &BunClient{config: config}

	sqlDb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(c.config.Dns)))

	c.Db = bun.NewDB(sqlDb, pgdialect.New())

	return c, nil
}
