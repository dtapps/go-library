package wechatminiprogram

import (
	"go.dtapp.net/golog"
	"go.dtapp.net/library/utils/gorequest"
	"go.dtapp.net/library/utils/gojson"
	"gorm.io/datatypes"
)

// 记录日志
func (app *App) postgresqlLog(request gorequest.Response) {
	body := golog.ApiPostgresqlLog{}
	body.RequestTime = golog.TimeString{Time: request.RequestTime}                         //【请求】时间
	body.RequestUri = request.RequestUri                                                   //【请求】链接
	body.RequestUrl = gorequest.UriParse(request.RequestUri).Url                           //【请求】链接
	body.RequestApi = gorequest.UriParse(request.RequestUri).Path                          //【请求】接口
	body.RequestMethod = request.RequestMethod                                             //【请求】方式
	body.RequestParams = datatypes.JSON(gojson.JsonEncodeNoError(request.RequestParams))   //【请求】参数
	body.RequestHeader = datatypes.JSON(gojson.JsonEncodeNoError(request.RequestHeader))   //【请求】头部
	body.ResponseHeader = datatypes.JSON(gojson.JsonEncodeNoError(request.ResponseHeader)) //【返回】头部
	body.ResponseStatusCode = request.ResponseStatusCode                                   //【返回】状态码
	body.ResponseContentLength = request.ResponseContentLength                             //【返回】大小
	body.ResponseTime = golog.TimeString{Time: request.ResponseTime}                       //【返回】时间
	if request.ResponseHeader.Get("Content-Type") == "image/jpeg" || request.ResponseHeader.Get("Content-Type") == "image/png" {
	} else {
		body.ResponseBody = request.ResponseBody //【返回】内容
	}
	app.log.Record(body)
}
