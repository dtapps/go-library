package goip

import (
	"go.dtapp.net/library/utils/goip/czdb"
	"go.dtapp.net/library/utils/goip/geoip"
	"go.dtapp.net/library/utils/goip/ip2region"
)

func (c *Client) GetCzdb() *czdb.Client {
	return c.czdbClient
}

func (c *Client) GetGeoIp() *geoip.Client {
	return c.geoIpClient
}

func (c *Client) GetIpRegion() *ip2region.Client {
	return c.ipRegionClient
}
