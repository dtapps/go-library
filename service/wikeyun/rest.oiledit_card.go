package wikeyun

import "go.dtapp.net/library/utils/gorequest"

// RestOilCardEdit 编辑充值卡
func (c *Client) RestOilCardEdit(notMustParams ...gorequest.Params) (body []byte, err error) {
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	// 请求
	request, err := c.request(apiUrl+"/rest/Oil/editCard", params)
	return request.ResponseBody, err
}
