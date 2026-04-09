package czdb

import (
	"net"

	"github.com/tagphi/czdb-search-golang/pkg/db"
)

// QueryResult 返回
type QueryResult struct {
	Ip      string `json:"ip,omitempty"`       // ip
	RawData string `json:"raw_data,omitempty"` // 原始字符串
}

// Query ip地址查询对应归属地信息
func (c *Client) Query(ipAddress string) (result QueryResult, err error) {

	var searcher *db.DBSearcher

	// 判断 ip 类型
	ip := net.ParseIP(ipAddress)
	if ip.To4() == nil {
		searcher = c.v6Db
	} else {
		searcher = c.v4Db
	}

	// 查询
	region, err := db.Search(ipAddress, searcher)
	if err != nil {
		return QueryResult{}, err
	}

	// ip
	result.Ip = ipAddress

	// 原始字符串
	result.RawData = region

	return result, nil
}
