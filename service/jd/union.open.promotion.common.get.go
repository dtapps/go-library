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
	body     []byte                                    // 内容
	Err      error                                     // 错误
}

func NewUnionOpenPromotionCommonGetResult(responce UnionOpenPromotionCommonGetResultResponse, result UnionOpenPromotionCommonGetGetResult, body []byte, err error) *UnionOpenPromotionCommonGetResult {
	return &UnionOpenPromotionCommonGetResult{Responce: responce, Result: result, body: body, Err: err}
}

// UnionOpenPromotionCommonGet
// 网站/APP来获取的推广链接，功能同宙斯接口的自定义链接转换、 APP领取代码接口通过商品链接、活动链接获取普通推广链接，支持传入subunionid参数，可用于区分媒体自身的用户ID，该参数可在订单查询接口返回，需向cps-qxsq@jd.com申请权限。
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
