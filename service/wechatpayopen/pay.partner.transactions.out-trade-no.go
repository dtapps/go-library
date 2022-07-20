package wechatpayopen

import (
	"encoding/json"
	"fmt"
	"github.com/dtapps/go-library/utils/gorequest"
	"net/http"
)

type PayPartnerTransactionsOutTradeNoResponse struct {
	SpAppid        string `json:"sp_appid"`         // 服务商应用ID
	SpMchid        string `json:"sp_mchid"`         // 服务商户号
	SubAppid       string `json:"sub_appid"`        // 子商户应用ID
	SubMchid       string `json:"sub_mchid"`        // 子商户号
	OutTradeNo     string `json:"out_trade_no"`     // 商户订单号
	TransactionId  string `json:"transaction_id"`   // 微信支付订单号
	TradeType      string `json:"trade_type"`       // 交易类型
	TradeState     string `json:"trade_state"`      // 交易状态
	TradeStateDesc string `json:"trade_state_desc"` // 交易状态描述
	BankType       string `json:"bank_type"`        // 付款银行
	Attach         string `json:"attach"`           // 附加数据
	SuccessTime    string `json:"success_time"`     // 支付完成时间
	Payer          struct {
		SpOpenid  string `json:"sp_openid"`  // 用户服务标识
		SubOpenid string `json:"sub_openid"` // 用户子标识
	} `json:"payer"` // 支付者
	Amount struct {
		Total         int    `json:"total"`          // 总金额
		PayerTotal    int    `json:"payer_total"`    // 用户支付金额
		Currency      string `json:"currency"`       // 货币类型
		PayerCurrency string `json:"payer_currency"` // 用户支付币种
	} `json:"amount"` // 订单金额
	SceneInfo struct {
		DeviceId string `json:"device_id"` // 商户端设备号
	} `json:"scene_info"` // 场景信息
	PromotionDetail []struct {
		CouponId            string `json:"coupon_id"`            // 券ID
		Name                string `json:"name"`                 // 优惠名称
		Scope               string `json:"scope"`                // 优惠范围
		Type                string `json:"type"`                 // 优惠类型
		Amount              int    `json:"amount"`               // 优惠券面额
		StockId             string `json:"stock_id"`             // 活动ID
		WechatpayContribute int    `json:"wechatpay_contribute"` // 微信出资
		MerchantContribute  int    `json:"merchant_contribute"`  // 商户出资
		OtherContribute     int    `json:"other_contribute"`     // 其他出资
		Currency            string `json:"currency"`             // 优惠币种
		GoodsDetail         []struct {
			GoodsId        string `json:"goods_id"`        // 商品编码
			Quantity       int    `json:"quantity"`        // 商品数量
			UnitPrice      int    `json:"unit_price"`      // 商品单价
			DiscountAmount int    `json:"discount_amount"` // 商品优惠金额
			GoodsRemark    string `json:"goods_remark"`    // 商品备注
		} `json:"goods_detail"` // 单品列表
	} `json:"promotion_detail"` // 优惠功能
}

type PayPartnerTransactionsOutTradeNoResult struct {
	Result   PayPartnerTransactionsOutTradeNoResponse // 结果
	Body     []byte                                   // 内容
	Http     gorequest.Response                       // 请求
	Err      error                                    // 错误
	ApiError ApiError                                 // 接口错误
}

func newPayPartnerTransactionsOutTradeNoResult(result PayPartnerTransactionsOutTradeNoResponse, body []byte, http gorequest.Response, err error, apiError ApiError) *PayPartnerTransactionsOutTradeNoResult {
	return &PayPartnerTransactionsOutTradeNoResult{Result: result, Body: body, Http: http, Err: err, ApiError: apiError}
}

// PayPartnerTransactionsOutTradeNo 商户订单号查询
// https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter4_5_2.shtml
func (c *Client) PayPartnerTransactionsOutTradeNo(outTradeNo string) *PayPartnerTransactionsOutTradeNoResult {
	// 参数
	params := gorequest.NewParams()
	// 请求
	request, err := c.request(fmt.Sprintf(apiUrl+"/v3/pay/partner/transactions/out-trade-no/%s?sp_mchid=%s&sub_mchid=%s", outTradeNo, c.config.SpMchId, c.config.SubMchId), params, http.MethodGet)
	if err != nil {
		return newPayPartnerTransactionsOutTradeNoResult(PayPartnerTransactionsOutTradeNoResponse{}, request.ResponseBody, request, err, ApiError{})
	}
	// 定义
	var response PayPartnerTransactionsOutTradeNoResponse
	err = json.Unmarshal(request.ResponseBody, &response)
	// 错误
	var apiError ApiError
	err = json.Unmarshal(request.ResponseBody, &apiError)
	return newPayPartnerTransactionsOutTradeNoResult(response, request.ResponseBody, request, err, apiError)
}
