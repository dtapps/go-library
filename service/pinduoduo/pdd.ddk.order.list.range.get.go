package pinduoduo

import (
	"context"

	"go.dtapp.net/library/utils/gorequest"
)

type OrderListRangeGet struct {
	OrderListGetResponse struct {
		LastOrderID string                              `json:"last_order_id"` // last_order_id
		OrderList   []OrderDetailGetOrderDetailResponse `json:"order_list"`    // 多多进宝推广位对象列表
		RequestID   string                              `json:"request_id"`
	} `json:"order_list_get_response"`
}

// OrderListRangeGet 用时间段查询推广订单接口
// https://jinbao.pinduoduo.com/third-party/api-detail?apiName=pdd.ddk.order.list.range.get
func (c *Client) OrderListRangeGet(ctx context.Context, startTime string, endTime string, lastOrderID string, pageSize int64, notMustParams ...*gorequest.Params) (response OrderListRangeGet, err error) {

	// 参数
	params := NewParamsWithType("pdd.ddk.order.list.range.get", notMustParams...)
	params.Set("start_time", startTime)      // 支付起始时间，格式: "yyyy-MM-dd HH:mm:ss" ，比如 "2020-12-01 00:00:00"
	params.Set("end_time", endTime)          // 支付结束时间，格式: "yyyy-MM-dd HH:mm:ss" ，比如 "2020-12-01 00:00:00"
	params.Set("last_order_id", lastOrderID) // 上一次的迭代器id(第一次不填)
	if pageSize <= 0 {
		params.Set("page_size", pageSize) // 每次请求多少条，建议300
	}

	// 请求
	err = c.request(ctx, params, &response)
	return
}
