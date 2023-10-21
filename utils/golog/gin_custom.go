package golog

import (
	"context"
	"github.com/dtapps/go-library/utils/goip"
)

func NewGinCustomClient(ctx context.Context, config *GinCustomClientConfig) (*GinCustomClient, error) {

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
