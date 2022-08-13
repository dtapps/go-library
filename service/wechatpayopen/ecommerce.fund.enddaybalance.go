package wechatpayopen

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/dtapps/go-library/utils/gorequest"
	"net/http"
)

type EcommerceFundEndDayBalanceResponse struct {
	SubMchid        string `json:"sub_mchid"`        // 二级商户号
	AvailableAmount int64  `json:"available_amount"` // 可用余额
	PendingAmount   int64  `json:"pending_amount"`   // 不可用余额
}

type EcommerceFundEndDayBalanceResult struct {
	Result EcommerceFundEndDayBalanceResponse // 结果
	Body   []byte                             // 内容
	Http   gorequest.Response                 // 请求
	Err    error                              // 错误
}

func newEcommerceFundEndDayBalanceResult(result EcommerceFundEndDayBalanceResponse, body []byte, http gorequest.Response, err error) *EcommerceFundEndDayBalanceResult {
	return &EcommerceFundEndDayBalanceResult{Result: result, Body: body, Http: http, Err: err}
}

// EcommerceFundEndDayBalance 查询二级商户账户日终余额API
// date 日期 示例值：2019-08-17
// https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter7_7_2.shtml
func (c *Client) EcommerceFundEndDayBalance(ctx context.Context, date string) *EcommerceFundEndDayBalanceResult {
	// 参数
	params := gorequest.NewParams()
	// 请求
	request, err := c.request(ctx, fmt.Sprintf(apiUrl+"/v3/ecommerce/fund/enddaybalance/%s?date=%s", c.config.SubMchId, date), params, http.MethodGet)
	if err != nil {
		return newEcommerceFundEndDayBalanceResult(EcommerceFundEndDayBalanceResponse{}, request.ResponseBody, request, err)
	}
	// 定义
	var response EcommerceFundEndDayBalanceResponse
	err = json.Unmarshal(request.ResponseBody, &response)
	return newEcommerceFundEndDayBalanceResult(response, request.ResponseBody, request, err)
}
