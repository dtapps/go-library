package wechatpayapiv3

func (c *Client) SetAppId(appId string) {
	c.config.AppId = appId
}

func (c *Client) SetAppSecret(appSecret string) {
	c.config.AppSecret = appSecret
}
