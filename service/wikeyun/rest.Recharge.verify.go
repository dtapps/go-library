package wikeyun

import (
	"context"
	"go.dtapp.net/library/utils/gorequest"
)

type RestRechargeVerifyResponse struct {
	Code string `json:"code"` // 0000代表正常，其他代表不可下单。
	Msg  string `json:"msg"`
	Time string `json:"time"`
}

type RestRechargeVerifyResult struct {
	Result RestRechargeVerifyResponse // 结果
	Body   []byte                     // 内容
	Http   gorequest.Response         // 请求
}

func newRestRechargeVerifyResult(result RestRechargeVerifyResponse, body []byte, http gorequest.Response) *RestRechargeVerifyResult {
	return &RestRechargeVerifyResult{Result: result, Body: body, Http: http}
}

// RestRechargeVerify 话费充值验证
// mobile = 需要充值的手机号
// amount = 需要充值的金额
// recharge_type = 充值类型
// https://open.wikeyun.cn/#/apiDocument/9/document/405
func (c *Client) RestRechargeVerify(ctx context.Context, mobile string, amount int64, rechargeType int64, notMustParams ...*gorequest.Params) (*RestRechargeVerifyResult, error) {

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	params.Set("mobile", mobile)              // 需要充值的手机号
	params.Set("amount", amount)              //	需要充值的金额
	params.Set("recharge_type", rechargeType) // 充值类型

	// 请求
	var response RestRechargeVerifyResponse
	request, err := c.request(ctx, "rest/Recharge/verify", params, &response)
	return newRestRechargeVerifyResult(response, request.ResponseBody, request), err
}
