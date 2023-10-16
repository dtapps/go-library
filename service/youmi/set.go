package youmi

import "github.com/dtapps/go-library/utils/gorequest"

// SetHttp 配置请求
func (c *Client) SetHttp(app *gorequest.App) {
	c.requestClient = app
	c.requestClientStatus = true
}

func (c *Client) SetApiURL(apiURL string) {
	c.config.apiURL = apiURL
}

func (c *Client) SetUserID(userID int64) {
	c.config.userID = userID
}

func (c *Client) SetApiKey(apiKey string) {
	c.config.apiKey = apiKey
}
