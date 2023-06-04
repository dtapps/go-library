package leshuazf

func (c *Client) GetEnvironment() string {
	return c.config.environment
}

func (c *Client) GetAgentId() string {
	return c.config.agentId
}

func (c *Client) GetKeyAgent() string {
	return c.config.keyAgent
}
