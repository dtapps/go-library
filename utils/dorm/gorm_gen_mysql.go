package dorm

import (
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
)

func NewGormGenMysqlClient(config *GormGenClientConfig) (*GormGenClient, error) {

	c := &GormGenClient{config: config}

	c.Generator = gen.NewGenerator(gen.Config{
		OutPath: "../query",
		Mode:    gen.WithoutContext | gen.WithDefaultQuery | gen.WithQueryInterface, // generate mode
	})

	c.Db, _ = gorm.Open(mysql.Open(c.config.Dns), &gorm.Config{})
	c.Generator.UseDB(c.Db)

	return c, nil
}
