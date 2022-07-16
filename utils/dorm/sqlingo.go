package dorm

import (
	"github.com/lqs/sqlingo"
)

type ConfigSqLiNgoClient struct {
	Dns string // 地址
}

// SqLiNgoClient
// https://github.com/lqs/sqlingo
type SqLiNgoClient struct {
	Db     sqlingo.Database     // 驱动
	config *ConfigSqLiNgoClient // 配置
}
