package kashangwl

import (
	"go.dtapp.net/library/utils/gorequest"
)

func (c *Client) GetCustomerId() int64 {
	return c.config.customerId
}

func (c *Client) SetCustomerId(v int64) *Client {
	c.config.customerId = v
	return c
}

func (c *Client) GetCustomerKey() string {
	return c.config.customerKey
}

func (c *Client) SetCustomerKey(v string) *Client {
	c.config.customerKey = v
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
