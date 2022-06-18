package jd

import (
	"encoding/json"
	"go.dtapp.net/library/utils/gorequest"
)

type UnionOpenOrderRowQueryResultResponse struct {
	JdUnionOpenOrderRowQueryResponce struct {
		Code        string `json:"code"`
		QueryResult string `json:"queryResult"`
	} `json:"jd_union_open_order_row_query_responce"`
}

type UnionOpenOrderRowQueryQueryResult struct {
	Code int `json:"code"`
	Data []struct {
		ActualCosPrice float64 `json:"actualCosPrice"`
		ActualFee      float64 `json:"actualFee"`
		BalanceExt     string  `json:"balanceExt"`
		CategoryInfo   struct {
			Cid1 int `json:"cid1"`
			Cid2 int `json:"cid2"`
			Cid3 int `json:"cid3"`
		} `json:"categoryInfo"`
		ChannelId           int     `json:"channelId"`
		Cid1                int     `json:"cid1"`
		Cid2                int     `json:"cid2"`
		Cid3                int     `json:"cid3"`
		CommissionRate      float64 `json:"commissionRate"`
		CpActId             int     `json:"cpActId"`
		EstimateCosPrice    float64 `json:"estimateCosPrice"`
		EstimateFee         float64 `json:"estimateFee"`
		ExpressStatus       int     `json:"expressStatus"`
		Ext1                string  `json:"ext1"`
		FinalRate           float64 `json:"finalRate"`
		FinishTime          string  `json:"finishTime"`
		GiftCouponKey       string  `json:"giftCouponKey"`
		GiftCouponOcsAmount float64 `json:"giftCouponOcsAmount"`
		GoodsInfo           struct {
			MainSkuId int `json:"mainSkuId"`
			ProductId int `json:"productId"`
			ShopId    int `json:"shopId"`
		} `json:"goodsInfo"`
		Id             string  `json:"id"`
		ModifyTime     string  `json:"modifyTime"`
		OrderEmt       int     `json:"orderEmt"`
		OrderId        int64   `json:"orderId"`
		OrderTime      string  `json:"orderTime"`
		ParentId       int     `json:"parentId"`
		PayMonth       int     `json:"payMonth"`
		Pid            string  `json:"pid"`
		Plus           int     `json:"plus"`
		PopId          int     `json:"popId"`
		PositionId     int     `json:"positionId"`
		Price          float64 `json:"price"`
		ProPriceAmount float64 `json:"proPriceAmount"`
		Rid            int     `json:"rid"`
		SiteId         int     `json:"siteId"`
		SkuFrozenNum   int     `json:"skuFrozenNum"`
		SkuId          int64   `json:"skuId"`
		SkuName        string  `json:"skuName"`
		SkuNum         int     `json:"skuNum"`
		SkuReturnNum   int     `json:"skuReturnNum"`
		SubSideRate    float64 `json:"subSideRate"`
		SubUnionId     string  `json:"subUnionId"`
		SubsidyRate    float64 `json:"subsidyRate"`
		TraceType      int     `json:"traceType"`
		UnionAlias     string  `json:"unionAlias"`
		UnionId        int     `json:"unionId"`
		UnionRole      int     `json:"unionRole"`
		UnionTag       string  `json:"unionTag"`
		ValidCode      int     `json:"validCode"`
	} `json:"data"`
	HasMore   bool   `json:"hasMore"`
	Message   string `json:"message"`
	RequestId string `json:"requestId"`
}

type UnionOpenOrderRowQueryResult struct {
	Responce UnionOpenOrderRowQueryResultResponse // 结果
	Result   UnionOpenOrderRowQueryQueryResult    // 结果
	Body     []byte                               // 内容
	Http     gorequest.Response                   // 请求
	Err      error                                // 错误
}

func NewUnionOpenOrderRowQueryResult(responce UnionOpenOrderRowQueryResultResponse, result UnionOpenOrderRowQueryQueryResult, body []byte, http gorequest.Response, err error) *UnionOpenOrderRowQueryResult {
	return &UnionOpenOrderRowQueryResult{Responce: responce, Result: result, Body: body, Http: http, Err: err}
}

// UnionOpenOrderRowQuery 订单行查询接口
// https://union.jd.com/openplatform/api/v2?apiName=jd.union.open.order.row.query
func (app *App) UnionOpenOrderRowQuery(notMustParams ...Params) *UnionOpenOrderRowQueryResult {
	// 参数
	params := NewParamsWithType("jd.union.open.order.row.query", notMustParams...)
	// 请求
	request, err := app.request(params)
	// 定义
	var responce UnionOpenOrderRowQueryResultResponse
	var result UnionOpenOrderRowQueryQueryResult
	err = json.Unmarshal(request.ResponseBody, &responce)
	err = json.Unmarshal([]byte(responce.JdUnionOpenOrderRowQueryResponce.QueryResult), &result)
	return NewUnionOpenOrderRowQueryResult(responce, result, request.ResponseBody, request, err)
}
