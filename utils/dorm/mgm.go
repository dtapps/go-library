package dorm

import (
	"github.com/kamva/mgm/v3"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ConfigMgmClient struct {
	Opts         *options.ClientOptions
	DatabaseName string // 库名
}

type MgmClient struct {
	config *ConfigMgmClient // 配置
}

func NewMgmClient(config *ConfigMgmClient) (*MgmClient, error) {

	c := &MgmClient{config: config}

	_ = mgm.SetDefaultConfig(nil, c.config.DatabaseName, c.config.Opts)

	return c, nil
}
