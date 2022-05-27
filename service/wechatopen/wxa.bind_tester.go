package wechatopen

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type WxaBindTesterResponse struct {
	Errcode int    `json:"errcode"` // 错误码
	Errmsg  string `json:"errmsg"`  // 错误信息
	Userstr string `json:"userstr"` // 人员对应的唯一字符串
}

type WxaBindTesterResult struct {
	Result WxaBindTesterResponse // 结果
	Body   []byte                // 内容
	Err    error                 // 错误
}

func NewWxaBindTesterResult(result WxaBindTesterResponse, body []byte, err error) *WxaBindTesterResult {
	return &WxaBindTesterResult{Result: result, Body: body, Err: err}
}

// WxaBindTester 绑定微信用户为体验者
// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/Mini_Program_AdminManagement/Admin.html
func (app *App) WxaBindTester(wechatid string) *WxaBindTesterResult {
	app.authorizerAccessToken = app.GetAuthorizerAccessToken()
	// 参数
	params := NewParams()
	params["wechatid"] = wechatid
	// 请求
	body, err := app.request(fmt.Sprintf("https://api.weixin.qq.com/wxa/bind_tester?access_token=%s", app.authorizerAccessToken), params, http.MethodPost)
	// 定义
	var response WxaBindTesterResponse
	err = json.Unmarshal(body, &response)
	return NewWxaBindTesterResult(response, body, err)
}

// ErrcodeInfo 错误描述
func (resp *WxaBindTesterResult) ErrcodeInfo() string {
	switch resp.Result.Errcode {
	case 85001:
		return "微信号不存在或微信号设置为不可搜索"
	case 85002:
		return "小程序绑定的体验者数量达到上限"
	case 85003:
		return "微信号绑定的小程序体验者达到上限"
	case 85004:
		return "微信号已经绑定"
	}
	return "系统繁忙"
}