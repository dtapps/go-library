package goip

import (
	"github.com/dtapps/go-library/utils/goip/geoip"
	"github.com/dtapps/go-library/utils/goip/ip2region"
	"github.com/dtapps/go-library/utils/goip/ip2region_v2"
	"github.com/dtapps/go-library/utils/goip/ipv6wry"
	"github.com/dtapps/go-library/utils/goip/qqwry"
)

type ClientConfig struct {
	Ip2regionPath    string
	Ip2regionByte    []byte
	Ip2regionV2Path  string
	Ip2regionV2Byte  []byte
	QqwryPath        string
	QqwryByte        []byte
	Ipv6wryPath      string
	Ipv6wryByte      []byte
	GeoipAsnPath     string
	GeoipAsnByte     []byte
	GeoipCityPath    string
	GeoipCityByte    []byte
	GeoipCountryPath string
	GeoipCountryByte []byte
}

type Client struct {
	ip2regionV2Client *ip2region_v2.Client
	ip2regionClient   *ip2region.Client
	qqwryClient       *qqwry.Client
	geoIpClient       *geoip.Client
	ipv6wryClient     *ipv6wry.Client
}

// NewIp 实例化
func NewIp(config ClientConfig) (*Client, error) {

	var err error
	c := &Client{}

	if config.Ip2regionV2Path == "" {
		c.ip2regionV2Client, err = ip2region_v2.NewBuff(config.Ip2regionV2Byte)
		if err != nil {
			return nil, err
		}
	} else {
		c.ip2regionV2Client, err = ip2region_v2.New(config.Ip2regionV2Path)
		if err != nil {
			return nil, err
		}
	}

	if config.Ip2regionPath == "" {
		c.ip2regionClient, err = ip2region.NewBuff(config.Ip2regionByte)
		if err != nil {
			return nil, err
		}
	} else {
		c.ip2regionClient, err = ip2region.New(config.Ip2regionPath)
		if err != nil {
			return nil, err
		}
	}

	if config.QqwryPath == "" {
		c.qqwryClient, err = qqwry.NewBuff(config.QqwryByte)
		if err != nil {
			return nil, err
		}
	} else {
		c.qqwryClient, err = qqwry.New(config.QqwryPath)
		if err != nil {
			return nil, err
		}
	}

	if config.GeoipAsnPath == "" || config.GeoipCityPath == "" || config.GeoipCountryPath == "" {
		c.geoIpClient, err = geoip.NewBuff(config.GeoipAsnByte, config.GeoipCityByte, config.GeoipCountryByte)
		if err != nil {
			return nil, err
		}
	} else {
		c.geoIpClient, err = geoip.New(config.GeoipAsnPath, config.GeoipCityPath, config.GeoipCountryPath)
		if err != nil {
			return nil, err
		}
	}

	if config.Ipv6wryPath == "" {
		c.ipv6wryClient, err = ipv6wry.NewBuff(config.Ipv6wryByte)
		if err != nil {
			return nil, err
		}
	} else {
		c.ipv6wryClient, err = ipv6wry.New(config.Ipv6wryPath)
		if err != nil {
			return nil, err
		}
	}

	return c, nil
}

func (c *Client) Close() {
	c.geoIpClient.Close()
}
