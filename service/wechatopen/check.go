package wechatopen

func (c *Client) checkComponentIsConfig() error {
	componentAppId := c.GetComponentAppId()
	if componentAppId == "" {
		return componentAppIdNoConfig
	}
	componentAppSecret := c.GetComponentAppSecret()
	if componentAppSecret == "" {
		return componentAppSecretNoConfig
	}
	return nil
}

func (c *Client) checkAuthorizerIsConfig() error {
	authorizerAppid := c.GetAuthorizerAppid()
	if authorizerAppid == "" {
		return authorizerAppidNoConfig
	}
	return nil
}
