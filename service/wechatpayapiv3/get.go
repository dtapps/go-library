package wechatpayapiv3

import "github.com/dtapps/go-library/utils/golog"

func (c *Client) GetAppId() string {
	return c.config.appId
}

func (c *Client) GetAppSecret() string {
	return c.config.appSecret
}

func (c *Client) GetMchId() string {
	return c.config.mchId
}

func (c *Client) GetAesKey() string {
	return c.config.aesKey
}

func (c *Client) GetApiV3() string {
	return c.config.apiV3
}

func (c *Client) GetMchSslKey() string {
	return c.config.mchSslKey
}

func (c *Client) GetMchSslSerialNo() string {
	return c.config.mchSslSerialNo
}

func (c *Client) GetLog() *golog.ApiClient {
	return c.log.client
}
