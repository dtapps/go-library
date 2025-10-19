package aswzk

import (
	"context"
	"net/http"

	"go.dtapp.net/library/utils/gorequest"
)

type PackageListResponse struct {
	Code int    `json:"code"`
	Info string `json:"info"`
	Data []struct {
		RechargeMoney        float64 `json:"recharge_money"`                   // 充值金额
		RechargeType         string  `json:"recharge_type"`                    // 充值类型
		RechargeOperatorType string  `json:"recharge_operator_type,omitempty"` // 充值运营商类型
	} `json:"data,omitempty"`
	Time    int    `json:"time"`
	TraceId string `json:"trace_id"`
}

// PackageList 套餐列表
// package_type = 套餐类型 phone_bill=话费 electricity=电费)
func (c *Client) PackageList(ctx context.Context, notMustParams ...*gorequest.Params) (response PackageListResponse, err error) {

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)

	// 请求
	err = c.request(ctx, "package/list", params, http.MethodGet, &response)
	return
}
