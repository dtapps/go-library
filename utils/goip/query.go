package goip

import (
	"errors"
	"github.com/dtapps/go-library/utils/goip/geoip"
	"github.com/dtapps/go-library/utils/goip/ip2region"
	"github.com/dtapps/go-library/utils/goip/ip2region_v2"
	"github.com/dtapps/go-library/utils/goip/ipv6wry"
	"github.com/dtapps/go-library/utils/goip/qqwry"
	"net"
)

var (
	QueryIncorrect = errors.New("ip地址不正确")
)

// QueryQqWry 纯真IP库
// https://www.cz88.net/
func (c *Client) QueryQqWry(ipAddress net.IP) (result qqwry.QueryResult, err error) {
	if ipAddress.To4() == nil {
		return result, QueryIncorrect
	}

	query, err := c.qqwryClient.Query(ipAddress)
	if err != nil {
		return qqwry.QueryResult{}, err
	}

	return query, err
}

// QueryIp2Region ip2region
// https://github.com/lionsoul2014/ip2region
func (c *Client) QueryIp2Region(ipAddress net.IP) (result ip2region.QueryResult, err error) {
	if ipAddress.To4() == nil {
		return result, QueryIncorrect
	}

	query, err := c.ip2regionClient.Query(ipAddress)
	if err != nil {
		return ip2region.QueryResult{}, err
	}

	return query, err
}

// QueryIp2RegionV2 ip2region
// https://github.com/lionsoul2014/ip2region
func (c *Client) QueryIp2RegionV2(ipAddress net.IP) (result ip2region_v2.QueryResult, err error) {
	if ipAddress.To4() == nil {
		return result, QueryIncorrect
	}

	query, err := c.ip2regionV2Client.Query(ipAddress)
	if err != nil {
		return ip2region_v2.QueryResult{}, err
	}

	return query, nil
}

// QueryGeoIp ip2region
// https://www.maxmind.com/
func (c *Client) QueryGeoIp(ipAddress net.IP) (result geoip.QueryCityResult, err error) {
	if ipAddress.String() == "<nil>" {
		return result, QueryIncorrect
	}

	query, err := c.geoIpClient.QueryCity(ipAddress)
	if err != nil {
		return geoip.QueryCityResult{}, err
	}

	return query, nil
}

// QueryIpv6wry ip2region
// https://ip.zxinc.org
func (c *Client) QueryIpv6wry(ipAddress net.IP) (result ipv6wry.QueryResult, err error) {
	if ipAddress.To16() == nil {
		return result, QueryIncorrect
	}

	query, err := c.ipv6wryClient.Query(ipAddress)
	if err != nil {
		return ipv6wry.QueryResult{}, err
	}

	return query, nil
}
