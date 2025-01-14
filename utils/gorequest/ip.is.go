package gorequest

import (
	"net"
	"strings"
)

// IpIs 是否ip
func IpIs(ipStr string) string {

	if ipStr == "" {
		return ""
	}

	// ipv4
	if strings.Contains(ipStr, "/32") {
		cidr, _, _ := net.ParseCIDR(ipStr)
		if cidr != nil {
			return cidr.String()
		}
	}

	// ipv6
	if strings.Contains(ipStr, "/128") {
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

// IpIsConsistent 两个ip是否一致
func IpIsConsistent(ipStr1, ipStr2 string) bool {

	ip1Result := IpIs(ipStr1)
	ip2Result := IpIs(ipStr2)

	if ip1Result != "" && ip2Result != "" {
		if ip1Result == ip2Result {
			return true
		}
	}

	return false
}
