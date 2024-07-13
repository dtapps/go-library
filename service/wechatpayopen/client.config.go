package wechatpayopen

import (
	"go.dtapp.net/library/utils/gorequest"
)

func (c *Client) GetSpAppid() string {
	return c.config.spAppid
}

func (c *Client) SetSpAppid(v string) *Client {
	c.config.spAppid = v
	return c
}

func (c *Client) GetSpMchId() string {
	return c.config.spMchId
}

func (c *Client) SetSpMchId(v string) *Client {
	c.config.spMchId = v
	return c
}

func (c *Client) GetSubAppid() string {
	return c.config.subAppid
}

func (c *Client) SetSubAppid(v string) *Client {
	c.config.subAppid = v
	return c
}

func (c *Client) GetSubMchId() string {
	return c.config.subMchId
}

func (c *Client) SetSubMchId(v string) *Client {
	c.config.subMchId = v
	return c
}

func (c *Client) GetApiV2() string {
	return c.config.apiV2
}

func (c *Client) SetApiV2(v string) *Client {
	c.config.apiV2 = v
	return c
}

func (c *Client) GetApiV3() string {
	return c.config.apiV3
}

func (c *Client) SetApiV3(v string) *Client {
	c.config.apiV3 = v
	return c
}

func (c *Client) GetSerialNo() string {
	return c.config.serialNo
}

func (c *Client) SetSerialNo(v string) *Client {
	c.config.serialNo = v
	return c
}

func (c *Client) GetMchSslSerialNo() string {
	return c.config.mchSslSerialNo
}

func (c *Client) SetMchSslSerialNo(v string) *Client {
	c.config.mchSslSerialNo = v
	return c
}

func (c *Client) GetMchSslCer() string {
	return c.config.mchSslCer
}

func (c *Client) SetMchSslCer(v string) *Client {
	c.config.mchSslCer = v
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
