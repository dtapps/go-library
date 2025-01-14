package wechatpayapiv3

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

func (c *Client) GetAppSecret() string {
	return c.config.appSecret
}

func (c *Client) SetAppSecret(v string) *Client {
	c.config.appSecret = v
	return c
}

func (c *Client) GetMchId() string {
	return c.config.mchId
}

func (c *Client) SetMchId(v string) *Client {
	c.config.mchId = v
	return c
}

func (c *Client) GetAesKey() string {
	return c.config.aesKey
}

func (c *Client) SetAesKey(v string) *Client {
	c.config.aesKey = v
	return c
}

func (c *Client) GetApiV3() string {
	return c.config.apiV3
}

func (c *Client) SetApiV3(v string) *Client {
	c.config.apiV3 = v
	return c
}

func (c *Client) GetMchSslSerialNo() string {
	return c.config.mchSslSerialNo
}

func (c *Client) SetMchSslSerialNo(v string) *Client {
	c.config.mchSslSerialNo = v
	return c
}

func (c *Client) GetMchSslKey() string {
	return c.config.mchSslKey
}

func (c *Client) SetMchSslKey(v string) *Client {
	c.config.mchSslKey = v
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
