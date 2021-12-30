package jd

import "encoding/json"

type UnionOpenGoodsPromotionGoodsInfoQueryResultResponse struct {
	JdUnionOpenGoodsPromotiongoodsinfoQueryResponce struct {
		Code        string `json:"code"`
		QueryResult string `json:"queryResult"`
	} `json:"jd_union_open_goods_promotiongoodsinfo_query_responce"`
}

type UnionOpenGoodsPromotionGoodsInfoQueryQueryResult struct {
	Code int `json:"code"`
	Data []struct {
		UnitPrice         float64 `json:"unitPrice"`
		MaterialUrl       string  `json:"materialUrl"`
		EndDate           int64   `json:"endDate"`
		IsFreeFreightRisk int     `json:"isFreeFreightRisk"`
		IsFreeShipping    int     `json:"isFreeShipping"`
		CommisionRatioWl  float64 `json:"commisionRatioWl"`
		CommisionRatioPc  float64 `json:"commisionRatioPc"`
		ImgUrl            string  `json:"imgUrl"`
		Vid               int     `json:"vid"`
		CidName           string  `json:"cidName"`
		WlUnitPrice       float64 `json:"wlUnitPrice"`
		Cid2Name          string  `json:"cid2Name"`
		IsSeckill         int     `json:"isSeckill"`
		Cid2              int     `json:"cid2"`
		Cid3Name          string  `json:"cid3Name"`
		Unt               int     `json:"unt"`
		Cid3              int     `json:"cid3"`
		ShopId            int     `json:"shopId"`
		IsJdSale          int     `json:"isJdSale"`
		GoodsName         string  `json:"goodsName"`
		SkuId             int64   `json:"skuId"`
		StartDate         int64   `json:"startDate"`
		Cid               int64   `json:"cid"`
	} `json:"data"`
	Message   string `json:"message"`
	RequestId string `json:"requestId"`
}

type UnionOpenGoodsPromotionGoodsInfoQueryResult struct {
	Responce UnionOpenGoodsPromotionGoodsInfoQueryResultResponse // 结果
	Result   UnionOpenGoodsPromotionGoodsInfoQueryQueryResult    // 结果
	body     []byte                                              // 内容
	Err      error                                               // 错误
}

func NewUnionOpenGoodsPromotionGoodsInfoQueryResult(responce UnionOpenGoodsPromotionGoodsInfoQueryResultResponse, result UnionOpenGoodsPromotionGoodsInfoQueryQueryResult, body []byte, err error) *UnionOpenGoodsPromotionGoodsInfoQueryResult {
	return &UnionOpenGoodsPromotionGoodsInfoQueryResult{Responce: responce, Result: result, body: body, Err: err}
}

// UnionOpenGoodsPromotionGoodsInfoQuery
// 通过SKUID查询推广商品的名称、主图、类目、价格、物流、是否自营、30天引单数量等详细信息，支持批量获取。通常用于在媒体侧展示商品详情。
// https://union.jd.com/openplatform/api/v2?apiName=jd.union.open.goods.promotiongoodsinfo.query
func (app *App) UnionOpenGoodsPromotionGoodsInfoQuery(notMustParams ...Params) *UnionOpenGoodsPromotionGoodsInfoQueryResult {
	// 参数
	params := NewParamsWithType("jd.union.open.goods.promotiongoodsinfo.query", notMustParams...)
	// 请求
	body, err := app.request(params)
	// 定义
	var responce UnionOpenGoodsPromotionGoodsInfoQueryResultResponse
	var result UnionOpenGoodsPromotionGoodsInfoQueryQueryResult
	err = json.Unmarshal(body, &responce)
	err = json.Unmarshal([]byte(responce.JdUnionOpenGoodsPromotiongoodsinfoQueryResponce.QueryResult), &result)
	return NewUnionOpenGoodsPromotionGoodsInfoQueryResult(responce, result, body, err)
}
