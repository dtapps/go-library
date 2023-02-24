package dorm

import (
	"gorm.io/driver/postgres"
	"gorm.io/gen"
	"gorm.io/gorm"
)

func NewGormGenPostgresqlClient(config *GormGenClientConfig) (*GormGenClient, error) {

	c := &GormGenClient{config: config}

	c.generator = gen.NewGenerator(config.Config)

	c.db, _ = gorm.Open(postgres.Open(c.config.Dns), &gorm.Config{})
	c.generator.UseDB(c.db)

	return c, nil
}
