package wikeyun

import "go.dtapp.net/library/utils/gorequest"

// RestOilCardInfo 油卡详情
func (c *Client) RestOilCardInfo(notMustParams ...gorequest.Params) (body []byte, err error) {
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	// 请求
	request, err := c.request(apiUrl+"/rest/Oil/cardInfo", params)
	return request.ResponseBody, err
}
