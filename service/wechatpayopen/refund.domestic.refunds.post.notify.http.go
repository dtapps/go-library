package wechatpayopen

import (
	"context"
	"encoding/xml"
	"go.dtapp.net/library/utils/gojson"
	"net/http"
	"time"
)

// RefundDomesticRefundsPostNotifyHttpRequest 退款结果通知API - 请求参数
type RefundDomesticRefundsPostNotifyHttpRequest struct {
	Id           string `json:"status" xml:"id"`                   // 通知ID
	CreateTime   string `json:"create_time" xml:"create_time"`     // 通知创建时间
	EventType    string `json:"event_type" xml:"event_type"`       // 通知类型
	ResourceType string `json:"resource_type" xml:"resource_type"` // 通知数据类型
	Resource     struct {
		Algorithm      string `json:"algorithm" xml:"algorithm"`                                 // 加密算法类型
		Ciphertext     string `json:"ciphertext" xml:"ciphertext"`                               // 数据密文
		AssociatedData string `json:"associated_data,omitempty" xml:"associated_data,omitempty"` // 附加数据
		OriginalType   string `json:"original_type" xml:"original_type"`                         // 原始类型
		Nonce          string `json:"nonce" xml:"nonce"`                                         // 随机串
	} `json:"resource" xml:"resource"` // 通知数据
	Summary string `json:"summary" xml:"summary"` // 回调摘要
}

// RefundDomesticRefundsPostNotifyHttp 退款结果通知API
// https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter4_4_11.shtml
func (c *Client) RefundDomesticRefundsPostNotifyHttp(ctx context.Context, w http.ResponseWriter, r *http.Request) (validateXml RefundDomesticRefundsPostNotifyHttpRequest, response RefundDomesticRefundsPostNotifyHttpResponse, gcm []byte, err error) {

	// 解析
	_ = xml.NewDecoder(r.Body).Decode(&validateXml)

	gcm, err = c.decryptGCM(c.GetApiV3(), validateXml.Resource.Nonce, validateXml.Resource.Ciphertext, validateXml.Resource.AssociatedData)
	if err != nil {
		return validateXml, response, gcm, err
	}

	err = gojson.Unmarshal(gcm, &response)
	return validateXml, response, gcm, err
}

// RefundDomesticRefundsPostNotifyHttpResponse 退款结果通知API - 解密后数据
type RefundDomesticRefundsPostNotifyHttpResponse struct {
	SpMchid             string    `json:"sp_mchid"`              // 服务商户号
	SubMchid            string    `json:"sub_mchid"`             // 子商户号
	TransactionId       string    `json:"transaction_id"`        // 商户订单号
	OutTradeNo          string    `json:"out_trade_no"`          // 微信支付订单号
	RefundId            string    `json:"refund_id"`             // 商户退款单号
	OutRefundNo         string    `json:"out_refund_no"`         // 微信支付退款单号
	RefundStatus        string    `json:"refund_status"`         // 退款状态
	SuccessTime         time.Time `json:"success_time"`          // 退款成功时间
	UserReceivedAccount string    `json:"user_received_account"` // 退款入账账户
	Amount              struct {
		Total       int `json:"total"`        // 订单金额
		Refund      int `json:"refund"`       // 退款金额
		PayerTotal  int `json:"payer_total"`  // 用户支付金额
		PayerRefund int `json:"payer_refund"` // 用户退款金额
	} `json:"amount"` // 金额信息
}
