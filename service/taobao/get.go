package taobao

func (c *Client) GetAppKey() string {
	return c.config.AppKey
}

func (c *Client) GetAppSecret() string {
	return c.config.AppSecret
}

func (c *Client) GetAdzoneId() int64 {
	return c.config.AdzoneId
}
