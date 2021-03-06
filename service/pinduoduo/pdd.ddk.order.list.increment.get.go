package pinduoduo

import (
	"encoding/json"
	"github.com/dtapps/go-library/utils/gorequest"
)

type OrderListIncrementGetResponse struct {
	OrderListGetResponse struct {
		TotalCount int `json:"total_count"`
		OrderList  []struct {
			SepMarketFee          int    `json:"sep_market_fee"`
			GoodsPrice            int64  `json:"goods_price"`
			SepDuoId              int    `json:"sep_duo_id"`
			PromotionRate         int64  `json:"promotion_rate"`
			Type                  int    `json:"type"`
			SubsidyDuoAmountLevel int    `json:"subsidy_duo_amount_level"`
			CatIds                []int  `json:"cat_ids"`
			OrderStatus           int    `json:"order_status"`
			OrderCreateTime       int64  `json:"order_create_time"`
			IsDirect              int    `json:"is_direct"`
			OrderGroupSuccessTime int    `json:"order_group_success_time"`
			MallId                int    `json:"mall_id"`
			OrderAmount           int64  `json:"order_amount"`
			PriceCompareStatus    int    `json:"price_compare_status"`
			OrderModifyAt         int    `json:"order_modify_at"`
			AuthDuoId             int    `json:"auth_duo_id"`
			CpaNew                int    `json:"cpa_new"`
			GoodsName             string `json:"goods_name"`
			BatchNo               string `json:"batch_no"`
			RedPacketType         int    `json:"red_packet_type"`
			GoodsQuantity         int    `json:"goods_quantity"`
			FailReason            string `json:"fail_reason,omitempty"`
			GoodsId               int64  `json:"goods_id"`
			SepParameters         string `json:"sep_parameters"`
			SepRate               int    `json:"sep_rate"`
			SubsidyType           int    `json:"subsidy_type"`
			CustomParameters      string `json:"custom_parameters"`
			GoodsThumbnailUrl     string `json:"goods_thumbnail_url"`
			ShareRate             int    `json:"share_rate"`
			PromotionAmount       int64  `json:"promotion_amount"`
			OrderPayTime          int64  `json:"order_pay_time"`
			OrderReceiveTime      int64  `json:"order_receive_time"`
			OrderSettleTime       int64  `json:"order_settle_time"`
			ActivityTags          []int  `json:"activity_tags"`
			GroupId               int64  `json:"group_id"`
			SepPid                string `json:"sep_pid"`
			OrderStatusDesc       string `json:"order_status_desc"`
			ShareAmount           int    `json:"share_amount"`
			OrderId               string `json:"order_id"`
			GoodsSign             string `json:"goods_sign"`
			OrderSn               string `json:"order_sn"`
			OrderVerifyTime       int64  `json:"order_verify_time"`
			PId                   string `json:"p_id"`
			ZsDuoId               int    `json:"zs_duo_id"`
		} `json:"order_list"`
		RequestId string `json:"request_id"`
	} `json:"order_list_get_response"`
}

type OrderListIncrementGetResult struct {
	Result OrderListIncrementGetResponse // ??????
	Body   []byte                        // ??????
	Http   gorequest.Response            // ??????
	Err    error                         // ??????
}

func newOrderListIncrementGetResult(result OrderListIncrementGetResponse, body []byte, http gorequest.Response, err error) *OrderListIncrementGetResult {
	return &OrderListIncrementGetResult{Result: result, Body: body, Http: http, Err: err}
}

// OrderListIncrementGet ???????????????????????????????????????????????????
// https://jinbao.pinduoduo.com/third-party/api-detail?apiName=pdd.ddk.order.list.increment.get
func (c *Client) OrderListIncrementGet(notMustParams ...Params) *OrderListIncrementGetResult {
	// ??????
	params := NewParamsWithType("pdd.ddk.order.list.increment.get", notMustParams...)
	// ??????
	request, err := c.request(params)
	// ??????
	var response OrderListIncrementGetResponse
	err = json.Unmarshal(request.ResponseBody, &response)
	return newOrderListIncrementGetResult(response, request.ResponseBody, request, err)
}
