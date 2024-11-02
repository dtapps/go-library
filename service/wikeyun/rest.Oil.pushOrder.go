package wikeyun

import (
	"context"
	"go.dtapp.net/library/utils/gorequest"
)

type RestOilPushOrderResponse struct {
	Code string `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		OrderNumber string  `json:"order_number"`
		Amount      int64   `json:"amount"`
		Fanli       float64 `json:"fanli"`
		CostPrice   float64 `json:"cost_price"`
	} `json:"data"`
}

type RestOilPushOrderResult struct {
	Result RestOilPushOrderResponse // 结果
	Body   []byte                   // 内容
	Http   gorequest.Response       // 请求
}

func newRestOilPushOrderResult(result RestOilPushOrderResponse, body []byte, http gorequest.Response) *RestOilPushOrderResult {
	return &RestOilPushOrderResult{Result: result, Body: body, Http: http}
}

// RestOilPushOrder 油卡充值
// order_no = 商户单号
// amount = 充值金额
// recharge_type = 充值类型 1快充 0慢充
// notify_url = 回调通知地址，用于订单状态通知
// cardId = 卡号ID，通过新增获取
// https://open.wikeyun.cn/#/apiDocument/9/document/367
func (c *Client) RestOilPushOrder(ctx context.Context, notMustParams ...gorequest.Params) (*RestOilPushOrderResult, error) {

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	params.Set("store_id", c.GetStoreId()) // 店铺ID

	// 请求
	var response RestOilPushOrderResponse
	request, err := c.request(ctx, "rest/Oil/pushOrder", params, &response)
	return newRestOilPushOrderResult(response, request.ResponseBody, request), err
}
