package dorm

import (
	"github.com/uptrace/bun"
)

// UptraceClientFun *UptraceClient 驱动
type UptraceClientFun func() *UptraceClient

// UptraceClientTableFun *UptraceClient 驱动
// string 表名
type UptraceClientTableFun func() (*UptraceClient, string)

type UptraceClientConfig struct {
	Dns string // 地址
}

// UptraceClient
// https://bun.uptrace.dev/
type UptraceClient struct {
	db     *bun.DB              // 驱动
	config *UptraceClientConfig // 配置
}
