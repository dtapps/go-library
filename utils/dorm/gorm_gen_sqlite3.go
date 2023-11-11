package dorm

import (
	// "gorm.io/driver/sqlite"
	"context"
	"github.com/glebarez/sqlite"
	"gorm.io/gen"
	"gorm.io/gorm"
)

func NewGormGenSqlite3Client(ctx context.Context, config *GormGenClientConfig) (*GormGenClient, error) {

	c := &GormGenClient{config: config}

	c.generator = gen.NewGenerator(config.Config)

	if c.config.Dns != "" {
		c.db, _ = gorm.Open(sqlite.Open(c.config.Dns), &gorm.Config{})
		c.generator.UseDB(c.db)
	} else {
		c.generator.UseDB(c.config.Db)
	}

	return c, nil
}
