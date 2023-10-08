package praise_goodness

import "github.com/dtapps/go-library/utils/gorequest"

// SetHttp 配置请求
func (c *Client) SetHttp(app *gorequest.App) {
	c.requestClient = app
	c.requestClientStatus = true
}

func (c *Client) SetApiURL(apiURL string) {
	c.config.apiURL = apiURL
}

func (c *Client) SetMchID(mchID int64) {
	c.config.mchID = mchID
}

func (c *Client) SetAppKey(key string) {
	c.config.Key = key
}
