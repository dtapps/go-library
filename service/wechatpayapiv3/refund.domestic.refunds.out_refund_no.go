package wechatpayapiv3

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/dtapps/go-library/utils/gorequest"
	"net/http"
)

type RefundDomesticRefundsOutRefundNoResponse struct {
	RefundId            string `json:"refund_id"`               // 微信支付退款单号
	OutRefundNo         string `json:"out_refund_no"`           // 商户退款单号
	TransactionId       string `json:"transaction_id"`          // 微信支付订单号
	OutTradeNo          string `json:"out_trade_no"`            // 商户订单号
	Channel             string `json:"channel"`                 // 退款渠道
	UserReceivedAccount string `json:"user_received_account"`   // 退款入账账户
	SuccessTime         string `json:"success_time,omitempty"`  // 退款成功时间
	CreateTime          string `json:"create_time"`             // 退款创建时间
	Status              string `json:"status"`                  // 退款状态
	FundsAccount        string `json:"funds_account,omitempty"` // 资金账户
	Amount              struct {
		Total  int `json:"total"`  // 订单金额
		Refund int `json:"refund"` // 退款金额
		From   []struct {
			Account string `json:"account"` // 出资账户类型
			Amount  int    `json:"amount"`  // 出资金额
		} `json:"from,omitempty"`
		PayerTotal       int    `json:"payer_Total"`       // 用户支付金额
		PayerRefund      int    `json:"payer_Refund"`      // 用户退款金额
		SettlementRefund int    `json:"settlement_Refund"` // 应结退款金额
		SettlementTotal  int    `json:"settlement_total"`  // 应结订单金额
		DiscountRefund   int    `json:"discount_refund"`   // 优惠退款金额
		Currency         string `json:"currency"`          // 退款币种
	} `json:"amount"` // 金额信息
	PromotionDetail []struct {
		PromotionId  string `json:"promotion_id"`  // 券ID
		Scope        string `json:"scope"`         // 优惠范围
		Type         string `json:"type"`          // 优惠类型
		Amount       int    `json:"amount"`        // 优惠券面额
		RefundAmount int    `json:"refund_amount"` // 优惠退款金额
		GoodsDetail  []struct {
			MerchantGoodsId  string `json:"merchant_goods_id"`            // 商户侧商品编码
			WechatpayGoodsId string `json:"wechatpay_goods_id,omitempty"` // 微信侧商品编码
			GoodsName        string `json:"goods_name,omitempty"`         // 商品名称
			UnitPrice        int    `json:"unit_price"`                   // 商品单价
			RefundAmount     int    `json:"refund_amount"`                // 商品退款金额
			RefundQuantity   int    `json:"refund_quantity"`              // 商品退货数量
		} `json:"goods_detail"`
	} `json:"promotion_detail,omitempty"` // 优惠退款信息
}

type RefundDomesticRefundsOutRefundNoResult struct {
	Result RefundDomesticRefundsOutRefundNoResponse // 结果
	Body   []byte                                   // 内容
	Http   gorequest.Response                       // 请求
	Err    error                                    // 错误
}

func newRefundDomesticRefundsOutRefundNoResult(result RefundDomesticRefundsOutRefundNoResponse, body []byte, http gorequest.Response, err error) *RefundDomesticRefundsOutRefundNoResult {
	return &RefundDomesticRefundsOutRefundNoResult{Result: result, Body: body, Http: http, Err: err}
}

// RefundDomesticRefundsOutRefundNo 查询单笔退款API https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter4_5_10.shtml
func (c *Client) RefundDomesticRefundsOutRefundNo(ctx context.Context, outRefundNo string) *RefundDomesticRefundsOutRefundNoResult {
	// 请求
	request, err := c.request(ctx, fmt.Sprintf(apiUrl+"/v3/refund/domestic/refunds/%s", outRefundNo), map[string]interface{}{}, http.MethodGet, true)
	if err != nil {
		return newRefundDomesticRefundsOutRefundNoResult(RefundDomesticRefundsOutRefundNoResponse{}, request.ResponseBody, request, err)
	}
	// 定义
	var response RefundDomesticRefundsOutRefundNoResponse
	err = json.Unmarshal(request.ResponseBody, &response)
	return newRefundDomesticRefundsOutRefundNoResult(response, request.ResponseBody, request, err)
}
