package wechatopen

import (
	"context"
	"net/http"

	"go.dtapp.net/library/utils/gorequest"
)

type ModifyThirdpartyServerDomainResponse struct {
	APIResponse                     // 错误
	PublishedWxaServerDomain string `json:"published_wxa_server_domain"` // 目前生效的 “全网发布版”第三方平台“小程序服务器域名”。如果修改失败，该字段不会返回。如果没有已发布的第三方平台，该字段也不会返回。
	TestingWxaServerDomain   string `json:"testing_wxa_server_domain"`   // 目前生效的 “测试版”第三方平台“小程序服务器域名”。如果修改失败，该字段不会返回
	InvalidWxaServerDomain   string `json:"invalid_wxa_server_domain"`   // 未通过验证的域名。如果不存在未通过验证的域名，该字段不会返回。
}

// ModifyThirdpartyServerDomain 设置第三方平台服务器域名
// https://developers.weixin.qq.com/doc/oplatform/openApi/OpenApiDoc/thirdparty-management/domain-mgnt/modifyThirdpartyServerDomain.html
func (c *Client) ModifyThirdpartyServerDomain(ctx context.Context, action string, notMustParams ...*gorequest.Params) (response ModifyThirdpartyServerDomainResponse, err error) {

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	params.Set("action", action)

	// 请求
	err = c.request(ctx, "cgi-bin/component/modify_wxa_server_domain?access_token="+c.GetComponentAccessToken(), params, http.MethodPost, &response)
	return
}
