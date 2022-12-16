package dorm

import (
	"xorm.io/xorm"
)

type XormClientConfigXorm struct {
	Dns string // 地址
}

// XormClient
// https://xorm.io/
type XormClient struct {
	Db     *xorm.Engine          // 驱动
	config *XormClientConfigXorm // 配置
}
