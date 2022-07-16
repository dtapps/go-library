package dorm

import (
	"xorm.io/xorm"
)

type ConfigXormClient struct {
	Dns string // 地址
}

type XormClient struct {
	Db     *xorm.Engine      // 驱动
	config *ConfigXormClient // 配置
}
