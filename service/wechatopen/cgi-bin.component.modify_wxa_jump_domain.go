package wechatopen

import (
	"context"
	"go.dtapp.net/library/utils/gorequest"
	"net/http"
)

type ModifyThirdpartyJumpDomainResponse struct {
	Errcode                  int    `json:"errcode"`                      // 返回码
	Errmsg                   string `json:"errmsg"`                       // 返回码信息
	PublishedWxaJumpH5Domain string `json:"published_wxa_jump_h5_domain"` // 目前生效的 “全网发布版”第三方平台“小程序业务域名”。如果修改失败，该字段不会返回。如果没有已发布的第三方平台，该字段也不会返回。
	TestingWxaJumpH5Domain   string `json:"testing_wxa_jump_h5_domain"`   // 目前生效的 “测试版”第三方平台“小程序业务域名”。如果修改失败，该字段不会返回
	InvalidWxaJumpH5Domain   string `json:"invalid_wxa_jump_h5_domain"`   // 未通过验证的域名。如果不存在未通过验证的域名，该字段不会返回。
}

type ModifyThirdpartyJumpDomainResult struct {
	Result ModifyThirdpartyJumpDomainResponse // 结果
	Body   []byte                             // 内容
	Http   gorequest.Response                 // 请求
}

func newModifyThirdpartyJumpDomainResult(result ModifyThirdpartyJumpDomainResponse, body []byte, http gorequest.Response) *ModifyThirdpartyJumpDomainResult {
	return &ModifyThirdpartyJumpDomainResult{Result: result, Body: body, Http: http}
}

// ModifyThirdpartyJumpDomain 设置第三方平台业务域名
// https://developers.weixin.qq.com/doc/oplatform/openApi/OpenApiDoc/thirdparty-management/domain-mgnt/modifyThirdpartyJumpDomain.html
func (c *Client) ModifyThirdpartyJumpDomain(ctx context.Context, componentAccessToken string, action string, notMustParams ...gorequest.Params) (*ModifyThirdpartyJumpDomainResult, error) {

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	params.Set("action", action)

	// 请求
	var response ModifyThirdpartyJumpDomainResponse
	request, err := c.request(ctx, "cgi-bin/component/modify_wxa_jump_domain?access_token="+componentAccessToken, params, http.MethodPost, &response)
	return newModifyThirdpartyJumpDomainResult(response, request.ResponseBody, request), err
}
