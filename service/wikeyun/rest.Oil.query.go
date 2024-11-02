package wikeyun

import (
	"context"
	"go.dtapp.net/library/utils/gorequest"
)

type RestOilQueryResponse struct {
	Code string `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		OrderNumber string  `json:"order_number"`
		OrderNo     string  `json:"order_no"`
		CardId      int64   `json:"card_id"`
		Amount      int64   `json:"amount"`
		CostPrice   float64 `json:"cost_price"`
		Fanli       float64 `json:"fanli"`
		Status      int64   `json:"status"`
	} `json:"data"`
}

type RestOilQueryResult struct {
	Result RestOilQueryResponse // 结果
	Body   []byte               // 内容
	Http   gorequest.Response   // 请求
}

func newRestOilQueryResult(result RestOilQueryResponse, body []byte, http gorequest.Response) *RestOilQueryResult {
	return &RestOilQueryResult{Result: result, Body: body, Http: http}
}

// RestOilQuery 油卡订单查询
// order_number = 平台单号，与商户单号二选一
// order_no = 商户单号
// https://open.wikeyun.cn/#/apiDocument/9/document/368
func (c *Client) RestOilQuery(ctx context.Context, notMustParams ...gorequest.Params) (*RestOilQueryResult, error) {

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	params.Set("store_id", c.GetStoreId()) // 店铺ID

	// 请求
	var response RestOilQueryResponse
	request, err := c.request(ctx, "rest/Oil/query", params, &response)
	return newRestOilQueryResult(response, request.ResponseBody, request), err
}
