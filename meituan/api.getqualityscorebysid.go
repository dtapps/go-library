package meituan

import (
	"encoding/json"
	gorequest "go.dtapp.net/library/gorequest"
	"go.dtapp.net/library/gotime"
	"net/http"
)

type ApiGetQuaLitYsCoreBySidResponse struct {
	Status int    `json:"status"`
	Des    string `json:"des"`
	Data   struct {
		DataList []struct {
			Appkey         string `json:"appkey"`         // appkey
			Sid            string `json:"sid"`            // 推广位sid
			Date           string `json:"date"`           // 质量分归属日期
			QualityGrade   string `json:"qualityGrade"`   // 质量分
			RepurchaseRate string `json:"repurchaseRate"` // sid维度的七日复购率
		} `json:"dataList"`
		Total int `json:"total"`
	} `json:"data"`
}
type ApiGetQuaLitYsCoreBySidResult struct {
	Result ApiGetQuaLitYsCoreBySidResponse // 结果
	Body   []byte                          // 内容
	Http   gorequest.Response              // 请求
	Err    error                           // 错误
}

func NewApiGetQuaLitYsCoreBySidResult(result ApiGetQuaLitYsCoreBySidResponse, body []byte, http gorequest.Response, err error) *ApiGetQuaLitYsCoreBySidResult {
	return &ApiGetQuaLitYsCoreBySidResult{Result: result, Body: body, Http: http, Err: err}
}

// ApiGetQuaLitYsCoreBySid 优选sid质量分&复购率查询
// https://union.meituan.com/v2/apiDetail?id=28
func (app *App) ApiGetQuaLitYsCoreBySid(notMustParams ...gorequest.Params) *ApiGetQuaLitYsCoreBySidResult {
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	// 请求时刻10位时间戳(秒级)，有效期60s
	params["ts"] = gotime.Current().Timestamp()
	params["appkey"] = app.appKey
	params["sign"] = app.getSign(app.secret, params)
	// 请求
	request, err := app.request("https://openapi.meituan.com/api/getqualityscorebysid", params, http.MethodGet)
	// 定义
	var response ApiGetQuaLitYsCoreBySidResponse
	err = json.Unmarshal(request.ResponseBody, &response)
	return NewApiGetQuaLitYsCoreBySidResult(response, request.ResponseBody, request, err)
}
