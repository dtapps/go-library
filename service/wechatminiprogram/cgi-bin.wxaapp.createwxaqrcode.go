package wechatminiprogram

import (
	"encoding/json"
	"fmt"
	"gopkg.in/dtapps/go-library.v3/utils/gohttp"
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
	Http   gohttp.Response                     // 请求
	Err    error                               // 错误
}

func NewCgiBinWxaAppCreateWxaQrCodeResult(result CgiBinWxaAppCreateWxaQrCodeResponse, body []byte, http gohttp.Response, err error) *CgiBinWxaAppCreateWxaQrCodeResult {
	return &CgiBinWxaAppCreateWxaQrCodeResult{Result: result, Body: body, Http: http, Err: err}
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
	request, err := app.request(fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/wxaapp/createwxaqrcode?access_token=%s", app.AccessToken), params, http.MethodPost)
	// 定义
	var response CgiBinWxaAppCreateWxaQrCodeResponse
	err = json.Unmarshal(request.Body, &response)
	return NewCgiBinWxaAppCreateWxaQrCodeResult(response, request.Body, request, err)
}
