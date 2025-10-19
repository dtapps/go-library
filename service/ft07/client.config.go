package ft07

import (
	"fmt"
)

func (c *Client) SetUrl(baseURL string) *Client {
	if baseURL != "" {
		c.config.baseURL = baseURL
	}
	return c
}

func (c *Client) SetSendKey(sendKey string) *Client {
	c.config.sendKey = sendKey
	if len(c.config.sendKey) >= 4 && c.config.sendKey[:4] == "sctp" {
		c.config.baseURL = fmt.Sprintf("https://%s.push.ft07.com/send", c.config.sendKey)
	} else {
		c.config.baseURL = fmt.Sprintf("https://sctapi.ftqq.com/%s.send", c.config.sendKey)
	}
	return c
}
