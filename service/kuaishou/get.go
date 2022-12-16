package kuaishou

import "github.com/dtapps/go-library/utils/golog"

func (c *Client) GetLog() *golog.ApiClient {
	return c.log.client
}
