package pinduoduo

import (
	"context"
	"go.dtapp.net/library/utils/gorequest"
)

type OrderDetailGetResponse struct {
	OrderDetailResponse struct {
		SepMarketFee          int64   `json:"sep_market_fee"`
		GoodsPrice            int64   `json:"goods_price"`
		SepDuoId              int64   `json:"sep_duo_id"`
		Pid                   string  `json:"pid"`
		PromotionRate         int64   `json:"promotion_rate"`
		CpsSign               string  `json:"cps_sign"`
		Type                  int64   `json:"type"`
		SubsidyDuoAmountLevel int64   `json:"subsidy_duo_amount_level"`
		OrderStatus           int64   `json:"order_status"`
		CatIds                []int64 `json:"cat_ids"`
		OrderCreateTime       int64   `json:"order_create_time"`
		IsDirect              int64   `json:"is_direct"`
		OrderGroupSuccessTime int64   `json:"order_group_success_time"`
		MallId                int64   `json:"mall_id"`
		OrderAmount           int64   `json:"order_amount"`
		PriceCompareStatus    int64   `json:"price_compare_status"`
		MallName              string  `json:"mall_name"`
		OrderModifyAt         int64   `json:"order_modify_at"`
		AuthDuoId             int64   `json:"auth_duo_id"`
		CpaNew                int64   `json:"cpa_new"`
		GoodsName             string  `json:"goods_name"`
		BatchNo               string  `json:"batch_no"`
		RedPacketType         int64   `json:"red_packet_type"`
		UrlLastGenerateTime   int64   `json:"url_last_generate_time"`
		GoodsQuantity         int64   `json:"goods_quantity"`
		GoodsId               int64   `json:"goods_id"`
		SepParameters         string  `json:"sep_parameters"`
		SepRate               int64   `json:"sep_rate"`
		SubsidyType           int64   `json:"subsidy_type"`
		ShareRate             int64   `json:"share_rate"`
		CustomParameters      string  `json:"custom_parameters"`
		GoodsThumbnailUrl     string  `json:"goods_thumbnail_url"`
		PromotionAmount       int64   `json:"promotion_amount"`
		OrderPayTime          int64   `json:"order_pay_time"`
		GroupId               int64   `json:"group_id"`
		SepPid                string  `json:"sep_pid"`
		ReturnStatus          int64   `json:"return_status"`
		OrderStatusDesc       string  `json:"order_status_desc"`
		ShareAmount           int64   `json:"share_amount"`
		GoodsCategoryName     string  `json:"goods_category_name"`
		RequestId             string  `json:"request_id"`
		GoodsSign             string  `json:"goods_sign"`
		OrderSn               string  `json:"order_sn"`
		ZsDuoId               int64   `json:"zs_duo_id"`
	} `json:"order_detail_response"`
}

type OrderDetailGetResult struct {
	Result OrderDetailGetResponse // 结果
	Body   []byte                 // 内容
	Http   gorequest.Response     // 请求
}

func newOrderDetailGetResult(result OrderDetailGetResponse, body []byte, http gorequest.Response) *OrderDetailGetResult {
	return &OrderDetailGetResult{Result: result, Body: body, Http: http}
}

// OrderDetailGet 多多进宝商品查询 https://jinbao.pinduoduo.com/third-party/api-detail?apiName=pdd.ddk.order.detail.get
func (c *Client) OrderDetailGet(ctx context.Context, orderSn string, notMustParams ...gorequest.Params) (*OrderDetailGetResult, error) {

	// 参数
	params := NewParamsWithType("pdd.ddk.order.detail.get", notMustParams...)
	params.Set("order_sn", orderSn)

	// 请求
	var response OrderDetailGetResponse
	request, err := c.request(ctx, params, &response)
	return newOrderDetailGetResult(response, request.ResponseBody, request), err
}
