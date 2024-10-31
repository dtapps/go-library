package pushdeer

import (
	"fmt"
	"go.dtapp.net/library/utils/gorequest"
)

func (c *Client) SetSendKey(sendKey string) *Client {
	c.config.sendKey = sendKey
	if len(c.config.sendKey) >= 4 && c.config.sendKey[:4] == "sctp" {
		c.config.url = fmt.Sprintf("https://%s.push.ft07.com/send", c.config.sendKey)
	} else {
		c.config.url = fmt.Sprintf("https://sctapi.ftqq.com/%s.send", c.config.sendKey)
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
