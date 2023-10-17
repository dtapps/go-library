package wikeyun

import (
	"context"
	"github.com/dtapps/go-library/utils/gojson"
	"github.com/dtapps/go-library/utils/gorequest"
)

type RestRechargePushOrderResponse struct {
	Code string `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		OrderNumber string `json:"order_number"`
	} `json:"data"`
}

type RestRechargePushOrderResult struct {
	Result RestRechargePushOrderResponse // 结果
	Body   []byte                        // 内容
	Http   gorequest.Response            // 请求
}

func newRestRechargePushOrderResult(result RestRechargePushOrderResponse, body []byte, http gorequest.Response) *RestRechargePushOrderResult {
	return &RestRechargePushOrderResult{Result: result, Body: body, Http: http}
}

// RestRechargePushOrder 话费充值推送
// mobile = 充值手机号,虚拟号,协号转网不支持充值
// order_no = 第三方单号
// money = 充值金额，目前有50，100，200三种，具体联系客服咨询
// recharge_type = 类型 1快充 0慢充
// notify_url = 回调通知地址，用于订单状态通知
// change = 失败更换渠道充值 0 否 1是 不传系统根据设置判断
// source = 是否强制渠道，因为每个渠道价格不同，不同用户提交的业务不同，默认不强制，具体渠道价格联系客服
// https://open.wikeyun.cn/#/apiDocument/9/document/298
func (c *Client) RestRechargePushOrder(ctx context.Context, mobile string, orderNo string, money int64, rechargeType int64, notifyUrl string, notMustParams ...gorequest.Params) (*RestRechargePushOrderResult, error) {
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	params.Set("store_id", c.GetStoreId())    // 店铺ID
	params.Set("mobile", mobile)              // 充值手机号,虚拟号,协号转网不支持充值
	params.Set("order_no", orderNo)           // 第三方单号
	params.Set("money", money)                // 充值金额，目前有50，100，200三种，具体联系客服咨询
	params.Set("recharge_type", rechargeType) // 类型 1快充 0慢充
	params.Set("notify_url", notifyUrl)       // 回调通知地址，用于订单状态通知
	// 请求
	request, err := c.request(ctx, apiUrl+"/rest/Recharge/pushOrder", params)
	if err != nil {
		return newRestRechargePushOrderResult(RestRechargePushOrderResponse{}, request.ResponseBody, request), err
	}
	// 定义
	var response RestRechargePushOrderResponse
	err = gojson.Unmarshal(request.ResponseBody, &response)
	return newRestRechargePushOrderResult(response, request.ResponseBody, request), err
}
