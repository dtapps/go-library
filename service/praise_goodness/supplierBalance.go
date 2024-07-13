package praise_goodness

import (
	"context"
	"go.dtapp.net/library/utils/gorequest"
	"net/http"
)

type SupplierBalanceResponse struct {
	Code int    `json:"code"` // 1：请求成功 -1：请求失败
	Msg  string `json:"msg"`  // 返回说明
	Time string `json:"time"` // 时间戳
	Data struct {
		Balance string `json:"balance"` // 用户余额
	} `json:"data"`
}

type SupplierBalanceResult struct {
	Result SupplierBalanceResponse // 结果
	Body   []byte                  // 内容
	Http   gorequest.Response      // 请求
}

func newSupplierBalanceResult(result SupplierBalanceResponse, body []byte, http gorequest.Response) *SupplierBalanceResult {
	return &SupplierBalanceResult{Result: result, Body: body, Http: http}
}

// SupplierBalance 用户余额查询接口
func (c *Client) SupplierBalance(ctx context.Context, notMustParams ...gorequest.Params) (*SupplierBalanceResult, error) {

	// OpenTelemetry链路追踪
	ctx = c.TraceStartSpan(ctx, "api/order/supplierBalance")
	defer c.TraceEndSpan()

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	params.Set("mch_id", c.GetMchID()) // 商户编号 (平台提供)

	// 响应
	var response SupplierBalanceResponse

	// 请求
	request, err := c.request(ctx, "api/order/supplierBalance", params, http.MethodPost, &response)
	return newSupplierBalanceResult(response, request.ResponseBody, request), err
}
