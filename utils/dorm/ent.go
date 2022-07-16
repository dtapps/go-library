package dorm

import (
	"entgo.io/ent"
)

type ConfigEntClient struct {
	Dns string // 地址
}

// EntClient
// https://entgo.io/
type EntClient struct {
	Db     *ent.Edge        // 驱动
	config *ConfigEntClient // 配置
}
