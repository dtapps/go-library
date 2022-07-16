package dorm

import (
	"github.com/beego/beego/v2/client/orm"
)

type ConfigFlyClient struct {
	Dns string // 地址
}

// FlyClient
// https://github.com/daodao97/fly
type FlyClient struct {
	Db     *orm.Ormer       // 驱动
	config *ConfigFlyClient // 配置
}
