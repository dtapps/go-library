package kashangwl

import (
	"encoding/json"
	"go.dtapp.net/library/utils/gorequest"
)

type ApiProductRechargeParamsResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Data    struct {
		RechargeAccountLabel string `json:"recharge_account_label"`
		RechargeParams       []struct {
			Name    string `json:"name"`
			Type    string `json:"type"`
			Options string `json:"options"`
		} `json:"recharge_params"`
	} `json:"data"`
}

type ApiProductRechargeParamsResult struct {
	Result ApiProductRechargeParamsResponse // 结果
	Body   []byte                           // 内容
	Http   gorequest.Response               // 请求
	Err    error                            // 错误
}

func NewApiProductRechargeParamsResult(result ApiProductRechargeParamsResponse, body []byte, http gorequest.Response, err error) *ApiProductRechargeParamsResult {
	return &ApiProductRechargeParamsResult{Result: result, Body: body, Http: http, Err: err}
}

// ApiProductRechargeParams 接口说明
// 获取商品的充值参数（仅支持充值类商品）
// http://doc.cqmeihu.cn/sales/ProductParams.html
func (app App) ApiProductRechargeParams(notMustParams ...Params) *ApiProductRechargeParamsResult {
	// 参数
	params := app.NewParamsWith(notMustParams...)
	// 请求
	request, err := app.request("http://www.kashangwl.com/api/product/recharge-params", params)
	// 定义
	var response ApiProductRechargeParamsResponse
	err = json.Unmarshal(request.ResponseBody, &response)
	return NewApiProductRechargeParamsResult(response, request.ResponseBody, request, err)
}