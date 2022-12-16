package kashangwl

type SetConfigConfig struct {
	CustomerId  int64  // 商家编号
	CustomerKey string // 商家密钥
}

// SetConfig 配置
func (c *Client) SetConfig(config *SetConfigConfig) *Client {
	c.config.customerId = config.CustomerId
	c.config.customerKey = config.CustomerKey
	return c
}
