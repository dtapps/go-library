package wechatopen

import (
	"context"
	"time"
)

// GetPreAuthCodeCacheKeyName 第三方平台预授权码 缓存名称
func GetPreAuthCodeCacheKeyName(ctx context.Context, c *Client) string {
	return c.cache.preAuthCodePrefix + c.GetComponentAppId(ctx)
}

// GetPreAuthCode 第三方平台预授权码 获取
func GetPreAuthCode(ctx context.Context, c *Client) string {
	if c.cache.redisClient.GetDb() == nil {
		return c.config.authorizerAccessToken
	}
	result, _ := c.cache.redisClient.Get(ctx, GetPreAuthCodeCacheKeyName(ctx, c)).Result()
	return result
}

// SetPreAuthCode 第三方平台预授权码 设置
func SetPreAuthCode(ctx context.Context, c *Client, preAuthCode string) string {
	if preAuthCode == "" {
		return ""
	}
	c.cache.redisClient.Set(ctx, GetPreAuthCodeCacheKeyName(ctx, c), preAuthCode, time.Second*1700)
	return GetPreAuthCode(ctx, c)
}

// DelPreAuthCode 第三方平台预授权码 删除
func DelPreAuthCode(ctx context.Context, c *Client) error {
	return c.cache.redisClient.Del(ctx, GetPreAuthCodeCacheKeyName(ctx, c)).Err()
}

// MonitorPreAuthCode 第三方平台预授权码 监控
func MonitorPreAuthCode(ctx context.Context, c *Client) (string, error) {
	// 查询
	preAuthCode := GetPreAuthCode(ctx, c)
	// 判断
	if preAuthCode != "" {
		return preAuthCode, nil
	}
	// 重新获取
	resp, err := c.CgiBinComponentApiCreatePreAuthCoden(ctx)
	return SetPreAuthCode(ctx, c, resp.Result.PreAuthCode), err
}
