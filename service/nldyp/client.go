package nldyp

import (
	"github.com/dtapps/go-library/utils/golog"
	"github.com/dtapps/go-library/utils/gorequest"
)

// ClientConfig 实例配置
type ClientConfig struct {
	Vendor string
}

// Client 实例
type Client struct {
	requestClient *gorequest.App // 请求服务
	config        struct {
		vendor string
	}
	slog struct {
		status bool           // 状态
		client *golog.ApiSLog // 日志服务
	}
}
