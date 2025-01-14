package kuaidi100

import "go.dtapp.net/library/utils/gorequest"

func (c *Client) GetCustomer() string {
	return c.config.customer
}

func (c *Client) SetAesCustomer(v string) *Client {
	c.config.customer = v
	return c
}

func (c *Client) GetKey() string {
	return c.config.key
}

func (c *Client) SetKey(v string) *Client {
	c.config.key = v
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
