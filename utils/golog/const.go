package golog

import "github.com/dtapps/go-library/utils/goip"

// ApiSLog 接口日志
type ApiSLog struct {
	slog struct {
		status bool  // 状态
		client *SLog // 日志服务
	}
}

// ApiSLogFun  接口日志驱动
type ApiSLogFun func() *ApiSLog

// GinClient 框架日志
type GinClient struct {
	ipService *goip.Client // IP服务
	slog      struct {
		status bool  // 状态
		client *SLog // 日志服务
	}
}

// GinClientFun  框架日志驱动
type GinClientFun func() *GinClient

// GinClientConfig 框架日志配置
type GinClientConfig struct {
	IpService *goip.Client // IP服务
}

// GinCustomClient 框架自定义日志
type GinCustomClient struct {
	ipService *goip.Client // IP服务
	slog      struct {
		status bool  // 状态
		client *SLog // 日志服务
	}
}

// GinCustomClientFun  框架自定义日志驱动
type GinCustomClientFun func() *GinCustomClient

// GinCustomClientConfig 框架自定义日志配置
type GinCustomClientConfig struct {
	IpService *goip.Client // IP服务
	CurrentIp string       // 当前IP
}
