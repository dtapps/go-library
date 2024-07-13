package wechatpayopen

import (
	"context"
	"encoding/xml"
	"go.dtapp.net/library/utils/gojson"
	"go.opentelemetry.io/otel/codes"
	"net/http"
	"time"
)

// PayPartnerTransactionsNativePostNotifyHttpRequest 支付通知API - 请求参数
type PayPartnerTransactionsNativePostNotifyHttpRequest struct {
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

// PayPartnerTransactionsNativePostNotifyHttp 支付通知API
// https://pay.weixin.qq.com/wiki/doc/apiv3_partner/apis/chapter4_4_5.shtml
func (c *Client) PayPartnerTransactionsNativePostNotifyHttp(ctx context.Context, w http.ResponseWriter, r *http.Request) (validateXml PayPartnerTransactionsNativePostNotifyHttpRequest, response PayPartnerTransactionsNativePostNotifyHttpResponse, gcm []byte, err error) {

	// OpenTelemetry链路追踪
	ctx = c.TraceStartSpan(ctx, "PayPartnerTransactionsNativePostNotifyHttp")
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

// PayPartnerTransactionsNativePostNotifyHttpResponse 支付通知API - 解密后数据
type PayPartnerTransactionsNativePostNotifyHttpResponse struct {
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
