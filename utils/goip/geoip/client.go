package geoip

import (
	"github.com/oschwald/geoip2-golang"
)

type Client struct {
	asnFilepath     string
	asnDb           *geoip2.Reader
	cityFilepath    string
	cityDb          *geoip2.Reader
	countryFilepath string
	countryDb       *geoip2.Reader
}

func New(asnFilepath string, cityFilepath string, countryFilepath string) (*Client, error) {

	var err error
	c := &Client{
		asnFilepath:     asnFilepath,
		cityFilepath:    cityFilepath,
		countryFilepath: countryFilepath,
	}

	if asnFilepath != "" {
		c.asnDb, err = geoip2.Open(asnFilepath)
		if err != nil {
			return nil, err
		}
	}

	c.cityDb, err = geoip2.Open(cityFilepath)
	if err != nil {
		return nil, err
	}

	if countryFilepath != "" {
		c.countryDb, err = geoip2.Open(countryFilepath)
		if err != nil {
			return nil, err
		}
	}

	return c, err
}

func (c *Client) Close() {
	if c.asnFilepath != "" {
		c.asnDb.Close()
	}
	c.cityDb.Close()
	if c.countryFilepath != "" {
		c.countryDb.Close()
	}
}
