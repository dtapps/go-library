package kashangwl

import (
	"context"
	"go.dtapp.net/library/utils/gorequest"
)

type ApiProductRechargeParamsResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Data    struct {
		RechargeAccountLabel string `json:"recharge_account_label"` // 充值账号类型名称
		RechargeParams       []struct {
			Name    string `json:"name"`
			Type    string `json:"type"`
			Options string `json:"options"`
		} `json:"recharge_params"` // 	充值参数
	} `json:"data"`
}

type ApiProductRechargeParamsResult struct {
	Result ApiProductRechargeParamsResponse // 结果
	Body   []byte                           // 内容
	Http   gorequest.Response               // 请求
}

func newApiProductRechargeParamsResult(result ApiProductRechargeParamsResponse, body []byte, http gorequest.Response) *ApiProductRechargeParamsResult {
	return &ApiProductRechargeParamsResult{Result: result, Body: body, Http: http}
}

// ApiProductRechargeParams 获取商品的充值参数（仅支持充值类商品）
// product_id = 商品编号
// http://doc.cqmeihu.cn/sales/recharge-params.html
func (c *Client) ApiProductRechargeParams(ctx context.Context, productID int64, notMustParams ...gorequest.Params) (*ApiProductRechargeParamsResult, error) {

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	params.Set("product_id", productID) // 商品编号

	// 请求
	var response ApiProductRechargeParamsResponse
	request, err := c.request(ctx, "api/product/recharge-params", params, &response)
	return newApiProductRechargeParamsResult(response, request.ResponseBody, request), err
}
