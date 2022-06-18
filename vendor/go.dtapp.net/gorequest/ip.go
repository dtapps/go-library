package gorequest

import (
	"net"
	"net/http"
	"strings"
)

// ClientIp 尽最大努力实现获取客户端 IP 的算法。
// 解析 X-Real-IP 和 X-Forwarded-For 以便于反向代理（nginx 或 haproxy）可以正常工作。
func ClientIp(r *http.Request) string {

	// 转发IP
	xForwardedFor := r.Header.Get("X-Forwarded-For")
	ip := strings.TrimSpace(strings.Split(xForwardedFor, ",")[0])
	if ip != "" {
		return ip
	}

	// 真实Ip
	ip = strings.TrimSpace(r.Header.Get("X-Real-Ip"))
	if ip != "" {
		return ip
	}

	// HTTP客户端IP
	httpClientIp := r.Header.Get("HTTP_CLIENT_IP")
	ip = strings.TrimSpace(strings.Split(httpClientIp, ",")[0])
	if ip != "" {
		return ip
	}

	// HTTP转发IP
	HttpXForwardedFor := r.Header.Get("HTTP_X_FORWARDED_FOR")
	ip = strings.TrimSpace(strings.Split(HttpXForwardedFor, ",")[0])
	if ip != "" {
		return ip
	}

	// 系统
	if ip, _, err := net.SplitHostPort(strings.TrimSpace(r.RemoteAddr)); err == nil {
		return ip
	}

	return ""
}
