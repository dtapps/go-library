package wechatopen

import (
	"context"
	"go.dtapp.net/library/utils/gorequest"
	"net/http"
)

// RequestAuthorizer 发起小程序
func (c *Client) RequestAuthorizer(ctx context.Context, url string, accessToken string, notMustParams ...*gorequest.Params) (gorequest.Response, error) {

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)

	// 请求
	var response WxaApiWxaembeddedAddEmbeddedResponse
	request, err := c.request(ctx, url+"?access_token="+accessToken, params, http.MethodPost, &response)
	return request, err
}
