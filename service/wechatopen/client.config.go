package wechatopen

import (
	"go.dtapp.net/library/utils/gorequest"
)

func (c *Client) GetComponentAppId() string {
	return c.config.componentAppId
}

func (c *Client) SetComponentAppId(v string) *Client {
	c.config.componentAppId = v
	return c
}

func (c *Client) GetComponentAppSecret() string {
	return c.config.componentAppSecret
}

func (c *Client) SetComponentAppSecret(v string) *Client {
	c.config.componentAppSecret = v
	return c
}

func (c *Client) GetMessageToken() string {
	return c.config.messageToken
}

func (c *Client) SetMessageToken(v string) *Client {
	c.config.messageToken = v
	return c
}

func (c *Client) GetMessageKey() string {
	return c.config.messageKey
}

func (c *Client) SetMessageKey(v string) *Client {
	c.config.messageKey = v
	return c
}

func (c *Client) GetComponentAccessToken() string {
	return c.config.componentAccessToken
}

func (c *Client) SetComponentAccessToken(v string) *Client {
	c.config.componentAccessToken = v
	return c
}

func (c *Client) GetComponentVerifyTicket() string {
	return c.config.componentVerifyTicket
}

func (c *Client) SetComponentVerifyTicket(v string) *Client {
	c.config.componentVerifyTicket = v
	return c
}

func (c *Client) GetComponentPreAuthCode() string {
	return c.config.componentPreAuthCode
}

func (c *Client) SetComponentPreAuthCode(v string) *Client {
	c.config.componentPreAuthCode = v
	return c
}

func (c *Client) GetAuthorizerAppid() string {
	return c.config.authorizerAppid
}

func (c *Client) SetAuthorizerAppid(v string) *Client {
	c.config.authorizerAppid = v
	return c
}

func (c *Client) GetAuthorizerAccessToken() string {
	return c.config.authorizerAccessToken
}

func (c *Client) SetAuthorizerAccessToken(v string) *Client {
	c.config.authorizerAccessToken = v
	return c
}

func (c *Client) GetAuthorizerRefreshToken() string {
	return c.config.authorizerRefreshToken
}

func (c *Client) SetAuthorizerRefreshToken(v string) *Client {
	c.config.authorizerRefreshToken = v
	return c
}

// SetClientIP 配置
func (c *Client) SetClientIP(clientIP string) *Client {
	c.clientIP = clientIP
	if c.httpClient != nil {
		c.httpClient.SetClientIP(clientIP)
	}
	return c
}

// SetLogFun 设置日志记录函数
func (c *Client) SetLogFun(logFun gorequest.LogFunc) {
	if c.httpClient != nil {
		c.httpClient.SetLogFunc(logFun)
	}
}
