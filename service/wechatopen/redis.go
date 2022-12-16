package wechatopen

import (
	"context"
	"time"
)

// 微信后台推送的ticke
func (c *Client) getComponentVerifyTicketCacheKeyName() string {
	return c.cache.componentVerifyTicketPrefix + c.GetComponentAppId()
}

// SetComponentVerifyTicket 设置微信后台推送的ticke
func (c *Client) SetComponentVerifyTicket(ctx context.Context, componentVerifyTicket string) string {
	if componentVerifyTicket == "" {
		return ""
	}
	c.cache.redisClient.Set(ctx, c.getComponentVerifyTicketCacheKeyName(), componentVerifyTicket, time.Hour*12)
	return c.GetComponentVerifyTicket(ctx)
}

// GetComponentVerifyTicket 获取微信后台推送的ticke
func (c *Client) GetComponentVerifyTicket(ctx context.Context) string {
	if c.cache.redisClient.Db == nil {
		return c.config.componentVerifyTicket
	}
	result, _ := c.cache.redisClient.Get(ctx, c.getComponentVerifyTicketCacheKeyName()).Result()
	return result
}

// 令牌
func (c *Client) getComponentAccessTokenCacheKeyName() string {
	return c.cache.componentAccessTokenPrefix + c.GetComponentAppId()
}

// SetComponentAccessToken 设置令牌
func (c *Client) SetComponentAccessToken(ctx context.Context, componentAccessToken string) string {
	if componentAccessToken == "" {
		return ""
	}
	c.cache.redisClient.Set(ctx, c.getComponentAccessTokenCacheKeyName(), componentAccessToken, time.Second*7200)
	return c.GetComponentAccessToken(ctx)
}

// GetComponentAccessToken 获取令牌
func (c *Client) GetComponentAccessToken(ctx context.Context) string {
	if c.cache.redisClient.Db == nil {
		return c.config.componentAccessToken
	}
	result, _ := c.cache.redisClient.Db.Get(ctx, c.getComponentAccessTokenCacheKeyName()).Result()
	return result
}

// MonitorComponentAccessToken 监控令牌
func (c *Client) MonitorComponentAccessToken(ctx context.Context) (string, error) {
	// 查询
	componentAccessToken := c.GetComponentAccessToken(ctx)
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
	return c.SetComponentAccessToken(ctx, resp.Result.ComponentAccessToken), err
}

// 授权方令牌
func (c *Client) getAuthorizerAccessTokenCacheKeyName() string {
	return c.cache.authorizerAccessTokenPrefix + c.GetComponentAppId() + ":" + c.GetAuthorizerAppid()
}

// SetAuthorizerAccessToken 设置授权方令牌
func (c *Client) SetAuthorizerAccessToken(ctx context.Context, authorizerAccessToken string) string {
	if authorizerAccessToken == "" {
		return ""
	}
	c.cache.redisClient.Set(ctx, c.getAuthorizerAccessTokenCacheKeyName(), authorizerAccessToken, time.Hour*2)
	return c.GetComponentAccessToken(ctx)
}

// GetAuthorizerAccessToken 获取授权方令牌
func (c *Client) GetAuthorizerAccessToken(ctx context.Context) string {
	if c.cache.redisClient.Db == nil {
		return c.config.authorizerAccessToken
	}
	result, _ := c.cache.redisClient.Get(ctx, c.getAuthorizerAccessTokenCacheKeyName()).Result()
	return result
}

// MonitorAuthorizerAccessToken 监控授权方令牌
func (c *Client) MonitorAuthorizerAccessToken(ctx context.Context, authorizerRefreshToken string) (string, error) {
	// 查询
	authorizerAccessToken := c.GetAuthorizerAccessToken(ctx)
	// 判断
	if authorizerAccessToken != "" {
		return authorizerAccessToken, nil
	}
	// 重新获取
	resp, err := c.CgiBinComponentApiAuthorizerToken(ctx, authorizerRefreshToken)
	return c.SetAuthorizerAccessToken(ctx, resp.Result.AuthorizerAccessToken), err
}

// 预授权码
func (c *Client) getPreAuthCodeCacheKeyName() string {
	return c.cache.preAuthCodePrefix + c.GetComponentAppId()
}

// SetPreAuthCode 设置预授权码
func (c *Client) SetPreAuthCode(ctx context.Context, preAuthCode string) string {
	if preAuthCode == "" {
		return ""
	}
	c.cache.redisClient.Set(ctx, c.getPreAuthCodeCacheKeyName(), preAuthCode, time.Second*1700)
	return c.GetComponentAccessToken(ctx)
}

// GetPreAuthCode 获取预授权码
func (c *Client) GetPreAuthCode(ctx context.Context) string {
	if c.cache.redisClient.Db == nil {
		return c.config.authorizerAccessToken
	}
	result, _ := c.cache.redisClient.Get(ctx, c.getPreAuthCodeCacheKeyName()).Result()
	return result
}

// DelPreAuthCode 删除预授权码
func (c *Client) DelPreAuthCode(ctx context.Context) error {
	return c.cache.redisClient.Del(ctx, c.getPreAuthCodeCacheKeyName()).Err()
}

// MonitorPreAuthCode 监控预授权码
func (c *Client) MonitorPreAuthCode(ctx context.Context) (string, error) {
	// 查询
	preAuthCode := c.GetPreAuthCode(ctx)
	// 判断
	if preAuthCode != "" {
		return preAuthCode, nil
	}
	// 重新获取
	resp, err := c.CgiBinComponentApiCreatePreAuthCoden(ctx)
	return c.SetPreAuthCode(ctx, resp.Result.PreAuthCode), err
}
