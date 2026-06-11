package wechatopen

import (
	"context"
	"net/http"

	"go.dtapp.net/library/utils/gorequest"
)

type GetEffectiveServerDomainResponse struct {
	APIResponse // 错误
	MpDomain    struct {
		Requestdomain   []any `json:"requestdomain"`
		Wsrequestdomain []any `json:"wsrequestdomain"`
		Uploaddomain    []any `json:"uploaddomain"`
		Downloaddomain  []any `json:"downloaddomain"`
		Udpdomain       []any `json:"udpdomain"`
		Tcpdomain       []any `json:"tcpdomain"`
	} `json:"mp_domain"`
	ThirdDomain struct {
		Requestdomain   []any `json:"requestdomain"`
		Wsrequestdomain []any `json:"wsrequestdomain"`
		Uploaddomain    []any `json:"uploaddomain"`
		Downloaddomain  []any `json:"downloaddomain"`
		Udpdomain       []any `json:"udpdomain"`
		Tcpdomain       []any `json:"tcpdomain"`
	} `json:"third_domain"`
	DirectDomain struct {
		Requestdomain   []any `json:"requestdomain"`
		Wsrequestdomain []any `json:"wsrequestdomain"`
		Uploaddomain    []any `json:"uploaddomain"`
		Downloaddomain  []any `json:"downloaddomain"`
		Udpdomain       []any `json:"udpdomain"`
		Tcpdomain       []any `json:"tcpdomain"`
	} `json:"direct_domain"`
}

// GetEffectiveServerDomain 获取发布后生效服务器域名列表
// https://developers.weixin.qq.com/doc/oplatform/openApi/miniprogram-management/domain-management/api_geteffectiveserverdomain.html
func (c *Client) GetEffectiveServerDomain(ctx context.Context, notMustParams ...*gorequest.Params) (response GetEffectiveServerDomainResponse, err error) {

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)

	// 请求
	err = c.request(ctx, "wxa/get_effective_domain?access_token="+c.GetAuthorizerAccessToken(), params, http.MethodPost, &response)
	return
}
