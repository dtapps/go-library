package praise_goodness

import (
	"context"
	"go.dtapp.net/library/utils/gorequest"
	"net/http"
)

type AddProductOrderResponse struct {
	Code int      `json:"code"` // 1：请求成功 -1：请求失败
	Msg  string   `json:"msg"`  // 返回说明
	Time string   `json:"time"` // 时间戳
	Data struct{} `json:"data"`
}

type AddProductOrderResult struct {
	Result AddProductOrderResponse // 结果
	Body   []byte                  // 内容
	Http   gorequest.Response      // 请求
}

func newAddProductOrderResult(result AddProductOrderResponse, body []byte, http gorequest.Response) *AddProductOrderResult {
	return &AddProductOrderResult{Result: result, Body: body, Http: http}
}

// AddProductOrder 下单接口
// type = 固定值:1
// mobile = 充值账号
// trade_id = 商户订单号，10-32位
// amount = 充值金额，单位 (元) ，只传整 数不带小数点，例如:100
// official = 运营商 1 :中国移动 2 :中国联通 3 :中国电信
// area = 归属地，填省份，如:山东，北京 ，内蒙古，黑龙江，重庆 填:auto,系统自动识别归属地
// notifyurl = 回调通知地址
func (c *Client) AddProductOrder(ctx context.Context, Type int64, mobile string, tradeID string, amount int64, official int64, area string, notifyurl string, notMustParams ...*gorequest.Params) (*AddProductOrderResult, error) {

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	params.Set("mch_id", c.GetMchID()) // 商户编号 (平台提供)
	params.Set("type", Type)           // 固定值:1
	params.Set("mobile", mobile)       // 充值账号
	params.Set("trade_id", tradeID)    // 商户订单号，10-32位
	params.Set("amount", amount)       // 充值金额，单位 (元) ，只传整 数不带小数点，例如:100
	params.Set("official", official)   // 运营商 1 :中国移动 2 :中国联通 3 :中国电信
	params.Set("area", area)           // 归属地，填省份，如:山东，北京 ，内蒙古，黑龙江，重庆 填:auto,系统自动识别归属地
	params.Set("notifyurl", notifyurl) // 回调通知地址

	// 响应
	var response AddProductOrderResponse

	// 请求
	request, err := c.request(ctx, "api/order/addProductOrder", params, http.MethodPost, &response)
	return newAddProductOrderResult(response, request.ResponseBody, request), err
}
