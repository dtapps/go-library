package wechatminiprogram

import (
	"encoding/json"
	"fmt"
	"go.dtapp.net/library/utils/gorequest"
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
	Http   gorequest.Response                  // 请求
	Err    error                               // 错误
}

func NewCgiBinWxaAppCreateWxaQrCodeResult(result CgiBinWxaAppCreateWxaQrCodeResponse, body []byte, http gorequest.Response, err error) *CgiBinWxaAppCreateWxaQrCodeResult {
	return &CgiBinWxaAppCreateWxaQrCodeResult{Result: result, Body: body, Http: http, Err: err}
}

// CgiBinWxaAppCreateWxaQrCode 获取小程序二维码，适用于需要的码数量较少的业务场景。通过该接口生成的小程序码，永久有效，有数量限制
// https://developers.weixin.qq.com/miniprogram/dev/api-backend/open-api/qr-code/wxacode.createQRCode.html
func (c *Client) CgiBinWxaAppCreateWxaQrCode(notMustParams ...Params) *CgiBinWxaAppCreateWxaQrCodeResult {
	// 参数
	params := c.NewParamsWith(notMustParams...)
	// 请求
	request, err := c.request(fmt.Sprintf("https://api.weixin.qq.com/cgi-bin/wxaapp/createwxaqrcode?access_token=%s", c.getAccessToken()), params, http.MethodPost)
	// 定义
	var response CgiBinWxaAppCreateWxaQrCodeResponse
	err = json.Unmarshal(request.ResponseBody, &response)
	return NewCgiBinWxaAppCreateWxaQrCodeResult(response, request.ResponseBody, request, err)
}
