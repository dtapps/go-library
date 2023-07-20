package kuaidi100

func (c *Client) GetCustomer() string {
	return c.config.customer
}

func (c *Client) GetKey() string {
	return c.config.key
}
