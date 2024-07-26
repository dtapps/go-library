package gorequest

import "time"

type LogResponse struct {
	TraceID string `json:"trace_id"` // 追踪编号

	RequestID          string    `json:"request_id"`           // 请求编号
	RequestTime        time.Time `json:"request_time"`         // 请求时间
	RequestHost        string    `json:"request_host"`         // 请求主机
	RequestPath        string    `json:"request_path"`         // 请求地址
	RequestQuery       string    `json:"request_query"`        // 请求参数
	RequestMethod      string    `json:"request_method"`       // 请求方式
	RequestScheme      string    `json:"request_scheme"`       // 请求协议
	RequestContentType string    `json:"request_content_type"` // 请求类型
	RequestBody        string    `json:"request_body"`         // 请求内容
	RequestClientIP    string    `json:"request_client_ip"`    // 请求IP
	RequestUserAgent   string    `json:"request_user_agent"`   // 请求UA
	RequestHeader      string    `json:"request_header"`       // 请求头
	RequestCostTime    int64     `json:"request_cost_time"`    // 请求消耗时长

	ResponseTime       time.Time `json:"response_time"`        // 返回时间
	ResponseHeader     string    `json:"response_header"`      // 返回头部
	ResponseStatusCode int       `json:"response_status_code"` // 返回状态码
	ResponseBody       string    `json:"response_body"`        // 返回Json数据
	ResponseBodyJson   string    `json:"response_body_json"`   // 返回Json数据
	ResponseBodyXml    string    `json:"response_body_xml"`    // 返回Xml数据

	GoVersion  string `json:"go_version"`  // 程序GoVersion
	SdkVersion string `json:"sdk_version"` // 程序SdkVersion
}
