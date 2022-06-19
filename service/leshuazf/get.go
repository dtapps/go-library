package leshuazf

func (c *Client) GetEnvironment() string {
	return c.config.Environment
}

func (c *Client) GetAgentId() string {
	return c.config.AgentId
}

func (c *Client) GetKeyAgent() string {
	return c.config.KeyAgent
}
