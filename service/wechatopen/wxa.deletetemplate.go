package wechatopen

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type WxaDeleteTemplateResponse struct {
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

type WxaDeleteTemplateResult struct {
	Result WxaDeleteTemplateResponse // 结果
	Body   []byte                    // 内容
	Err    error                     // 错误
}

func NewWxaDeleteTemplateResult(result WxaDeleteTemplateResponse, body []byte, err error) *WxaDeleteTemplateResult {
	return &WxaDeleteTemplateResult{Result: result, Body: body, Err: err}
}

// WxaDeleteTemplate 删除指定代码模板
// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/ThirdParty/code_template/deletetemplate.html
func (app *App) WxaDeleteTemplate(templateId string) *WxaDeleteTemplateResult {
	app.componentAccessToken = app.GetComponentAccessToken()
	// 参数
	params := NewParams()
	params.Set("template_id", templateId)
	// 请求
	body, err := app.request(fmt.Sprintf("https://api.weixin.qq.com/wxa/deletetemplate?access_token=%s", app.componentAccessToken), params, http.MethodPost)
	// 定义
	var response WxaDeleteTemplateResponse
	err = json.Unmarshal(body, &response)
	return NewWxaDeleteTemplateResult(response, body, err)
}

// ErrcodeInfo 错误描述
func (resp *WxaDeleteTemplateResult) ErrcodeInfo() string {
	switch resp.Result.Errcode {
	case 85064:
		return "找不到模板，请检查模板id是否输入正确"
	}
	return "系统繁忙"
}
