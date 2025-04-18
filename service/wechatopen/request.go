package wechatopen

import (
	"context"
	"go.dtapp.net/library/utils/gorequest"
	"net/http"
)

type RequestAuthorizerParams struct {
	Url         string
	AccessToken string
	Response    any
}

// RequestAuthorizer 发起小程序
func (c *Client) RequestAuthorizer(ctx context.Context, param RequestAuthorizerParams, notMustParams ...*gorequest.Params) (gorequest.Response, error) {
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	// 请求
	request, err := c.request(ctx, param.Url+"?access_token="+param.AccessToken, params, http.MethodPost, &param.Response)
	return request, err
}
