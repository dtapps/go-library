package wechatpayopen

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/dtapps/go-library/utils/gorequest"
	"net/http"
)

type PayPartnerTransactionsIdResponse struct {
	Appid          string `json:"appid"`
	Mchid          string `json:"mchid"`
	OutTradeNo     string `json:"out_trade_no"`
	TransactionId  string `json:"transaction_id,omitempty"`
	TradeType      string `json:"trade_type,omitempty"`
	TradeState     string `json:"trade_state"`
	TradeStateDesc string `json:"trade_state_desc"`
	BankType       string `json:"bank_type,omitempty"`
	Attach         string `json:"attach,omitempty"`
	SuccessTime    string `json:"success_time,omitempty"`
	Payer          struct {
		Openid string `json:"openid"`
	} `json:"payer"`
	Amount struct {
		Total         int    `json:"total,omitempty"`
		PayerTotal    int    `json:"payer_total,omitempty"`
		Currency      string `json:"currency,omitempty"`
		PayerCurrency string `json:"payer_currency,omitempty"`
	} `json:"amount,omitempty"`
	SceneInfo struct {
		DeviceId string `json:"device_id,omitempty"`
	}
	PromotionDetail []struct {
		CouponId            string `json:"coupon_id"`
		Name                string `json:"name,omitempty"`
		Scope               string `json:"scope,omitempty"`
		Type                string `json:"type,omitempty"`
		Amount              int    `json:"amount"`
		StockId             string `json:"stock_id,omitempty"`
		WechatpayContribute int    `json:"wechatpay_contribute,omitempty"`
		MerchantContribute  int    `json:"merchant_contribute,omitempty"`
		OtherContribute     int    `json:"other_contribute,omitempty"`
		Currency            string `json:"currency,omitempty"`
		GoodsDetail         []struct {
			GoodsId        string `json:"goods_id"`
			Quantity       int    `json:"quantity"`
			UnitPrice      int    `json:"unit_price"`
			DiscountAmount int    `json:"discount_amount"`
			GoodsRemark    string `json:"goods_remark,omitempty"`
		} `json:"goods_detail"`
	}
}

type PayPartnerTransactionsIdResult struct {
	Result PayPartnerTransactionsIdResponse // 结果
	Body   []byte                           // 内容
	Http   gorequest.Response               // 请求
	Err    error                            // 错误
}

func newPayPartnerTransactionsIdResult(result PayPartnerTransactionsIdResponse, body []byte, http gorequest.Response, err error) *PayPartnerTransactionsIdResult {
	return &PayPartnerTransactionsIdResult{Result: result, Body: body, Http: http, Err: err}
}

// PayPartnerTransactionsId 微信支付订单号查询
// https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter4_5_2.shtml
func (c *Client) PayPartnerTransactionsId(ctx context.Context, transactionId string) *PayPartnerTransactionsIdResult {
	// 参数
	params := gorequest.NewParams()
	// 请求
	request, err := c.request(ctx, fmt.Sprintf(apiUrl+"/v3/pay/partner/transactions/id/%s?sp_mchid=%s&sub_mchid=%s", transactionId, c.config.SpMchId, c.config.SubMchId), params, http.MethodGet)
	if err != nil {
		return newPayPartnerTransactionsIdResult(PayPartnerTransactionsIdResponse{}, request.ResponseBody, request, err)
	}
	// 定义
	var response PayPartnerTransactionsIdResponse
	err = json.Unmarshal(request.ResponseBody, &response)
	return newPayPartnerTransactionsIdResult(response, request.ResponseBody, request, err)
}
