package ejiaofei

import "github.com/dtapps/go-library/utils/golog"

func (c *Client) GetUserId() string {
	return c.config.userId
}

func (c *Client) GetPwd() string {
	return c.config.pwd
}

func (c *Client) GetKey() string {
	return c.config.key
}

func (c *Client) GetLog() *golog.ApiClient {
	return c.log.client
}
