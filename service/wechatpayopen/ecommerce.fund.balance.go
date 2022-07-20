package wechatpayopen

import (
	"encoding/json"
	"fmt"
	"github.com/dtapps/go-library/utils/gorequest"
	"net/http"
)

type EcommerceFundBalanceResponse struct {
	SubMchid        string `json:"sub_mchid"`        // 二级商户号
	AccountType     string `json:"account_type"`     // 账户类型
	AvailableAmount int64  `json:"available_amount"` // 可用余额
	PendingAmount   int64  `json:"pending_amount"`   // 不可用余额
}

type EcommerceFundBalanceResult struct {
	Result EcommerceFundBalanceResponse // 结果
	Body   []byte                       // 内容
	Http   gorequest.Response           // 请求
	Err    error                        // 错误
}

func newEcommerceFundBalanceResult(result EcommerceFundBalanceResponse, body []byte, http gorequest.Response, err error) *EcommerceFundBalanceResult {
	return &EcommerceFundBalanceResult{Result: result, Body: body, Http: http, Err: err}
}

// EcommerceFundBalance 查询二级商户账户实时余额API
// accountType 账户类型 BASIC：基本账户 OPERATION：运营账户 FEES：手续费账户
// https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter7_7_1.shtml
func (c *Client) EcommerceFundBalance(accountType string) *EcommerceFundBalanceResult {
	// 参数
	params := gorequest.NewParams()
	// 请求
	request, err := c.request(fmt.Sprintf(apiUrl+"/v3/ecommerce/fund/balance/%s?account_type=%s", c.config.SubMchId, accountType), params, http.MethodGet)
	if err != nil {
		return newEcommerceFundBalanceResult(EcommerceFundBalanceResponse{}, request.ResponseBody, request, err)
	}
	// 定义
	var response EcommerceFundBalanceResponse
	err = json.Unmarshal(request.ResponseBody, &response)
	return newEcommerceFundBalanceResult(response, request.ResponseBody, request, err)
}
