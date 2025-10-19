package aswzk

import (
	"context"
	"net/http"

	"go.dtapp.net/library/utils/gorequest"
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

// AssetsBalance 余额查询
func (c *Client) AssetsBalance(ctx context.Context, notMustParams ...*gorequest.Params) (response AssetsBalanceResponse, err error) {

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)

	// 请求
	err = c.request(ctx, "assets/balance", params, http.MethodGet, &response)
	return
}
