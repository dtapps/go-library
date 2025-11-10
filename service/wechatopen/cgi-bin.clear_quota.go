package wechatopen

import (
	"context"
	"net/http"

	"go.dtapp.net/library/utils/gorequest"
)

// ClearQuota 重置API调用次数
// https://developers.weixin.qq.com/doc/oplatform/openApi/OpenApiDoc/openapi/clearQuota.html
func (c *Client) ClearQuota(ctx context.Context, authorizerAccessToken, appid string, notMustParams ...*gorequest.Params) (response APIResponse, err error) {

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	params.Set("appid", appid)

	// 请求
	err = c.request(ctx, "cgi-bin/clear_quota?access_token="+authorizerAccessToken, params, http.MethodPost, &response)
	return
}
