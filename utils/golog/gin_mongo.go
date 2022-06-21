package golog

import (
	"go.dtapp.net/library/utils/dorm"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// GinMongoLog 结构体
type GinMongoLog struct {
	LogId             primitive.ObjectID `json:"log_id,omitempty" bson:"_id,omitempty"`                              //【记录】编号
	TraceId           string             `json:"trace_id,omitempty" bson:"trace_id,omitempty"`                       //【系统】链编号
	RequestTime       dorm.BsonTime      `json:"request_time,omitempty" bson:"request_time,omitempty"`               //【请求】时间
	RequestUri        string             `json:"request_uri,omitempty" bson:"request_uri,omitempty"`                 //【请求】请求链接 域名+路径+参数
	RequestUrl        string             `json:"request_url,omitempty" bson:"request_url,omitempty"`                 //【请求】请求链接 域名+路径
	RequestApi        string             `json:"request_api,omitempty" bson:"request_api,omitempty"`                 //【请求】请求接口 路径
	RequestMethod     string             `json:"request_method,omitempty" bson:"request_method,omitempty"`           //【请求】请求方式
	RequestProto      string             `json:"request_proto,omitempty" bson:"request_proto,omitempty"`             //【请求】请求协议
	RequestUa         string             `json:"request_ua,omitempty" bson:"request_ua,omitempty"`                   //【请求】请求UA
	RequestReferer    string             `json:"request_referer,omitempty" bson:"request_referer,omitempty"`         //【请求】请求referer
	RequestBody       interface{}        `json:"request_body,omitempty" bson:"request_body,omitempty"`               //【请求】请求主体
	RequestUrlQuery   interface{}        `json:"request_url_query,omitempty" bson:"request_url_query,omitempty"`     //【请求】请求URL参数
	RequestIp         string             `json:"request_ip,omitempty" bson:"request_ip,omitempty"`                   //【请求】请求客户端Ip
	RequestIpCountry  string             `json:"request_ip_country,omitempty" bson:"request_ip_country,omitempty"`   //【请求】请求客户端城市
	RequestIpRegion   string             `json:"request_ip_region,omitempty" bson:"request_ip_region,omitempty"`     //【请求】请求客户端区域
	RequestIpProvince string             `json:"request_ip_province,omitempty" bson:"request_ip_province,omitempty"` //【请求】请求客户端省份
	RequestIpCity     string             `json:"request_ip_city,omitempty" bson:"request_ip_city,omitempty"`         //【请求】请求客户端城市
	RequestIpIsp      string             `json:"request_ip_isp,omitempty" bson:"request_ip_isp,omitempty"`           //【请求】请求客户端运营商
	RequestHeader     interface{}        `json:"request_header,omitempty" bson:"request_header,omitempty"`           //【请求】请求头
	ResponseTime      dorm.BsonTime      `json:"response_time,omitempty" bson:"response_time,omitempty"`             //【返回】时间
	ResponseCode      int                `json:"response_code,omitempty" bson:"response_code,omitempty"`             //【返回】状态码
	ResponseMsg       string             `json:"response_msg,omitempty" bson:"response_msg,omitempty"`               //【返回】描述
	ResponseData      interface{}        `json:"response_data,omitempty" bson:"response_data,omitempty"`             //【返回】数据
	CostTime          int64              `json:"cost_time,omitempty" bson:"cost_time,omitempty"`                     //【系统】花费时间
	SystemHostName    string             `json:"system_host_name,omitempty" bson:"system_host_name,omitempty"`       //【系统】主机名
	SystemInsideIp    string             `json:"system_inside_ip,omitempty" bson:"system_inside_ip,omitempty"`       //【系统】内网ip
	GoVersion         string             `json:"go_version,omitempty" bson:"go_version,omitempty"`                   //【程序】Go版本
}
