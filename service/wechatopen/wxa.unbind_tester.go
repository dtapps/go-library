package wechatopen

import (
	"context"
	"net/http"

	"go.dtapp.net/library/utils/gorequest"
)

// UnbindTester 解除绑定体验者
// https://developers.weixin.qq.com/doc/oplatform/openApi/OpenApiDoc/miniprogram-management/member-management/unbindTester.html
func (c *Client) UnbindTester(ctx context.Context, wechatid string, userstr string, notMustParams ...*gorequest.Params) (response APIResponse, err error) {

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	if wechatid != "" {
		params.Set("wechatid", wechatid)
	}
	params.Set("userstr", userstr)

	// 请求
	err = c.request(ctx, "wxa/unbind_tester?access_token="+c.GetAuthorizerAccessToken(), params, http.MethodPost, &response)
	return
}
