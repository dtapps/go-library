package kuaidi100

func (c *Client) GetCustomer() string {
	return c.config.customer
}

func (c *Client) SetAesCustomer(v string) *Client {
	c.config.customer = v
	return c
}

func (c *Client) GetKey() string {
	return c.config.key
}

func (c *Client) SetKey(v string) *Client {
	c.config.key = v
	return c
}
