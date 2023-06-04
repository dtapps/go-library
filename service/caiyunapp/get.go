package caiyunapp

func (c *Client) GetToken() string {
	return c.config.token
}

func (c *Client) getApiUrl() string {
	return apiUrl + "/" + c.config.token
}
