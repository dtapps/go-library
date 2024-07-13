package aswzk

import (
	"context"
	"go.dtapp.net/library/utils/gorequest"
	"net/http"
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

type PackageListResult struct {
	Result PackageListResponse // 结果
	Body   []byte              // 内容
	Http   gorequest.Response  // 请求
}

func newPackageListResult(result PackageListResponse, body []byte, http gorequest.Response) *PackageListResult {
	return &PackageListResult{Result: result, Body: body, Http: http}
}

// PackageList 套餐列表
// package_type = 套餐类型 phone_bill=话费 electricity=电费)
func (c *Client) PackageList(ctx context.Context, notMustParams ...gorequest.Params) (*PackageListResult, error) {

	// OpenTelemetry链路追踪
	ctx = c.TraceStartSpan(ctx, "package/list")
	defer c.TraceEndSpan()

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)

	// 请求
	var response PackageListResponse
	request, err := c.request(ctx, "package/list", params, http.MethodGet, &response)
	return newPackageListResult(response, request.ResponseBody, request), err
}
