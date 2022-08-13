package wikeyun

import (
	"context"
	"github.com/dtapps/go-library/utils/gorequest"
)

// RestPowerEditCard 编辑电费充值卡
// https://open.wikeyun.cn/#/apiDocument/9/document/329
func (c *Client) RestPowerEditCard(ctx context.Context, notMustParams ...gorequest.Params) (body []byte, err error) {
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	// 请求
	request, err := c.request(ctx, apiUrl+"/rest/Power/editCard", params)
	return request.ResponseBody, err
}
