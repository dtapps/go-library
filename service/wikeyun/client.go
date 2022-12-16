package wikeyun

import (
	"github.com/dtapps/go-library/utils/golog"
	"github.com/dtapps/go-library/utils/gorequest"
)

// ClientConfig 实例配置
type ClientConfig struct {
	StoreId   int    // 店铺ID
	AppKey    int    // key
	AppSecret string // secret
	CurrentIp string // 当前ip
}

// Client 实例
type Client struct {
	requestClient *gorequest.App // 请求服务
	config        struct {
		clientIp  string // 当前Ip
		storeId   int    // 店铺ID
		appKey    int    // key
		appSecret string // secret
	}
	log struct {
		status bool             // 状态
		client *golog.ApiClient // 日志服务
	}
}

// NewClient 创建实例化
func NewClient(config *ClientConfig) (*Client, error) {

	c := &Client{}

	c.config.clientIp = config.CurrentIp

	c.config.storeId = config.StoreId
	c.config.appKey = config.AppKey
	c.config.appSecret = config.AppSecret

	c.requestClient = gorequest.NewHttp()

	return c, nil
}
