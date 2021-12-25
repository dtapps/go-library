package gorequest

import (
	"errors"
	"math/rand"
	"net"
	"net/http"
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
	return userAgentList[r.Intn(len(userAgentList))]
}

var userAgentList = []string{

	"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/92.0.4515.107 Safari/537.36 Edg/92.0.902.55", // Edge 2021-07-28
	"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/92.0.4515.131 Safari/537.36 Edg/92.0.902.67", // Edge 2021-08-09
	"Mozilla/5.0 (Windows NT 10.0) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/92.0.4515.131 Safari/537.36 Edg/92.0.902.67",                   // Edge 2021-08-11

	"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_6) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147.135 Safari/537.36",
	"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/92.0.4515.107 Safari/537.36", // Chrome 2021-07-28
	"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/92.0.4515.131 Safari/537.36", // Chrome 2021-08-11
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/92.0.4515.131 Safari/537.36",       // Chrome 2021-08-11

	"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/15.0 Safari/605.1.15", // Safari 2021-07-28

	"Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:90.0) Gecko/20100101 Firefox/90.0",  // Firefox 2021-07-28
	"Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:91.0) Gecko/20100101 Firefox/91.0 ", // Firefox 2021-08-11
}

// GetRandomUserAgentMobile 获取用户UA
func GetRandomUserAgentMobile() string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return userAgentMobileList[r.Intn(len(userAgentMobileList))]
}

var userAgentMobileList = []string{
	"Mozilla/5.0 (iPhone; CPU iPhone OS 15_0 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) EdgiOS/94.0.972.2 Version/15.0 Mobile/15E148 Safari/604.1", // Edge 2021-08-11

	"Mozilla/5.0 (iPhone; CPU iPhone OS 15_0 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) CriOS/92.0.4515.90 Mobile/15E148 Safari/604.1", // Chrome 2021-08-11

	"Mozilla/5.0 (iPhone; CPU iPhone OS 15_0 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/15.0 Mobile/15E148 Safari/604.1", // Safari 2021-08-11

	"Mozilla/5.0 (Linux; U; Android 11; zh-cn; M2011K2C Build/RKQ1.200928.002) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/79.0.3945.147 Mobile Safari/537.36 XiaoMi/MiuiBrowser/15.1.12", // XIAOMI 2021-08-11

	"Mozilla/5.0 (iPhone; CPU iPhone OS 15_0 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Mobile/15E148 SP-engine/2.34.0 main%2F1.0 baiduboxapp/12.21.1.10 (Baidu; P2 15.0) NABar/1.0 themeUA=Theme/default webCore=0x1337abb70", // BAIDU 2021-08-11

	"Mozilla/5.0 (iPhone; CPU iPhone OS 15_0 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Mobile/19A5318f ChannelId(29) Ariver/1.1.0 AliApp(AP/10.2.28.6000) Nebula WK RVKType(1) AlipayDefined(nt:WIFI,ws:414|832|3.0,ac:T) AlipayClient/10.2.28.6000 Language/zh-Hans Region/CN NebulaX/1.0.0", // ALIPAY 2021-08-11

	"Mozilla/5.0 (iPhone; CPU iPhone OS 15_0 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Mobile/19A5261w AliApp(DingTalk/6.0.23) com.laiwang.DingTalk/15108471 Channel/201200 language/zh-Hans-CN UT4Aplus/0.0.6 WK", // DINGDING 2021-08-11

	"Mozilla/5.0 (iPhone; CPU iPhone OS 15_0 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Mobile/15E148 MicroMessenger/8.0.10(0x18000a24) NetType/WIFI Language/zh_CN", // WECHAT 2021-08-11

	"Mozilla/5.0 (iPhone; CPU iPhone OS 15_0 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Mobile/15E148 wxwork/3.1.11 MicroMessenger/7.0.1 Language/zh ColorScheme/Dark", // WECHATWORK 2021-08-11

	"Mozilla/5.0 (iPhone; CPU iPhone OS 15_0 like Mac OS X) AppleWebKit/605.1.15 (KHTML, like Gecko) Mobile/19A5318f QQ/8.8.17.612 V1_IPH_SQ_8.8.17_1_APP_A Pixel/1242 MiniAppEnable SimpleUISwitch/0 StudyMode/0 QQTheme/1102 Core/WKWebView Device/Apple(iPhone 11 Pro Max) NetType/WIFI QBWebViewType/1 WKType/1", // QQ 2021-08-11

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
