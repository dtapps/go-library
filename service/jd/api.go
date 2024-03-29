package jd

import (
	"context"
	"github.com/dtapps/go-library/utils/gorequest"
)

func (c *Client) Post(ctx context.Context, _method string, notMustParams ...gorequest.Params) ([]byte, error) {
	// 参数
	params := NewParamsWithType(_method, notMustParams...)
	// 请求
	request, err := c.request(ctx, params)
	// 定义
	return request.ResponseBody, err
}
