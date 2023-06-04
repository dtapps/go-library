package kashangwl

func (c *Client) GetCustomerId() int64 {
	return c.config.customerId
}

func (c *Client) GetCustomerKey() string {
	return c.config.customerKey
}
