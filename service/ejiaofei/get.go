package ejiaofei

func (c *Client) GetUserId() string {
	return c.config.userId
}

func (c *Client) GetPwd() string {
	return c.config.pwd
}

func (c *Client) GetKey() string {
	return c.config.key
}
