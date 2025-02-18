package gorequest

import "time"

type LogResponse struct {
	TraceID string `json:"trace_id"` // 追踪编号

	RequestID       string              `json:"request_id"`        // 请求编号
	RequestTime     time.Time           `json:"request_time"`      // 请求时间
	RequestHost     string              `json:"request_host"`      // 请求主机
	RequestPath     string              `json:"request_path"`      // 请求地址
	RequestQuery    map[string][]string `json:"request_query"`     // 请求参数
	RequestMethod   string              `json:"request_method"`    // 请求方式
	RequestBody     map[string]any      `json:"request_body"`      // 请求内容
	RequestIP       string              `json:"request_ip"`        // 请求IP
	RequestHeader   map[string][]string `json:"request_header"`    // 请求头
	RequestCostTime int64               `json:"request_cost_time"` // 请求消耗时长

	ResponseTime     time.Time           `json:"response_time"`      // 返回时间
	ResponseHeader   map[string][]string `json:"response_header"`    // 返回头部
	ResponseCode     int                 `json:"response_code"`      // 返回状态码
	ResponseBody     string              `json:"response_body"`      // 返回Json数据
	ResponseBodyJson map[string]any      `json:"response_body_json"` // 返回Json数据

}
