package goip

import (
	"go.dtapp.net/goip/ip2region"
	v4 "go.dtapp.net/goip/v4"
	v6 "go.dtapp.net/goip/v6"
	"log"
	"strings"
)

type Client struct {
	V4Region ip2region.Ip2Region // IPV4
	V4db     v4.Pointer          // IPV4
	V6db     v6.Pointer          // IPV6
}

// NewIp 实例化
func NewIp() *Client {
	app := &Client{}
	v4Num := app.V4db.InitIPV4Data()
	log.Printf("IPV4 库加载完成 共加载:%d 条 IP 记录\n", v4Num)
	v6Num := app.V6db.InitIPV4Data()
	log.Printf("IPV6 库加载完成 共加载:%d 条 IP 记录\n", v6Num)
	return app
}

func (c *Client) Ipv4(ip string) (res v4.Result, resInfo ip2region.IpInfo) {
	res = c.V4db.Find(ip)
	resInfo, _ = c.V4Region.MemorySearch(ip)
	return res, resInfo
}

func (c *Client) Ipv6(ip string) (res v6.Result) {
	res = c.V6db.Find(ip)
	return res
}

func (c *Client) isIpv4OrIpv6(ip string) string {
	if len(ip) < 7 {
		return ""
	}
	arrIpv4 := strings.Split(ip, ".")
	if len(arrIpv4) == 4 {
		//. 判断IPv4
		for _, val := range arrIpv4 {
			if !c.CheckIpv4(val) {
				return ""
			}
		}
		return ipv4
	}
	arrIpv6 := strings.Split(ip, ":")
	if len(arrIpv6) == 8 {
		// 判断Ipv6
		for _, val := range arrIpv6 {
			if !c.CheckIpv6(val) {
				return "Neither"
			}
		}
		return ipv6
	}
	return ""
}
