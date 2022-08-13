package jd

import (
	"context"
	"encoding/json"
	"github.com/dtapps/go-library/utils/gorequest"
)

type UnionOpenCategoryGoodsGetResultResponse struct {
	JdUnionOpenCategoryGoodsGetResponce struct {
		Code      string `json:"code"`
		GetResult string `json:"getResult"`
	} `json:"jd_union_open_category_goods_get_responce"`
}

type UnionOpenCategoryGoodsGetQueryResult struct {
	Code int `json:"code"`
	Data []struct {
		Grade    int    `json:"grade"`    // 类目级别(类目级别 0，1，2 代表一、二、三级类目)
		Name     string `json:"name"`     // 类目名称
		Id       int    `json:"id"`       // 类目Id
		ParentId int    `json:"parentId"` // 父类目Id
	} `json:"data"`
	Message   string `json:"message"`
	RequestId string `json:"requestId"`
}

type UnionOpenCategoryGoodsGetResult struct {
	Responce UnionOpenCategoryGoodsGetResultResponse // 结果
	Result   UnionOpenCategoryGoodsGetQueryResult    // 结果
	Body     []byte                                  // 内容
	Http     gorequest.Response                      // 请求
	Err      error                                   // 错误
}

func newUnionOpenCategoryGoodsGetResult(responce UnionOpenCategoryGoodsGetResultResponse, result UnionOpenCategoryGoodsGetQueryResult, body []byte, http gorequest.Response, err error) *UnionOpenCategoryGoodsGetResult {
	return &UnionOpenCategoryGoodsGetResult{Responce: responce, Result: result, Body: body, Http: http, Err: err}
}

// UnionOpenCategoryGoodsGet 商品类目查询接口
// https://union.jd.com/openplatform/api/v2?apiName=jd.union.open.category.goods.get
func (c *Client) UnionOpenCategoryGoodsGet(ctx context.Context, notMustParams ...Params) *UnionOpenCategoryGoodsGetResult {
	// 参数
	params := NewParamsWithType("jd.union.open.category.goods.get", notMustParams...)
	// 请求
	request, err := c.request(ctx, params)
	// 定义
	var responce UnionOpenCategoryGoodsGetResultResponse
	var result UnionOpenCategoryGoodsGetQueryResult
	err = json.Unmarshal(request.ResponseBody, &responce)
	err = json.Unmarshal([]byte(responce.JdUnionOpenCategoryGoodsGetResponce.GetResult), &result)
	return newUnionOpenCategoryGoodsGetResult(responce, result, request.ResponseBody, request, err)
}
