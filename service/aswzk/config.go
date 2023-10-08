package aswzk

import (
	"github.com/dtapps/go-library/utils/golog"
	"github.com/dtapps/go-library/utils/gorequest"
)

// ConfigApp 配置
func (c *Client) ConfigApp(userID string, apiKey string) *Client {
	c.config.userID = userID
	c.config.apiKey = apiKey
	return c
}

// ConfigSLogClientFun 日志配置
func (c *Client) ConfigSLogClientFun(apiSLogFun golog.ApiSLogFun) {
	apiSLog := apiSLogFun()
	if apiSLog != nil {
		c.slog.client = apiSLog
		c.slog.status = true
	}
}

// SetHttp 配置请求
func (c *Client) SetHttp(app *gorequest.App) {
	c.requestClient = app
	c.requestClientStatus = true
	c.requestClient.Uri = apiUrl
}

// DefaultHttp 默认请求
func (c *Client) DefaultHttp() {
	c.requestClient = gorequest.NewHttp()
	c.requestClientStatus = true
	c.requestClient.Uri = apiUrl
}
