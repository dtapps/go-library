package dorm

import (
	"context"
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
)

func NewGormGenMysqlClient(ctx context.Context, config *GormGenClientConfig) (*GormGenClient, error) {

	c := &GormGenClient{config: config}

	c.generator = gen.NewGenerator(config.Config)

	if c.config.Dns != "" {
		c.db, _ = gorm.Open(mysql.Open(c.config.Dns), &gorm.Config{})
		c.generator.UseDB(c.db)
	} else {
		c.generator.UseDB(c.config.Db)
	}

	return c, nil
}
