package wikeyun

// RestOilCardInfo 油卡详情
func (c *Client) RestOilCardInfo(notMustParams ...Params) (body []byte, err error) {
	// 参数
	params := c.NewParamsWith(notMustParams...)
	// 请求
	request, err := c.request("https://router.wikeyun.cn/rest/Oil/cardInfo", params)
	return request.ResponseBody, err
}
