package wechatpayopen

// SubConfig 子商户配置
func (c *Client) SubConfig(subAppid, subMchId string) *Client {
	c.config.subAppid = subAppid
	c.config.subMchId = subMchId
	return c
}
