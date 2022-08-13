package wechatoffice

import "context"

func (c *Client) GetAppId() string {
	return c.config.AppId
}

func (c *Client) GetAppSecret() string {
	return c.config.AppSecret
}

func (c *Client) getAccessToken(ctx context.Context) string {
	c.config.AccessToken = c.GetAccessToken(ctx)
	return c.config.AccessToken
}

func (c *Client) getJsapiTicket(ctx context.Context) string {
	c.config.JsapiTicket = c.GetJsapiTicket(ctx)
	return c.config.JsapiTicket
}
