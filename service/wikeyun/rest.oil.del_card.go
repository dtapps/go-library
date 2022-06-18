package wikeyun

// RestOilCardDel 油卡删除
func (c *Client) RestOilCardDel(notMustParams ...Params) (body []byte, err error) {
	// 参数
	params := c.NewParamsWith(notMustParams...)
	// 请求
	request, err := c.request("https://router.wikeyun.cn/rest/Oil/delCard", params)
	return request.ResponseBody, err
}
