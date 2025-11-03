package wechatpayopen

import (
	"context"
	"encoding/json"
	"encoding/xml"
	"net/http"
)

// ProviderNotifyHttpRequest 支付通知API - 请求参数
type ProviderNotifyHttpRequest struct {
	Id           string `json:"status" xml:"id"`                   // 通知ID
	CreateTime   string `json:"create_time" xml:"create_time"`     // 通知创建时间
	ResourceType string `json:"resource_type" xml:"resource_type"` // 通知数据类型
	EventType    string `json:"event_type" xml:"event_type"`       // 通知类型
	Summary      string `json:"summary" xml:"summary"`             // 回调摘要
	Resource     struct {
		Algorithm      string `json:"algorithm" xml:"algorithm"`                                 // 加密算法类型
		Ciphertext     string `json:"ciphertext" xml:"ciphertext"`                               // 数据密文
		AssociatedData string `json:"associated_data,omitempty" xml:"associated_data,omitempty"` // 附加数据
		OriginalType   string `json:"original_type" xml:"original_type"`                         // 原始类型
		Nonce          string `json:"nonce" xml:"nonce"`                                         // 随机串
	} `json:"resource" xml:"resource"` // 通知数据
}

// ProviderNotifyHttp 合作伙伴订阅
// https://pay.weixin.qq.com/doc/v3/partner/4016022266
func (c *Client) ProviderNotifyHttp(ctx context.Context, w http.ResponseWriter, r *http.Request) (validateXml ProviderNotifyHttpRequest, response ProviderNotifyHttpResponse, gcm []byte, err error) {

	// 解析
	_ = xml.NewDecoder(r.Body).Decode(&validateXml)

	gcm, err = c.decryptGCM(c.GetApiV3(), validateXml.Resource.Nonce, validateXml.Resource.Ciphertext, validateXml.Resource.AssociatedData)
	if err != nil {
		return validateXml, response, gcm, err
	}

	err = json.Unmarshal(gcm, &response)
	return validateXml, response, gcm, err
}

// ProviderNotifyHttpResponse 接收订阅消息 - 解密后数据
type ProviderNotifyHttpResponse struct {
	// MessageContent 消息内容
	MessageContent struct {
		// MerchantCode 商户号
		MerchantCode string `json:"merchant_code"`
		// MerchantCompanyName 商户公司名称
		MerchantCompanyName string `json:"merchant_company_name"`
		// BusinessTime 业务时间
		BusinessTime string `json:"business_time"`
		// BusinessCode 业务代码
		BusinessCode string `json:"business_code"`
		// BusinessState 业务状态
		BusinessState string `json:"business_state"`
	} `json:"message_content"`
	// TopicName 主题名称
	TopicName struct {
		// TopicEnglishName 主题英文名称
		TopicEnglishName string `json:"topic_english_name"`
		// TopicChineseName 主题中文名称
		TopicChineseName string `json:"topic_chinese_name"`
	} `json:"topic_name"`
}
