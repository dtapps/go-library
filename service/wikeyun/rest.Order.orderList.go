package wikeyun

import (
	"context"
	"go.dtapp.net/library/utils/gorequest"
)

type RestOrderOrderListResponse struct {
	Code string `json:"code"`
	Msg  string `json:"msg"`
	Data struct {
		Id     string `json:"id"`
		Avatar string `json:"avatar"`
		Money  string `json:"money"`
		Mobile string `json:"mobile"`
	} `json:"data"`
}

type RestOrderOrderListResult struct {
	Result RestOrderOrderListResponse // 结果
	Body   []byte                     // 内容
	Http   gorequest.Response         // 请求
}

func newRestOrderOrderListResult(result RestOrderOrderListResponse, body []byte, http gorequest.Response) *RestOrderOrderListResult {
	return &RestOrderOrderListResult{Result: result, Body: body, Http: http}
}

// RestOrderOrderList 导购订单列表【系统内】
// order_type = 订单类型，0=淘宝，1=天猫，2=拼多多，3=京东,4=ele,5=美团，6=电影票 7话费 8油卡 9=肯德基 10=美团生活，11=会员卡 12=滴滴 13=电费，14=生活团购 15=唯品会
// status = 订单状态 0=订单付款，1=订单结算，2=订单失效
// query_type = 查询时间类型 1创建时间 2更新时间 3结算时间
// page_size = 每页大小,默认20 ，最大500
// p = 页数，1
// start_time = 开始时间，linux时间戳10位
// end_time = 结束时间，linux时间戳10位，跟开始时间相差不能超过2小时
// https://open.wikeyun.cn/#/apiDocument/13/document/364
func (c *Client) RestOrderOrderList(ctx context.Context, notMustParams ...*gorequest.Params) (*RestOrderOrderListResult, error) {

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)

	// 请求
	var response RestOrderOrderListResponse
	request, err := c.request(ctx, "rest/Order/orderList", params, &response)
	return newRestOrderOrderListResult(response, request.ResponseBody, request), err
}
