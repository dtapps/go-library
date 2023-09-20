package alipayopen

import (
	"context"
	"github.com/dtapps/go-library/utils/gorequest"
)

func (c *Client) Get(ctx context.Context, _method string, notMustParams ...*gorequest.Params) ([]byte, error) {
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	// 请求
	request, err := c.request(ctx, c.newParamsWithType(_method, params))
	if err != nil {
		return nil, err
	}
	return request.ResponseBody, err
}
