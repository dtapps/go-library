package pinduoduo

import (
	"context"
	"go.dtapp.net/library/utils/gorequest"
)

type OrderListIncrementGetResponse struct {
	OrderListGetResponse struct {
		TotalCount int64 `json:"total_count"`
		OrderList  []struct {
			SepMarketFee          int64   `json:"sep_market_fee"`
			GoodsPrice            int64   `json:"goods_price"`
			SepDuoId              int64   `json:"sep_duo_id"`
			PromotionRate         int64   `json:"promotion_rate"`
			Type                  int64   `json:"type"`
			SubsidyDuoAmountLevel int64   `json:"subsidy_duo_amount_level"`
			CatIds                []int64 `json:"cat_ids"`
			OrderStatus           int64   `json:"order_status"`
			OrderCreateTime       int64   `json:"order_create_time"`
			IsDirect              int64   `json:"is_direct"`
			OrderGroupSuccessTime int64   `json:"order_group_success_time"`
			MallId                int64   `json:"mall_id"`
			OrderAmount           int64   `json:"order_amount"`
			PriceCompareStatus    int64   `json:"price_compare_status"`
			OrderModifyAt         int64   `json:"order_modify_at"`
			AuthDuoId             int64   `json:"auth_duo_id"`
			CpaNew                int64   `json:"cpa_new"`
			GoodsName             string  `json:"goods_name"`
			BatchNo               string  `json:"batch_no"`
			RedPacketType         int64   `json:"red_packet_type"`
			GoodsQuantity         int64   `json:"goods_quantity"`
			FailReason            string  `json:"fail_reason,omitempty"`
			GoodsId               int64   `json:"goods_id"`
			SepParameters         string  `json:"sep_parameters"`
			SepRate               int64   `json:"sep_rate"`
			SubsidyType           int64   `json:"subsidy_type"`
			CustomParameters      string  `json:"custom_parameters"`
			GoodsThumbnailUrl     string  `json:"goods_thumbnail_url"`
			ShareRate             int64   `json:"share_rate"`
			PromotionAmount       int64   `json:"promotion_amount"`
			OrderPayTime          int64   `json:"order_pay_time"`
			OrderReceiveTime      int64   `json:"order_receive_time"`
			OrderSettleTime       int64   `json:"order_settle_time"`
			ActivityTags          []int64 `json:"activity_tags"`
			GroupId               int64   `json:"group_id"`
			SepPid                string  `json:"sep_pid"`
			OrderStatusDesc       string  `json:"order_status_desc"`
			ShareAmount           int64   `json:"share_amount"`
			OrderId               string  `json:"order_id"`
			GoodsSign             string  `json:"goods_sign"`
			OrderSn               string  `json:"order_sn"`
			OrderVerifyTime       int64   `json:"order_verify_time"`
			PId                   string  `json:"p_id"`
			ZsDuoId               int64   `json:"zs_duo_id"`
		} `json:"order_list"`
		RequestId string `json:"request_id"`
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
func (c *Client) OrderListIncrementGet(ctx context.Context, notMustParams ...gorequest.Params) (*OrderListIncrementGetResult, error) {

	// OpenTelemetry链路追踪
	ctx, span := TraceStartSpan(ctx, "pdd.ddk.order.list.increment.get")
	defer span.End()

	// 参数
	params := NewParamsWithType("pdd.ddk.order.list.increment.get", notMustParams...)

	// 请求
	var response OrderListIncrementGetResponse
	request, err := c.request(ctx, span, params, &response)
	return newOrderListIncrementGetResult(response, request.ResponseBody, request), err
}
