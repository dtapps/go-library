package meituan

import (
	"encoding/json"
	"go.dtapp.net/library/utils/gorequest"
	"net/http"
)

type ApiMiniCodeResponse struct {
	Status int    `json:"status"`         // 状态值，0为成功，非0为异常
	Des    string `json:"des,omitempty"`  // 异常描述信息
	Data   string `json:"data,omitempty"` // 小程序二维码图片地址
}

type ApiMiniCodeResult struct {
	Result ApiMiniCodeResponse // 结果
	Body   []byte              // 内容
	Http   gorequest.Response  // 请求
	Err    error               // 错误
}

func NewApiMiniCodeResult(result ApiMiniCodeResponse, body []byte, http gorequest.Response, err error) *ApiMiniCodeResult {
	return &ApiMiniCodeResult{Result: result, Body: body, Http: http, Err: err}
}

// ApiMiniCode 小程序生成二维码（新版）
// https://union.meituan.com/v2/apiDetail?id=26
func (app *App) ApiMiniCode(actId int64, sid string) *ApiMiniCodeResult {
	// 参数
	param := gorequest.NewParams()
	param.Set("appkey", app.appKey)
	param.Set("sid", sid)
	param.Set("actId", actId)
	// 转换
	params := gorequest.NewParamsWith(param)
	params["sign"] = app.getSign(app.secret, params)
	// 请求
	request, err := app.request("https://openapi.meituan.com/api/miniCode", params, http.MethodGet)
	// 定义
	var response ApiMiniCodeResponse
	err = json.Unmarshal(request.ResponseBody, &response)
	return NewApiMiniCodeResult(response, request.ResponseBody, request, err)
}
