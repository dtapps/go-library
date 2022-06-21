package taobao

import (
	"go.dtapp.net/library/utils/dorm"
	"go.dtapp.net/library/utils/golog"
	"go.dtapp.net/library/utils/gorequest"
)

func (c *Client) mongoLog(api string, request gorequest.Response) {
	c.log.MongoRecord(golog.ApiMongoLog{
		RequestTime:           dorm.BsonTime(request.RequestTime),         //【请求】时间
		RequestUri:            request.RequestUri,                         //【请求】链接
		RequestUrl:            gorequest.UriParse(request.RequestUri).Url, //【请求】链接
		RequestApi:            api,                                        //【请求】接口
		RequestMethod:         request.RequestMethod,                      //【请求】方式
		RequestParams:         request.RequestParams,                      //【请求】参数
		RequestHeader:         request.RequestHeader,                      //【请求】头部
		ResponseHeader:        request.ResponseHeader,                     //【返回】头部
		ResponseStatusCode:    request.ResponseStatusCode,                 //【返回】状态码
		ResponseBody:          request.ResponseBody,                       //【返回】内容
		ResponseContentLength: request.ResponseContentLength,              //【返回】大小
		ResponseTime:          dorm.BsonTime(request.ResponseTime),        //【返回】时间
	})
}
