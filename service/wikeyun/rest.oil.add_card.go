package wikeyun

// RestOilCardAdd 添加充值卡
func (c *Client) RestOilCardAdd(notMustParams ...Params) (body []byte, err error) {
	// 参数
	params := c.NewParamsWith(notMustParams...)
	// 请求
	request, err := c.request("https://router.wikeyun.cn/rest/Oil/addCard", params)
	return request.ResponseBody, err
}
