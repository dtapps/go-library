package dorm

import (
	"github.com/shenghui0779/yiigo"
)

func NewYiiGoMongoDbClient(config *ConfigYiiGoClient) (*YiiGoClient, error) {

	c := &YiiGoClient{config: config}

	yiigo.Init(
		yiigo.WithMongo(yiigo.Default, c.config.Dns),
	)

	c.MDb = yiigo.Mongo()

	return c, nil
}
