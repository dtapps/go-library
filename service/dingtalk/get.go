package dingtalk

import "github.com/dtapps/go-library/utils/golog"

func (c *Client) GetSecret() string {
	return c.config.secret
}

func (c *Client) GetAccessToken() string {
	return c.config.accessToken
}

func (c *Client) GetLog() *golog.ApiClient {
	return c.log.client
}
