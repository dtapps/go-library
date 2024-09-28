package pinduoduo

import (
	"context"
	"go.dtapp.net/library/utils/gorequest"
)

type PddDdkOauthOrderDetailGetResponse struct {
	OrderDetailGetResponse struct {
		SepMarketFee          int    `json:"sep_market_fee"`
		OrderPrice            int    `json:"Order_price"`
		SepDuoId              int    `json:"sep_duo_id"`
		Pid                   string `json:"pid"`
		PromotionRate         int    `json:"promotion_rate"`
		CpsSign               string `json:"cps_sign"`
		Type                  int    `json:"type"`
		SubsidyDuoAmountLevel int    `json:"subsidy_duo_amount_level"`
		OrderStatus           int    `json:"order_status"`
		CatIds                []int  `json:"cat_ids"`
		OrderCreateTime       int64  `json:"order_create_time"`
		IsDirect              int    `json:"is_direct"`
		OrderGroupSuccessTime int    `json:"order_group_success_time"`
		MallId                int    `json:"mall_id"`
		OrderAmount           int64  `json:"order_amount"`
		PriceCompareStatus    int    `json:"price_compare_status"`
		MallName              string `json:"mall_name"`
		OrderModifyAt         int    `json:"order_modify_at"`
		AuthDuoId             int    `json:"auth_duo_id"`
		CpaNew                int    `json:"cpa_new"`
		OrderName             string `json:"Order_name"`
		BatchNo               string `json:"batch_no"`
		RedPacketType         int    `json:"red_packet_type"`
		UrlLastGenerateTime   int    `json:"url_last_generate_time"`
		OrderQuantity         int    `json:"Order_quantity"`
		OrderId               int64  `json:"Order_id"`
		SepParameters         string `json:"sep_parameters"`
		SepRate               int    `json:"sep_rate"`
		SubsidyType           int    `json:"subsidy_type"`
		ShareRate             int    `json:"share_rate"`
		CustomParameters      string `json:"custom_parameters"`
		OrderThumbnailUrl     string `json:"Order_thumbnail_url"`
		PromotionAmount       int64  `json:"promotion_amount"`
		OrderPayTime          int    `json:"order_pay_time"`
		GroupId               int64  `json:"group_id"`
		SepPid                string `json:"sep_pid"`
		ReturnStatus          int    `json:"return_status"`
		OrderStatusDesc       string `json:"order_status_desc"`
		ShareAmount           int    `json:"share_amount"`
		OrderCategoryName     string `json:"Order_category_name"`
		RequestId             string `json:"request_id"`
		OrderSign             string `json:"Order_sign"`
		OrderSn               string `json:"order_sn"`
		ZsDuoId               int    `json:"zs_duo_id"`
	} `json:"order_DetailGet_response"`
}

type PddDdkOauthOrderDetailGetResult struct {
	Result PddDdkOauthOrderDetailGetResponse // 结果
	Body   []byte                            // 内容
	Http   gorequest.Response                // 请求
}

func newPddDdkOauthOrderDetailGetResult(result PddDdkOauthOrderDetailGetResponse, body []byte, http gorequest.Response) *PddDdkOauthOrderDetailGetResult {
	return &PddDdkOauthOrderDetailGetResult{Result: result, Body: body, Http: http}
}

// OauthOrderDetailGet 获取订单详情
// https://jinbao.pinduoduo.com/third-party/api-detail?apiName=pdd.ddk.oauth.order.detail.get
func (c *Client) OauthOrderDetailGet(ctx context.Context, notMustParams ...gorequest.Params) (*PddDdkOauthOrderDetailGetResult, error) {

	// OpenTelemetry链路追踪
	ctx, span := TraceStartSpan(ctx, "pdd.ddk.oauth.order.detail.get")
	defer span.End()

	// 参数
	params := NewParamsWithType("pdd.ddk.oauth.order.detail.get", notMustParams...)

	// 请求
	var response PddDdkOauthOrderDetailGetResponse
	request, err := c.request(ctx, span, params, &response)
	return newPddDdkOauthOrderDetailGetResult(response, request.ResponseBody, request), err
}
