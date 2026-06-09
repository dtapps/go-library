package bark

import (
	"strings"
)

func (c *Client) SetUrl(baseURL string) *Client {
	if baseURL != "" {
		c.config.baseURL = baseURL
	}
	return c
}

func (c *Client) SetPushKey(pushKey string) *Client {
	if after, ok := strings.CutPrefix(pushKey, c.config.baseURL); ok {
		c.config.pushKey = after
	} else {
		c.config.pushKey = pushKey
	}
	return c
}
