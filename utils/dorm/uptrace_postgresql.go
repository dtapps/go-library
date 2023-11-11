package dorm

//import (
//	"context"
//	"database/sql"
//	"github.com/uptrace/bun"
//	"github.com/uptrace/bun/dialect/pgdialect"
//	"github.com/uptrace/bun/driver/pgdriver"
//)
//
//func NewUptracePgsqlClient(ctx context.Context, config *UptraceClientConfig) (*UptraceClient, error) {
//
//	c := &UptraceClient{config: config}
//
//	sqlDb := sql.OpenDB(pgdriver.NewConnector(pgdriver.WithDSN(c.config.Dns)))
//
//	c.db = bun.NewDB(sqlDb, pgdialect.New())
//
//	return c, nil
//}
