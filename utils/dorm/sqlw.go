package dorm

import (
	"github.com/lesismal/sqlw"
)

type ConfigSqlWClient struct {
	Dns string // 地址
}

// SqlWClient
// https://github.com/lesismal/sqlw
type SqlWClient struct {
	Db     *sqlw.DB          // 驱动
	config *ConfigSqlWClient // 配置
}
