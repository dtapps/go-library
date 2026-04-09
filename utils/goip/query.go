package goip

import (
	"errors"
	"net"

	"go.dtapp.net/library/utils/goip/czdb"
	"go.dtapp.net/library/utils/goip/geoip"
	"go.dtapp.net/library/utils/goip/ip2region"
)

var (
	QueryIncorrect = errors.New("ip地址不正确")
)

// QueryCzdb cz88 ip库
// https://www.cz88.net/
func (c *Client) QueryCzdb(ipAddress string) (result czdb.QueryResult, err error) {
	ip := net.ParseIP(ipAddress)
	if ip.To4() == nil {
		return result, QueryIncorrect
	}
	query, err := c.czdbClient.Query(ipAddress)
	if err != nil {
		return czdb.QueryResult{}, err
	}
	return query, err
}

// QueryGeoIp GeoIP
// https://www.maxmind.com/
func (c *Client) QueryGeoIp(ipAddress string) (result geoip.QueryCityResult, err error) {
	query, err := c.geoIpClient.QueryCity(ipAddress)
	if err != nil {
		return geoip.QueryCityResult{}, err
	}
	return query, nil
}

// QueryIpRegion ip2region
// https://ip2region.net/
func (c *Client) QueryIpRegion(ipAddress string) (result ip2region.QueryResult, err error) {
	query, err := c.ipRegionClient.Query(ipAddress)
	if err != nil {
		return ip2region.QueryResult{}, err
	}
	return query, nil
}
