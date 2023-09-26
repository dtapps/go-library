package wechatpayapiv3

import (
	"context"
	"github.com/dtapps/go-library/utils/gorequest"
	"net/http"
)

func (c *Client) Get(ctx context.Context, _method string, commonParams bool, notMustParams ...gorequest.Params) ([]byte, error) {
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	// 请求
	request, err := c.request(ctx, apiUrl+_method, params, http.MethodGet, commonParams)
	// 定义
	return request.ResponseBody, err
}

func (c *Client) Post(ctx context.Context, _method string, commonParams bool, notMustParams ...gorequest.Params) ([]byte, error) {
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	// 请求
	request, err := c.request(ctx, apiUrl+_method, params, http.MethodPost, commonParams)
	// 定义
	return request.ResponseBody, err
}
