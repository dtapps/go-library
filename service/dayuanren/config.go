package dayuanren

import (
	"github.com/dtapps/go-library/utils/golog"
	"github.com/dtapps/go-library/utils/gorequest"
)

// ConfigClient 配置
func (c *Client) ConfigClient(config *ClientConfig) {
	c.config.apiURL = config.ApiURL
	c.config.userID = config.UserID
	c.config.apiKey = config.ApiKey
}

// ConfigSLogClientFun 日志配置
func (c *Client) ConfigSLogClientFun(sLogFun golog.SLogFun) {
	sLog := sLogFun()
	if sLog != nil {
		c.slog.client = sLog
		c.slog.status = true
	}
}

// DefaultHttp 默认请求
func (c *Client) DefaultHttp() {
	c.requestClient = gorequest.NewHttp()
	c.requestClientStatus = true
}
