package wechatqy

func (c *Client) GetKey() string {
	return c.config.Key
}

func (c *Client) GetAppId() string {
	return c.config.AppId
}

func (c *Client) GetAgentId() int {
	return c.config.AgentId
}

func (c *Client) GetSecret() string {
	return c.config.Secret
}

func (c *Client) GetRedirectUri() string {
	return c.config.RedirectUri
}
