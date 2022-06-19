package ejiaofei

func (c *Client) getUserId() string {
	return c.config.UserId
}

func (c *Client) getPwd() string {
	return c.config.Pwd
}

func (c *Client) getKey() string {
	return c.config.Key
}
