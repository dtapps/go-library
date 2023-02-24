package dorm

import (
	"gitee.com/chunanyong/zorm"
)

type ConfigZormClient struct {
	Dns string // 地址
}

// ZormClient
// https://zorm.cn/
// https://www.yuque.com/u27016943/nrgi00
type ZormClient struct {
	db     *zorm.DBDao       // 驱动
	config *ConfigZormClient // 配置
}
