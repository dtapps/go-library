package jd

import (
	"context"
	"github.com/dtapps/go-library/utils/gojson"
	"github.com/dtapps/go-library/utils/gorequest"
)

type UnionOpenGoodsBigfieldQueryResultResponse struct {
	JdUnionOpenGoodsBigfieldQueryResponce struct {
		Code        string `json:"code"`
		QueryResult string `json:"queryResult"`
	} `json:"jd_union_open_goods_bigfield_query_responce"`
}

type UnionOpenGoodsBigfieldQueryQueryResult struct {
	Code int `json:"code"`
	Data []struct {
		BaseBigFieldInfo struct {
			WareQD string `json:"wareQD"`
			Wdis   string `json:"wdis"`
		} `json:"baseBigFieldInfo"`
		CategoryInfo struct {
			Cid1     int    `json:"cid1"`
			Cid1Name string `json:"cid1Name"`
			Cid2     int    `json:"cid2"`
			Cid2Name string `json:"cid2Name"`
			Cid3     int    `json:"cid3"`
			Cid3Name string `json:"cid3Name"`
		} `json:"categoryInfo"`
		DetailImages string `json:"detailImages"`
		ImageInfo    struct {
			ImageList []struct {
				Url string `json:"url"`
			} `json:"imageList"`
		} `json:"imageInfo"`
		MainSkuId int64  `json:"mainSkuId"`
		Owner     string `json:"owner"`
		ProductId int64  `json:"productId"`
		SkuId     int64  `json:"skuId"`
		SkuName   string `json:"skuName"`
		SkuStatus int    `json:"skuStatus"`
	} `json:"data"`
	Message   string `json:"message"`
	RequestId string `json:"requestId"`
}

type UnionOpenGoodsBigfieldQueryResult struct {
	Responce UnionOpenGoodsBigfieldQueryResultResponse // 结果
	Result   UnionOpenGoodsBigfieldQueryQueryResult    // 结果
	Body     []byte                                    // 内容
	Http     gorequest.Response                        // 请求
	Err      error                                     // 错误
}

func newUnionOpenGoodsBigfieldQueryResult(responce UnionOpenGoodsBigfieldQueryResultResponse, result UnionOpenGoodsBigfieldQueryQueryResult, body []byte, http gorequest.Response, err error) *UnionOpenGoodsBigfieldQueryResult {
	return &UnionOpenGoodsBigfieldQueryResult{Responce: responce, Result: result, Body: body, Http: http, Err: err}
}

// UnionOpenGoodsBigfieldQuery 商品详情查询接口
// https://union.jd.com/openplatform/api/v2?apiName=jd.union.open.goods.bigfield.query
func (c *Client) UnionOpenGoodsBigfieldQuery(ctx context.Context, notMustParams ...*gorequest.Params) *UnionOpenGoodsBigfieldQueryResult {
	// 参数
	params := NewParamsWithType("jd.union.open.goods.bigfield.query", notMustParams...)
	// 请求
	request, err := c.request(ctx, params)
	// 定义
	var responce UnionOpenGoodsBigfieldQueryResultResponse
	var result UnionOpenGoodsBigfieldQueryQueryResult
	err = gojson.Unmarshal(request.ResponseBody, &responce)
	err = gojson.Unmarshal([]byte(responce.JdUnionOpenGoodsBigfieldQueryResponce.QueryResult), &result)
	return newUnionOpenGoodsBigfieldQueryResult(responce, result, request.ResponseBody, request, err)
}
