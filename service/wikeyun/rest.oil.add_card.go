package wikeyun

import (
	"context"
	"github.com/dtapps/go-library/utils/gorequest"
)

// RestOilCardAdd 添加充值卡
func (c *Client) RestOilCardAdd(ctx context.Context, notMustParams ...gorequest.Params) (body []byte, err error) {
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	// 请求
	request, err := c.request(ctx, apiUrl+"/rest/Oil/addCard", params)
	return request.ResponseBody, err
}
