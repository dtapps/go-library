package wechatopen

import (
	"context"
	"github.com/dtapps/go-library/utils/gojson"
	"github.com/dtapps/go-library/utils/gorequest"
	"net/http"
)

type WxaDeleteTemplateResponse struct {
	Errcode int    `json:"errcode"`
	Errmsg  string `json:"errmsg"`
}

type WxaDeleteTemplateResult struct {
	Result WxaDeleteTemplateResponse // 结果
	Body   []byte                    // 内容
	Http   gorequest.Response        // 请求
}

func newWxaDeleteTemplateResult(result WxaDeleteTemplateResponse, body []byte, http gorequest.Response) *WxaDeleteTemplateResult {
	return &WxaDeleteTemplateResult{Result: result, Body: body, Http: http}
}

// WxaDeleteTemplate 删除指定代码模板
// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/ThirdParty/code_template/deletetemplate.html
func (c *Client) WxaDeleteTemplate(ctx context.Context, componentAccessToken, templateId string, notMustParams ...gorequest.Params) (*WxaDeleteTemplateResult, error) {
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	params.Set("template_id", templateId)
	// 请求
	request, err := c.request(ctx, apiUrl+"/wxa/deletetemplate?access_token="+componentAccessToken, params, http.MethodPost)
	if err != nil {
		return newWxaDeleteTemplateResult(WxaDeleteTemplateResponse{}, request.ResponseBody, request), err
	}
	// 定义
	var response WxaDeleteTemplateResponse
	err = gojson.Unmarshal(request.ResponseBody, &response)
	return newWxaDeleteTemplateResult(response, request.ResponseBody, request), err
}

// ErrcodeInfo 错误描述
func (resp *WxaDeleteTemplateResult) ErrcodeInfo() string {
	switch resp.Result.Errcode {
	case 85064:
		return "找不到模板，请检查模板id是否输入正确"
	}
	return "系统繁忙"
}
