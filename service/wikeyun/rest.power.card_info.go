package wikeyun

import (
	"context"
	"github.com/dtapps/go-library/utils/gorequest"
)

// PowerCardInfo 电费充值卡详情
// https://open.wikeyun.cn/#/apiDocument/9/document/333
func (c *Client) PowerCardInfo(ctx context.Context, notMustParams ...gorequest.Params) (body []byte, err error) {
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	// 请求
	request, err := c.request(ctx, apiUrl+"/rest/Power/cardInfo", params)
	return request.ResponseBody, err
}
