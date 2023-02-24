package dorm

import (
	"github.com/uptrace/bun"
)

type ConfigBunClient struct {
	Dns string // 地址
}

// BunClient
// https://bun.uptrace.dev/
type BunClient struct {
	Db     *bun.DB          // 驱动
	config *ConfigBunClient // 配置
}
