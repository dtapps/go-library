package gorequest

import "net"

func IsIPV4(s string) bool {
	ip := net.ParseIP(s)
	if ip == nil {
		// 不是合法的IP地址
		return false
	}

	if ip.To4() != nil {
		return true
	} else if ip.To16() != nil {
		return false
	}
	return false
}

func IsIPv4Public(ip net.IP) bool {
	// 判断是否为回环地址或私有地址
	if ip.IsLoopback() {
		return false
	}

	// 判断是否为保留地址
	reserved := []net.IPNet{
		{IP: net.ParseIP("0.0.0.0"), Mask: net.CIDRMask(8, 32)},       // 0.0.0.0/8
		{IP: net.ParseIP("10.0.0.0"), Mask: net.CIDRMask(8, 32)},      // 10.0.0.0/8
		{IP: net.ParseIP("127.0.0.0"), Mask: net.CIDRMask(8, 32)},     // 127.0.0.0/8
		{IP: net.ParseIP("169.254.0.0"), Mask: net.CIDRMask(16, 32)},  // 169.254.0.0/16
		{IP: net.ParseIP("172.16.0.0"), Mask: net.CIDRMask(12, 32)},   // 172.16.0.0/12
		{IP: net.ParseIP("192.0.0.0"), Mask: net.CIDRMask(24, 32)},    // 192.0.0.0/24
		{IP: net.ParseIP("192.0.2.0"), Mask: net.CIDRMask(24, 32)},    // 192.0.2.0/24
		{IP: net.ParseIP("192.88.99.0"), Mask: net.CIDRMask(24, 32)},  // 192.88.99.0/24
		{IP: net.ParseIP("192.168.0.0"), Mask: net.CIDRMask(16, 32)},  // 192.168.0.0/16
		{IP: net.ParseIP("198.18.0.0"), Mask: net.CIDRMask(15, 32)},   // 198.18.0.0/15
		{IP: net.ParseIP("198.51.100.0"), Mask: net.CIDRMask(24, 32)}, // 198.51.100.0/24
		{IP: net.ParseIP("203.0.113.0"), Mask: net.CIDRMask(24, 32)},  // 203.0.113.0/24
		{IP: net.ParseIP("224.0.0.0"), Mask: net.CIDRMask(4, 32)},     // 224.0.0.0/4
	}

	for _, r := range reserved {
		if r.Contains(ip) {
			return false
		}
	}

	// 如果不是回环地址、私有地址或保留地址，则认为是公网地址
	return true
}

func IsIPV6(s string) bool {
	ip := net.ParseIP(s)
	if ip == nil {
		// 不是合法的IP地址
		return false
	}

	if ip.To4() != nil {
		return false
	} else if ip.To16() != nil {
		return true
	}
	return false
}

func IsIPv6Public(ip net.IP) bool {
	// 获取IPv6地址的前缀
	prefix := ip[0] & 0xF0

	// 判断前缀是否位于公网地址范围内
	return prefix == 0x20 || // 2000::/3
		(prefix == 0x20 && ip[1] == 0x01) || // 2001::/16
		(prefix == 0x20 && ip[1]&0xF0 == 0x20) // 2000::/4
}
