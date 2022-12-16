package goip

import (
	"github.com/dtapps/go-library/utils/goip/geoip"
	"github.com/dtapps/go-library/utils/goip/ip2region"
	"github.com/dtapps/go-library/utils/goip/ip2region_v2"
	"github.com/dtapps/go-library/utils/goip/ipv6wry"
	"github.com/dtapps/go-library/utils/goip/qqwry"
)

type Client struct {
	ip2regionV2Client *ip2region_v2.Client
	ip2regionClient   *ip2region.Client
	qqwryClient       *qqwry.Client
	geoIpClient       *geoip.Client
	ipv6wryClient     *ipv6wry.Client
}

// NewIp 实例化
func NewIp() *Client {

	c := &Client{}

	c.ip2regionV2Client, _ = ip2region_v2.New()

	c.ip2regionClient = ip2region.New()

	c.qqwryClient = qqwry.New()

	c.geoIpClient, _ = geoip.New()

	c.ipv6wryClient = ipv6wry.New()

	return c
}

func (c *Client) Close() {
	c.geoIpClient.Close()
}
