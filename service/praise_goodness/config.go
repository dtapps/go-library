package praise_goodness

import (
	"github.com/dtapps/go-library/utils/golog"
	"github.com/dtapps/go-library/utils/gorequest"
)

// ConfigSLogClientFun 日志配置
func (c *Client) ConfigSLogClientFun(apiSLogFun golog.ApiSLogFun) {
	apiSLog := apiSLogFun()
	if apiSLog != nil {
		c.slog.client = apiSLog
		c.slog.status = true
	}
}

// DefaultHttp 默认请求
func (c *Client) DefaultHttp() {
	c.requestClient = gorequest.NewHttp()
	c.requestClientStatus = true
}

func (c *Client) ConfigClient(config *ClientConfig) {
	c.config.apiURL = config.ApiURL
	c.config.mchID = config.MchID
	c.config.Key = config.Key
}
