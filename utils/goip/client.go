package goip

import (
	"errors"
	"go.dtapp.net/library/utils/goip/geoip"
	"go.dtapp.net/library/utils/goip/qqwry"
)

type ClientConfig struct {
	GeoipAsnPath     string
	GeoipCityPath    string
	GeoipCountryPath string
	QqwryPath        string
}

type Client struct {
	geoIpClient *geoip.Client
	qqwryClient *qqwry.Client
	config      *ClientConfig
}

// NewIp 实例化
func NewIp(config *ClientConfig) (*Client, error) {

	var err error
	c := &Client{config: config}

	if config.GeoipCityPath == "" {
		return nil, errors.New("请配置 GeoipCityPath 文件")
	}
	c.geoIpClient, err = geoip.New(config.GeoipAsnPath, config.GeoipCityPath, config.GeoipCountryPath)
	if err != nil {
		return nil, err
	}

	if config.QqwryPath != "" {
		c.qqwryClient, err = qqwry.New(config.QqwryPath)
		if err != nil {
			return nil, err
		}
	}

	return c, nil
}

func (c *Client) Close() {
	c.geoIpClient.Close()
}
