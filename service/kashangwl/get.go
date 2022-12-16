package kashangwl

import "github.com/dtapps/go-library/utils/golog"

func (c *Client) GetCustomerId() int64 {
	return c.config.customerId
}

func (c *Client) GetCustomerKey() string {
	return c.config.customerKey
}

func (c *Client) GetLog() *golog.ApiClient {
	return c.log.client
}
