package wechatopen

import (
	"context"
	"fmt"
	"go.dtapp.net/library/utils/gorequest"
)

type RequestAuthorizerParams struct {
	Url         string
	UrlMethod   string
	AccessToken string
	Response    any
}

// RequestAuthorizer 发起小程序
func (c *Client) RequestAuthorizer(ctx context.Context, param RequestAuthorizerParams, notMustParams ...*gorequest.Params) (gorequest.Response, error) {
	if param.Url == "" {
		return gorequest.Response{}, fmt.Errorf("url is empty")
	}
	if param.UrlMethod == "" {
		return gorequest.Response{}, fmt.Errorf("url_method is empty")
	}
	if param.AccessToken == "" {
		return gorequest.Response{}, fmt.Errorf("access_token is empty")
	}
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	// 请求
	request, err := c.request(ctx, param.Url+"?access_token="+param.AccessToken, params, param.UrlMethod, &param.Response)
	return request, err
}
