package jd

import (
	"context"
	"github.com/dtapps/go-library/utils/gojson"
	"github.com/dtapps/go-library/utils/gorequest"
)

type UnionOpenPromotionCommonGetResultResponse struct {
	JdUnionOpenPromotionCommonGetResponce struct {
		Code      string `json:"code"`
		GetResult string `json:"getResult"`
	} `json:"jd_union_open_promotion_common_get_responce"`
}

type UnionOpenPromotionCommonGetGetResult struct {
	Code int `json:"code"`
	Data struct {
		ClickURL string `json:"clickURL"`
		JCommand string `json:"jCommand"`
	} `json:"data"`
	Message   string `json:"message"`
	RequestId string `json:"requestId"`
}

type UnionOpenPromotionCommonGetResult struct {
	Responce UnionOpenPromotionCommonGetResultResponse // 结果
	Result   UnionOpenPromotionCommonGetGetResult      // 结果
	Body     []byte                                    // 内容
	Http     gorequest.Response                        // 请求
}

func newUnionOpenPromotionCommonGetResult(responce UnionOpenPromotionCommonGetResultResponse, result UnionOpenPromotionCommonGetGetResult, body []byte, http gorequest.Response) *UnionOpenPromotionCommonGetResult {
	return &UnionOpenPromotionCommonGetResult{Responce: responce, Result: result, Body: body, Http: http}
}

// UnionOpenPromotionCommonGet 网站/APP获取推广链接接口
// https://union.jd.com/openplatform/api/v2?apiName=jd.union.open.promotion.common.get
func (c *Client) UnionOpenPromotionCommonGet(ctx context.Context, notMustParams ...gorequest.Params) (*UnionOpenPromotionCommonGetResult, error) {
	// 参数
	params := NewParamsWithType("jd.union.open.promotion.common.get", notMustParams...)
	// 请求
	request, err := c.request(ctx, params)
	if err != nil {
		return newUnionOpenPromotionCommonGetResult(UnionOpenPromotionCommonGetResultResponse{}, UnionOpenPromotionCommonGetGetResult{}, request.ResponseBody, request), err
	}
	// 定义
	var responce UnionOpenPromotionCommonGetResultResponse
	var result UnionOpenPromotionCommonGetGetResult
	err = gojson.Unmarshal(request.ResponseBody, &responce)
	err = gojson.Unmarshal([]byte(responce.JdUnionOpenPromotionCommonGetResponce.GetResult), &result)
	return newUnionOpenPromotionCommonGetResult(responce, result, request.ResponseBody, request), err
}
