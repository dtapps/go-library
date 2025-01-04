package gorequest

import (
	"context"
	"net"
)

// 定义一个私有类型的键，以避免与其他包中的键冲突。
type contextKey string

const requestIPKey contextKey = "request_ip"

// SetRequestIP 为给定的上下文添加请求IP信息（支持 net.IP 和 string 类型）
func SetRequestIP(ctx context.Context, ip any) context.Context {
	switch v := ip.(type) {
	case net.IP:
		return context.WithValue(ctx, requestIPKey, v)
	case string:
		return context.WithValue(ctx, requestIPKey, v)
	default:
		return ctx // 如果类型不匹配，则返回原始的上下文
	}
}

// GetRequestIP 从给定的上下文中获取请求IP信息并尝试转换为 net.IP 类型
func GetRequestIP(ctx context.Context) (net.IP, bool) {
	if ip, ok := ctx.Value(requestIPKey).(net.IP); ok {
		return ip, true
	} else if ipStr, ok := ctx.Value(requestIPKey).(string); ok {
		ip := net.ParseIP(ipStr)
		return ip, ip != nil
	}
	return nil, false
}

// GetRequestIPStr 从给定的上下文中获取请求IP信息并尝试转换为 string 类型
func GetRequestIPStr(ctx context.Context) (string, bool) {
	if ipStr, ok := ctx.Value(requestIPKey).(string); ok {
		return ipStr, true
	} else if ip, ok := ctx.Value(requestIPKey).(net.IP); ok {
		return ip.String(), true
	}
	return "", false
}
