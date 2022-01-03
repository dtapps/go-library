package wechatminiprogram

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type CgiBinWxaAppCreateWxaQrCodeResponse struct {
	Errcode     int         `json:"errcode"`
	Errmsg      string      `json:"errmsg"`
	ContentType string      `json:"contentType"`
	Buffer      interface{} `json:"buffer"`
}

type CgiBinWxaAppCreateWxaQrCodeResult struct {
	Result CgiBinWxaAppCreateWxaQrCodeResponse // 结果
	Body   []byte                              // 内容
	Err    error                               // 错误
}

func NewCgiBinWxaAppCreateWxaQrCodeResult(result CgiBinWxaAppCreateWxaQrCodeResponse, body []byte, err error) *CgiBinWxaAppCreateWxaQrCodeResult {
	return &CgiBinWxaAppCreateWxaQrCodeResult{Result: result, Body: body, Err: err}
}

// CgiBinWxaAppCreateWxaQrCode 获取小程序二维码，适用于需要的码数量较少的业务场景。通过该接口生成的小程序码，永久有效，有数量限制
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/qr-code/wxacode.createQRCode.html
func (app *App) CgiBinWxaAppCreateWxaQrCode(path string, width int) *CgiBinWxaAppCreateWxaQrCodeResult {
	// 参数
	param := NewParams()
	param.Set("path", path)
	param.Set("width", width)
	params := app.NewParamsWith(param)
	// 请求
	body, err := app.request(fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/wxaapp/createwxaqrcode?access_token=%s", app.AccessToken), params, http.MethodPost)
	// 定义
	var response CgiBinWxaAppCreateWxaQrCodeResponse
	err = json.Unmarshal(body, &response)
	return NewCgiBinWxaAppCreateWxaQrCodeResult(response, body, err)
}
