package wechatpayopen

import (
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
)

// PayPartnerTransactionsJsapiNotifyGinRequest JSAPI下单 - 回调通知 - 请求参数
type PayPartnerTransactionsJsapiNotifyGinRequest struct {
	Id           string `form:"id" json:"status" xml:"id" uri:"id" binding:"required"`                                         // 通知ID
	CreateTime   string `form:"create_time" json:"create_time" xml:"create_time" uri:"create_time" binding:"required"`         // 通知创建时间
	EventType    string `form:"event_type" json:"event_type" xml:"event_type" uri:"event_type" binding:"required"`             // 通知类型
	ResourceType string `form:"resource_type" json:"resource_type" xml:"resource_type" uri:"resource_type" binding:"required"` // 通知数据类型
	Resource     struct {
		Algorithm      string `form:"algorithm" json:"algorithm" xml:"algorithm" uri:"algorithm" binding:"required"`                          // 加密算法类型
		Ciphertext     string `form:"ciphertext" json:"ciphertext" xml:"ciphertext" uri:"ciphertext" binding:"required"`                      // 数据密文
		AssociatedData string `form:"associated_data" json:"associated_data" xml:"associated_data" uri:"associated_data" binding:"omitempty"` // 附加数据
		OriginalType   string `form:"original_type" json:"original_type" xml:"original_type" uri:"original_type" binding:"required"`          // 原始类型
		Nonce          string `form:"nonce" json:"nonce" xml:"nonce" uri:"nonce" binding:"required"`                                          // 随机串
	} `form:"resource" json:"resource" xml:"resource" uri:"resource" binding:"required"` // 通知数据
	Summary string `form:"summary" json:"summary" xml:"summary" uri:"summary" binding:"required"` // 回调摘要
}

// PayPartnerTransactionsJsapiNotifyGin JSAPI下单 - 回调通知
// https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter4_1_5.shtml
func (c *Client) PayPartnerTransactionsJsapiNotifyGin(ctx context.Context, ginCtx *gin.Context) (validateJson PayPartnerTransactionsJsapiNotifyGinRequest, response PayPartnerTransactionsJsapiNotifyGinResponse, gcm []byte, err error) {

	// 解析
	err = ginCtx.ShouldBind(&validateJson)

	gcm, err = c.decryptGCM(c.GetApiV3(), validateJson.Resource.Nonce, validateJson.Resource.Ciphertext, validateJson.Resource.AssociatedData)
	if err != nil {
		return validateJson, response, gcm, err
	}

	err = json.Unmarshal(gcm, &response)

	return validateJson, response, gcm, err
}

// PayPartnerTransactionsJsapiNotifyGinResponse JSAPI下单 - 回调通知 - 解密后数据
type PayPartnerTransactionsJsapiNotifyGinResponse struct {
	SpAppid        string `json:"sp_appid"`            // 服务商应用ID
	SpMchid        string `json:"sp_mchid"`            // 服务商户号
	SubAppid       string `json:"sub_appid,omitempty"` // 子商户应用ID
	SubMchid       string `json:"sub_mchid"`           // 子商户号
	OutTradeNo     string `json:"out_trade_no"`        // 商户订单号
	TransactionId  string `json:"transaction_id"`      // 微信支付订单号
	TradeType      string `json:"trade_type"`          // 交易类型
	TradeState     string `json:"trade_state"`         // 交易状态
	TradeStateDesc string `json:"trade_state_desc"`    // 交易状态描述
	BankType       string `json:"bank_type"`           // 付款银行
	Attach         string `json:"attach,omitempty"`    // 附加数据
	SuccessTime    string `json:"success_time"`        // 支付完成时间
	Payer          struct {
		Openid    string `json:"openid"`               // 用户服务标识
		SpOpenid  string `json:"sp_openid,omitempty"`  // 用户服务标识
		SubOpenid string `json:"sub_openid,omitempty"` // 用户子标识
	} `json:"payer"` // -支付者
	Amount struct {
		Total         int    `json:"total"`          // 总金额
		PayerTotal    int    `json:"payer_total"`    // 用户支付金额
		Currency      string `json:"currency"`       // 货币类型
		PayerCurrency string `json:"payer_currency"` // 用户支付币种
	} `json:"amount"` // 订单金额
	SceneInfo struct {
		DeviceId string `json:"device_id,omitempty"` //商户端设备号
	} `json:"scene_info,omitempty"` // 场景信息
	PromotionDetail []struct {
		CouponId            string `json:"coupon_id"`                      // 券ID
		Name                string `json:"name,omitempty"`                 // 优惠名称
		Scope               string `json:"scope,omitempty"`                // 优惠范围
		Type                string `json:"type,omitempty"`                 // 优惠类型
		Amount              int    `json:"amount"`                         // 优惠券面额
		StockId             string `json:"stock_id,omitempty"`             // 活动ID
		WechatpayContribute int    `json:"wechatpay_contribute,omitempty"` // 微信出资
		MerchantContribute  int    `json:"merchant_contribute,omitempty"`  // 商户出资
		OtherContribute     int    `json:"other_contribute,omitempty"`     // 其他出资
		Currency            string `json:"currency,omitempty"`             // 优惠币种
		GoodsDetail         []struct {
			GoodsId        string `json:"goods_id"`               // 商品编码
			Quantity       int    `json:"quantity"`               // 商品数量
			UnitPrice      int    `json:"unit_price"`             // 商品单价
			DiscountAmount int    `json:"discount_amount"`        // 商品优惠金额
			GoodsRemark    string `json:"goods_remark,omitempty"` // 商品备注
		} `json:"goods_detail,omitempty"` // 单品列表
	} `json:"promotion_detail,omitempty"` // 优惠功能
}
