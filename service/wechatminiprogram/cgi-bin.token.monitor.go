package wechatminiprogram

import (
	"context"
	"errors"
	"time"
)

func (c *Client) GetAccessTokenMonitor() (string, error) {
	if c.config.RedisClient.Db == nil {
		return "", errors.New("驱动没有初始化")
	}
	result := c.GetCallBackIp()
	if len(result.Result.IpList) <= 0 {
		token := c.CgiBinToken()
		c.config.RedisClient.Db.Set(context.Background(), c.getAccessTokenCacheKeyName(), token.Result.AccessToken, time.Second*7000)
		return token.Result.AccessToken, nil
	}
	return c.config.AccessToken, nil
}
