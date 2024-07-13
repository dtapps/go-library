package goip

import (
	"context"
	"go.dtapp.net/library/utils/gorequest"
)

// GetInsideIp 内网ip
func GetInsideIp(ctx context.Context) string {
	return gorequest.GetInsideIp(ctx)
}

// Ips 获取全部网卡的全部IP
func Ips(ctx context.Context) (map[string]string, error) {
	return gorequest.Ips(ctx)
}

// GetOutsideIp 外网ip
func GetOutsideIp(ctx context.Context) string {
	return gorequest.GetOutsideIp(ctx)
}

// GetMacAddr 获取Mac地址
func GetMacAddr(ctx context.Context) (arrays []string) {
	return gorequest.GetMacAddr(ctx)
}
