package wechatpayopen

import (
	"encoding/json"
	"fmt"
	"github.com/dtapps/go-library/utils/gorequest"
	"net/http"
)

type MerchantFundBalanceResponse struct {
	AvailableAmount int64 `json:"available_amount"` // 可用余额
	PendingAmount   int64 `json:"pending_amount"`   // 不可用余额
}

type MerchantFundBalanceResult struct {
	Result MerchantFundBalanceResponse // 结果
	Body   []byte                      // 内容
	Http   gorequest.Response          // 请求
	Err    error                       // 错误
}

func newMerchantFundBalanceResult(result MerchantFundBalanceResponse, body []byte, http gorequest.Response, err error) *MerchantFundBalanceResult {
	return &MerchantFundBalanceResult{Result: result, Body: body, Http: http, Err: err}
}

// MerchantFundBalance 查询电商平台账户实时余额API
// accountType 账户类型 BASIC：基本账户 OPERATION：运营账户 FEES：手续费账户
// https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter7_7_3.shtml
func (c *Client) MerchantFundBalance(accountType string) *MerchantFundBalanceResult {
	// 参数
	params := gorequest.NewParams()
	// 请求
	request, err := c.request(fmt.Sprintf(apiUrl+"/v3/merchant/fund/balance/%s", accountType), params, http.MethodGet)
	if err != nil {
		return newMerchantFundBalanceResult(MerchantFundBalanceResponse{}, request.ResponseBody, request, err)
	}
	// 定义
	var response MerchantFundBalanceResponse
	err = json.Unmarshal(request.ResponseBody, &response)
	return newMerchantFundBalanceResult(response, request.ResponseBody, request, err)
}
