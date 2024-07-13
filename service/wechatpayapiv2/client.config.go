package wechatpayapiv2

import (
	"crypto/tls"
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

func (c *Client) GetMchKey() string {
	return c.config.mchKey
}

func (c *Client) SetMchKey(v string) *Client {
	c.config.mchKey = v
	return c
}

func (c *Client) GetCertString() string {
	return c.config.certString
}

func (c *Client) SetCertString(v string) *Client {
	c.config.certString = v
	return c
}

func (c *Client) GetKeyString() string {
	return c.config.keyString
}

func (c *Client) SetKeyString(v string) *Client {
	c.config.keyString = v
	return c
}

func (c *Client) P12ToPem() (*tls.Certificate, error) {
	pemCert, err := tls.X509KeyPair([]byte(c.GetCertString()), []byte(c.GetKeyString()))
	return &pemCert, err
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
