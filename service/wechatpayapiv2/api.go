package wechatpayapiv2

import (
	"context"
	"crypto/tls"
	"github.com/dtapps/go-library/utils/gorequest"
)

func (c *Client) Post(ctx context.Context, _method string, certStatus bool, cert *tls.Certificate, notMustParams ...*gorequest.Params) ([]byte, error) {
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	// 请求
	request, err := c.request(ctx, apiUrl+_method, params, true, cert)
	// 定义
	return request.ResponseBody, err
}
