package pinduoduo

func (c *Client) GetClientId() string {
	return c.config.clientId
}

func (c *Client) SetClientId(v string) *Client {
	c.config.clientId = v
	return c
}

func (c *Client) GetClientSecret() string {
	return c.config.clientSecret
}

func (c *Client) SetClientSecret(v string) *Client {
	c.config.clientSecret = v
	return c
}

func (c *Client) GetMediaId() string {
	return c.config.mediaId
}

func (c *Client) SetMediaId(v string) *Client {
	c.config.mediaId = v
	return c
}

func (c *Client) GetPid() string {
	return c.config.pid
}

func (c *Client) SetPid(v string) *Client {
	c.config.pid = v
	return c
}

// GetAccessToken 通过code获取的access_token(无需授权的接口，该字段不参与sign签名运算)
func (c *Client) GetAccessToken() string {
	return c.config.accessToken
}

// SetAccessToken 通过code获取的access_token(无需授权的接口，该字段不参与sign签名运算)
func (c *Client) SetAccessToken(accessToken string) *Client {
	c.config.accessToken = accessToken
	return c
}

// GetAccessTokenScope 授权范围
func (c *Client) GetAccessTokenScope() []string {
	return c.config.accessTokenScope
}

// SetAccessTokenScope 授权范围
func (c *Client) SetAccessTokenScope(accessTokenScope []string) *Client {
	c.config.accessTokenScope = accessTokenScope
	return c
}
