package dorm

import (
	"context"
	"errors"
	"fmt"
	"github.com/dtapps/go-library/utils/dorm/ent"
)

const EntPostgresDriver = "postgres"

func NewEntPostgresClient(ctx context.Context, config *EntClientConfig) (*EntClient, error) {

	var err error
	c := &EntClient{config: config}

	c.db, err = ent.Open(EntPostgresDriver, c.config.Dns)
	if err != nil {
		return nil, errors.New(fmt.Sprintf("连接失败：%v", err))
	}
	defer c.db.Close()

	// 运行自动迁移工具
	if c.config.AutoMigration {
		if err = c.db.Schema.Create(ctx); err != nil {
			return nil, errors.New(fmt.Sprintf("创建架构失败：%v", err))
		}
	}

	return c, nil
}
