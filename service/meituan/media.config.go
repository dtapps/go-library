package meituan

import (
	"go.dtapp.net/library/utils/gorequest"
)

func (c *MediaClient) GetAppKey() string {
	return c.config.appKey
}

func (c *MediaClient) SetAppKey(v string) *MediaClient {
	c.config.appKey = v
	return c
}

func (c *MediaClient) GetAppSecret() string {
	return c.config.appSecret
}

func (c *MediaClient) SetAppSecret(v string) *MediaClient {
	c.config.appSecret = v
	return c
}

// SetClientIP 配置
func (c *MediaClient) SetClientIP(clientIP string) *MediaClient {
	c.clientIP = clientIP
	if c.httpClient != nil {
		c.httpClient.SetClientIP(clientIP)
	}
	return c
}

// SetLogFun 设置日志记录函数
func (c *MediaClient) SetLogFun(logFun gorequest.LogFunc) {
	if c.httpClient != nil {
		c.httpClient.SetLogFunc(logFun)
	}
}
