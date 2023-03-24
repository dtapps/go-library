package pinduoduo

import (
	"context"
	"github.com/dtapps/go-library/utils/gojson"
	"github.com/dtapps/go-library/utils/gorequest"
)

type PddDdkOauthPidMediaIdBindResponse struct {
	OrderMediaIdBindResponse struct {
		SepMarketFee          int    `json:"sep_market_fee"`
		PidPrice              int    `json:"Pid_price"`
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
		PidName               string `json:"Pid_name"`
		BatchNo               string `json:"batch_no"`
		RedPacketType         int    `json:"red_packet_type"`
		UrlLastGenerateTime   int    `json:"url_last_generate_time"`
		PidQuantity           int    `json:"Pid_quantity"`
		PidId                 int64  `json:"Pid_id"`
		SepParameters         string `json:"sep_parameters"`
		SepRate               int    `json:"sep_rate"`
		SubsidyType           int    `json:"subsidy_type"`
		ShareRate             int    `json:"share_rate"`
		CustomParameters      string `json:"custom_parameters"`
		PidThumbnailUrl       string `json:"Pid_thumbnail_url"`
		PromotionAmount       int64  `json:"promotion_amount"`
		OrderPayTime          int    `json:"order_pay_time"`
		GroupId               int64  `json:"group_id"`
		SepPid                string `json:"sep_pid"`
		ReturnStatus          int    `json:"return_status"`
		OrderStatusDesc       string `json:"order_status_desc"`
		ShareAmount           int    `json:"share_amount"`
		PidCategoryName       string `json:"Pid_category_name"`
		RequestId             string `json:"request_id"`
		PidSign               string `json:"Pid_sign"`
		OrderSn               string `json:"order_sn"`
		ZsDuoId               int    `json:"zs_duo_id"`
	} `json:"order_MediaIdBind_response"`
}

type PddDdkOauthPidMediaIdBindResult struct {
	Result PddDdkOauthPidMediaIdBindResponse // 结果
	Body   []byte                            // 内容
	Http   gorequest.Response                // 请求
	Err    error                             // 错误
}

func newPddDdkOauthPidMediaIdBindResult(result PddDdkOauthPidMediaIdBindResponse, body []byte, http gorequest.Response, err error) *PddDdkOauthPidMediaIdBindResult {
	return &PddDdkOauthPidMediaIdBindResult{Result: result, Body: body, Http: http, Err: err}
}

// MediaIdBind 批量绑定推广位的媒体id
// https://jinbao.pinduoduo.com/third-party/api-detail?apiName=pdd.ddk.oauth.pid.mediaid.bind
func (c *PddDdkOauthPidApi) MediaIdBind(ctx context.Context, notMustParams ...Params) *PddDdkOauthPidMediaIdBindResult {
	// 参数
	params := NewParamsWithType("pdd.ddk.oauth.pid.mediaid.bind", notMustParams...)
	// 请求
	request, err := c.client.request(ctx, params)
	// 定义
	var response PddDdkOauthPidMediaIdBindResponse
	err = gojson.Unmarshal(request.ResponseBody, &response)
	return newPddDdkOauthPidMediaIdBindResult(response, request.ResponseBody, request, err)
}
