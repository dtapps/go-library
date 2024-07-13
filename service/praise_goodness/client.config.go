package praise_goodness

import "go.dtapp.net/library/utils/gorequest"

func (c *Client) GetApiURL() string {
	return c.config.apiURL
}

func (c *Client) SetApiURL(v string) *Client {
	c.config.apiURL = v
	return c
}

func (c *Client) GetMchID() int64 {
	return c.config.mchID
}

func (c *Client) SetMchID(v int64) *Client {
	c.config.mchID = v
	return c
}

func (c *Client) GetKey() string {
	return c.config.Key
}

func (c *Client) SetKey(v string) *Client {
	c.config.Key = v
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
