package gorequest

import (
	"net"
	"net/http"
	"strings"
)

// ClientIp 尽最大努力实现获取客户端 IP 的算法。
// 解析 X-Real-IP 和 X-Forwarded-For 以便于反向代理（nginx 或 haproxy）可以正常工作。
func ClientIp(r *http.Request) string {

	// CloudFlare
	CfConnectingIp := strings.TrimSpace(r.Header.Get("Cf-Connecting-Ip"))
	if CfConnectingIp != "" {
		return CfConnectingIp
	}

	// 转发IP
	xForwardedFor := strings.TrimSpace(strings.Split(r.Header.Get("X-Forwarded-For"), ",")[0])
	if xForwardedFor != "" {
		return xForwardedFor
	}

	// 真实Ip
	XRealIp := strings.TrimSpace(r.Header.Get("X-Real-Ip"))
	if XRealIp != "" {
		return XRealIp
	}

	// HTTP客户端IP
	HttpClientIp := strings.TrimSpace(strings.Split(r.Header.Get("HTTP_CLIENT_IP"), ",")[0])
	if HttpClientIp != "" {
		return HttpClientIp
	}

	// HTTP转发IP
	HttpXForwardedFor := strings.TrimSpace(strings.Split(r.Header.Get("HTTP_X_FORWARDED_FOR"), ",")[0])
	if HttpXForwardedFor != "" {
		return HttpXForwardedFor
	}

	// 系统
	ip, _, err := net.SplitHostPort(strings.TrimSpace(r.RemoteAddr))
	if err == nil {
		return ip
	}

	return ""
}
