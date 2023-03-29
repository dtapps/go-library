package golog

import (
	"context"
	"errors"
	"github.com/dtapps/go-library/utils/dorm"
	"github.com/dtapps/go-library/utils/goip"
)

type GinCustomClient struct {
	gormClient *dorm.GormClient // 数据库驱动
	ipService  *goip.Client     // ip服务
	config     struct {
		systemHostname  string // 主机名
		systemOs        string // 系统类型
		systemKernel    string // 系统内核
		systemInsideIp  string // 内网ip
		systemOutsideIp string // 外网ip
		goVersion       string // go版本
		sdkVersion      string // sdk版本
	}
	gormConfig struct {
		stats     bool   // 状态
		tableName string // 表名
	}
}

type ConfigGinCustomClient struct {
	IpService     *goip.Client            // ip服务
	GormClientFun dorm.GormClientTableFun // 日志配置
}

func NewGinCustomClient(config *ConfigGinCustomClient) (*GinCustomClient, error) {

	var ctx = context.Background()

	c := &GinCustomClient{}

	c.config.systemOutsideIp = goip.IsIp(c.config.systemOutsideIp)
	if c.config.systemOutsideIp == "" {
		return nil, currentIpNoConfig
	}

	c.ipService = config.IpService

	// 配置信息
	c.setConfig(ctx)

	gormClient, gormTableName := config.GormClientFun()

	if gormClient == nil || gormClient.GetDb() == nil {
		return nil, dbClientFunNoConfig
	}

	// 配置关系数据库
	if gormClient != nil || gormClient.GetDb() != nil {

		c.gormClient = gormClient

		if gormTableName == "" {
			return nil, errors.New("没有设置表名")
		} else {
			c.gormConfig.tableName = gormTableName
		}

		err := c.autoMigrate(ctx)
		if err != nil {
			return nil, err
		}

		c.gormConfig.stats = true
	}

	return c, nil
}
