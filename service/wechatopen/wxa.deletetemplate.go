package wechatopen

import (
	"context"
	"net/http"

	"go.dtapp.net/library/utils/gorequest"
)

// WxaDeleteTemplate 删除指定代码模板
// https://developers.weixin.qq.com/doc/oplatform/Third-party_Platforms/2.0/api/ThirdParty/code_template/deletetemplate.html
func (c *Client) WxaDeleteTemplate(ctx context.Context, componentAccessToken, templateId string, notMustParams ...*gorequest.Params) (response APIResponse, err error) {

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	params.Set("template_id", templateId)

	// 请求
	err = c.request(ctx, "wxa/deletetemplate?access_token="+componentAccessToken, params, http.MethodPost, &response)
	return
}

// ErrcodeInfo 错误描述
func GetWxaDeleteTemplateErrcodeInfo(errcode int, errmsg string) string {
	switch errcode {
	case 85064:
		return "找不到模板，请检查模板id是否输入正确"
	default:
		return errmsg
	}
}
