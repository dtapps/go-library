package wechatopen

import "github.com/dtapps/go-library/utils/golog"

//func (c *Client) GetComponentAccessToken() string {
//	return c.config.componentAccessToken
//}

//func (c *Client) GetComponentVerifyTicket() string {
//	return c.config.componentVerifyTicket
//}

//func (c *Client) GetPreAuthCode() string {
//	return c.config.preAuthCode
//}

//func (c *Client) GetAuthorizerAccessToken() string {
//	return c.config.authorizerAccessToken
//}

func (c *Client) GetAuthorizerRefreshToken() string {
	return c.config.authorizerRefreshToken
}

func (c *Client) GetAuthorizerAppid() string {
	return c.config.authorizerAppid
}

func (c *Client) GetComponentAppId() string {
	return c.config.componentAppId
}

func (c *Client) GetComponentAppSecret() string {
	return c.config.componentAppSecret
}

func (c *Client) GetMessageToken() string {
	return c.config.messageToken
}

func (c *Client) GetMessageKey() string {
	return c.config.messageKey
}

func (c *Client) GetLog() *golog.ApiClient {
	return c.log.client
}
