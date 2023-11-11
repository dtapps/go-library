package dorm

import (
	"gitee.com/chunanyong/zorm"
)

// ZormClientFun *ZormClient 驱动
type ZormClientFun func() *ZormClient

// ZormClientTableFun *ZormClient 驱动
// string 表名
type ZormClientTableFun func() (*ZormClient, string)

type ZormClientConfig struct {
	Dns string // 地址
}

// ZormClient
// https://zorm.cn/
// https://www.yuque.com/u27016943/nrgi00
type ZormClient struct {
	db     *zorm.DBDao       // 驱动
	config *ZormClientConfig // 配置
}
