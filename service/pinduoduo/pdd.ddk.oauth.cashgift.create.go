package pinduoduo

import (
	"context"
	"github.com/dtapps/go-library/utils/gojson"
	"github.com/dtapps/go-library/utils/gorequest"
)

type PddDdkOauthCashGiftCreateResponse struct {
	CreateCashgiftResponse struct {
		CashGiftId float64 `json:"cash_gift_id"` // 礼金ID
		Success    bool    `json:"success"`      // 创建结果
	} `json:"create_cashgift_response"`
}

type PddDdkOauthCashGiftCreateResult struct {
	Result PddDdkOauthCashGiftCreateResponse // 结果
	Body   []byte                            // 内容
	Http   gorequest.Response                // 请求
	Err    error                             // 错误
}

func newPddDdkOauthCashGiftCreateResult(result PddDdkOauthCashGiftCreateResponse, body []byte, http gorequest.Response, err error) *PddDdkOauthCashGiftCreateResult {
	return &PddDdkOauthCashGiftCreateResult{Result: result, Body: body, Http: http, Err: err}
}

// Create 创建多多礼金
// https://jinbao.pinduoduo.com/third-party/api-detail?apiName=pdd.ddk.oauth.cashgift.create
func (c *PddDdkOauthCashGiftApi) Create(ctx context.Context, notMustParams ...*gorequest.Params) *PddDdkOauthCashGiftCreateResult {
	// 参数
	params := NewParamsWithType("pdd.ddk.oauth.cashgift.create", notMustParams...)
	// 请求
	request, err := c.client.request(ctx, params)
	// 定义
	var response PddDdkOauthCashGiftCreateResponse
	err = gojson.Unmarshal(request.ResponseBody, &response)
	return newPddDdkOauthCashGiftCreateResult(response, request.ResponseBody, request, err)
}
