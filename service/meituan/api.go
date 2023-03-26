package meituan

import (
	"context"
	"github.com/dtapps/go-library/utils/gorequest"
	"net/http"
)

func (c *Client) Get(ctx context.Context, _method string, notMustParams ...gorequest.Params) ([]byte, error) {
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	// 请求
	request, err := c.request(ctx, apiUrl+_method, params, http.MethodGet)
	// 定义
	return request.ResponseBody, err
}

func (c *Client) Post(ctx context.Context, _method string, notMustParams ...gorequest.Params) ([]byte, error) {
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	// 请求
	request, err := c.request(ctx, apiUrl+_method, params, http.MethodPost)
	// 定义
	return request.ResponseBody, err
}
