package wechatopen

import (
	"context"
	"github.com/dtapps/go-library/utils/gojson"
	"github.com/dtapps/go-library/utils/gorequest"
	"net/http"
)

type WxaAddToTemplateResponse struct {
	Errcode int    `json:"errcode"` // 错误码
	Errmsg  string `json:"errmsg"`  // 错误信息
}

type WxaAddToTemplateResult struct {
	Result WxaAddToTemplateResponse // 结果
	Body   []byte                   // 内容
	Http   gorequest.Response       // 请求
}

func newWxaAddToTemplateResult(result WxaAddToTemplateResponse, body []byte, http gorequest.Response) *WxaAddToTemplateResult {
	return &WxaAddToTemplateResult{Result: result, Body: body, Http: http}
}

// WxaAddToTemplate 将草稿添加到代码模板库
// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/ThirdParty/code_template/addtotemplate.html
func (c *Client) WxaAddToTemplate(ctx context.Context, draftId string, templateType int, notMustParams ...gorequest.Params) (*WxaAddToTemplateResult, error) {
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	params.Set("draft_id", draftId)
	params.Set("template_type", templateType)
	// 请求
	request, err := c.request(ctx, apiUrl+"/wxa/addtotemplate?access_token="+GetComponentAccessToken(ctx, c), params, http.MethodPost)
	if err != nil {
		return newWxaAddToTemplateResult(WxaAddToTemplateResponse{}, request.ResponseBody, request), err
	}
	// 定义
	var response WxaAddToTemplateResponse
	err = gojson.Unmarshal(request.ResponseBody, &response)
	return newWxaAddToTemplateResult(response, request.ResponseBody, request), err
}

// ErrcodeInfo 错误描述
func (resp *WxaAddToTemplateResult) ErrcodeInfo() string {
	switch resp.Result.Errcode {
	case 85064:
		return "找不到草稿"
	case 85065:
		return "模板库已满"
	}
	return "系统繁忙"
}
