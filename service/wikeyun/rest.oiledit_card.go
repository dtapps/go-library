package wikeyun

// RestOilCardEdit 编辑充值卡
func (c *Client) RestOilCardEdit(notMustParams ...Params) (body []byte, err error) {
	// 参数
	params := c.NewParamsWith(notMustParams...)
	// 请求
	request, err := c.request("https://router.wikeyun.cn/rest/Oil/editCard", params)
	return request.ResponseBody, err
}
