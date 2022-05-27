package wechatopen

import (
	"encoding/json"
	"fmt"
	"go.dtapp.net/library/gorequest"
	"net/http"
)

type CgiBinWxOpenQrCodeJumpPublishResponse struct {
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

type CgiBinWxOpenQrCodeJumpPublishResult struct {
	Result CgiBinWxOpenQrCodeJumpPublishResponse // 结果
	Body   []byte                                // 内容
	Http   gorequest.Response                    // 请求
	Err    error                                 // 错误
}

func NewCgiBinWxOpenQrCodeJumpPublishResult(result CgiBinWxOpenQrCodeJumpPublishResponse, body []byte, http gorequest.Response, err error) *CgiBinWxOpenQrCodeJumpPublishResult {
	return &CgiBinWxOpenQrCodeJumpPublishResult{Result: result, Body: body, Http: http, Err: err}
}

// CgiBinWxOpenQrCodeJumpPublish 发布已设置的二维码规则
// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/qrcode/qrcodejumppublish.html
func (app *App) CgiBinWxOpenQrCodeJumpPublish(prefix string) *CgiBinWxOpenQrCodeJumpPublishResult {
	// 参数
	params := NewParams()
	params["prefix"] = prefix
	// 请求
	request, err := app.request(fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/wxopen/qrcodejumppublish?access_token=%s", app.GetAuthorizerAccessToken()), params, http.MethodPost)
	// 定义
	var response CgiBinWxOpenQrCodeJumpPublishResponse
	err = json.Unmarshal(request.ResponseBody, &response)
	return NewCgiBinWxOpenQrCodeJumpPublishResult(response, request.ResponseBody, request, err)
}
