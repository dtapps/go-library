package nldyp

import (
	"github.com/dtapps/go-library/utils/golog"
	"github.com/dtapps/go-library/utils/gorequest"
)

// ClientConfig 实例配置
type ClientConfig struct {
	vendor string // 秘钥
	AppKey string // 渠道标记
}

// Client 实例
type Client struct {
	requestClient *gorequest.App // 请求服务
	config        struct {
		secret string // 秘钥
		appKey string // 渠道标记
	}
	log struct {
		status bool             // 状态
		client *golog.ApiClient // 日志服务
	}
}
