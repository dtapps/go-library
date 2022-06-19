package kashangwl

func (c *Client) GetCustomerId() int {
	return c.config.CustomerId
}

func (c *Client) GetCustomerKey() string {
	return c.config.CustomerKey
}
