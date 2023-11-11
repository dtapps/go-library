package dorm

import (
	"github.com/dtapps/go-library/utils/dorm/ent"
)

// EntClientFun *EntClient 驱动
type EntClientFun func() *EntClient

// EntClientTableFun *EntClient 驱动
// string 表名
type EntClientTableFun func() (*EntClient, string)

type EntClientConfig struct {
	Dns           string // 地址
	AutoMigration bool   // 自动迁移
}

// EntClient
// https://entgo.io/zh/docs/getting-started
// https://ent.ryansu.tech/#/zh-cn/getting-started
type EntClient struct {
	db     *ent.Client      // 驱动
	config *EntClientConfig // 配置
}
