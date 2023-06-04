package wechatqy

func (c *Client) GetKey() string {
	return c.config.key
}

func (c *Client) GetAppId() string {
	return c.config.appId
}

func (c *Client) GetAgentId() int {
	return c.config.agentId
}

func (c *Client) GetSecret() string {
	return c.config.secret
}

func (c *Client) GetRedirectUri() string {
	return c.config.redirectUri
}
