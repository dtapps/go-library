package goip

import (
	"github.com/dtapps/go-library/utils/gostring"
	"net"
	"strings"
)

var (
	ipv4 = "IPV4"
	ipv6 = "IPV6"
)

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

// IsIp 是否ip
func IsIp(ipStr string) string {

	if ipStr == "" {
		return ""
	}

	// ipv4
	if gostring.Contains(ipStr, "/32") {
		cidr, _, _ := net.ParseCIDR(ipStr)
		if cidr != nil {
			return cidr.String()
		}
	}

	// ipv6
	if gostring.Contains(ipStr, "/128") {
		cidr, _, _ := net.ParseCIDR(ipStr)
		if cidr != nil {
			return cidr.String()
		}
	}

	// 解析
	result := net.ParseIP(ipStr).String()
	if result != "<nil>" {
		return result
	}

	return ""
}

// IsIpConsistent 两个ip是否一致
func IsIpConsistent(ipStr1, ipStr2 string) bool {

	ip1Result := IsIp(ipStr1)
	ip2Result := IsIp(ipStr2)

	if ip1Result != "" && ip2Result != "" {
		if ip1Result == ip2Result {
			return true
		}
	}

	return false
}
