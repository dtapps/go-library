package wikeyun

import (
	"encoding/json"
	"github.com/dtapps/go-library/utils/gohttp"
	"github.com/dtapps/go-library/utils/gotime"
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
	return "wikeyun_" + gotime.Current().SetFormat("200601")
}

func (app *App) mongoLog(url string, params map[string]interface{}, method string, request gohttp.Response) {
	if app.Mongo.Db == nil {
		return
	}
	var body map[string]interface{}
	_ = json.Unmarshal(request.Body, &body)
	app.Mongo.Model(&mongoZap{}).InsertOne(mongoZap{
		Url:           url,
		Params:        params,
		Method:        method,
		Header:        request.Header,
		Status:        request.Status,
		StatusCode:    request.StatusCode,
		Body:          body,
		ContentLength: request.ContentLength,
		CreateTime:    gotime.Current().Format(),
	})
}
