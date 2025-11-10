package wechatopen

import (
	"context"
	"net/http"

	"go.dtapp.net/library/utils/gorequest"
)

// WxaAddToTemplate 将草稿添加到代码模板库
// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/ThirdParty/code_template/addtotemplate.html
func (c *Client) WxaAddToTemplate(ctx context.Context, componentAccessToken, draftId string, templateType int, notMustParams ...*gorequest.Params) (response APIResponse, err error) {

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	params.Set("draft_id", draftId)
	params.Set("template_type", templateType)

	// 请求
	err = c.request(ctx, "wxa/addtotemplate?access_token="+componentAccessToken, params, http.MethodPost, &response)
	return
}

// ErrcodeInfo 错误描述
func GetWxaAddToTemplateErrcodeInfo(errcode int, errmsg string) string {
	switch errcode {
	case 85064:
		return "找不到草稿"
	case 85065:
		return "模板库已满"
	default:
		return errmsg
	}
}
