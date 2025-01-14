package wechatpayapiv3

import (
	"context"
	"fmt"
	"go.dtapp.net/library/utils/gorequest"
	"net/http"
)

type PayTransactionsOutTradeNoResponse struct {
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
	} `json:"scene_info"`
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
	} `json:"promotion_detail"`
}

type PayTransactionsOutTradeNoResult struct {
	Result PayTransactionsOutTradeNoResponse // 结果
	Body   []byte                            // 内容
	Http   gorequest.Response                // 请求
}

func newPayTransactionsOutTradeNoResult(result PayTransactionsOutTradeNoResponse, body []byte, http gorequest.Response) *PayTransactionsOutTradeNoResult {
	return &PayTransactionsOutTradeNoResult{Result: result, Body: body, Http: http}
}

// PayTransactionsOutTradeNo 商户订单号查询 https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter3_1_2.shtml
func (c *Client) PayTransactionsOutTradeNo(ctx context.Context, outTradeNo string, notMustParams ...*gorequest.Params) (*PayTransactionsOutTradeNoResult, error) {

	// 参数
	params := gorequest.NewParamsWith(notMustParams...)

	// 请求
	var response PayTransactionsOutTradeNoResponse
	request, err := c.request(ctx, fmt.Sprintf("v3/pay/transactions/out-trade-no/%s?mchid=%s", outTradeNo, c.GetMchId()), params, http.MethodGet, true, &response)
	return newPayTransactionsOutTradeNoResult(response, request.ResponseBody, request), err
}
