package geoip

import (
	_ "embed"
	"github.com/oschwald/geoip2-golang"
)

//go:embed GeoLite2-ASN.mmdb
var asnBuff []byte

//go:embed GeoLite2-City.mmdb
var cityBuff []byte

//go:embed GeoLite2-Country.mmdb
var countryBuff []byte

type Client struct {
	asnDb     *geoip2.Reader
	cityDb    *geoip2.Reader
	countryDb *geoip2.Reader
}

func New() (*Client, error) {

	var err error
	c := &Client{}

	c.asnDb, err = geoip2.FromBytes(asnBuff)
	if err != nil {
		return nil, err
	}

	c.cityDb, err = geoip2.FromBytes(cityBuff)
	if err != nil {
		return nil, err
	}

	c.countryDb, err = geoip2.FromBytes(countryBuff)
	if err != nil {
		return nil, err
	}

	return c, err
}

func (c *Client) Close() {

	c.asnDb.Close()
	c.cityDb.Close()
	c.countryDb.Close()

}
