package pinduoduo

import "go.dtapp.net/library/utils/gorequest"

func (c *Client) GetClientId() string {
	return c.config.clientId
}

func (c *Client) SetClientId(v string) *Client {
	c.config.clientId = v
	return c
}

func (c *Client) GetClientSecret() string {
	return c.config.clientSecret
}

func (c *Client) SetClientSecret(v string) *Client {
	c.config.clientSecret = v
	return c
}

func (c *Client) GetMediaId() string {
	return c.config.mediaId
}

func (c *Client) SetMediaId(v string) *Client {
	c.config.mediaId = v
	return c
}

func (c *Client) GetPid() string {
	return c.config.pid
}

func (c *Client) SetPid(v string) *Client {
	c.config.pid = v
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
