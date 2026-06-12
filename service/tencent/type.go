package tencent

// 获取域名的解析记录列表
// https://cloud.tencent.com/document/api/1427/56166
type ResponseDescribeRecordList struct {
	Response struct {
		RequestId       string `json:"RequestId"`
		RecordCountInfo struct {
			SubdomainCount int `json:"SubdomainCount"`
			TotalCount     int `json:"TotalCount"`
			ListCount      int `json:"ListCount"`
		} `json:"RecordCountInfo"`
		RecordList []struct {
			RecordId int64  `json:"RecordId"` // 记录Id
			Value    string `json:"Value"`    // 记录值
			Status   string `json:"Status"`   // 记录状态，启用：ENABLE，暂停：DISABLE
			Name     string `json:"Name"`     // 主机名
			Line     string `json:"Line"`     // 记录线路
			Type     string `json:"Type"`     // 记录类型
			Remark   string `json:"Remark"`   // 记录备注说明
		} `json:"RecordList"`
	} `json:"Response"`
}

// 添加记录
// https://cloud.tencent.com/document/api/1427/56180
type ResponseCreateRecord struct {
	Response struct {
		RequestId string `json:"RequestId"`
		RecordId  int    `json:"RecordId,omitempty"` // 记录Id
	} `json:"Response"`
}

// 修改记录
// https://cloud.tencent.com/document/api/1427/56157
type ResponseModifyRecord struct {
	Response struct {
		RequestId string `json:"RequestId"`
		RecordId  int    `json:"RecordId,omitempty"` // 记录Id
	} `json:"Response"`
}

// 查询加速域名列表
// https://cloud.tencent.com/document/api/1552/86336
type ResponseQueryDomain struct {
	Response struct {
		RequestId           string `json:"RequestId"`  // 唯一请求ID
		TotalCount          int    `json:"TotalCount"` // 记录总数
		AccelerationDomains []struct {
			ZoneId       string `json:"ZoneId"` // 站点 ID
			DomainId     string `json:"DomainId"`
			DomainName   string `json:"DomainName"`   // 加速域名名称
			DomainStatus string `json:"DomainStatus"` // 加速域名状态
			OriginDetail struct {
				OriginType string `json:"OriginType"` // 源站类型
				Origin     string `json:"Origin"`     // 源站地址
			} `json:"OriginDetail"`
			OriginProtocol  string `json:"OriginProtocol,omitempty"`  // 回源协议
			HttpOriginPort  int    `json:"HttpOriginPort,omitempty"`  // HTTP回源端口
			HttpsOriginPort int    `json:"HttpsOriginPort,omitempty"` // HTTPS回源端口
			Cname           string `json:"Cname"`                     // CNAME 地址
		} `json:"AccelerationDomains"`
	} `json:"Response"`
}

// 创建加速域名
// https://cloud.tencent.com/document/api/1552/86338
type ResponseCreateDomain struct {
	Response struct {
		RequestId string `json:"RequestId"` // 唯一请求ID
	} `json:"Response"`
}

// 修改加速域名信息
// https://cloud.tencent.com/document/api/1552/86335
type ResponseUpdateDomain struct {
	Response struct {
		RequestId string `json:"RequestId"` // 唯一请求ID
	} `json:"Response"`
}

// 获取源站组列表
// https://cloud.tencent.com/document/api/1552/80594
type ResponseQueryOriginGroup struct {
	Response struct {
		RequestId    string `json:"RequestId"`  // 唯一请求ID
		TotalCount   int    `json:"TotalCount"` // 记录总数
		OriginGroups []struct {
			GroupId string `json:"GroupId"` // 源站组ID
			Name    string `json:"Name"`    // 源站组名称
			Type    string `json:"Type"`    // 源站组类型
			Records []struct {
				Record   string `json:"Record"`   // 源站记录值
				Type     string `json:"Type"`     // 源站类型
				RecordId string `json:"RecordId"` // 源站记录ID
			} `json:"Records"` // 源站记录信息
		} `json:"OriginGroups"` // 源站组信息
	} `json:"Response"`
}

// 创建源站组
// https://cloud.tencent.com/document/api/1552/80598
type ResponseCreateOriginGroup struct {
	Response struct {
		RequestId     string `json:"RequestId"`               // 唯一请求ID
		OriginGroupId int    `json:"OriginGroupId,omitempty"` // 源站组ID
	} `json:"Response"`
}

// 修改源站组
// https://cloud.tencent.com/document/api/1552/80592
type ResponseUpdateOriginGroup struct {
	Response struct {
		RequestId string `json:"RequestId"` // 唯一请求ID
	} `json:"Response"`
}
