package goip

import (
	"github.com/tagphi/czdb-search-golang/pkg/db"
	"go.dtapp.net/library/utils/goip/czdb"
	"go.dtapp.net/library/utils/goip/geoip"
	"go.dtapp.net/library/utils/goip/ip2region"
)

type ClientConfig struct {
	CzdbV4Path       string
	CzdbV6Path       string
	CzdbKey          string
	GeoipAsnPath     string
	GeoipCityPath    string
	GeoipCountryPath string
	IpRegionV4Path   string
	IpRegionV6Path   string
}

type Client struct {
	config         *ClientConfig
	czdbClient     *czdb.Client
	geoIpClient    *geoip.Client
	ipRegionClient *ip2region.Client
}

// NewIp 实例化
func NewIp(config *ClientConfig) (*Client, error) {

	var err error
	c := &Client{config: config}

	if (config.CzdbV4Path != "" || config.CzdbV6Path != "") && config.CzdbKey != "" {
		c.czdbClient, err = czdb.New(config.CzdbV4Path, config.CzdbV6Path, config.CzdbKey, db.MEMORY)
		if err != nil {
			return nil, err
		}
	}

	if config.GeoipCityPath != "" {
		c.geoIpClient, err = geoip.New(config.GeoipAsnPath, config.GeoipCityPath, config.GeoipCountryPath)
		if err != nil {
			return nil, err
		}
	}

	if config.IpRegionV4Path != "" || config.IpRegionV6Path != "" {
		c.ipRegionClient, err = ip2region.New(config.IpRegionV4Path, config.IpRegionV6Path)
		if err != nil {
			return nil, err
		}
	}

	return c, nil
}

func (c *Client) Close() {
	if c.config.GeoipCityPath != "" {
		c.geoIpClient.Close()
	}
	if (c.config.CzdbV4Path != "" || c.config.CzdbV6Path != "") && c.config.CzdbKey != "" {
		c.czdbClient.Close()
	}
	if c.config.IpRegionV4Path != "" || c.config.IpRegionV6Path != "" {
		c.ipRegionClient.Close()
	}
}
