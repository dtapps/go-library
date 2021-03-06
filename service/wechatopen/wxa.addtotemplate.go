package wechatopen

import (
	"encoding/json"
	"fmt"
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
	Err    error                    // 错误
}

func newWxaAddToTemplateResult(result WxaAddToTemplateResponse, body []byte, http gorequest.Response, err error) *WxaAddToTemplateResult {
	return &WxaAddToTemplateResult{Result: result, Body: body, Http: http, Err: err}
}

// WxaAddToTemplate 将草稿添加到代码模板库
// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/ThirdParty/code_template/addtotemplate.html
func (c *Client) WxaAddToTemplate(draftId string, templateType int) *WxaAddToTemplateResult {
	// 参数
	params := gorequest.NewParams()
	params["draft_id"] = draftId
	params["template_type"] = templateType
	// 请求
	request, err := c.request(fmt.Sprintf(apiUrl+"/wxa/addtotemplate?access_token=%s", c.GetComponentAccessToken()), params, http.MethodPost)
	// 定义
	var response WxaAddToTemplateResponse
	err = json.Unmarshal(request.ResponseBody, &response)
	return newWxaAddToTemplateResult(response, request.ResponseBody, request, err)
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
