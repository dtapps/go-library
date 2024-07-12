package gorequest

import (
	"math/rand"
	"runtime"
	"time"
)

const (
	linux   = "linux"
	windows = "windows"
	mac     = "mac"
)

// GetRandomUserAgent 获取随机UA
func GetRandomUserAgent() string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	return userAgentList[r.Intn(len(userAgentList))]
}

var userAgentList = []string{
	"Mozilla/5.0 (Windows NT 10.0; WOW64; Trident/7.0; rv:11.0) like Gecko",                                                                                   // IE浏览器 2023-12-30 http://api.dtapp.net/ip?ua=widows_ie
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36",                                         // Chrome浏览器 2023-12-30 http://api.dtapp.net/ip?ua=widows_chrome
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36 Edg/120.0.0.0",                           // Edge浏览器 2023-12-30 http://api.dtapp.net/ip?ua=widows_edge
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:121.0) Gecko/20100101 Firefox/121.0",                                                                        // Firefox浏览器 2023-12-30 http://api.dtapp.net/ip?ua=widows_firefox
	"Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/94.0.4606.71 Safari/537.36 Core/1.94.218.400 QQBrowser/12.1.5496.400", // QQ浏览器 2023-12-30 http://api.dtapp.net/ip?ua=widows_qq
	"Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.5735.289 Safari/537.36",                                         // 360浏览器 2023-12-30 http://api.dtapp.net/ip?ua=widows_360
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/119.0.6045.160 Safari/537.36",                                    // 360极速浏览器 2023-12-30 http://api.dtapp.net/ip?ua=widows_360js
	"Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/94.0.4606.71 Safari/537.36 SE 2.X MetaSr 1.0",                         // 搜狗浏览器 2023-12-30 http://api.dtapp.net/ip?ua=widows_sg

	"Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/111.0.0.0 Safari/537.36",                   // Chrome 2023-03-22
	"Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/111.0.0.0 Safari/537.36 Edg/111.0.1661.51", // Edge 2023-03-22
	"Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:109.0) Gecko/20100101 Firefox/111.0",                                          // Firefox 2023-03-22

	"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.3 Safari/605.1.15",                         // Safari浏览器 2023-12-30 https://api.dtapp.net/ip?ua=mac_safari
	"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36",                         // Chrome浏览器 2023-12-30 https://api.dtapp.net/ip?ua=mac_chrome
	"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/120.0.0.0 Safari/537.36 Edg/120.0.0.0",           // Edge浏览器 2023-12-30 https://api.dtapp.net/ip?ua=mac_edge
	"Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:121.0) Gecko/20100101 Firefox/121.0",                                                          // Firefox浏览器 2023-12-30 http://api.dtapp.net/ip?ua=mac_firefox
	"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/98.0.4758.102 Safari/537.36 QQBrowser/5.0.4.209", // QQ浏览器 2023-12-30 http://api.dtapp.net/ip?ua=mac_qq
	"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/108.0.5359.95 Safari/537.36",                     // 360极速浏览器 2023-12-30 http://api.dtapp.net/ip?ua=mac_360js
}

// GetRandomUserAgentSystem 获取系统随机UA
func GetRandomUserAgentSystem() string {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	switch runtime.GOOS {
	case linux:
		return userAgentListLinux[r.Intn(len(userAgentListLinux))]
	case windows:
		return userAgentListWindows[r.Intn(len(userAgentListWindows))]
	case mac:
		return userAgentListMac[r.Intn(len(userAgentListMac))]
	default:
		return userAgentListMac[r.Intn(len(userAgentListMac))]
	}
}

var userAgentListWindows = []string{
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/123.0.0.0 Safari/537.36",                                         // Chrome浏览器 2024-04-12
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/123.0.0.0 Safari/537.36 Edg/123.0.0.0",                           // Edge浏览器 2024-04-12
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64; rv:124.0) Gecko/20100101 Firefox/124.0",                                                                        // Firefox浏览器 2024-04-12
	"Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/94.0.4606.71 Safari/537.36 Core/1.94.236.400 QQBrowser/12.4.5604.400", // QQ浏览器 2024-04-12
	"Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.5735.289 Safari/537.36",                                         // 360安全浏览器 2024-04-12
	"Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/86.0.4240.198 Safari/537.36",                                          // 360极速浏览器 2024-04-12
	"Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/94.0.4606.71 Safari/537.36 SE 2.X MetaSr 1.0",                         // 搜狗浏览器 2024-04-12
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.6261.95 Safari/537.36",                                     // 360极速浏览器X 2024-04-12
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.5735.289 Safari/537.36",                                    // 360Ai浏览器 2024-04-12
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/123.0.0.0 Safari/537.36 OPR/109.0.0.0",                           // Opera浏览器 2024-04-12
	"Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/102.0.0.0 Safari/537.36",                                              // Cent浏览器 2024-04-12
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/109.0.0.0 Safari/537.36",                                         // 傲游浏览器 2024-04-12
	"Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.0.0 Safari/537.36",                                         // Vivaldi浏览器 2024-04-12
}

var userAgentListLinux = []string{
	"Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/111.0.0.0 Safari/537.36",                   // Chrome 2023-03-22
	"Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/111.0.0.0 Safari/537.36 Edg/111.0.1661.51", // Edge 2023-03-22
	"Mozilla/5.0 (X11; Ubuntu; Linux x86_64; rv:109.0) Gecko/20100101 Firefox/111.0",                                          // Firefox 2023-03-22
}

var userAgentListMac = []string{
	"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/605.1.15 (KHTML, like Gecko) Version/17.3 Safari/605.1.15",                         // Safari浏览器 2024-04-12
	"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/123.0.0.0 Safari/537.36",                         // Chrome浏览器 2024-04-12
	"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/123.0.0.0 Safari/537.36 Edg/123.0.0.0",           // Edge浏览器 2024-04-12
	"Mozilla/5.0 (Macintosh; Intel Mac OS X 10.15; rv:124.0) Gecko/20100101 Firefox/124.0",                                                          // Firefox浏览器 2024-04-12
	"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/98.0.4758.102 Safari/537.36 QQBrowser/5.0.4.211", // QQ浏览器 2024-04-12
	"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/108.0.5359.95 Safari/537.36",                     // 360极速浏览器Pro 2024-04-12
	"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/123.0.0.0 Safari/537.36 OPR/109.0.0.0",           // Opera浏览器 2024-04-12
	"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/109.0.0.0 Safari/537.36",                         // 傲游浏览器 2024-04-12
	"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.0.0 Safari/537.36",                         // Vivaldi浏览器 2024-04-12
	"Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_7) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/123.0.0.0 Safari/537.36",                         // Arc浏览器 2024-04-12
}
