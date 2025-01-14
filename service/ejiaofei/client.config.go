package ejiaofei

import "go.dtapp.net/library/utils/gorequest"

func (c *Client) GetUserId() string {
	return c.config.userId
}

func (c *Client) SetUserId(v string) *Client {
	c.config.userId = v
	return c
}

func (c *Client) GetPwd() string {
	return c.config.pwd
}

func (c *Client) SetPwd(v string) *Client {
	c.config.pwd = v
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
