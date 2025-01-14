package x7s

import "go.dtapp.net/library/utils/gorequest"

func (c *Client) GetApiURL() string {
	return c.config.apiURL
}

func (c *Client) SetApiURL(v string) *Client {
	c.config.apiURL = v
	return c
}

func (c *Client) GetPartnerID() int64 {
	return c.config.partnerID
}

func (c *Client) SetPartnerID(v int64) *Client {
	c.config.partnerID = v
	return c
}

func (c *Client) GetApiKey() string {
	return c.config.apiKey
}

func (c *Client) SetApiKey(v string) *Client {
	c.config.apiKey = v
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
