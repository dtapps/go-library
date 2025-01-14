package meituan

import (
	"go.dtapp.net/library/utils/gorequest"
)

func (c *Client) GetSecret() string {
	return c.config.secret
}

func (c *Client) SetSecret(v string) *Client {
	c.config.secret = v
	return c
}

func (c *Client) GetAppKey() string {
	return c.config.appKey
}

func (c *Client) SetAppKey(v string) *Client {
	c.config.appKey = v
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
