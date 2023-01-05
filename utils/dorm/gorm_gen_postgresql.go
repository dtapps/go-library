package dorm

import (
	"gorm.io/driver/postgres"
	"gorm.io/gen"
	"gorm.io/gorm"
)

func NewGormGenPostgresClient(config *GormGenClientConfig) (*GormGenClient, error) {

	c := &GormGenClient{config: config}

	c.Generator = gen.NewGenerator(gen.Config{
		OutPath: "../query",
		Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface, // generate mode
	})

	c.Db, _ = gorm.Open(postgres.Open(c.config.Dns), &gorm.Config{})
	c.Generator.UseDB(c.Db)

	return c, nil
}

func NewGormGenPostgresqlClient(config *GormGenClientConfig) (*GormGenClient, error) {

	c := &GormGenClient{config: config}

	c.Generator = gen.NewGenerator(gen.Config{
		OutPath: "../query",
		Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface, // generate mode
	})

	c.Db, _ = gorm.Open(postgres.Open(c.config.Dns), &gorm.Config{})
	c.Generator.UseDB(c.Db)

	return c, nil
}
