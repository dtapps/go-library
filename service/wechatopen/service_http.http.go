package wechatopen

import (
	"context"
	"encoding/xml"
	"go.dtapp.net/library/utils/gojson"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
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

	// OpenTelemetry链路追踪
	ctx, span := TraceStartSpan(ctx, "ServeHttpHttp")
	defer span.End()

	query := r.URL.Query()

	// 解析请求体
	var validateXml struct {
		AppId   string `json:"AppId,omitempty" xml:"AppId,omitempty"`     // 第三方平台 appid
		Encrypt string `json:"Encrypt,omitempty" xml:"Encrypt,omitempty"` // 加密内容
	}
	err := xml.NewDecoder(r.Body).Decode(&validateXml)
	if err != nil {
		span.RecordError(err, trace.WithStackTrace(true))
		span.SetStatus(codes.Error, err.Error())
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

	span.SetAttributes(attribute.String("http.params", gojson.JsonEncodeNoError(validateXml)))

	span.SetAttributes(attribute.String("http.query.msg_signature", response.MsgSignature))
	span.SetAttributes(attribute.String("http.query.timestamp", response.Timestamp))
	span.SetAttributes(attribute.String("http.query.nonce", response.Nonce))
	span.SetAttributes(attribute.String("http.query.signature", response.Signature))
	span.SetAttributes(attribute.String("http.query.encrypt_type", response.EncryptType))

	return response, err
}
