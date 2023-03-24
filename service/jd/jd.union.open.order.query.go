package jd

import (
	"context"
	"github.com/dtapps/go-library/utils/gojson"
	"github.com/dtapps/go-library/utils/gorequest"
)

type UnionOpenOrderQueryResultResponse struct {
	JdUnionOpenOrderQueryResponce struct {
		Code        string `json:"code"`
		QueryResult string `json:"queryResult"`
	} `json:"jd_union_open_order_query_responce"`
}

type UnionOpenOrderQueryQueryResult struct {
	Code int `json:"code"`
	Data []struct {
		Ext1       string `json:"ext1"`
		FinishTime int64  `json:"finishTime"`
		OrderEmt   int    `json:"orderEmt"`
		OrderId    int64  `json:"orderId"`
		OrderTime  int64  `json:"orderTime"`
		ParentId   int    `json:"parentId"`
		PayMonth   int    `json:"payMonth"`
		Plus       int    `json:"plus"`
		PopId      int    `json:"popId"`
		SkuList    []struct {
			ActualCosPrice      float64 `json:"actualCosPrice"`
			ActualFee           float64 `json:"actualFee"`
			Cid1                int     `json:"cid1"`
			Cid2                int     `json:"cid2"`
			Cid3                int     `json:"cid3"`
			CommissionRate      float64 `json:"commissionRate"`
			CpActId             int     `json:"cpActId"`
			EstimateCosPrice    float64 `json:"estimateCosPrice"`
			EstimateFee         float64 `json:"estimateFee"`
			Ext1                string  `json:"ext1"`
			FinalRate           float64 `json:"finalRate"`
			FrozenSkuNum        int     `json:"frozenSkuNum"`
			GiftCouponKey       string  `json:"giftCouponKey"`
			GiftCouponOcsAmount float64 `json:"giftCouponOcsAmount"`
			PayMonth            int     `json:"payMonth"`
			Pid                 string  `json:"pid"`
			PopId               int     `json:"popId"`
			PositionId          int     `json:"positionId"`
			Price               float64 `json:"price"`
			ProPriceAmount      float64 `json:"proPriceAmount"`
			SiteId              int     `json:"siteId"`
			SkuId               int64   `json:"skuId"`
			SkuName             string  `json:"skuName"`
			SkuNum              int     `json:"skuNum"`
			SkuReturnNum        int     `json:"skuReturnNum"`
			SubSideRate         float64 `json:"subSideRate"`
			SubUnionId          string  `json:"subUnionId"`
			SubsidyRate         float64 `json:"subsidyRate"`
			TraceType           int     `json:"traceType"`
			UnionAlias          string  `json:"unionAlias"`
			UnionRole           int     `json:"unionRole"`
			UnionTag            string  `json:"unionTag"`
			UnionTrafficGroup   int     `json:"unionTrafficGroup"`
			ValidCode           int     `json:"validCode"`
		} `json:"skuList"`
		UnionId   int `json:"unionId"`
		ValidCode int `json:"validCode"`
	} `json:"data"`
	HasMore   bool   `json:"hasMore"`
	Message   string `json:"message"`
	RequestId string `json:"requestId"`
}

type UnionOpenOrderQueryResult struct {
	Responce UnionOpenOrderQueryResultResponse // 结果
	Result   UnionOpenOrderQueryQueryResult    // 结果
	Body     []byte                            // 内容
	Http     gorequest.Response                // 请求
	Err      error                             // 错误
}

func newUnionOpenOrderQueryResult(responce UnionOpenOrderQueryResultResponse, result UnionOpenOrderQueryQueryResult, body []byte, http gorequest.Response, err error) *UnionOpenOrderQueryResult {
	return &UnionOpenOrderQueryResult{Responce: responce, Result: result, Body: body, Http: http, Err: err}
}

// UnionOpenOrderQuery 订单查询接口
// https://union.jd.com/openplatform/api/v2?apiName=jd.union.open.order.query
func (c *Client) UnionOpenOrderQuery(ctx context.Context, notMustParams ...Params) *UnionOpenOrderQueryResult {
	// 参数
	params := NewParamsWithType("jd.union.open.order.query", notMustParams...)
	// 请求
	request, err := c.request(ctx, params)
	// 定义
	var responce UnionOpenOrderQueryResultResponse
	var result UnionOpenOrderQueryQueryResult
	err = gojson.Unmarshal(request.ResponseBody, &responce)
	err = gojson.Unmarshal([]byte(responce.JdUnionOpenOrderQueryResponce.QueryResult), &result)
	return newUnionOpenOrderQueryResult(responce, result, request.ResponseBody, request, err)
}
