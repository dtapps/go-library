package kashangwl

import "encoding/json"

type ProductRechargeParamsResponse struct {
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

type ProductRechargeParamsResult struct {
	Result ProductRechargeParamsResponse // 结果
	Body   []byte                        // 内容
	Err    error                         // 错误
}

func NewProductRechargeParamsResult(result ProductRechargeParamsResponse, body []byte, err error) *ProductRechargeParamsResult {
	return &ProductRechargeParamsResult{Result: result, Body: body, Err: err}
}

// ProductRechargeParams 接口说明
// 获取商品的充值参数（仅支持充值类商品）
// http://doc.cqmeihu.cn/sales/ProductParams.html
func (app App) ProductRechargeParams(notMustParams ...Params) *ProductRechargeParamsResult {
	// 参数
	params := app.NewParamsWith(notMustParams...)
	// 请求
	body, err := app.request("http://www.kashangwl.com/api/product/recharge-params", params)
	// 定义
	var response ProductRechargeParamsResponse
	err = json.Unmarshal(body, &response)
	return NewProductRechargeParamsResult(response, body, err)
}
