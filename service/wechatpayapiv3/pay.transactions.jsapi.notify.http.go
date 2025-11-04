package wechatpayapiv3

import (
	"context"
	"encoding/json"
	"net/http"
)

// PayTransactionsJsapiNotifyHttpRequest JSAPI下单 - 回调通知 - 请求参数
type PayTransactionsJsapiNotifyHttpRequest struct {
	Id           string `json:"status"`        // 【通知ID】回调通知的唯一编号
	CreateTime   string `json:"create_time"`   // 【通知创建时间】
	EventType    string `json:"event_type"`    // 【通知的类型】微信支付回调通知的类型
	ResourceType string `json:"resource_type"` // 【通知数据类型】通知的资源数据类型
	Resource     struct {
		Algorithm      string `json:"algorithm"`                 // 【加密算法类型】回调数据密文的加密算法类型，目前为AEAD_AES_256_GCM，开发者需要使用同样类型的数据进行解密
		Ciphertext     string `json:"ciphertext"`                // 【数据密文】Base64编码后的回调数据密文，商户需Base64解码并使用APIV3密钥解密
		AssociatedData string `json:"associated_data,omitempty"` // 【附加数据】参与解密的附加数据，该字段可能为空
		OriginalType   string `json:"original_type"`             // 【原始回调类型】加密前的对象类型
		Nonce          string `json:"nonce"`                     // 【随机串】参与解密的随机串
	} `json:"resource"` // 【通知数据】通知资源数据
	Summary string `json:"summary"` // 【回调摘要】微信支付对回调内容的摘要备注
}

// PayTransactionsJsapiNotifyHttp 支付成功回调通知
// https://pay.weixin.qq.com/doc/v3/merchant/4012791902
func (c *Client) PayTransactionsJsapiNotifyHttp(ctx context.Context, w http.ResponseWriter, r *http.Request) (validateJson PayTransactionsJsapiNotifyHttpRequest, response PayTransactionsJsapiNotifyHttpResponse, gcm []byte, err error) {

	// 解析
	_ = json.NewDecoder(r.Body).Decode(&validateJson)

	gcm, err = c.decryptGCM(c.GetApiV3(), validateJson.Resource.Nonce, validateJson.Resource.Ciphertext, validateJson.Resource.AssociatedData)
	if err != nil {
		return validateJson, response, gcm, err
	}

	err = json.Unmarshal(gcm, &response)
	return validateJson, response, gcm, err
}

// PayTransactionsJsapiNotifyHttpResponse JSAPI下单 - 回调通知 - 解密后数据
type PayTransactionsJsapiNotifyHttpResponse struct {
	Appid          string `json:"appid"`            // 应用ID
	Mchid          string `json:"mchid"`            // 商户号
	OutTradeNo     string `json:"out_trade_no"`     // 商户订单号
	TransactionId  string `json:"transaction_id"`   // 微信支付订单号
	TradeType      string `json:"trade_type"`       // 交易类型
	TradeState     string `json:"trade_state"`      // 交易状态
	TradeStateDesc string `json:"trade_state_desc"` // 交易状态描述
	BankType       string `json:"bank_type"`        // 付款银行
	Attach         string `json:"attach,omitempty"` // 附加数据
	SuccessTime    string `json:"success_time"`     // 支付完成时间
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
