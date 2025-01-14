package wechatpayopen

import (
	"context"
	"encoding/xml"
	"go.dtapp.net/library/utils/gojson"
	"net/http"
)

// ProfitSharingReceiversNotifyHttpRequest 分账动账通知API - 回调通知 - 请求参数
type ProfitSharingReceiversNotifyHttpRequest struct {
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

// ProfitSharingReceiversNotifyHttp 分账动账通知API - 回调通知
// https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter4_1_5.shtml
func (c *Client) ProfitSharingReceiversNotifyHttp(ctx context.Context, w http.ResponseWriter, r *http.Request) (validateXml ProfitSharingReceiversNotifyHttpRequest, response ProfitSharingReceiversNotifyHttpResponse, gcm []byte, err error) {

	// 解析
	_ = xml.NewDecoder(r.Body).Decode(&validateXml)

	gcm, err = c.decryptGCM(c.GetApiV3(), validateXml.Resource.Nonce, validateXml.Resource.Ciphertext, validateXml.Resource.AssociatedData)
	if err != nil {
		return validateXml, response, gcm, err
	}

	err = gojson.Unmarshal(gcm, &response)
	return validateXml, response, gcm, err
}

// ProfitSharingReceiversNotifyHttpResponse 分账动账通知API - 回调通知 - 解密后数据
type ProfitSharingReceiversNotifyHttpResponse struct {
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
