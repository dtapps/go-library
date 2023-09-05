package geoip

import (
	"github.com/oschwald/geoip2-golang"
	"os"
)

var asnBuff []byte

var cityBuff []byte

var countryBuff []byte

type Client struct {
	asnDb     *geoip2.Reader
	cityDb    *geoip2.Reader
	countryDb *geoip2.Reader
}

func New(asnFilepath string, cityFilepath string, countryFilepath string) (*Client, error) {

	var err error
	c := &Client{}

	asnBuff, err = os.ReadFile(asnFilepath)
	if err != nil {
		return nil, err
	}
	c.asnDb, err = geoip2.FromBytes(asnBuff)
	if err != nil {
		return nil, err
	}

	cityBuff, err = os.ReadFile(cityFilepath)
	if err != nil {
		return nil, err
	}
	c.cityDb, err = geoip2.FromBytes(cityBuff)
	if err != nil {
		return nil, err
	}

	countryBuff, err = os.ReadFile(countryFilepath)
	if err != nil {
		return nil, err
	}
	c.countryDb, err = geoip2.FromBytes(countryBuff)
	if err != nil {
		return nil, err
	}

	return c, err
}

func NewBuff(asnFile []byte, cityFile []byte, countryFile []byte) (*Client, error) {

	var err error
	c := &Client{}

	asnBuff = asnFile
	c.asnDb, err = geoip2.FromBytes(asnFile)
	if err != nil {
		return nil, err
	}

	cityBuff = cityFile
	c.cityDb, err = geoip2.FromBytes(cityFile)
	if err != nil {
		return nil, err
	}

	countryBuff = countryFile
	c.countryDb, err = geoip2.FromBytes(countryFile)
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
