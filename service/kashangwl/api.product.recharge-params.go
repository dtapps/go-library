package kashangwl

import (
	"context"
	"github.com/dtapps/go-library/utils/gojson"
	"github.com/dtapps/go-library/utils/gorequest"
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

// ApiProductRechargeParams 接口说明
// 获取商品的充值参数（仅支持充值类商品）
// http://doc.cqmeihu.cn/sales/ProductParams.html
func (c *Client) ApiProductRechargeParams(ctx context.Context, notMustParams ...gorequest.Params) (*ApiProductRechargeParamsResult, error) {
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	// 请求
	request, err := c.request(ctx, apiUrl+"/api/product/recharge-params", params)
	if err != nil {
		return newApiProductRechargeParamsResult(ApiProductRechargeParamsResponse{}, request.ResponseBody, request), err
	}
	// 定义
	var response ApiProductRechargeParamsResponse
	err = gojson.Unmarshal(request.ResponseBody, &response)
	return newApiProductRechargeParamsResult(response, request.ResponseBody, request), err
}
