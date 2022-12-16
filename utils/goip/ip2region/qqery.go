package ip2region

import (
	"errors"
	"net"
	"strconv"
)

type QueryResult struct {
	Ip       string `json:"ip,omitempty"`       // ip
	CityId   int64  `json:"city_id,omitempty"`  // 城市代码
	Country  string `json:"country,omitempty"`  // 国家
	Region   string `json:"region,omitempty"`   // 区域
	Province string `json:"province,omitempty"` // 省份
	City     string `json:"city,omitempty"`     // 城市
	Isp      string `json:"isp,omitempty"`      // 运营商
}

func (ip QueryResult) String() string {
	return ip.Ip + "|" + strconv.FormatInt(ip.CityId, 10) + "|" + ip.Country + "|" + ip.Region + "|" + ip.Province + "|" + ip.City + "|" + ip.Isp
}

// Query memory算法：整个数据库全部载入内存，单次查询都在0.1x毫秒内
func (c *Client) Query(ipAddress net.IP) (result QueryResult, err error) {

	result.Ip = ipAddress.String()

	if c.totalBlocks == 0 {

		if err != nil {

			return QueryResult{}, err
		}

		c.firstIndexPtr = getLong(dbBuff, 0)
		c.lastIndexPtr = getLong(dbBuff, 4)
		c.totalBlocks = (c.lastIndexPtr-c.firstIndexPtr)/IndexBlockLength + 1
	}

	ip, err := ip2long(result.Ip)
	if err != nil {
		return QueryResult{}, err
	}

	h := c.totalBlocks
	var dataPtr, l int64
	for l <= h {

		m := (l + h) >> 1
		p := c.firstIndexPtr + m*IndexBlockLength
		sip := getLong(dbBuff, p)
		if ip < sip {
			h = m - 1
		} else {
			eip := getLong(dbBuff, p+4)
			if ip > eip {
				l = m + 1
			} else {
				dataPtr = getLong(dbBuff, p+8)
				break
			}
		}
	}
	if dataPtr == 0 {
		return QueryResult{}, errors.New("not found")
	}

	dataLen := (dataPtr >> 24) & 0xFF
	dataPtr = dataPtr & 0x00FFFFFF
	result = getIpInfo(result.Ip, getLong(dbBuff, dataPtr), dbBuff[(dataPtr)+4:dataPtr+dataLen])

	return result, nil
}
