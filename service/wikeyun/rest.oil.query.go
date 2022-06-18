package wikeyun

// RestOilOrderQuery 订单查询
func (c *Client) RestOilOrderQuery(notMustParams ...Params) (body []byte, err error) {
	// 参数
	params := c.NewParamsWith(notMustParams...)
	// 请求
	request, err := c.request("https://router.wikeyun.cn/rest/Oil/query", params)
	return request.ResponseBody, err
}
