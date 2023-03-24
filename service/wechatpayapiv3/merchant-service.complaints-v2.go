package wechatpayapiv3

import (
	"context"
	"github.com/dtapps/go-library/utils/gojson"
	"github.com/dtapps/go-library/utils/gorequest"
	"net/http"
)

type MerchantServiceComplaintsV2Response struct {
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

type MerchantServiceComplaintsV2Result struct {
	Result MerchantServiceComplaintsV2Response // 结果
	Body   []byte                              // 内容
	Http   gorequest.Response                  // 请求
	Err    error                               // 错误
}

func newMerchantServiceComplaintsV2Result(result MerchantServiceComplaintsV2Response, body []byte, http gorequest.Response, err error) *MerchantServiceComplaintsV2Result {
	return &MerchantServiceComplaintsV2Result{Result: result, Body: body, Http: http, Err: err}
}

// MerchantServiceComplaintsV2 查询投诉单列表API
// https://pay.weixin.qq.com/wiki/doc/apiv3/apis/chapter10_2_11.shtml
func (c *Client) MerchantServiceComplaintsV2(ctx context.Context, notMustParams ...gorequest.Params) *MerchantServiceComplaintsV2Result {
	// 参数
	params := gorequest.NewParamsWith(notMustParams...)
	// 请求
	request, err := c.request(ctx, apiUrl+"/v3/merchant-service/complaints-v2", params, http.MethodGet, false)
	if err != nil {
		return newMerchantServiceComplaintsV2Result(MerchantServiceComplaintsV2Response{}, request.ResponseBody, request, err)
	}
	// 定义
	var response MerchantServiceComplaintsV2Response
	err = gojson.Unmarshal(request.ResponseBody, &response)
	return newMerchantServiceComplaintsV2Result(response, request.ResponseBody, request, err)
}
