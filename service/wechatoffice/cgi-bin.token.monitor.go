package wechatoffice

import (
	"context"
	"errors"
	"time"
)

func (c *Client) GetAccessTokenMonitor(ctx context.Context) (string, error) {
	if c.cache.redisClient.GetDb() == nil {
		return "", errors.New("驱动没有初始化")
	}
	result := c.GetCallBackIp(ctx)
	if len(result.Result.IpList) > 0 {
		return c.config.accessToken, nil
	}
	token := c.CgiBinToken(ctx)
	c.cache.redisClient.Set(ctx, c.getAccessTokenCacheKeyName(), token.Result.AccessToken, time.Second*7000)
	return token.Result.AccessToken, nil
}
