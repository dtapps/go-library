package wechatopen

import (
	"context"
	"time"
)

// GetComponentAccessTokenCacheKeyName 第三方平台access_token 缓存名称
func GetComponentAccessTokenCacheKeyName(ctx context.Context, c *Client) string {
	return c.cache.componentAccessTokenPrefix + c.GetComponentAppId(ctx)
}

// GetComponentAccessToken 第三方平台access_token 获取
func GetComponentAccessToken(ctx context.Context, c *Client) string {
	if c.cache.redisClient.GetDb() == nil {
		return c.config.componentAccessToken
	}
	result, _ := c.cache.redisClient.GetDb().Get(ctx, GetComponentAccessTokenCacheKeyName(ctx, c)).Result()
	return result
}

// SetComponentAccessToken 第三方平台access_token 设置
func SetComponentAccessToken(ctx context.Context, c *Client, componentAccessToken string) string {
	if componentAccessToken == "" {
		return ""
	}
	c.cache.redisClient.Set(ctx, GetComponentAccessTokenCacheKeyName(ctx, c), componentAccessToken, time.Second*7200)
	return GetComponentAccessToken(ctx, c)
}

// DelComponentAccessToken 第三方平台access_token 删除
func DelComponentAccessToken(ctx context.Context, c *Client) error {
	return c.cache.redisClient.Del(ctx, GetComponentAccessTokenCacheKeyName(ctx, c)).Err()
}

// MonitorComponentAccessToken 第三方平台access_token 监控
func MonitorComponentAccessToken(ctx context.Context, c *Client) (string, error) {
	// 查询
	componentAccessToken := GetComponentAccessToken(ctx, c)
	// 判断
	result, err := c.CgiBinGetApiDomainIp(ctx, componentAccessToken)
	if err != nil {
		return "", err
	}
	if len(result.Result.IpList) > 0 {
		return componentAccessToken, err
	}
	// 重新获取
	resp, err := c.CgiBinComponentApiComponentToken(ctx)
	return SetComponentAccessToken(ctx, c, resp.Result.ComponentAccessToken), err
}
