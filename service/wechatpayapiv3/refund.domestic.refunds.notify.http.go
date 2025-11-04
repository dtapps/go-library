package wechatpayapiv3

import (
	"context"
	"encoding/json"
	"encoding/xml"
	"net/http"
)

// RefundDomesticRefundsNotifyHttpRequest 申请退款API - 回调通知 - 请求参数
type RefundDomesticRefundsNotifyHttpRequest struct {
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

// RefundDomesticRefundsNotifyHttp 退款结果回调通知
// https://pay.weixin.qq.com/doc/v3/merchant/4012791906
func (c *Client) RefundDomesticRefundsNotifyHttp(ctx context.Context, w http.ResponseWriter, r *http.Request) (validateJson RefundDomesticRefundsNotifyHttpRequest, response RefundDomesticRefundsNotifyHttpResponse, gcm []byte, err error) {

	// 解析
	_ = xml.NewDecoder(r.Body).Decode(&validateJson)

	gcm, err = c.decryptGCM(c.GetApiV3(), validateJson.Resource.Nonce, validateJson.Resource.Ciphertext, validateJson.Resource.AssociatedData)
	if err != nil {
		return validateJson, response, gcm, err
	}

	err = json.Unmarshal(gcm, &response)
	return validateJson, response, gcm, err
}

// RefundDomesticRefundsNotifyHttpResponse 申请退款API - 回调通知 - 解密后数据
type RefundDomesticRefundsNotifyHttpResponse struct {
	Mchid               string `json:"mchid"`                 // 直连商户号
	OutTradeNo          string `json:"out_trade_no"`          // 商户订单号
	TransactionId       string `json:"transaction_id"`        // 微信支付订单号
	OutRefundNo         string `json:"out_refund_no"`         // 商户退款单号
	RefundId            string `json:"refund_id"`             // 微信支付退款单号
	RefundStatus        string `json:"refund_status"`         // 退款状态
	SuccessTime         string `json:"success_time"`          // 退款成功时间
	UserReceivedAccount string `json:"user_received_account"` // 退款入账账户
	Amount              struct {
		Total       int `json:"total"`        // 订单金额
		Refund      int `json:"refund"`       // 退款金额
		PayerTotal  int `json:"payer_total"`  // 用户支付金额
		PayerRefund int `json:"payer_refund"` // 用户退款金额
	} `json:"amount"` // 金额信息
}
