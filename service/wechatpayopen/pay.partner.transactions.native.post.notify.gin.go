package wechatpayopen

import (
	"context"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"time"
)

// PayPartnerTransactionsNativePostNotifyGinRequest 支付通知API - 请求参数
type PayPartnerTransactionsNativePostNotifyGinRequest struct {
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

// PayPartnerTransactionsNativePostNotifyGin 支付通知API
// https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter4_4_5.shtml
func (c *Client) PayPartnerTransactionsNativePostNotifyGin(ctx context.Context, ginCtx *gin.Context) (validateJson PayPartnerTransactionsNativePostNotifyGinRequest, response PayPartnerTransactionsNativePostNotifyGinResponse, gcm []byte, err error) {

	// 解析
	err = ginCtx.ShouldBind(&validateJson)

	gcm, err = c.decryptGCM(c.GetApiV3(), validateJson.Resource.Nonce, validateJson.Resource.Ciphertext, validateJson.Resource.AssociatedData)
	if err != nil {
		return validateJson, response, gcm, err
	}

	err = json.Unmarshal(gcm, &response)

	return validateJson, response, gcm, err
}

// PayPartnerTransactionsNativePostNotifyGinResponse 支付通知API - 解密后数据
type PayPartnerTransactionsNativePostNotifyGinResponse struct {
	SpAppid        string    `json:"sp_appid"`         // 服务商应用ID
	SpMchid        string    `json:"sp_mchid"`         // 服务商户号
	SubAppid       string    `json:"sub_appid"`        // 子商户应用ID
	SubMchid       string    `json:"sub_mchid"`        // 子商户号
	OutTradeNo     string    `json:"out_trade_no"`     // 商户订单号
	TradeStateDesc string    `json:"trade_state_desc"` // 交易状态描述
	TradeType      string    `json:"trade_type"`       // 交易类型
	Attach         string    `json:"attach"`           // 附加数据
	TransactionId  string    `json:"transaction_id"`   // 微信支付订单号
	TradeState     string    `json:"trade_state"`      // 交易状态
	BankType       string    `json:"bank_type"`        // 付款银行
	SuccessTime    time.Time `json:"success_time"`     // 支付完成时间
	Amount         struct {
		Total         int    `json:"total"`          // 总金额
		PayerTotal    int    `json:"payer_total"`    // 用户支付金额
		Currency      string `json:"currency"`       // 货币类型
		PayerCurrency string `json:"payer_currency"` // 用户支付币种
	} `json:"amount"` // 订单金额
	Payer struct {
		SpOpenid  string `json:"sp_openid"`  // 用户服务标识
		SubOpenid string `json:"sub_openid"` // 用户子标识
	} `json:"payer"` // 支付者
	SceneInfo struct {
		DeviceId string `json:"device_id"` // 商户端设备号
	} `json:"scene_info"` // 场景信息
}
