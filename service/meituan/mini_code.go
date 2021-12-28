package meituan

import (
	"encoding/json"
	"net/http"
)

// MiniCodeResult 返回参数
type MiniCodeResult struct {
	Status int    `json:"status"`         // 状态值，0为成功，非0为异常
	Des    string `json:"des,omitempty"`  // 异常描述信息
	Data   string `json:"data,omitempty"` // 小程序二维码图片地址
}

// MiniCode 小程序二维码生成 https://union.meituan.com/v2/apiDetail?id=26
func (app *App) MiniCode(actId int64, sid string) (result MiniCodeResult, err error) {

	// 参数
	param := NewParams()
	param.Set("appkey", app.AppKey)
	param.Set("sid", sid)
	param.Set("actId", actId)

	// 转换
	params := app.NewParamsWith(param)
	params["sign"] = app.getSign(app.Secret, params)

	// 请求
	body, err := app.request("https://openapi.meituan.com/api/miniCode", params, http.MethodGet)
	if err != nil {
		return
	}

	// 解析
	if err = json.Unmarshal(body, &result); err != nil {
		return
	}
	return
}
