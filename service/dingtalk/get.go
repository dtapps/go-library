package dingtalk

func (c *Client) GetSecret() string {
	return c.config.secret
}

func (c *Client) GetAccessToken() string {
	return c.config.accessToken
}
