package wechatpayopen

import (
	"context"
	"github.com/dtapps/go-library/utils/gojson"
	"github.com/gin-gonic/gin"
)

// ProfitSharingReceiversNotifyGinRequest 分账动账通知API - 回调通知 - 请求参数
type ProfitSharingReceiversNotifyGinRequest struct {
	Id           string `form:"id" json:"status" xml:"id" uri:"id" binding:"required"`                                         // 通知ID
	CreateTime   string `form:"create_time" json:"create_time" xml:"create_time" uri:"create_time" binding:"required"`         // 通知创建时间
	EventType    string `form:"event_type" json:"event_type" xml:"event_type" uri:"event_type" binding:"required"`             // 通知类型
	Summary      string `form:"summary" json:"summary" xml:"summary" uri:"summary" binding:"required"`                         // 通知简要说明
	ResourceType string `form:"resource_type" json:"resource_type" xml:"resource_type" uri:"resource_type" binding:"required"` // 通知数据类型
	Resource     struct {
		Algorithm      string `form:"algorithm" json:"algorithm" xml:"algorithm" uri:"algorithm" binding:"required"`                          // 加密算法类型
		Ciphertext     string `form:"ciphertext" json:"ciphertext" xml:"ciphertext" uri:"ciphertext" binding:"required"`                      // 数据密文
		AssociatedData string `form:"associated_data" json:"associated_data" xml:"associated_data" uri:"associated_data" binding:"omitempty"` // 附加数据
		OriginalType   string `form:"original_type" json:"original_type" xml:"original_type" uri:"original_type" binding:"required"`          // 原始类型
		Nonce          string `form:"nonce" json:"nonce" xml:"nonce" uri:"nonce" binding:"required"`                                          // 随机串
	} `form:"resource" json:"resource" xml:"resource" uri:"resource" binding:"required"` // 通知数据
}

// ProfitSharingReceiversNotifyGin 分账动账通知API - 回调通知
// https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter4_1_5.shtml
func (c *Client) ProfitSharingReceiversNotifyGin(ctx context.Context, ginCtx *gin.Context) (validateJson ProfitSharingReceiversNotifyGinRequest, response ProfitSharingReceiversNotifyGinResponse, gcm []byte, err error) {

	// 解析
	err = ginCtx.ShouldBind(&validateJson)

	gcm, err = c.decryptGCM(c.GetApiV3(), validateJson.Resource.Nonce, validateJson.Resource.Ciphertext, validateJson.Resource.AssociatedData)
	if err != nil {
		return validateJson, response, gcm, err
	}

	err = gojson.Unmarshal(gcm, &response)

	return validateJson, response, gcm, err
}

// ProfitSharingReceiversNotifyGinResponse 分账动账通知API - 回调通知 - 解密后数据
type ProfitSharingReceiversNotifyGinResponse struct {
	SpMchid       string `json:"sp_mchid"`       // 服务商商户号
	SubMchid      string `json:"sub_mchid"`      // 子商户号
	TransactionId string `json:"transaction_id"` // 微信订单号
	OrderId       string `json:"order_id"`       // 微信分账/回退单号
	OutOrderNo    string `json:"out_order_no"`   // 商户分账/回退单号
	Receiver      struct {
		Type        string `json:"type"`        // 分账接收方类型
		Account     string `json:"account"`     // 分账接收方账号
		Amount      int    `json:"amount"`      // 分账动账金额
		Description string `json:"description"` // 分账/回退描述
	} `json:"receiver"` // 分账接收方列表
	SuccessTime string `json:"success_time"` // 成功时间
}
