package bark

import (
	"go.dtapp.net/library/utils/gorequest"
	"strings"
)

func (c *Client) SetPushKey(pushKey string) *Client {
	if strings.HasPrefix(pushKey, c.config.url) {
		c.config.pushKey = strings.TrimPrefix(pushKey, c.config.url)
	} else {
		c.config.pushKey = pushKey
	}
	return c
}

func (c *Client) SetUrl(url string) *Client {
	if url != "" {
		c.config.url = url
	}
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
