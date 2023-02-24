package wechatpayopen

import (
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"time"
)

// RefundDomesticRefundsPostNotifyGinRequest 退款结果通知API - 请求参数
type RefundDomesticRefundsPostNotifyGinRequest struct {
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

// RefundDomesticRefundsPostNotifyGin 退款结果通知API
// https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter4_4_11.shtml
func (c *Client) RefundDomesticRefundsPostNotifyGin(ctx context.Context, ginCtx *gin.Context) (validateJson RefundDomesticRefundsPostNotifyGinRequest, response RefundDomesticRefundsPostNotifyGinResponse, gcm []byte, err error) {

	// 解析
	err = ginCtx.ShouldBind(&validateJson)

	gcm, err = c.decryptGCM(c.GetApiV3(), validateJson.Resource.Nonce, validateJson.Resource.Ciphertext, validateJson.Resource.AssociatedData)
	if err != nil {
		return validateJson, response, gcm, err
	}

	err = json.Unmarshal(gcm, &response)

	return validateJson, response, gcm, err
}

// RefundDomesticRefundsPostNotifyGinResponse 退款结果通知API - 解密后数据
type RefundDomesticRefundsPostNotifyGinResponse struct {
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
