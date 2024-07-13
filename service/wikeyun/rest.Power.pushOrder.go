package wikeyun

import (
	"context"
	"go.dtapp.net/library/utils/gorequest"
)

type RestPowerPushOrderResponse struct {
	Code string `json:"code"`
	Msg  string `json:"msg"`
	Time string `json:"time"`
	Data struct {
		OrderNumber string `json:"order_number"`
	} `json:"data"`
}

type RestPowerPushOrderResult struct {
	Result RestPowerPushOrderResponse // 结果
	Body   []byte                     // 内容
	Http   gorequest.Response         // 请求
}

func newRestPowerPushOrderResult(result RestPowerPushOrderResponse, body []byte, http gorequest.Response) *RestPowerPushOrderResult {
	return &RestPowerPushOrderResult{Result: result, Body: body, Http: http}
}

// RestPowerPushOrder 电费充值API
// cardId = 充值卡ID，通过创建充值卡接口获取
// order_no = 第三方单号
// amount = 充值金额，支持100,200,300,400,500,600,800,1000
// recharge_type = 类型 1快充 0慢充
// notify_url = 回调通知地址，用于订单状态通知
// change = 是否开启更换渠道补单，1开启0关闭
// https://open.wikeyun.cn/#/apiDocument/9/document/311
func (c *Client) RestPowerPushOrder(ctx context.Context, cardID int64, orderNo string, amount int64, rechargeType int64, notMustParams ...gorequest.Params) (*RestPowerPushOrderResult, error) {

	// OpenTelemetry链路追踪
	ctx = c.TraceStartSpan(ctx, "rest/Power/pushOrder")
	defer c.TraceEndSpan()

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	params.Set("cardId", cardID)              // 充值卡ID，通过创建充值卡接口获取
	params.Set("store_id", c.GetStoreId())    // 店铺ID
	params.Set("order_no", orderNo)           // 第三方单号
	params.Set("amount", amount)              // 充值金额，支持100,200,300,400,500,600,800,1000
	params.Set("recharge_type", rechargeType) //  类型 1快充 0慢充

	// 请求
	var response RestPowerPushOrderResponse
	request, err := c.request(ctx, "rest/Power/pushOrder", params, &response)
	return newRestPowerPushOrderResult(response, request.ResponseBody, request), err
}
