package taobao

import "encoding/json"

type TbkShopRecommendGetResponse struct {
	TbkShopRecommendGetResponse struct {
		Results struct {
			NTbkShop []struct {
				UserId     int    `json:"user_id"`
				ShopTitle  string `json:"shop_title"`
				ShopType   string `json:"shop_type"`
				SellerNick string `json:"seller_nick"`
				PictUrl    string `json:"pict_url"`
				ShopUrl    string `json:"shop_url"`
			} `json:"n_tbk_shop"`
		} `json:"results"`
	} `json:"tbk_shop_recommend_get_response"`
}

type TbkShopRecommendGetResult struct {
	Result TbkShopRecommendGetResponse // 结果
	Body   []byte                      // 内容
	Err    error                       // 错误
}

func NewTbkShopRecommendGetResult(result TbkShopRecommendGetResponse, body []byte, err error) *TbkShopRecommendGetResult {
	return &TbkShopRecommendGetResult{Result: result, Body: body, Err: err}
}

// TbkShopRecommendGet 淘宝客-公用-店铺关联推荐
// https://open.taobao.com/api.htm?docId=24522&docType=2
func (app *App) TbkShopRecommendGet(notMustParams ...Params) *TbkShopRecommendGetResult {
	// 参数
	params := NewParamsWithType("taobao.tbk.shop.recommend.get", notMustParams...)
	// 请求
	body, err := app.request(params)
	// 定义
	var response TbkShopRecommendGetResponse
	err = json.Unmarshal(body, &response)
	return NewTbkShopRecommendGetResult(response, body, err)
}
