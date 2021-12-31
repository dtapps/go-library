package meituan

import (
	"encoding/json"
	"net/http"
)

type MiniCodeResponse struct {
	Status int    `json:"status"`         // 状态值，0为成功，非0为异常
	Des    string `json:"des,omitempty"`  // 异常描述信息
	Data   string `json:"data,omitempty"` // 小程序二维码图片地址
}

type MiniCodeResult struct {
	Result MiniCodeResponse // 结果
	Body   []byte           // 内容
	Err    error            // 错误
}

func NewMiniCodeResult(result MiniCodeResponse, body []byte, err error) *MiniCodeResult {
	return &MiniCodeResult{Result: result, Body: body, Err: err}
}

// MiniCode 小程序二维码生成
// https://union.meituan.com/v2/apiDetail?id=26
func (app *App) MiniCode(actId int64, sid string) *MiniCodeResult {
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
	// 定义
	var response MiniCodeResponse
	err = json.Unmarshal(body, &response)
	return NewMiniCodeResult(response, body, err)
}
