package golog

import "github.com/dtapps/go-library/utils/goip"

// ApiSLog 接口日志
type ApiSLog struct {
	systemConfig struct {
		systemHostname  string // 主机名
		systemOs        string // 系统类型
		systemKernel    string // 系统内核
		systemInsideIp  string // 内网ip
		systemOutsideIp string // 外网ip
		goVersion       string // go版本
		sdkVersion      string // sdk版本
	}
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
	config    struct {
		systemHostname      string  // 主机名
		systemOs            string  // 系统类型
		systemVersion       string  // 系统版本
		systemKernel        string  // 系统内核
		systemKernelVersion string  // 系统内核版本
		systemBootTime      uint64  // 系统开机时间
		cpuCores            int     // CPU核数
		cpuModelName        string  // CPU型号名称
		cpuMhz              float64 // CPU兆赫
		systemInsideIp      string  // 内网ip
		systemOutsideIp     string  // 外网ip
		goVersion           string  // go版本
		sdkVersion          string  // sdk版本
	}
	slog struct {
		status bool  // 状态
		client *SLog // 日志服务
	}
}

// GinClientFun  框架日志驱动
type GinClientFun func() *GinClient

// GinClientConfig 框架日志配置
type GinClientConfig struct {
	IpService *goip.Client // IP服务
	CurrentIp string       // 当前IP
}

// GinCustomClient 框架自定义日志
type GinCustomClient struct {
	ipService *goip.Client // IP服务
	config    struct {
		systemHostname  string // 主机名
		systemOs        string // 系统类型
		systemKernel    string // 系统内核
		systemInsideIp  string // 内网ip
		systemOutsideIp string // 外网ip
		goVersion       string // go版本
		sdkVersion      string // sdk版本
	}
	slog struct {
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
