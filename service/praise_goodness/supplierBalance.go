package praise_goodness

import (
	"context"
	"github.com/dtapps/go-library/utils/gojson"
	"github.com/dtapps/go-library/utils/gorequest"
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
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	// 请求
	request, err := c.request(ctx, "api/order/supplierBalance", params, http.MethodPost)
	if err != nil {
		return newSupplierBalanceResult(SupplierBalanceResponse{}, request.ResponseBody, request), err
	}
	// 定义
	var response SupplierBalanceResponse
	err = gojson.Unmarshal(request.ResponseBody, &response)
	return newSupplierBalanceResult(response, request.ResponseBody, request), err
}
