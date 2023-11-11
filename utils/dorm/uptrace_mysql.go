package dorm

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/uptrace/bun"
	"github.com/uptrace/bun/dialect/mysqldialect"
)

const UptraceMysqlDriver = "mysql"

func NewUptraceMysqlClient(ctx context.Context, config *UptraceClientConfig) (*UptraceClient, error) {

	var err error
	c := &UptraceClient{config: config}

	sqlDb, err := sql.Open(UptraceMysqlDriver, c.config.Dns)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("加载驱动失败：%v", err))
	}

	c.db = bun.NewDB(sqlDb, mysqldialect.New())

	return c, nil
}
