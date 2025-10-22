package qxwlwagnt

type Config struct {
	BaseURL   string // 接口地址
	UserName  string // userName
	AppKey    string // appKey
	AppSecret string // appSecret
}

func (c *Client) Config(config Config) {
	c.config.baseURL = config.BaseURL
	c.config.userName = config.UserName
	c.config.appKey = config.AppKey
	c.config.appSecret = config.AppSecret

	// 设置基础 URL
	if c.config.baseURL != "" {
		c.httpClient.SetBaseURL(c.config.baseURL)
	}
}
