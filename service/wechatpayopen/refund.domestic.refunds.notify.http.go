package wechatpayopen

import (
	"context"
	"encoding/xml"
	"go.dtapp.net/library/utils/gojson"
	"go.opentelemetry.io/otel/codes"
	"net/http"
)

// RefundDomesticRefundsNoNotifyHttpRequest 申请退款API - 回调通知 - 请求参数
type RefundDomesticRefundsNoNotifyHttpRequest struct {
	Id           string `json:"status" xml:"id"`                   // 通知ID
	CreateTime   string `json:"create_time" xml:"create_time"`     // 通知创建时间
	EventType    string `json:"event_type" xml:"event_type"`       // 通知类型
	Summary      string `json:"summary" xml:"summary"`             // 通知简要说明
	ResourceType string `json:"resource_type" xml:"resource_type"` // 通知数据类型
	Resource     struct {
		Algorithm      string `json:"algorithm" xml:"algorithm"`                                 // 加密算法类型
		Ciphertext     string `json:"ciphertext" xml:"ciphertext"`                               // 数据密文
		AssociatedData string `json:"associated_data,omitempty" xml:"associated_data,omitempty"` // 附加数据
		OriginalType   string `json:"original_type" xml:"original_type"`                         // 原始类型
		Nonce          string `json:"nonce" xml:"nonce"`                                         // 随机串
	} `json:"resource" xml:"resource"` // 通知数据
}

// RefundDomesticRefundsNoNotifyHttp 申请退款API - 回调通知
// https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter4_1_11.shtml
func (c *Client) RefundDomesticRefundsNoNotifyHttp(ctx context.Context, w http.ResponseWriter, r *http.Request) (validateXml RefundDomesticRefundsNoNotifyHttpRequest, response RefundDomesticRefundsNoNotifyHttpResponse, gcm []byte, err error) {

	// OpenTelemetry链路追踪
	ctx = c.TraceStartSpan(ctx, "RefundDomesticRefundsNoNotifyHttp")
	defer c.TraceEndSpan()

	// 解析
	err = xml.NewDecoder(r.Body).Decode(&validateXml)
	if err != nil {
		c.TraceRecordError(err)
		c.TraceSetStatus(codes.Error, err.Error())
	}

	gcm, err = c.decryptGCM(c.GetApiV3(), validateXml.Resource.Nonce, validateXml.Resource.Ciphertext, validateXml.Resource.AssociatedData)
	if err != nil {
		c.TraceRecordError(err)
		c.TraceSetStatus(codes.Error, err.Error())
		return validateXml, response, gcm, err
	}

	err = gojson.Unmarshal(gcm, &response)
	if err != nil {
		c.TraceRecordError(err)
		c.TraceSetStatus(codes.Error, err.Error())
	}
	return validateXml, response, gcm, err
}

// RefundDomesticRefundsNoNotifyHttpResponse 申请退款API - 回调通知 - 解密后数据
type RefundDomesticRefundsNoNotifyHttpResponse struct {
	SpMchid             string `json:"sp_mchid"`              // 服务商户号
	SubMchid            string `json:"sub_mchid"`             // 子商户号
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
