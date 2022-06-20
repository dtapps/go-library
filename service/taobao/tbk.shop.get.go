package taobao

import (
	"encoding/json"
	"go.dtapp.net/library/utils/gorequest"
)

type TbkShopGetResponse struct {
	TbkShopGetResponse struct {
		Results struct {
			NTbkShop []struct {
				PictUrl    string `json:"pict_url"`
				SellerNick string `json:"seller_nick"`
				ShopTitle  string `json:"shop_title"`
				ShopType   string `json:"shop_type"`
				ShopUrl    string `json:"shop_url"`
				UserId     int64  `json:"user_id"`
			} `json:"n_tbk_shop"`
		} `json:"results"`
		TotalResults int    `json:"total_results"`
		RequestId    string `json:"request_id"`
	} `json:"tbk_shop_get_response"`
}

type TbkShopGetResult struct {
	Result TbkShopGetResponse // 结果
	Body   []byte             // 内容
	Http   gorequest.Response // 请求
	Err    error              // 错误
}

func newTbkShopGetResult(result TbkShopGetResponse, body []byte, http gorequest.Response, err error) *TbkShopGetResult {
	return &TbkShopGetResult{Result: result, Body: body, Http: http, Err: err}
}

// TbkShopGet 淘宝客-推广者-店铺搜索
// https://open.taobao.com/api.htm?docId=24521&docType=2
func (c *Client) TbkShopGet(notMustParams ...Params) *TbkShopGetResult {
	// 参数
	params := NewParamsWithType("taobao.tbk.shop.get", notMustParams...)
	// 请求
	request, err := c.request(params)
	// 定义
	var response TbkShopGetResponse
	err = json.Unmarshal(request.ResponseBody, &response)
	return newTbkShopGetResult(response, request.ResponseBody, request, err)
}
