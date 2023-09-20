package wikeyun

import (
	"context"
	"github.com/dtapps/go-library/utils/gorequest"
)

// RestOilOrderPush 充值下单
func (c *Client) RestOilOrderPush(ctx context.Context, notMustParams ...*gorequest.Params) (body []byte, err error) {
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	// 请求
	request, err := c.request(ctx, apiUrl+"/rest/Oil/pushOrder", params)
	return request.ResponseBody, err
}
