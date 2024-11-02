package x7s

import (
	"errors"
	"go.dtapp.net/library/utils/gorequest"
)

// ClientConfig 实例配置
type ClientConfig struct {
	ApiURL    string `json:"api_url"`    // 接口地址
	PartnerID int64  `json:"partner_id"` // 平台分配商户号
	ApiKey    string `json:"api_key"`    // 渠道分配的密钥
}

// Client 实例
type Client struct {
	config struct {
		apiURL    string // 接口地址
		partnerID int64  // 平台分配商户号
		apiKey    string // 渠道分配的密钥
	}
	httpClient *gorequest.App // HTTP请求客户端
	clientIP   string         // 客户端IP
}

// NewClient 创建实例化
func NewClient(config *ClientConfig) (*Client, error) {
	c := &Client{}

	if config.ApiURL == "" {
		return nil, errors.New("需要配置ApiURL")
	}

	c.httpClient = gorequest.NewHttp()

	c.config.apiURL = config.ApiURL
	c.config.partnerID = config.PartnerID
	c.config.apiKey = config.ApiKey

	return c, nil
}
