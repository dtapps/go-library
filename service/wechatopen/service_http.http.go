package wechatopen

import (
	"context"
	"encoding/xml"
	"net/http"
)

// ResponseServeHttpHttp 推送信息
type ResponseServeHttpHttp struct {
	MsgSignature string `json:"msg_signature"` // 签名串，对应 URL 参数的msg_signature
	Timestamp    string `json:"timestamp"`     // 时间戳，对应 URL 参数的timestamp
	Nonce        string `json:"nonce"`         // 随机串，对应 URL 参数的nonce
	Signature    string `json:"signature"`
	EncryptType  string `json:"encrypt_type"` // 加密类型
	AppId        string `json:"app_id"`       // 第三方平台 appid
	Encrypt      string `json:"encrypt"`      // 加密内容
}

// ServeHttpHttp 验证票据推送
func (c *Client) ServeHttpHttp(ctx context.Context, w http.ResponseWriter, r *http.Request) (ResponseServeHttpHttp, error) {

	query := r.URL.Query()

	// 解析请求体
	var validateXml struct {
		AppId   string `json:"AppId,omitempty" xml:"AppId,omitempty"`     // 第三方平台 appid
		Encrypt string `json:"Encrypt,omitempty" xml:"Encrypt,omitempty"` // 加密内容
	}
	err := xml.NewDecoder(r.Body).Decode(&validateXml)
	if err != nil {
		return ResponseServeHttpHttp{}, err
	}

	response := ResponseServeHttpHttp{
		MsgSignature: query.Get("msg_signature"),
		Timestamp:    query.Get("timestamp"),
		Nonce:        query.Get("nonce"),
		Signature:    query.Get("signature"),
		EncryptType:  query.Get("encrypt_type"),
		AppId:        validateXml.AppId,
		Encrypt:      validateXml.Encrypt,
	}

	return response, err
}
