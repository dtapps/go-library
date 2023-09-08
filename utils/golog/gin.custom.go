package golog

import (
	"context"
	"github.com/dtapps/go-library/utils/goip"
)

type GinCustomClient struct {
	ipService *goip.Client // IP服务
	config    struct {
		systemHostname  string // 主机名
		systemOs        string // 系统类型
		systemKernel    string // 系统内核
		systemInsideIp  string // 内网ip
		systemOutsideIp string // 外网ip
		goVersion       string // go版本
		sdkVersion      string // sdk版本
	}
	slog struct {
		status bool  // 状态
		client *SLog // 日志服务
	}
}

type ConfigGinCustomClient struct {
	IpService *goip.Client // IP服务
	CurrentIp string       // 当前IP
}

func NewGinCustomClient(ctx context.Context, config *ConfigGinCustomClient) (*GinCustomClient, error) {

	c := &GinCustomClient{}

	c.config.systemOutsideIp = config.CurrentIp

	c.config.systemOutsideIp = goip.IsIp(c.config.systemOutsideIp)
	if c.config.systemOutsideIp == "" {
		return nil, currentIpNoConfig
	}

	c.ipService = config.IpService

	// 配置信息
	c.setConfig(ctx)

	return c, nil
}
