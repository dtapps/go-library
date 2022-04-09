package meituan

import (
	"encoding/json"
	"github.com/dtapps/go-library/utils/gotime"
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
	Err    error                           // 错误
}

func NewApiGetQuaLitYsCoreBySidResult(result ApiGetQuaLitYsCoreBySidResponse, body []byte, err error) *ApiGetQuaLitYsCoreBySidResult {
	return &ApiGetQuaLitYsCoreBySidResult{Result: result, Body: body, Err: err}
}

// ApiGetQuaLitYsCoreBySid 优选sid质量分&复购率查询
// https://union.meituan.com/v2/apiDetail?id=28
func (app *App) ApiGetQuaLitYsCoreBySid(notMustParams ...Params) *ApiGetQuaLitYsCoreBySidResult {
	// 参数
	params := app.NewParamsWith(notMustParams...)
	// 请求时刻10位时间戳(秒级)，有效期60s
	params["ts"] = gotime.Current().Timestamp()
	params["appkey"] = app.AppKey
	params["sign"] = app.getSign(app.Secret, params)
	// 请求
	body, err := app.request("https://openapi.meituan.com/api/getqualityscorebysid", params, http.MethodGet)
	// 定义
	var response ApiGetQuaLitYsCoreBySidResponse
	err = json.Unmarshal(body, &response)
	return NewApiGetQuaLitYsCoreBySidResult(response, body, err)
}
