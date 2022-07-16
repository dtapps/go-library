package dorm

import (
	"github.com/beego/beego/v2/client/orm"
)

type ConfigBeegoClient struct {
	Dns string // 地址
}

// BeegoClient
// https://beego.vip/
type BeegoClient struct {
	Db     *orm.Ormer         // 驱动
	config *ConfigBeegoClient // 配置
}
