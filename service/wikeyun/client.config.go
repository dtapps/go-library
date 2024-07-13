package wikeyun

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

func (c *Client) GetStoreId() int64 {
	return c.config.storeId
}

func (c *Client) SetStoreId(v int64) *Client {
	c.config.storeId = v
	return c
}

func (c *Client) GetAppKey() int64 {
	return c.config.appKey
}

func (c *Client) SetAppKey(v int64) *Client {
	c.config.appKey = v
	return c
}

func (c *Client) GetAppSecret() string {
	return c.config.appSecret
}

func (c *Client) SetAppSecret(v string) *Client {
	c.config.appSecret = v
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
