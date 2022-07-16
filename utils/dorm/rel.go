package dorm

import (
	"github.com/go-rel/rel"
)

type ConfigRelClient struct {
	Dns string // 地址
}

// RelClient
// https://go-rel.github.io/
type RelClient struct {
	Db     *rel.Repository  // 驱动
	config *ConfigRelClient // 配置
}
