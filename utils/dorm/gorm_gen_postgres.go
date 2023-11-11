package dorm

import (
	"gorm.io/driver/postgres"
	"gorm.io/gen"
	"gorm.io/gorm"
)

func NewGormGenPostgresClient(config *GormGenClientConfig) (*GormGenClient, error) {

	c := &GormGenClient{config: config}

	c.generator = gen.NewGenerator(config.Config)

	if c.config.Dns != "" {
		c.db, _ = gorm.Open(postgres.Open(c.config.Dns), &gorm.Config{})
		c.generator.UseDB(c.db)
	} else {
		c.generator.UseDB(c.config.Db)
	}

	return c, nil
}
