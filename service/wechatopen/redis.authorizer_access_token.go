package wechatopen

import (
	"context"
	"errors"
	"fmt"
	"time"
)

// GetAuthorizerAccessTokenCacheKeyName 授权方access_token 缓存名称
func GetAuthorizerAccessTokenCacheKeyName(ctx context.Context, c *Client) string {
	return c.cache.authorizerAccessTokenPrefix + c.GetComponentAppId(ctx) + ":" + c.GetAuthorizerAppid(ctx)
}

// GetAuthorizerAccessToken 授权方access_token
func GetAuthorizerAccessToken(ctx context.Context, c *Client) string {
	if c.cache.redisClient.GetDb() == nil {
		return c.config.authorizerAccessToken
	}
	result, _ := c.cache.redisClient.Get(ctx, GetAuthorizerAccessTokenCacheKeyName(ctx, c)).Result()
	return result
}

// SetAuthorizerAccessToken 授权方access_token
func SetAuthorizerAccessToken(ctx context.Context, c *Client, authorizerAccessToken string) string {
	if authorizerAccessToken == "" {
		return ""
	}
	c.cache.redisClient.Set(ctx, GetAuthorizerAccessTokenCacheKeyName(ctx, c), authorizerAccessToken, time.Hour*2)
	return GetAuthorizerAccessToken(ctx, c)
}

// DelAuthorizerAccessToken 授权方access_token 删除
func DelAuthorizerAccessToken(ctx context.Context, c *Client) error {
	return c.cache.redisClient.Del(ctx, GetAuthorizerAccessTokenCacheKeyName(ctx, c)).Err()
}

// MonitorAuthorizerAccessToken 授权方access_token 监控
func MonitorAuthorizerAccessToken(ctx context.Context, c *Client, authorizerRefreshToken string) (string, error) {
	// 查询
	authorizerAccessToken := GetAuthorizerAccessToken(ctx, c)
	// 判断
	if authorizerAccessToken != "" {
		return authorizerAccessToken, nil
	}
	// 重新获取
	resp, err := c.CgiBinComponentApiAuthorizerToken(ctx, authorizerRefreshToken)
	if resp.Result.AuthorizerRefreshToken == "" {
		return authorizerAccessToken, errors.New(fmt.Sprintf("获取失败：%v", err))
	}
	return SetAuthorizerAccessToken(ctx, c, resp.Result.AuthorizerAccessToken), err
}
