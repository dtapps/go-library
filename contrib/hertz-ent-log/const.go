package hertz_ent_log

import "time"

// HertzLogData Hertz框架日志模型
type HertzLogData struct {
	TraceID         string              `json:"trace_id,omitempty"`          // 跟踪编号
	RequestID       string              `json:"request_id,omitempty"`        // 请求编号
	RequestTime     time.Time           `json:"request_time,omitempty"`      // 请求时间
	RequestHost     string              `json:"request_host,omitempty"`      // 请求主机
	RequestPath     string              `json:"request_path,omitempty"`      // 请求地址
	RequestQuery    map[string]any      `json:"request_query,omitempty"`     // 请求参数
	RequestMethod   string              `json:"request_method,omitempty"`    // 请求方式
	RequestBody     map[string]any      `json:"request_body,omitempty"`      // 请求内容
	RequestIP       string              `json:"request_ip,omitempty"`        // 请求IP
	RequestHeader   map[string][]string `json:"request_header,omitempty"`    // 请求头
	RequestCostTime int64               `json:"request_cost_time,omitempty"` // 请求消耗时长
	ResponseTime    time.Time           `json:"response_time,omitempty"`     // 响应时间
	ResponseHeader  map[string][]string `json:"response_header,omitempty"`   // 响应头
	ResponseCode    int                 `json:"response_code,omitempty"`     // 响应状态
	ResponseBody    map[string]any      `json:"response_body,omitempty"`     // 响应内容
}
