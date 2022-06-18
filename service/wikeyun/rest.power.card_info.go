package wikeyun

import "go.dtapp.net/library/utils/gorequest"

// PowerCardInfo 电费充值卡详情
// https://open.wikeyun.cn/#/apiDocument/9/document/333
func (c *Client) PowerCardInfo(notMustParams ...gorequest.Params) (body []byte, err error) {
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	// 请求
	request, err := c.request(apiUrl+"/rest/Power/cardInfo", params)
	return request.ResponseBody, err
}
