package gorequest

import "time"

type LogResponse struct {
	TraceID            string    `json:"trace_id"`             // 追踪编号
	RequestID          string    `json:"request_id"`           // 请求编号
	RequestTime        time.Time `json:"request_time"`         // 请求时间
	RequestUri         string    `json:"request_uri"`          // 请求链接
	RequestUrl         string    `json:"request_url"`          // 请求链接
	RequestApi         string    `json:"request_api"`          // 请求接口
	RequestMethod      string    `json:"request_method"`       // 请求方式
	RequestParams      string    `json:"request_params"`       // 请求参数
	RequestHeader      string    `json:"request_header"`       // 请求头部
	RequestIP          string    `json:"request_ip"`           // 请求请求IP
	RequestCostTime    int64     `json:"request_cost_time"`    // 请求消耗时长
	ResponseHeader     string    `json:"response_header"`      // 返回头部
	ResponseStatusCode int       `json:"response_status_code"` // 返回状态码
	ResponseBody       string    `json:"response_body"`        // 返回Json数据
	ResponseBodyJson   string    `json:"response_body_json"`   // 返回Json数据
	ResponseBodyXml    string    `json:"response_body_xml"`    // 返回Xml数据
	ResponseTime       time.Time `json:"response_time"`        // 返回时间
	GoVersion          string    `json:"go_version"`           // 程序GoVersion
	SdkVersion         string    `json:"sdk_version"`          // 程序SdkVersion
}
