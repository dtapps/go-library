package wnfuwu

import "github.com/dtapps/go-library/utils/golog"

func (c *Client) GetUserId() int64 {
	return c.config.userId
}

func (c *Client) GetApiKey() string {
	return c.config.apiKey
}

func (c *Client) GetLog() *golog.ApiClient {
	return c.log.client
}
