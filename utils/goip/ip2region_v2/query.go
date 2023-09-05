package ip2region_v2

import (
	_ "embed"
	"errors"
	"github.com/dtapps/go-library/utils/gostring"
	"net"
	"strings"
)

// QueryResult 返回
type QueryResult struct {
	Ip       string `json:"ip,omitempty"`       // ip
	Country  string `json:"country,omitempty"`  // 国家
	Province string `json:"province,omitempty"` // 省份
	City     string `json:"city,omitempty"`     // 城市
	Operator string `json:"operator,omitempty"` // 运营商
}

func (c *Client) Query(ipAddress net.IP) (result QueryResult, err error) {

	// 备注：并发使用，用整个 xdb 缓存创建的 searcher 对象可以安全用于并发。

	str, err := c.db.SearchByStr(ipAddress.String())
	if err != nil {
		return QueryResult{}, err
	}

	split := gostring.Split(str, "|")
	if len(split) <= 0 {
		return QueryResult{}, err
	}

	result.Ip = ipAddress.String()

	result.Country = split[0]
	if result.Country == "0" {
		result.Country = ""
	}
	result.Province = split[2]
	if result.Province == "0" {
		result.Province = ""
	}
	result.City = split[3]
	if result.City == "0" {
		result.City = ""
	}
	result.Operator = split[4]
	if result.Operator == "0" {
		result.Operator = ""
	}

	return result, err
}

// QueryIP ip地址查询对应归属地信息
func (c *Client) QueryIP(ipAddressStr string) (result QueryResult, err error) {
	arrIpv4 := strings.Split(ipAddressStr, ".")
	if len(arrIpv4) == 4 {
		return c.Query(net.ParseIP(ipAddressStr))
	}
	return QueryResult{}, errors.New("不是IPV4")
}
