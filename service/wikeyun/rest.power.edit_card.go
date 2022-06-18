package wikeyun

// RestPowerEditCard 编辑电费充值卡
// https://open.wikeyun.cn/#/apiDocument/9/document/329
func (c *Client) RestPowerEditCard(notMustParams ...Params) (body []byte, err error) {
	// 参数
	params := c.NewParamsWith(notMustParams...)
	// 请求
	request, err := c.request("https://router.wikeyun.cn/rest/Power/editCard", params)
	return request.ResponseBody, err
}
