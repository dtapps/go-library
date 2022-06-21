package golog

import "gorm.io/datatypes"

// GinPostgresqlLog 结构体
type GinPostgresqlLog struct {
	LogId             uint           `gorm:"primaryKey" json:"log_id"`                        //【记录】编号
	TraceId           string         `gorm:"type:text" json:"trace_id,omitempty"`             //【系统】链编号
	RequestTime       TimeString     `gorm:"index" json:"request_time,omitempty"`             //【请求】时间
	RequestUri        string         `gorm:"type:text" json:"request_uri,omitempty"`          //【请求】请求链接 域名+路径+参数
	RequestUrl        string         `gorm:"type:text" json:"request_url,omitempty"`          //【请求】请求链接 域名+路径
	RequestApi        string         `gorm:"type:text;index" json:"request_api,omitempty"`    //【请求】请求接口 路径
	RequestMethod     string         `gorm:"type:text;index" json:"request_method,omitempty"` //【请求】请求方式
	RequestProto      string         `gorm:"type:text" json:"request_proto,omitempty"`        //【请求】请求协议
	RequestUa         string         `gorm:"type:text" json:"request_ua,omitempty"`           //【请求】请求UA
	RequestReferer    string         `gorm:"type:text" json:"request_referer,omitempty"`      //【请求】请求referer
	RequestBody       datatypes.JSON `gorm:"type:jsonb" json:"request_body,omitempty"`        //【请求】请求主体
	RequestUrlQuery   datatypes.JSON `gorm:"type:jsonb" json:"request_url_query,omitempty"`   //【请求】请求URL参数
	RequestIp         string         `gorm:"type:text" json:"request_ip,omitempty"`           //【请求】请求客户端Ip
	RequestIpCountry  string         `gorm:"type:text" json:"request_ip_country,omitempty"`   //【请求】请求客户端城市
	RequestIpRegion   string         `gorm:"type:text" json:"request_ip_region,omitempty"`    //【请求】请求客户端区域
	RequestIpProvince string         `gorm:"type:text" json:"request_ip_province,omitempty"`  //【请求】请求客户端省份
	RequestIpCity     string         `gorm:"type:text" json:"request_ip_city,omitempty"`      //【请求】请求客户端城市
	RequestIpIsp      string         `gorm:"type:text" json:"request_ip_isp,omitempty"`       //【请求】请求客户端运营商
	RequestHeader     datatypes.JSON `gorm:"type:jsonb" json:"request_header,omitempty"`      //【请求】请求头
	ResponseTime      TimeString     `gorm:"index" json:"response_time,omitempty"`            //【返回】时间
	ResponseCode      int            `gorm:"type:bigint" json:"response_code,omitempty"`      //【返回】状态码
	ResponseMsg       string         `gorm:"type:text" json:"response_msg,omitempty"`         //【返回】描述
	ResponseData      datatypes.JSON `gorm:"type:jsonb" json:"response_data,omitempty"`       //【返回】数据
	CostTime          int64          `gorm:"type:bigint" json:"cost_time,omitempty"`          //【系统】花费时间
	SystemHostName    string         `gorm:"type:text" json:"system_host_name,omitempty"`     //【系统】主机名
	SystemInsideIp    string         `gorm:"type:text" json:"system_inside_ip,omitempty"`     //【系统】内网ip
	GoVersion         string         `gorm:"type:text" json:"go_version,omitempty"`           //【程序】Go版本
}
