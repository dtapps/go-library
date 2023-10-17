package praise_goodness

import (
	"context"
	"github.com/dtapps/go-library/utils/gojson"
	"github.com/dtapps/go-library/utils/gorequest"
	"net/http"
)

type QuerySupplierOrderInfoResponse struct {
	Code int    `json:"code"` // 1：请求成功 -1：请求失败
	Msg  string `json:"msg"`  // 返回说明
	Time string `json:"time"` // 时间戳
	Data struct {
		OrderID    string `json:"order_id"`   // 我方订单号
		TradeID    string `json:"trade_id"`   // 商户订单号
		Status     int    `json:"status"`     // 88：充值成功 22：充值失败 66：充值中 55：订单不存在
		Voucher    string `json:"voucher"`    // 充值凭证 (有可能为空值，不一定返回)
		Createtime int    `json:"createtime"` // 订单创建时间
		Proof      string `json:"proof"`      // 请求token
	} `json:"data"`
}

type QuerySupplierOrderInfoResult struct {
	Result QuerySupplierOrderInfoResponse // 结果
	Body   []byte                         // 内容
	Http   gorequest.Response             // 请求
}

func newQuerySupplierOrderInfoResult(result QuerySupplierOrderInfoResponse, body []byte, http gorequest.Response) *QuerySupplierOrderInfoResult {
	return &QuerySupplierOrderInfoResult{Result: result, Body: body, Http: http}
}

// QuerySupplierOrderInfo 订单查询接口
// trade_id = 商户订单号
func (c *Client) QuerySupplierOrderInfo(ctx context.Context, tradeID string, notMustParams ...gorequest.Params) (*QuerySupplierOrderInfoResult, error) {
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	params.Set("mch_id", c.GetMchID()) // 商户编号 (平台提供)
	params.Set("trade_id", tradeID)    // 商户订单号
	// 请求
	request, err := c.request(ctx, "api/order/querySupplierOrderInfo", params, http.MethodPost)
	if err != nil {
		return newQuerySupplierOrderInfoResult(QuerySupplierOrderInfoResponse{}, request.ResponseBody, request), err
	}
	// 定义
	var response QuerySupplierOrderInfoResponse
	err = gojson.Unmarshal(request.ResponseBody, &response)
	return newQuerySupplierOrderInfoResult(response, request.ResponseBody, request), err
}

func (QuerySupplierOrderInfoResponse) GetStatusDesc(status int) string {
	if status == 88 {
		return "充值成功"
	} else if status == 22 {
		return "充值失败"
	} else if status == 66 {
		return "充值中"
	} else if status == 55 {
		return "订单不存在"
	} else {
		return ""
	}
}
