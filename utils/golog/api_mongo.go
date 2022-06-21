package golog

import (
	"github.com/siddontang/go/bson"
	"go.dtapp.net/library/utils/dorm"
)

// ApiMongoLog 结构体
type ApiMongoLog struct {
	LogId                 bson.ObjectId          `json:"log_id" bson:"log_id"`                                                       //【记录】编号
	RequestTime           dorm.BsonTime          `json:"request_time,omitempty" bson:"request_time,omitempty"`                       //【请求】时间
	RequestUri            string                 `json:"request_uri,omitempty" bson:"request_uri,omitempty"`                         //【请求】链接
	RequestUrl            string                 `json:"request_url,omitempty" bson:"request_url,omitempty"`                         //【请求】链接
	RequestApi            string                 `json:"request_api,omitempty" bson:"request_api,omitempty"`                         //【请求】接口
	RequestMethod         string                 `json:"request_method,omitempty" bson:"request_method,omitempty"`                   //【请求】方式
	RequestParams         map[string]interface{} `json:"request_params,omitempty" bson:"request_params,omitempty"`                   //【请求】参数
	RequestHeader         map[string]string      `json:"request_header,omitempty" bson:"request_header,omitempty"`                   //【请求】头部
	ResponseHeader        map[string][]string    `json:"response_header,omitempty" bson:"response_header,omitempty"`                 //【返回】头部
	ResponseStatusCode    int                    `json:"response_status_code,omitempty" bson:"response_status_code,omitempty"`       //【返回】状态码
	ResponseBody          interface{}            `json:"response_body,omitempty" bson:"response_body,omitempty"`                     //【返回】内容
	ResponseContentLength int64                  `json:"response_content_length,omitempty" bson:"response_content_length,omitempty"` //【返回】大小
	ResponseTime          dorm.BsonTime          `json:"response_time,omitempty" bson:"response_time,omitempty"`                     //【返回】时间
	SystemHostName        string                 `json:"system_host_name,omitempty" bson:"system_host_name,omitempty"`               //【系统】主机名
	SystemInsideIp        string                 `json:"system_inside_ip,omitempty" bson:"system_inside_ip,omitempty"`               //【系统】内网ip
	GoVersion             string                 `json:"go_version,omitempty" bson:"go_version,omitempty"`                           //【程序】Go版本
}
