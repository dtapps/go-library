package dingdanxia

import (
	"go.dtapp.net/gomongo"
	"go.dtapp.net/library/utils/gorequest"
	"net/http"
)

// 日志
type mongoZap struct {
	RequestTime           gomongo.BsonTime       `json:"request_time" bson:"request_time"`                       //【请求】时间
	RequestUri            string                 `json:"request_uri" bson:"request_uri"`                         //【请求】链接
	RequestUrl            string                 `json:"request_url" bson:"request_url"`                         //【请求】链接
	RequestApi            string                 `json:"request_api" bson:"request_api"`                         //【请求】接口
	RequestMethod         string                 `json:"request_method" bson:"request_method"`                   //【请求】方式
	RequestParams         gorequest.Params       `json:"request_params" bson:"request_params"`                   //【请求】参数
	RequestHeader         gorequest.Headers      `json:"request_header" bson:"request_header"`                   //【请求】头部
	ResponseHeader        http.Header            `json:"response_header" bson:"response_header"`                 //【返回】头部
	ResponseStatusCode    int                    `json:"response_status_code" bson:"response_status_code"`       //【返回】状态码
	ResponseBody          map[string]interface{} `json:"response_body" bson:"response_body"`                     //【返回】内容
	ResponseContentLength int64                  `json:"response_content_length" bson:"response_content_length"` //【返回】大小
	ResponseTime          gomongo.BsonTime       `json:"response_time" bson:"response_time"`                     //【返回】时间
}

func (m *mongoZap) Database() string {
	return "zap_log"
}

func (m *mongoZap) TableName() string {
	return "dingdanxia"
}

func (app *App) mongoLog(request gorequest.Response) {
	_, _ = app.mongo.Model(&mongoZap{}).InsertOne(mongoZap{
		RequestTime:           gomongo.BsonTime(request.RequestTime),           //【请求】时间
		RequestUri:            request.RequestUri,                              //【请求】链接
		RequestUrl:            gorequest.UriParse(request.RequestUri).Url,      //【请求】链接
		RequestApi:            gorequest.UriParse(request.RequestUri).Path,     //【请求】接口
		RequestMethod:         request.RequestMethod,                           //【请求】方式
		RequestParams:         request.RequestParams,                           //【请求】参数
		RequestHeader:         request.RequestHeader,                           //【请求】头部
		ResponseHeader:        request.ResponseHeader,                          //【返回】头部
		ResponseStatusCode:    request.ResponseStatusCode,                      //【返回】状态码
		ResponseBody:          gomongo.JsonDecodeNoError(request.ResponseBody), //【返回】内容
		ResponseContentLength: request.ResponseContentLength,                   //【返回】大小
		ResponseTime:          gomongo.BsonTime(request.ResponseTime),          //【返回】时间
	})
}
