package jd

import (
	"context"
)

func (c *Client) Post(ctx context.Context, _method string, notMustParams ...Params) ([]byte, error) {
	// 参数
	params := NewParamsWithType(_method, notMustParams...)
	// 请求
	request, err := c.request(ctx, params)
	// 定义
	return request.ResponseBody, err
}
