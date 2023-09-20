package wikeyun

import (
	"context"
	"github.com/dtapps/go-library/utils/gorequest"
)

// RestOilCardInfo 油卡详情
func (c *Client) RestOilCardInfo(ctx context.Context, notMustParams ...*gorequest.Params) (body []byte, err error) {
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	// 请求
	request, err := c.request(ctx, apiUrl+"/rest/Oil/cardInfo", params)
	return request.ResponseBody, err
}
