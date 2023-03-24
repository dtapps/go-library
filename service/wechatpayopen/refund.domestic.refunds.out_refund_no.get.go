package wechatpayopen

import (
	"context"
	"github.com/dtapps/go-library/utils/gojson"
	"github.com/dtapps/go-library/utils/gorequest"
	"net/http"
	"time"
)

type RefundDomesticRefundsOutRefundNoGetResponse struct {
	RefundId            string    `json:"refund_id"`             // 微信支付退款单号
	OutRefundNo         string    `json:"out_refund_no"`         // 商户退款单号
	TransactionId       string    `json:"transaction_id"`        // 微信支付订单号
	OutTradeNo          string    `json:"out_trade_no"`          // 商户订单号
	Channel             string    `json:"channel"`               // 退款渠道
	UserReceivedAccount string    `json:"user_received_account"` // 退款入账账户
	SuccessTime         time.Time `json:"success_time"`          // 退款成功时间
	CreateTime          time.Time `json:"create_time"`           // 退款创建时间
	Status              string    `json:"status"`                // 退款状态
	FundsAccount        string    `json:"funds_account"`         // 资金账户
	Amount              struct {
		Total  int `json:"total"`  // 订单金额
		Refund int `json:"refund"` // 退款金额
		From   []struct {
			Account string `json:"account"` // 出资账户类型
			Amount  int    `json:"amount"`  // 出资金额
		} `json:"from"` // 退款出资账户及金额
		PayerTotal       int    `json:"payer_total"`       // 用户支付金额
		PayerRefund      int    `json:"payer_refund"`      // 用户退款金额
		SettlementRefund int    `json:"settlement_refund"` // 应结退款金额
		SettlementTotal  int    `json:"settlement_total"`  // 应结订单金额
		DiscountRefund   int    `json:"discount_refund"`   // 优惠退款金额
		Currency         string `json:"currency"`          // 退款币种
		RefundFee        int    `json:"refund_fee"`        // 手续费退款金额
	} `json:"amount"` // 金额信息
	PromotionDetail []struct {
		PromotionId  string `json:"promotion_id"`  // 券ID
		Scope        string `json:"scope"`         // 优惠范围
		Type         string `json:"type"`          // 优惠类型
		Amount       int    `json:"amount"`        // 优惠券面额
		RefundAmount int    `json:"refund_amount"` // 优惠退款金额
		GoodsDetail  []struct {
			MerchantGoodsId  string `json:"merchant_goods_id"`  // 商户侧商品编码
			WechatpayGoodsId string `json:"wechatpay_goods_id"` // 微信支付商品编码
			GoodsName        string `json:"goods_name"`         // 商品名称
			UnitPrice        int    `json:"unit_price"`         // 商品单价
			RefundAmount     int    `json:"refund_amount"`      // 商品退款金额
			RefundQuantity   int    `json:"refund_quantity"`    // 商品退货数量
		} `json:"goods_detail"` // 商品列表
	} `json:"promotion_detail"` // 优惠退款信息
}

type RefundDomesticRefundsOutRefundNoGetResult struct {
	Result   RefundDomesticRefundsOutRefundNoGetResponse // 结果
	Body     []byte                                      // 内容
	Http     gorequest.Response                          // 请求
	Err      error                                       // 错误
	ApiError ApiError                                    // 接口错误
}

func newRefundDomesticRefundsOutRefundNoGetResult(result RefundDomesticRefundsOutRefundNoGetResponse, body []byte, http gorequest.Response, err error, apiError ApiError) *RefundDomesticRefundsOutRefundNoGetResult {
	return &RefundDomesticRefundsOutRefundNoGetResult{Result: result, Body: body, Http: http, Err: err, ApiError: apiError}
}

// RefundDomesticRefundsOutRefundNoGet 查询单笔退款API
// https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter4_4_10.shtml
func (c *Client) RefundDomesticRefundsOutRefundNoGet(ctx context.Context, outRefundNo string, notMustParams ...gorequest.Params) *RefundDomesticRefundsOutRefundNoGetResult {
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	// 请求
	request, err := c.request(ctx, apiUrl+"/v3/refund/domestic/refunds/"+outRefundNo, params, http.MethodGet)
	if err != nil {
		return newRefundDomesticRefundsOutRefundNoGetResult(RefundDomesticRefundsOutRefundNoGetResponse{}, request.ResponseBody, request, err, ApiError{})
	}
	// 定义
	var response RefundDomesticRefundsOutRefundNoGetResponse
	err = gojson.Unmarshal(request.ResponseBody, &response)
	// 错误
	var apiError ApiError
	err = gojson.Unmarshal(request.ResponseBody, &apiError)
	return newRefundDomesticRefundsOutRefundNoGetResult(response, request.ResponseBody, request, err, apiError)
}
