package chengquan

import "github.com/dtapps/go-library/utils/gorequest"

// SetHttp 配置请求
func (c *Client) SetHttp(app *gorequest.App) {
	c.requestClient = app
	c.requestClientStatus = true
}

func (c *Client) SetApiURL(apiURL string) {
	c.config.apiURL = apiURL
}

func (c *Client) SetAppID(appID string) {
	c.config.appID = appID
}

func (c *Client) SetAppKey(appKey string) {
	c.config.appKey = appKey
}

func (c *Client) SetAesKey(aesKey string) {
	c.config.aesKey = aesKey
}

func (c *Client) SetAesIv(aesIv string) {
	c.config.aesIv = aesIv
}
