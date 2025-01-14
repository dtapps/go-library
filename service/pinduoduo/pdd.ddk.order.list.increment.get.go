package pinduoduo

import (
	"context"
	"go.dtapp.net/library/utils/gorequest"
)

type OrderListIncrementGetResponse struct {
	OrderListGetResponse struct {
		TotalCount int64                               `json:"total_count"` // 请求到的结果数
		OrderList  []OrderDetailGetOrderDetailResponse `json:"order_list"`  // 多多进宝推广位对象列表
		RequestID  string                              `json:"request_id"`
	} `json:"order_list_get_response"`
}

type OrderListIncrementGetResult struct {
	Result OrderListIncrementGetResponse // 结果
	Body   []byte                        // 内容
	Http   gorequest.Response            // 请求
}

func newOrderListIncrementGetResult(result OrderListIncrementGetResponse, body []byte, http gorequest.Response) *OrderListIncrementGetResult {
	return &OrderListIncrementGetResult{Result: result, Body: body, Http: http}
}

// OrderListIncrementGet 最后更新时间段增量同步推广订单信息
// https://jinbao.pinduoduo.com/third-party/api-detail?apiName=pdd.ddk.order.list.increment.get
func (c *Client) OrderListIncrementGet(ctx context.Context, startUpdateTime int64, endUpdateTime int64, page int64, pageSize int64, notMustParams ...*gorequest.Params) (*OrderListIncrementGetResult, error) {

	// 参数
	params := NewParamsWithType("pdd.ddk.order.list.increment.get", notMustParams...)
	params.Set("start_update_time", startUpdateTime) // 最近90天内多多进宝商品订单更新时间--查询时间开始。note：此时间为时间戳，指格林威治时间 1970 年01 月 01 日 00 时 00 分 00 秒(北京时间 1970 年 01 月 01 日 08 时 00 分 00 秒)起至现在的总秒数
	params.Set("end_update_time", endUpdateTime)     // 查询结束时间，和开始时间相差不能超过24小时。note：此时间为时间戳，指格林威治时间 1970 年01 月 01 日 00 时 00 分 00 秒(北京时间 1970 年 01 月 01 日 08 时 00 分 00 秒)起至现在的总秒数
	if page > 0 {
		params.Set("page", page) // 第几页，从1到10000，默认1，注：使用最后更新时间范围增量同步时，必须采用倒序的分页方式（从最后一页往回取）才能避免漏单问题。
	} else {
		params.Set("page", 1)
	}
	if pageSize > 0 {
		params.Set("page_size", pageSize) // 返回的每页结果订单数，默认为100，范围为10到100，建议使用40~50，可以提高成功率，减少超时数量。
	} else {
		params.Set("page_size", 100)
	}

	// 请求
	var response OrderListIncrementGetResponse
	request, err := c.request(ctx, params, &response)
	return newOrderListIncrementGetResult(response, request.ResponseBody, request), err
}
