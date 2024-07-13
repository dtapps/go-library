package wechatqy

import (
	"go.dtapp.net/library/utils/gorequest"
)

func (c *Client) GetAppId() string {
	return c.config.appId
}

func (c *Client) SetAppId(v string) *Client {
	c.config.appId = v
	return c
}

func (c *Client) GetAgentId() int {
	return c.config.agentId
}

func (c *Client) SetAgentId(v int) *Client {
	c.config.agentId = v
	return c
}

func (c *Client) GetSecret() string {
	return c.config.secret
}

func (c *Client) SetSecret(v string) *Client {
	c.config.secret = v
	return c
}

func (c *Client) GetRedirectUri() string {
	return c.config.redirectUri
}

func (c *Client) SetRedirectUri(v string) *Client {
	c.config.redirectUri = v
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
