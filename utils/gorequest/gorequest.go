package gorequest

import (
	"errors"
	"math/rand"
	"net"
	"net/http"
	"runtime"
	"strings"
	"time"
)

// ClientIp 尽最大努力实现获取客户端 IP 的算法。
// 解析 X-Real-IP 和 X-Forwarded-For 以便于反向代理（nginx 或 haproxy）可以正常工作。
func ClientIp(r *http.Request) string {
	// xForwardedFor
	xForwardedFor := r.Header.Get("X-Forwarded-For")
	ip := strings.TrimSpace(strings.Split(xForwardedFor, ",")[0])
	if ip != "" {
		return ip
	}
	// xRealIp
	ip = strings.TrimSpace(r.Header.Get("X-Real-Ip"))
	if ip != "" {
		return ip
	}
	// HTTPCLIENTIP
	HTTPCLIENTIP := r.Header.Get("HTTP_CLIENT_IP")
	ip = strings.TrimSpace(strings.Split(HTTPCLIENTIP, ",")[0])
	if ip != "" {
		return ip
	}
	// HTTPXFORWARDEDFOR
	HTTPXFORWARDEDFOR := r.Header.Get("HTTP_X_FORWARDED_FOR")
	ip = strings.TrimSpace(strings.Split(HTTPXFORWARDEDFOR, ",")[0])
	if ip != "" {
		return ip
	}
	// system
	if ip, _, err := net.SplitHostPort(strings.TrimSpace(r.RemoteAddr)); err == nil {
		return ip
	}

	return ""
}

// GetRandomUserAgent 获取用户UA
func GetRandomUserAgent() string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	switch runtime.GOOS {
	case "linux":
		return userAgentListLinux[r.Intn(len(userAgentListLinux))]
	case "windows":
		return userAgentListWindows[r.Intn(len(userAgentListWindows))]
	default:
		return userAgentListMac[r.Intn(len(userAgentListMac))]
	}
}

var userAgentListWindows = []string{
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/97.0.4692.71 Safari/537.36", // Chrome 2022-02-14
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:83.0) Gecko/20100101 Firefox/83.0",                                     // Firefox 2022-02-14
	"Mozilla/5.0 (Windows NT 10.0; WOW64; Trident/7.0; rv:11.0) like Gecko",                                              // IE 2022-02-14
}

var userAgentListLinux = []string{
	"Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/70.0.3538.80 Safari/537.36 HeyTapBrowser/40.7.35.1", // Chrome 2022-02-14
	"Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Safari/537.36 OnePlusBrowser/30.5.0.8",                     // Chrome 2022-02-14
	"Mozilla/5.0 (X11; U; Linux i686; en-US; rv:1.9a1) Gecko/20060814 Firefox/51.0",                                                    // Firefox 2022-02-14
}

var userAgentListMac = []string{
	"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/15.4 Safari/605.1.15",                     // Safari 2022-02-14
	"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/98.0.4758.82 Safari/537.36 Edg/98.0.1108.51", // Edge 2022-02-14
	"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/98.0.4758.80 Safari/537.36",                  // Chrome 2022-02-14
	"Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:97.0) Gecko/20100101 Firefox/97.0",                                                        // Firefox 2022-02-14
}

func ExternalIp() (string, error) {
	faces, err := net.Interfaces()
	if err != nil {
		return "", err
	}
	for _, iface := range faces {
		if iface.Flags&net.FlagUp == 0 {
			continue // interface down
		}
		if iface.Flags&net.FlagLoopback != 0 {
			continue // loopback interface
		}
		adders, err := iface.Addrs()
		if err != nil {
			return "", err
		}
		for _, addr := range adders {
			var ip net.IP
			switch v := addr.(type) {
			case *net.IPNet:
				ip = v.IP
			case *net.IPAddr:
				ip = v.IP
			}
			if ip == nil || ip.IsLoopback() {
				continue
			}
			ip = ip.To4()
			if ip == nil {
				continue // not an ipv4 address
			}
			return ip.String(), nil
		}
	}
	return "", errors.New("are you connected to the network")
}
