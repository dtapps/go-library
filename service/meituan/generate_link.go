package meituan

import (
	"encoding/json"
)

// GenerateLinkResult 返回参数
type GenerateLinkResult struct {
	Status int    `json:"status"`         // 状态值，0为成功，非0为异常
	Des    string `json:"des,omitempty"`  // 异常描述信息
	Data   string `json:"data,omitempty"` // 最终的推广链接
}

// GenerateLink 自助取链接口 https://union.meituan.com/v2/apiDetail?id=25
func (app *App) GenerateLink(actId int64, sid string, linkType, shortLink int) (result GenerateLinkResult, err error) {

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
	body, err := app.request("https://openapi.meituan.com/api/generateLink", params, "GET")
	if err != nil {
		return
	}

	// 解析
	if err = json.Unmarshal(body, &result); err != nil {
		return
	}
	return
}
