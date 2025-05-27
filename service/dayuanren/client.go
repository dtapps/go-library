package dayuanren

import (
	"resty.dev/v3"
)

// ClientConfig 实例配置
type ClientConfig struct {
	ApiURL string // 接口地址
	UserID int64  // 商户ID
	ApiKey string // 秘钥
	Debug  bool
}

// Client 实例
type Client struct {
	config struct {
		apiURL string // 接口地址
		userID int64  // 商户ID
		apiKey string // 秘钥
	}
	debug      bool
	httpClient *resty.Client // 请求客户端
}

// NewClient 创建实例化
func NewClient(config *ClientConfig, opts ...Option) (*Client, error) {

	options := NewOptions(opts)

	c := &Client{}

	if options.httpClient == nil {
		c.httpClient = resty.New().SetDebug(config.Debug)
	} else {
		c.httpClient = options.httpClient
	}

	c.config.apiURL = config.ApiURL
	c.config.userID = config.UserID
	c.config.apiKey = config.ApiKey

	return c, nil
}
