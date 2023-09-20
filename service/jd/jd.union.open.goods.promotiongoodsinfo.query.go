package jd

import (
	"context"
	"github.com/dtapps/go-library/utils/gojson"
	"github.com/dtapps/go-library/utils/gorequest"
)

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
	Body     []byte                                              // 内容
	Http     gorequest.Response                                  // 请求
	Err      error                                               // 错误
}

func newUnionOpenGoodsPromotionGoodsInfoQueryResult(responce UnionOpenGoodsPromotionGoodsInfoQueryResultResponse, result UnionOpenGoodsPromotionGoodsInfoQueryQueryResult, body []byte, http gorequest.Response, err error) *UnionOpenGoodsPromotionGoodsInfoQueryResult {
	return &UnionOpenGoodsPromotionGoodsInfoQueryResult{Responce: responce, Result: result, Body: body, Http: http, Err: err}
}

// UnionOpenGoodsPromotionGoodsInfoQuery 根据skuid查询商品信息接口
// https://union.jd.com/openplatform/api/v2?apiName=jd.union.open.goods.promotiongoodsinfo.query
func (c *Client) UnionOpenGoodsPromotionGoodsInfoQuery(ctx context.Context, notMustParams ...*gorequest.Params) *UnionOpenGoodsPromotionGoodsInfoQueryResult {
	// 参数
	params := NewParamsWithType("jd.union.open.goods.promotiongoodsinfo.query", notMustParams...)
	// 请求
	request, err := c.request(ctx, params)
	// 定义
	var responce UnionOpenGoodsPromotionGoodsInfoQueryResultResponse
	var result UnionOpenGoodsPromotionGoodsInfoQueryQueryResult
	err = gojson.Unmarshal(request.ResponseBody, &responce)
	err = gojson.Unmarshal([]byte(responce.JdUnionOpenGoodsPromotiongoodsinfoQueryResponce.QueryResult), &result)
	return newUnionOpenGoodsPromotionGoodsInfoQueryResult(responce, result, request.ResponseBody, request, err)
}
