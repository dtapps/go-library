package gorequest

import (
	"context"
	"errors"
	"net"
	"net/http"
	"strings"
	"time"
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

// Ips 获取全部网卡的全部IP
func Ips(ctx context.Context) (map[string]string, error) {

	ips := make(map[string]string)

	//返回 interface 结构体对象的列表，包含了全部网卡信息
	interfaces, err := net.Interfaces()
	if err != nil {
		return nil, err
	}

	//遍历全部网卡
	for _, i := range interfaces {

		// Addrs() 方法返回一个网卡上全部的IP列表
		address, err := i.Addrs()
		if err != nil {
			return nil, err
		}

		//遍历一个网卡上全部的IP列表，组合为一个字符串，放入对应网卡名称的map中
		for _, v := range address {
			ips[i.Name] += v.String() + " "
		}
	}
	return ips, nil
}

// GetMacAddr 获取Mac地址
func GetMacAddr(ctx context.Context) (arrays []string) {
	netInterfaces, err := net.Interfaces()
	if err != nil {
		return arrays
	}
	for _, netInterface := range netInterfaces {
		addr := netInterface.HardwareAddr.String()
		if len(addr) == 0 {
			continue
		}

		arrays = append(arrays, addr)
	}
	return arrays
}

// GetInsideIp 内网IP
func GetInsideIp(ctx context.Context) string {

	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		panic(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)
	return localAddr.IP.String()
}

// GetOutsideIp 获取公网 IP，任一服务成功即返回
// 成功返回 IP 和 nil；失败返回 0.0.0.0 和 error
func GetOutsideIp(ctx context.Context) (string, error) {

	// 总超时 5 秒
	ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
	defer cancel()

	client := &http.Client{
		Timeout: 3 * time.Second,
	}

	result := make(chan string, 1) // 缓冲 1，避免 goroutine 泄漏

	// 启动所有 goroutine
	for _, url := range ipServices {
		go fetchIP(ctx, client, url, result)
	}

	select {
	case ip := <-result:
		return ip, nil
	case <-ctx.Done():
		return "", errors.New("获取公网 IP 超时")
	}
}

// GetOutsideIPV4All 外网IPV4地址
func GetOutsideIPV4All(ctx context.Context) string {
	ipv4 := getIPV4_DtappNet(ctx)
	if ipv4 != "" {
		return ipv4
	}
	ipv4 = getIPV4_MyipIpipNet(ctx)
	if ipv4 != "" {
		return ipv4
	}
	ipv4 = getIPV4_DdnsOrayCom(ctx)
	if ipv4 != "" {
		return ipv4
	}
	ipv4 = getIPV4_Ip3322Net(ctx)
	if ipv4 != "" {
		return ipv4
	}
	ipv4 = getIPV4_4IpwCn(ctx)
	if ipv4 != "" {
		return ipv4
	}
	return getCmdIPV4()
}

// GetOutsideIPV6All 外网IPV6地址
func GetOutsideIPV6All(ctx context.Context) string {
	ipv6 := getIPV4_DtappNet(ctx)
	if ipv6 != "" {
		return ipv6
	}
	ipv6 = getIPV6_SpeedNeu6EduCn(ctx)
	if ipv6 != "" {
		return ipv6
	}
	ipv6 = getIPV6_V6IdentMe(ctx)
	if ipv6 != "" {
		return ipv6
	}
	ipv6 = getIPV6_6IpwCn(ctx)
	if ipv6 != "" {
		return ipv6
	}
	return getCmdIPV6()
}

// GetCmdOutsideIP 通过命令获取外网IP
func GetCmdOutsideIP() string {
	ipv6 := getCmdIPV6()
	if ipv6 != "" {
		return ipv6
	}
	return getCmdIPV4()
}
