package goip

import (
	"errors"
	"go.dtapp.net/library/utils/goip/geoip"
	"go.dtapp.net/library/utils/goip/qqwry"
	"net"
)

var (
	QueryIncorrect = errors.New("ip地址不正确")
)

// QueryQqWry 纯真IP库
// https://www.cz88.net/
func (c *Client) QueryQqWry(ipAddress string) (result qqwry.QueryResult, err error) {
	ip := net.ParseIP(ipAddress)
	if ip.To4() == nil {
		return result, QueryIncorrect
	}
	query, err := c.qqwryClient.Query(ipAddress)
	if err != nil {
		return qqwry.QueryResult{}, err
	}
	return query, err
}

// QueryGeoIp ip2region
// https://www.maxmind.com/
func (c *Client) QueryGeoIp(ipAddress string) (result geoip.QueryCityResult, err error) {
	query, err := c.geoIpClient.QueryCity(ipAddress)
	if err != nil {
		return geoip.QueryCityResult{}, err
	}
	return query, nil
}
