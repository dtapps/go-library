package wikeyun

// PowerCardInfo 电费充值卡详情
// https://open.wikeyun.cn/#/apiDocument/9/document/333
func (c *Client) PowerCardInfo(notMustParams ...Params) (body []byte, err error) {
	// 参数
	params := c.NewParamsWith(notMustParams...)
	// 请求
	request, err := c.request("https://router.wikeyun.cn/rest/Power/cardInfo", params)
	return request.ResponseBody, err
}
