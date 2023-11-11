package dorm

import (
	"github.com/dtapps/go-library/utils/dorm/ent"
	_ "github.com/go-sql-driver/mysql"
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
// https://Ent.io/
type EntClient struct {
	db     *ent.Client      // 驱动
	config *EntClientConfig // 配置
}
