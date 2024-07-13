package aswzk

import (
	"context"
	"go.dtapp.net/library/utils/gorequest"
	"net/http"
)

type AssetsBalanceResponse struct {
	Code int    `json:"code"`
	Info string `json:"info"`
	Data struct {
		Balance float64 `json:"balance"` // 余额
	} `json:"data,omitempty"`
	Time    int    `json:"time"`
	TraceId string `json:"trace_id"`
}

type AssetsBalanceResult struct {
	Result AssetsBalanceResponse // 结果
	Body   []byte                // 内容
	Http   gorequest.Response    // 请求
}

func newAssetsBalanceResult(result AssetsBalanceResponse, body []byte, http gorequest.Response) *AssetsBalanceResult {
	return &AssetsBalanceResult{Result: result, Body: body, Http: http}
}

// AssetsBalance 余额查询
func (c *Client) AssetsBalance(ctx context.Context, notMustParams ...gorequest.Params) (*AssetsBalanceResult, error) {

	// OpenTelemetry链路追踪
	ctx = c.TraceStartSpan(ctx, "assets/balance")
	defer c.TraceEndSpan()

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)

	// 请求
	var response AssetsBalanceResponse
	request, err := c.request(ctx, "assets/balance", params, http.MethodGet, &response)
	return newAssetsBalanceResult(response, request.ResponseBody, request), err
}
