package aswzk

import (
	"go.dtapp.net/library/utils/gorequest"
)

func (c *Client) GetApiUrl() string {
	return c.config.apiUrl
}

func (c *Client) SetApiUrl(v string) *Client {
	c.config.apiUrl = v
	return c
}

func (c *Client) GetUserID() string {
	return c.config.userID
}

func (c *Client) SetUserID(v string) *Client {
	c.config.userID = v
	return c
}
func (c *Client) GetApiKey() string {
	return c.config.apiKey
}

func (c *Client) SetApiKey(v string) *Client {
	c.config.apiKey = v
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
