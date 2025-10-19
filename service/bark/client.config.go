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
	if strings.HasPrefix(pushKey, c.config.baseURL) {
		c.config.pushKey = strings.TrimPrefix(pushKey, c.config.baseURL)
	} else {
		c.config.pushKey = pushKey
	}
	return c
}
