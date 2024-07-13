package pinduoduo

import (
	"context"
	"go.dtapp.net/library/utils/gorequest"
)

type PddDdkOauthMemberAuthorityQueryResponse struct {
	OrderAuthorityQueryResponse struct {
		SepMarketFee          int    `json:"sep_market_fee"`
		MemberPrice           int    `json:"Member_price"`
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
		MemberName            string `json:"Member_name"`
		BatchNo               string `json:"batch_no"`
		RedPacketType         int    `json:"red_packet_type"`
		UrlLastGenerateTime   int    `json:"url_last_generate_time"`
		MemberQuantity        int    `json:"Member_quantity"`
		MemberId              int64  `json:"Member_id"`
		SepParameters         string `json:"sep_parameters"`
		SepRate               int    `json:"sep_rate"`
		SubsidyType           int    `json:"subsidy_type"`
		ShareRate             int    `json:"share_rate"`
		CustomParameters      string `json:"custom_parameters"`
		MemberThumbnailUrl    string `json:"Member_thumbnail_url"`
		PromotionAmount       int64  `json:"promotion_amount"`
		OrderPayTime          int    `json:"order_pay_time"`
		GroupId               int64  `json:"group_id"`
		SepPid                string `json:"sep_pid"`
		ReturnStatus          int    `json:"return_status"`
		OrderStatusDesc       string `json:"order_status_desc"`
		ShareAmount           int    `json:"share_amount"`
		MemberCategoryName    string `json:"Member_category_name"`
		RequestId             string `json:"request_id"`
		MemberSign            string `json:"Member_sign"`
		OrderSn               string `json:"order_sn"`
		ZsDuoId               int    `json:"zs_duo_id"`
	} `json:"order_AuthorityQuery_response"`
}

type PddDdkOauthMemberAuthorityQueryResult struct {
	Result PddDdkOauthMemberAuthorityQueryResponse // 结果
	Body   []byte                                  // 内容
	Http   gorequest.Response                      // 请求
}

func newPddDdkOauthMemberAuthorityQueryResult(result PddDdkOauthMemberAuthorityQueryResponse, body []byte, http gorequest.Response) *PddDdkOauthMemberAuthorityQueryResult {
	return &PddDdkOauthMemberAuthorityQueryResult{Result: result, Body: body, Http: http}
}

// OauthMemberAuthorityQuery 查询是否绑定备案
// https://jinbao.pinduoduo.com/third-party/api-detail?apiName=pdd.ddk.oauth.member.authority.query
func (c *Client) OauthMemberAuthorityQuery(ctx context.Context, notMustParams ...gorequest.Params) (*PddDdkOauthMemberAuthorityQueryResult, error) {

	// OpenTelemetry链路追踪
	ctx = c.TraceStartSpan(ctx, "pdd.ddk.oauth.member.authority.query")
	defer c.TraceEndSpan()

	// 参数
	params := NewParamsWithType("pdd.ddk.oauth.member.authority.query", notMustParams...)

	// 请求
	var response PddDdkOauthMemberAuthorityQueryResponse
	request, err := c.request(ctx, params, &response)
	return newPddDdkOauthMemberAuthorityQueryResult(response, request.ResponseBody, request), err
}
