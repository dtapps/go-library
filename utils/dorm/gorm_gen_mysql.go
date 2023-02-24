package dorm

import (
	"gorm.io/driver/mysql"
	"gorm.io/gen"
	"gorm.io/gorm"
)

func NewGormGenMysqlClient(config *GormGenClientConfig) (*GormGenClient, error) {

	c := &GormGenClient{config: config}

	c.generator = gen.NewGenerator(config.Config)

	c.db, _ = gorm.Open(mysql.Open(c.config.Dns), &gorm.Config{})
	c.generator.UseDB(c.db)

	return c, nil
}
