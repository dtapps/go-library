package ip2region

import "strings"

// QueryResult 返回
type QueryResult struct {
	Ip       string `json:"ip,omitempty"`       // ip
	RawData  string `json:"raw_data,omitempty"` // 原始字符串
	Country  string `json:"country,omitempty"`  // 国家名称
	Province string `json:"province,omitempty"` // 省份名称
	City     string `json:"city,omitempty"`     // 城市名称
	Isp      string `json:"isp,omitempty"`      // 运营商 (Internet Service Provider)
}

// Query ip地址查询对应归属地信息
func (c *Client) Query(ipAddress string) (result QueryResult, err error) {

	// 查询
	region, err := c.service.Search(ipAddress)
	if err != nil {
		return QueryResult{}, err
	}

	// ip
	result.Ip = ipAddress

	// 原始字符串
	result.RawData = region

	// 详细信息
	country, province, city, isp := c.parseRegion(region)
	result.Country = country
	result.Province = province
	result.City = city
	result.Isp = isp

	return result, nil
}

// parseRegion 解析 ip2region 返回的区域字符串
// 标准格式: 国家|区域|省份|城市|运营商
// 例如: 中国|0|广东省|深圳市|电信
// 或者: 中国|广东省|深圳市|电信|CN
func (c *Client) parseRegion(region string) (country, province, city, isp string) {
	if region == "" {
		return
	}

	// 使用 | 分割
	parts := strings.Split(region, "|")

	// 过滤空字符串并清理空格 (可选，视数据干净程度而定)
	var cleanParts []string
	for _, p := range parts {
		p = strings.TrimSpace(p)
		if p != "" && p != "0" { // ip2region 中 "0" 通常表示未知或无信息
			cleanParts = append(cleanParts, p)
		}
	}

	length := len(cleanParts)
	if length == 0 {
		return
	}

	// 根据常见的 ip2region xdb 数据格式进行映射
	// 常见格式1: 国家|区域(省)|城市|运营商|...
	// 常见格式2: 国家|0|省份|城市|运营商

	// 这里采用一种比较通用的启发式解析，假设顺序为：
	// [0]国家, [1]大区/省, [2]城市, [3]运营商
	// 如果只有3个部分，可能是: 国家|省|城市 或 国家|城市|运营商

	if length >= 1 {
		country = cleanParts[0]
		// 如果国家是 "中国" 或 "CHINA"，通常后续字段才有意义
		// 如果是 "0" 或 "LAN" 等，可能没有后续地理位置
	}

	if length >= 2 {
		// 第二个字段可能是 "0" (已过滤)，或者是省份，或者是大区
		// 在标准 xdb 中，索引1通常是 Province (如果索引1是0，则索引2是Province)
		// 由于我们在上面过滤了 "0"，这里的逻辑需要调整

		// 重新使用原始 parts 进行更精确的位置映射可能更稳妥，
		// 但为了简单起见，我们假设过滤后的 cleanParts 如下：
		// Index 0: Country
		// Index 1: Province (或者大区)
		// Index 2: City
		// Index 3+: ISP / Other

		province = cleanParts[1]
	}

	if length >= 3 {
		city = cleanParts[2]
	}

	if length >= 4 {
		// 第4个部分通常是运营商，但也可能是其他信息
		// 简单的判断：如果包含 "电信", "联通", "移动" 等关键词，或者是英文ISP名称
		isp = cleanParts[3]
	}

	return
}
