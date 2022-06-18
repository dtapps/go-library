package wechatunion

func (c *Client) getAppId() string {
	return c.config.AppId
}

func (c *Client) getAppSecret() string {
	return c.config.AppSecret
}

func (c *Client) getAccessToken() string {
	c.config.AccessToken = c.GetAccessToken()
	return c.config.AccessToken
}

func (c *Client) getPid() string {
	return c.config.Pid
}
