package geoip

import (
	_ "embed"
	"net"
)

// QueryCityResult 返回
type QueryCityResult struct {
	Ip        string `json:"ip,omitempty"` // ip
	Continent struct {
		Code string `json:"code,omitempty"` // 大陆代码
		Name string `json:"name,omitempty"` // 大陆名称
	} `json:"continent,omitempty"`
	Country struct {
		Code string `json:"code,omitempty"` // 国家代码
		Name string `json:"name,omitempty"` // 国家名称
	} `json:"country,omitempty"`
	Province struct {
		Code string `json:"code,omitempty"` // 省份代码
		Name string `json:"name,omitempty"` // 省份名称
	} `json:"province,omitempty"`
	City struct {
		Name string `json:"name,omitempty"` // 城市名称
	} `json:"city,omitempty"`
	Location struct {
		TimeZone  string  `json:"time_zone,omitempty"` // 位置时区
		Latitude  float64 `json:"latitude,omitempty"`  // 坐标纬度
		Longitude float64 `json:"longitude,omitempty"` // 坐标经度
	} `json:"location,omitempty"`
}

func (c *Client) QueryCity(ipAddress net.IP) (result QueryCityResult, err error) {

	record, err := c.cityDb.City(ipAddress)
	if err != nil {
		return QueryCityResult{}, err
	}

	// ip
	result.Ip = ipAddress.String()

	// 大陆
	result.Continent.Code = record.Continent.Code
	result.Continent.Name = record.Continent.Names["zh-CN"]

	// 国家
	result.Country.Code = record.Country.IsoCode
	result.Country.Name = record.Country.Names["zh-CN"]

	// 省份
	if len(record.Subdivisions) > 0 {
		result.Province.Code = record.Subdivisions[0].IsoCode
		result.Province.Name = record.Subdivisions[0].Names["zh-CN"]
	}

	// 城市
	result.City.Name = record.City.Names["zh-CN"]

	// 位置
	result.Location.TimeZone = record.Location.TimeZone
	result.Location.Latitude = record.Location.Latitude
	result.Location.Longitude = record.Location.Longitude

	return result, err
}
