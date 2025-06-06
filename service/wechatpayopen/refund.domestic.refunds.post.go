package wechatpayopen

import (
	"context"
	"go.dtapp.net/library/utils/gorequest"
	"net/http"
	"time"
)

type RefundDomesticRefundsPostResponse struct {
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

type RefundDomesticRefundsPostResult struct {
	Result RefundDomesticRefundsPostResponse // 结果
	Body   []byte                            // 内容
	Http   gorequest.Response                // 请求
}

func newRefundDomesticRefundsPostResult(result RefundDomesticRefundsPostResponse, body []byte, http gorequest.Response) *RefundDomesticRefundsPostResult {
	return &RefundDomesticRefundsPostResult{Result: result, Body: body, Http: http}
}

// RefundDomesticRefundsPost 申请退款API
// https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter4_4_9.shtml
func (c *Client) RefundDomesticRefundsPost(ctx context.Context, outRefundNo string, notMustParams ...*gorequest.Params) (*RefundDomesticRefundsPostResult, ApiError, error) {

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	params.Set("sub_mchid", c.GetSubMchId()) // 子商户号

	// 请求
	var response RefundDomesticRefundsPostResponse
	var apiError ApiError
	request, err := c.request(ctx, "v3/refund/domestic/refunds", params, http.MethodPost, &response, &apiError)
	return newRefundDomesticRefundsPostResult(response, request.ResponseBody, request), apiError, err
}
