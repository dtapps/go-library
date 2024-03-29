package sendcloud

import (
	"context"
	"github.com/baidubce/bce-sdk-go/http"
	"github.com/dtapps/go-library/utils/gorequest"
)

func (c *Client) Get(ctx context.Context, _method string, notMustParams ...gorequest.Params) ([]byte, error) {
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	// 请求
	request, err := c.request(ctx, apiUrl+_method, params, http.GET)
	// 定义
	return request.ResponseBody, err
}

func (c *Client) Post(ctx context.Context, _method string, notMustParams ...gorequest.Params) ([]byte, error) {
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	// 请求
	request, err := c.request(ctx, apiUrl+_method, params, http.POST)
	// 定义
	return request.ResponseBody, err
}
