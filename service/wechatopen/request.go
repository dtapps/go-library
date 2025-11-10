package wechatopen

import (
	"context"
	"fmt"

	"go.dtapp.net/library/utils/gorequest"
)

type RequestAuthorizerParams struct {
	Url         string
	Method      string
	AccessToken string
	Response    any
}

// RequestAuthorizer 发起小程序
func (c *Client) RequestAuthorizer(ctx context.Context, param RequestAuthorizerParams, notMustParams ...*gorequest.Params) (err error) {
	if param.Url == "" {
		return fmt.Errorf("url is empty")
	}
	if param.Method == "" {
		return fmt.Errorf("method is empty")
	}
	if param.AccessToken == "" {
		return fmt.Errorf("access_token is empty")
	}
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)

	// 请求
	err = c.request(ctx, param.Url+"?access_token="+param.AccessToken, params, param.Method, &param.Response)
	return
}
