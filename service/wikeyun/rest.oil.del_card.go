package wikeyun

import "github.com/dtapps/go-library/utils/gorequest"

// RestOilCardDel 油卡删除
func (c *Client) RestOilCardDel(notMustParams ...gorequest.Params) (body []byte, err error) {
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	// 请求
	request, err := c.request(apiUrl+"/rest/Oil/delCard", params)
	return request.ResponseBody, err
}
