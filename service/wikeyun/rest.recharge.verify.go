package wikeyun

import (
	"context"
	"encoding/json"
	"github.com/dtapps/go-library/utils/gorequest"
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
	Err    error                      // 错误
}

func newRestRechargeVerifyResult(result RestRechargeVerifyResponse, body []byte, http gorequest.Response, err error) *RestRechargeVerifyResult {
	return &RestRechargeVerifyResult{Result: result, Body: body, Http: http, Err: err}
}

// RestRechargeVerify 话费充值验证
// https://open.wikeyun.cn/#/apiDocument/9/document/405
func (c *Client) RestRechargeVerify(ctx context.Context, mobile string, amount int64, rechargeType int) *RestRechargeVerifyResult {
	// 参数
	param := gorequest.NewParams()
	param.Set("mobile", mobile)              // 需要充值的手机号
	param.Set("amount", amount)              //	需要充值的金额
	param.Set("recharge_type", rechargeType) // 充值类型
	params := gorequest.NewParamsWith(param)
	// 请求
	request, err := c.request(ctx, apiUrl+"/rest/Recharge/verify", params)
	// 定义
	var response RestRechargeVerifyResponse
	err = json.Unmarshal(request.ResponseBody, &response)
	return newRestRechargeVerifyResult(response, request.ResponseBody, request, err)
}
