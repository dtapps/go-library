package ipv6wry

import (
	"errors"
	"github.com/dtapps/go-library/utils/gostring"
	"math/big"
	"net"
	"strings"
)

// QueryResult 返回
type QueryResult struct {
	Ip       string `json:"ip,omitempty"`       // ip
	Country  string `json:"country,omitempty"`  // 国家
	Province string `json:"province,omitempty"` // 省份
	City     string `json:"city,omitempty"`     // 城市
	Area     string `json:"area,omitempty"`     // 区域
	Isp      string `json:"isp,omitempty"`      // 运营商
}

// Query ip地址查询对应归属地信息
func (c *Client) Query(ipAddress net.IP) (result QueryResult, err error) {

	result.Ip = ipAddress.String()

	c.Offset = 0

	tp := big.NewInt(0)
	op := big.NewInt(0)
	tp.SetBytes(ipAddress.To16())
	op.SetString("18446744073709551616", 10)
	op.Div(tp, op)
	tp.SetString("FFFFFFFFFFFFFFFF", 16)
	op.And(op, tp)

	v6ip = op.Uint64()
	offset = c.searchIndex(v6ip)
	c.Offset = offset

	country, area = c.getAddr()

	// 解析地区数据
	info := strings.Split(string(country), "\t")
	if len(info) > 0 {
		i := 1
		for {
			if i > len(info) {
				break
			}
			switch i {
			case 1:
				result.Country = info[i-1]
				result.Country = gostring.SpaceAndLineBreak(result.Country)
			case 2:
				result.Province = info[i-1]
				result.Province = gostring.SpaceAndLineBreak(result.Province)
			case 3:
				result.City = info[i-1]
				result.City = gostring.SpaceAndLineBreak(result.City)
			case 4:
				result.Area = info[i-1]
				result.Area = gostring.SpaceAndLineBreak(result.Area)
			}
			i++ // 自增
		}
	} else {
		result.Country = string(country)
		result.Country = gostring.SpaceAndLineBreak(result.Country)
	}
	// 运营商
	result.Isp = string(area)

	// Delete ZX (防止不相关的信息产生干扰）
	if result.Isp == "ZX" || result.Isp == "" {
		result.Isp = ""
	} else {
		result.Isp = " " + result.Isp
	}

	result.Isp = gostring.SpaceAndLineBreak(result.Isp)

	return result, nil
}

// QueryIP ip地址查询对应归属地信息
func (c *Client) QueryIP(ipAddressStr string) (result QueryResult, err error) {
	arrIpv6 := strings.Split(ipAddressStr, ":")
	if len(arrIpv6) == 8 {
		return c.Query(net.ParseIP(ipAddressStr))
	}
	return QueryResult{}, errors.New("不是IPV6")
}
