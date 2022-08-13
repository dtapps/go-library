package wikeyun

import (
	"context"
	"github.com/dtapps/go-library/utils/gorequest"
)

// RestOilOrderQuery 订单查询
func (c *Client) RestOilOrderQuery(ctx context.Context, notMustParams ...gorequest.Params) (body []byte, err error) {
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	// 请求
	request, err := c.request(ctx, apiUrl+"/rest/Oil/query", params)
	return request.ResponseBody, err
}
