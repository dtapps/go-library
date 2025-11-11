package wechatopen

import (
	"context"
	"net/http"

	"go.dtapp.net/library/utils/gorequest"
)

// ClearComponentQuotaByAppSecret 使用AppSecret重置第三方平台 API 调用次数
// https://developers.weixin.qq.com/doc/oplatform/openApi/OpenApiDoc/openapi/clearComponentQuotaByAppSecret.html
func (c *Client) ClearComponentQuotaByAppSecret(ctx context.Context, notMustParams ...*gorequest.Params) (response APIResponse, err error) {

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	params.Set("appid", c.GetAuthorizerAppid())
	params.Set("component_appid", c.GetComponentAppId())
	params.Set("appsecret", c.GetComponentAppSecret())

	// 请求
	err = c.request(ctx, "cgi-bin/component/clear_quota/v2?access_token="+c.GetComponentAccessToken(), params, http.MethodPost, &response)
	return
}
