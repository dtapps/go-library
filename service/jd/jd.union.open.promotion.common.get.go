package jd

import "encoding/json"

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
	Err      error                                     // 错误
}

func NewUnionOpenPromotionCommonGetResult(responce UnionOpenPromotionCommonGetResultResponse, result UnionOpenPromotionCommonGetGetResult, body []byte, err error) *UnionOpenPromotionCommonGetResult {
	return &UnionOpenPromotionCommonGetResult{Responce: responce, Result: result, Body: body, Err: err}
}

// UnionOpenPromotionCommonGet 网站/APP获取推广链接接口
// https://union.jd.com/openplatform/api/v2?apiName=jd.union.open.promotion.common.get
func (app *App) UnionOpenPromotionCommonGet(notMustParams ...Params) *UnionOpenPromotionCommonGetResult {
	// 参数
	params := NewParamsWithType("jd.union.open.promotion.common.get", notMustParams...)
	// 请求
	body, err := app.request(params)
	// 定义
	var responce UnionOpenPromotionCommonGetResultResponse
	var result UnionOpenPromotionCommonGetGetResult
	err = json.Unmarshal(body, &responce)
	err = json.Unmarshal([]byte(responce.JdUnionOpenPromotionCommonGetResponce.GetResult), &result)
	return NewUnionOpenPromotionCommonGetResult(responce, result, body, err)
}
