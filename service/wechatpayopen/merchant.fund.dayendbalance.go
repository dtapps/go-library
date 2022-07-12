package wechatpayopen

import (
	"encoding/json"
	"fmt"
	"go.dtapp.net/library/utils/gorequest"
	"net/http"
)

type MerchantFundDayEndBalanceResponse struct {
	AvailableAmount int64 `json:"available_amount"` // 可用余额
	PendingAmount   int64 `json:"pending_amount"`   // 不可用余额
}

type MerchantFundDayEndBalanceResult struct {
	Result MerchantFundDayEndBalanceResponse // 结果
	Body   []byte                            // 内容
	Http   gorequest.Response                // 请求
	Err    error                             // 错误
}

func newMerchantFundDayEndBalanceResult(result MerchantFundDayEndBalanceResponse, body []byte, http gorequest.Response, err error) *MerchantFundDayEndBalanceResult {
	return &MerchantFundDayEndBalanceResult{Result: result, Body: body, Http: http, Err: err}
}

// MerchantFundDayEndBalance 查询电商平台账户日终余额API
// accountType 账户类型 BASIC：基本账户 OPERATION：运营账户 FEES：手续费账户
// date 日期 示例值：2019-08-17
// https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter7_7_4.shtml
func (c *Client) MerchantFundDayEndBalance(accountType, date string) *MerchantFundDayEndBalanceResult {
	// 参数
	params := gorequest.NewParams()
	// 请求
	request, err := c.request(fmt.Sprintf(apiUrl+"/v3/merchant/fund/dayendbalance/%s?date=%s", accountType, date), params, http.MethodGet)
	if err != nil {
		return newMerchantFundDayEndBalanceResult(MerchantFundDayEndBalanceResponse{}, request.ResponseBody, request, err)
	}
	// 定义
	var response MerchantFundDayEndBalanceResponse
	err = json.Unmarshal(request.ResponseBody, &response)
	return newMerchantFundDayEndBalanceResult(response, request.ResponseBody, request, err)
}
