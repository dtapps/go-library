package wikeyun

import "go.dtapp.net/library/utils/gorequest"

// RestPowerEditCard 编辑电费充值卡
// https://open.wikeyun.cn/#/apiDocument/9/document/329
func (c *Client) RestPowerEditCard(notMustParams ...gorequest.Params) (body []byte, err error) {
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	// 请求
	request, err := c.request(apiUrl+"/rest/Power/editCard", params)
	return request.ResponseBody, err
}
