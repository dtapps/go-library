package wechatopen

import (
	"context"
	"fmt"
	"time"
)

// 微信后台推送的ticke
func (c *Client) getComponentVerifyTicketCacheKeyName() string {
	return fmt.Sprintf("wechat_open:component_verify_ticket:%v", c.config.ComponentAppId)
}

// SetComponentVerifyTicket 设置微信后台推送的ticke
func (c *Client) SetComponentVerifyTicket(ctx context.Context, componentVerifyTicket string) string {
	if componentVerifyTicket == "" {
		return ""
	}
	c.config.RedisClient.Set(ctx, c.getComponentVerifyTicketCacheKeyName(), componentVerifyTicket, time.Hour*12)
	return c.GetComponentVerifyTicket(ctx)
}

// GetComponentVerifyTicket 获取微信后台推送的ticke
func (c *Client) GetComponentVerifyTicket(ctx context.Context) string {
	if c.config.RedisClient.Db == nil {
		return c.config.ComponentVerifyTicket
	}
	result, _ := c.config.RedisClient.Get(ctx, c.getComponentVerifyTicketCacheKeyName()).Result()
	return result
}

// 令牌
func (c *Client) getComponentAccessTokenCacheKeyName() string {
	return fmt.Sprintf("wechat_open:component_access_token:%v", c.config.ComponentAppId)
}

// SetComponentAccessToken 设置令牌
func (c *Client) SetComponentAccessToken(ctx context.Context, componentAccessToken string) string {
	if componentAccessToken == "" {
		return ""
	}
	c.config.RedisClient.Set(ctx, c.getComponentAccessTokenCacheKeyName(), componentAccessToken, time.Second*7200)
	return c.GetComponentAccessToken(ctx)
}

// GetComponentAccessToken 获取令牌
func (c *Client) GetComponentAccessToken(ctx context.Context) string {
	if c.config.RedisClient.Db == nil {
		return c.config.ComponentAccessToken
	}
	result, _ := c.config.RedisClient.Db.Get(ctx, c.getComponentAccessTokenCacheKeyName()).Result()
	return result
}

// MonitorComponentAccessToken 监控令牌
func (c *Client) MonitorComponentAccessToken(ctx context.Context) string {
	// 查询
	componentAccessToken := c.GetComponentAccessToken(ctx)
	// 判断
	result := c.CgiBinGetApiDomainIp(ctx, componentAccessToken)
	if len(result.Result.IpList) > 0 {
		return componentAccessToken
	}
	// 重新获取
	return c.SetComponentAccessToken(ctx, c.CgiBinComponentApiComponentToken(ctx).Result.ComponentAccessToken)
}

// 授权方令牌
func (c *Client) getAuthorizerAccessTokenCacheKeyName() string {
	return fmt.Sprintf("wechat_open:authorizer_access_token:%v:%v", c.config.ComponentAppId, c.config.AuthorizerAppid)
}

// SetAuthorizerAccessToken 设置授权方令牌
func (c *Client) SetAuthorizerAccessToken(ctx context.Context, authorizerAccessToken string) string {
	if authorizerAccessToken == "" {
		return ""
	}
	c.config.RedisClient.Set(ctx, c.getAuthorizerAccessTokenCacheKeyName(), authorizerAccessToken, time.Hour*2)
	return c.GetComponentAccessToken(ctx)
}

// GetAuthorizerAccessToken 获取授权方令牌
func (c *Client) GetAuthorizerAccessToken(ctx context.Context) string {
	if c.config.RedisClient.Db == nil {
		return c.config.AuthorizerAccessToken
	}
	result, _ := c.config.RedisClient.Get(ctx, c.getAuthorizerAccessTokenCacheKeyName()).Result()
	return result
}

// MonitorAuthorizerAccessToken 监控授权方令牌
func (c *Client) MonitorAuthorizerAccessToken(ctx context.Context, authorizerRefreshToken string) string {
	// 查询
	authorizerAccessToken := c.GetAuthorizerAccessToken(ctx)
	// 判断
	if authorizerAccessToken != "" {
		return authorizerAccessToken
	}
	// 重新获取
	return c.SetAuthorizerAccessToken(ctx, c.CgiBinComponentApiAuthorizerToken(ctx, authorizerRefreshToken).Result.AuthorizerAccessToken)
}

// 预授权码
func (c *Client) getPreAuthCodeCacheKeyName() string {
	return fmt.Sprintf("wechat_open:pre_auth_code:%v", c.config.ComponentAppId)
}

// SetPreAuthCode 设置预授权码
func (c *Client) SetPreAuthCode(ctx context.Context, preAuthCode string) string {
	if preAuthCode == "" {
		return ""
	}
	c.config.RedisClient.Set(ctx, c.getPreAuthCodeCacheKeyName(), preAuthCode, time.Second*1700)
	return c.GetComponentAccessToken(ctx)
}

// GetPreAuthCode 获取预授权码
func (c *Client) GetPreAuthCode(ctx context.Context) string {
	if c.config.RedisClient.Db == nil {
		return c.config.AuthorizerAccessToken
	}
	result, _ := c.config.RedisClient.Get(ctx, c.getPreAuthCodeCacheKeyName()).Result()
	return result
}

// DelPreAuthCode 删除预授权码
func (c *Client) DelPreAuthCode(ctx context.Context) error {
	return c.config.RedisClient.Del(ctx, c.getPreAuthCodeCacheKeyName()).Err()
}

// MonitorPreAuthCode 监控预授权码
func (c *Client) MonitorPreAuthCode(ctx context.Context) string {
	// 查询
	preAuthCode := c.GetPreAuthCode(ctx)
	// 判断
	if preAuthCode != "" {
		return preAuthCode
	}
	// 重新获取
	return c.SetPreAuthCode(ctx, c.CgiBinComponentApiCreatePreAuthCoden(ctx).Result.PreAuthCode)
}
