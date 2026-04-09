package geoip

import (
	"fmt"

	"github.com/oschwald/geoip2-golang/v2"
)

// Client 客户端
type Client struct {
	asnFilepath     string
	asnDb           *geoip2.Reader
	cityFilepath    string
	cityDb          *geoip2.Reader
	countryFilepath string
	countryDb       *geoip2.Reader
}

// New 创建
func New(asnFilepath string, cityFilepath string, countryFilepath string) (*Client, error) {

	var err error
	c := &Client{
		asnFilepath:     asnFilepath,
		cityFilepath:    cityFilepath,
		countryFilepath: countryFilepath,
	}

	if c.asnFilepath == "" && c.cityFilepath == "" && c.countryFilepath == "" {
		return nil, fmt.Errorf("asnFilepath, cityFilepath, and countryFilepath are empty")
	}

	if asnFilepath != "" {
		c.asnDb, err = geoip2.Open(asnFilepath)
		if err != nil {
			return nil, err
		}
	}

	if cityFilepath != "" {
		c.cityDb, err = geoip2.Open(cityFilepath)
		if err != nil {
			return nil, err
		}
	}

	if countryFilepath != "" {
		c.countryDb, err = geoip2.Open(countryFilepath)
		if err != nil {
			return nil, err
		}
	}

	return c, err
}

// Close 关闭
func (c *Client) Close() {
	if c.asnFilepath != "" {
		c.asnDb.Close()
	}
	if c.cityFilepath != "" {
		c.cityDb.Close()
	}
	if c.countryFilepath != "" {
		c.countryDb.Close()
	}
}
