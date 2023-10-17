package wikeyun

import (
	"context"
	"github.com/dtapps/go-library/utils/gojson"
	"github.com/dtapps/go-library/utils/gorequest"
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
// https://open.wikeyun.cn/#/apiDocument/9/document/311
func (c *Client) RestPowerPushOrder(ctx context.Context, cardID int64, orderNo string, amount int64, rechargeType int64, notMustParams ...gorequest.Params) (*RestPowerPushOrderResult, error) {
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	params.Set("cardId", cardID)              // 充值卡ID，通过创建充值卡接口获取
	params.Set("store_id", c.GetStoreId())    // 店铺ID
	params.Set("order_no", orderNo)           // 第三方单号
	params.Set("amount", amount)              // 充值金额，支持100,200,300,400,500,600,800,1000
	params.Set("recharge_type", rechargeType) //  类型 1快充 0慢充
	// 请求
	request, err := c.request(ctx, apiUrl+"/rest/Power/pushOrder", params)
	if err != nil {
		return newRestPowerPushOrderResult(RestPowerPushOrderResponse{}, request.ResponseBody, request), err
	}
	// 定义
	var response RestPowerPushOrderResponse
	err = gojson.Unmarshal(request.ResponseBody, &response)
	return newRestPowerPushOrderResult(response, request.ResponseBody, request), err
}
