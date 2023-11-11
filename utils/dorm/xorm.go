package dorm

import (
	"xorm.io/xorm"
)

// XormClientFun *XormClient 驱动
type XormClientFun func() *XormClient

// XormClientTableFun *XormClient 驱动
// string 表名
type XormClientTableFun func() (*XormClient, string)

type XormClientConfig struct {
	Dns string // 地址
}

// XormClient
// https://xorm.io/
type XormClient struct {
	db     *xorm.Engine      // 驱动
	config *XormClientConfig // 配置
}
