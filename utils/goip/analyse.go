package goip

import (
	"strconv"
)

var (
	ipv4 = "IPV4"
	ipv6 = "IPV6"
)

type AnalyseResult struct {
	IP       string `json:"ip,omitempty"`       // 输入的ip地址
	Country  string `json:"country,omitempty"`  // 国家或地区
	Province string `json:"province,omitempty"` // 省份
	City     string `json:"city,omitempty"`     // 城市
	Area     string `json:"area,omitempty"`     // 区域
	Isp      string `json:"isp,omitempty"`      // 运营商
}

func (c *Client) Analyse(item string) AnalyseResult {
	isIp := c.isIpv4OrIpv6(item)
	switch isIp {
	case ipv4:
		info := c.V4db.Find(item)
		search, err := c.V4Region.MemorySearch(item)
		if err != nil {
			return AnalyseResult{
				IP:      info.IP,
				Country: info.Country,
				Area:    info.Area,
			}
		} else {
			return AnalyseResult{
				IP:       search.IP,
				Country:  search.Country,
				Province: search.Province,
				City:     search.City,
				Isp:      info.Area,
			}
		}
	case ipv6:
		info := c.V6db.Find(item)
		return AnalyseResult{
			IP:       info.IP,
			Country:  info.Country,
			Province: info.Province,
			City:     info.City,
			Area:     info.Area,
			Isp:      info.Isp,
		}
	default:
		return AnalyseResult{}
	}
}

// CheckIpv4 检查数据是不是IPV4
func (c *Client) CheckIpv4(ips string) bool {
	if len(ips) > 3 {
		return false
	}
	nums, err := strconv.Atoi(ips)
	if err != nil {
		return false
	}
	if nums < 0 || nums > 255 {
		return false
	}
	if len(ips) > 1 && ips[0] == '0' {
		return false
	}
	return true
}

// CheckIpv6 检测是不是IPV6
func (c *Client) CheckIpv6(ips string) bool {
	if ips == "" {
		return true
	}
	if len(ips) > 4 {
		return false
	}
	for _, val := range ips {
		if !((val >= '0' && val <= '9') || (val >= 'a' && val <= 'f') || (val >= 'A' && val <= 'F')) {
			return false
		}
	}
	return true
}
