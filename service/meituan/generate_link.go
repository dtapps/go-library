package meituan

import (
	"encoding/json"
	"net/http"
)

type GenerateLinkResponse struct {
	Status int    `json:"status"`         // 状态值，0为成功，非0为异常
	Des    string `json:"des,omitempty"`  // 异常描述信息
	Data   string `json:"data,omitempty"` // 最终的推广链接
}

type GenerateLinkResult struct {
	Result GenerateLinkResponse // 结果
	Body   []byte               // 内容
	Err    error                // 错误
}

func NewGenerateLinkResult(result GenerateLinkResponse, body []byte, err error) *GenerateLinkResult {
	return &GenerateLinkResult{Result: result, Body: body, Err: err}
}

// GenerateLink 自助取链接口
// https://union.meituan.com/v2/apiDetail?id=25
func (app *App) GenerateLink(actId int64, sid string, linkType, shortLink int) *GenerateLinkResult {
	// 参数
	param := NewParams()
	param.Set("appkey", app.AppKey)
	param.Set("actId", actId)
	param.Set("sid", sid)
	param.Set("linkType", linkType)
	param.Set("shortLink", shortLink)
	// 转换
	params := app.NewParamsWith(param)
	params["sign"] = app.getSign(app.Secret, params)
	// 请求
	body, err := app.request("https://openapi.meituan.com/api/generateLink", params, http.MethodGet)
	// 定义
	var response GenerateLinkResponse
	err = json.Unmarshal(body, &response)
	return NewGenerateLinkResult(response, body, err)
}
