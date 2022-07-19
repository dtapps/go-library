package dorm

import (
	"github.com/shenghui0779/yiigo"
	"time"
)

func NewYiiGoMysqlClient(config *ConfigYiiGoClient) (*YiiGoClient, error) {

	c := &YiiGoClient{config: config}

	yiigo.Init(
		yiigo.WithMySQL(yiigo.Default, &yiigo.DBConfig{
			DSN: c.config.Dns,
			Options: &yiigo.DBOptions{
				MaxOpenConns:    20,
				MaxIdleConns:    10,
				ConnMaxLifetime: 10 * time.Minute,
				ConnMaxIdleTime: 5 * time.Minute,
			},
		}),
	)

	c.Db = yiigo.DB()

	return c, nil
}
