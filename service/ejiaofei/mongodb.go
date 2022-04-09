package ejiaofei

import (
	"dtapps/dta/library/utils/gohttp"
	"gitee.com/dtapps/go-library/utils/gotime"
)

// 日志
type mongoZap struct {
	Url           string      `json:"url" bson:"url"`
	Params        interface{} `json:"params" bson:"params"`
	Method        string      `json:"method" bson:"method"`
	Header        interface{} `json:"header" bson:"header"`
	Status        string      `json:"status" bson:"status"`
	StatusCode    int         `json:"status_code" bson:"status_code"`
	Body          interface{} `json:"body" bson:"body"`
	ContentLength int64       `json:"content_length" bson:"content_length"`
	CreateTime    string      `json:"create_time" bson:"create_time"`
}

func (m *mongoZap) Database() string {
	return "zap_logs"
}

func (m *mongoZap) TableName() string {
	return "ejiaofei_" + gotime.Current().SetFormat("200601")
}

func (app *App) mongoLog(url string, params map[string]interface{}, method string, request gohttp.Response) {
	if app.Mongo.Db == nil {
		return
	}
	app.Mongo.Model(&mongoZap{}).InsertOne(mongoZap{
		Url:           url,
		Params:        params,
		Method:        method,
		Header:        request.Header,
		Status:        request.Status,
		StatusCode:    request.StatusCode,
		Body:          string(request.Body),
		ContentLength: request.ContentLength,
		CreateTime:    gotime.Current().Format(),
	})
}
