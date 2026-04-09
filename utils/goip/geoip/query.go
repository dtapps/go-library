package geoip

import (
	"net/netip"
)

// QueryCityResult 返回
type QueryCityResult struct {
	Ip        string `json:"ip,omitempty"` // ip
	Continent struct {
		Code string `json:"code,omitempty"` // 大陆代码
		Name string `json:"name,omitempty"` // 大陆名称
	} `json:"continent"`
	Country struct {
		Code string `json:"code,omitempty"` // 国家代码
		Name string `json:"name,omitempty"` // 国家名称
	} `json:"country"`
	Province struct {
		Code string `json:"code,omitempty"` // 省份代码
		Name string `json:"name,omitempty"` // 省份名称
	} `json:"province"`
	City struct {
		Name string `json:"name,omitempty"` // 城市名称
	} `json:"city"`
	Location struct {
		TimeZone  string  `json:"time_zone,omitempty"` // 位置时区
		Latitude  float64 `json:"latitude,omitempty"`  // 坐标纬度
		Longitude float64 `json:"longitude,omitempty"` // 坐标经度
	} `json:"location"`
}

// QueryCity ip地址查询对应归属地信息
func (c *Client) QueryCity(ipAddress string) (result QueryCityResult, err error) {

	// 查询
	record, err := c.cityDb.City(netip.MustParseAddr(ipAddress))
	if err != nil {
		return QueryCityResult{}, err
	}

	// ip
	result.Ip = ipAddress

	// 大陆
	result.Continent.Code = record.Continent.Code
	result.Continent.Name = record.Continent.Names.SimplifiedChinese

	// 国家
	result.Country.Code = record.Country.ISOCode
	result.Country.Name = record.Country.Names.SimplifiedChinese

	// 省份
	if len(record.Subdivisions) > 0 {
		result.Province.Code = record.Subdivisions[0].ISOCode
		result.Province.Name = record.Subdivisions[0].Names.SimplifiedChinese
	}

	// 城市
	result.City.Name = record.City.Names.SimplifiedChinese

	// 位置
	result.Location.TimeZone = record.Location.TimeZone
	if record.Location.Latitude != nil {
		result.Location.Latitude = *record.Location.Latitude
	}
	if record.Location.Longitude != nil {
		result.Location.Longitude = *record.Location.Longitude
	}

	return result, err
}
