package pintoto

func (c *Client) GetAppKey() string {
	return c.config.AppKey
}
func (c *Client) GetAppSecret() string {
	return c.config.AppSecret
}
