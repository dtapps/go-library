package qqwry

import (
	"encoding/binary"
	"errors"
	"github.com/dtapps/go-library/utils/gostring"
	"golang.org/x/text/encoding/simplifiedchinese"
	"net"
	"strings"
)

// QueryResult 返回
type QueryResult struct {
	Ip      string `json:"ip,omitempty"`      // ip
	Country string `json:"country,omitempty"` // 国家或地区
	Area    string `json:"area,omitempty"`    // 区域
}

// Query ip地址查询对应归属地信息
func (c *Client) Query(ipAddress net.IP) (result QueryResult, err error) {

	c.Offset = 0

	// 偏移
	offset = c.searchIndex(binary.BigEndian.Uint32(ipAddress.To4()))
	if offset <= 0 {
		return QueryResult{}, errors.New("搜索失败")
	}

	result.Ip = ipAddress.String()

	c.Offset = offset + c.ItemLen

	country, area = c.getAddr()

	enc := simplifiedchinese.GBK.NewDecoder()

	result.Country, _ = enc.String(string(country))

	result.Country = gostring.SpaceAndLineBreak(result.Country)

	result.Area, _ = enc.String(string(area))

	// Delete CZ88.NET (防止不相关的信息产生干扰）
	if result.Area == " CZ88.NET" || result.Area == "" {
		result.Area = ""
	} else {
		result.Area = " " + result.Area
	}

	result.Area = gostring.SpaceAndLineBreak(result.Area)

	return result, nil
}

// QueryIP ip地址查询对应归属地信息
func (c *Client) QueryIP(ipAddressStr string) (result QueryResult, err error) {
	arrIpv4 := strings.Split(ipAddressStr, ".")
	if len(arrIpv4) == 4 {
		return c.Query(net.ParseIP(ipAddressStr))
	}
	return QueryResult{}, errors.New("不是IPV4")
}
