package wikeyun

import "go.dtapp.net/library/utils/gorequest"

// RestOilOrderQuery 订单查询
func (c *Client) RestOilOrderQuery(notMustParams ...gorequest.Params) (body []byte, err error) {
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	// 请求
	request, err := c.request(apiUrl+"/rest/Oil/query", params)
	return request.ResponseBody, err
}
