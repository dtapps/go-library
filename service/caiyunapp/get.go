package caiyunapp

import "github.com/dtapps/go-library/utils/golog"

func (c *Client) GetToken() string {
	return c.config.token
}

func (c *Client) GetLog() *golog.ApiClient {
	return c.log.client
}

func (c *Client) getApiUrl() string {
	return apiUrl + "/" + c.config.token
}
