package wikeyun

// RestOilOrderPush 充值下单
func (c *Client) RestOilOrderPush(notMustParams ...Params) (body []byte, err error) {
	// 参数
	params := c.NewParamsWith(notMustParams...)
	// 请求
	request, err := c.request("https://router.wikeyun.cn/rest/Oil/pushOrder", params)
	return request.ResponseBody, err
}
