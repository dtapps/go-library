package wechatopen

import (
	"context"
	"encoding/base64"
	"encoding/xml"
	"errors"
	"fmt"
	"github.com/mitchellh/mapstructure"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
	"io"
	"net/http"
	"strings"
)

// ResponseServeHttpVerifyTicket 验证票据推送
type ResponseServeHttpVerifyTicket struct {
	XMLName               xml.Name
	AppId                 string `xml:"appId" json:"appId"`                                 // 第三方平台 appid
	CreateTime            int64  `xml:"CreateTime" json:"CreateTime"`                       // 时间戳，单位：s
	InfoType              string `xml:"InfoType" json:"InfoType"`                           // 固定为："component_verify_ticket"
	ComponentVerifyTicket string `xml:"ComponentVerifyTicket" json:"ComponentVerifyTicket"` // Ticket 内容
}

type cipherRequestHttpBody struct {
	AppId   string `xml:"appId" json:"appId"`     // 第三方平台 appid
	Encrypt string `xml:"Encrypt" json:"Encrypt"` // 加密内容
}

// ServeHttpVerifyTicket 验证票据推送
func (c *Client) ServeHttpVerifyTicket(ctx context.Context, w http.ResponseWriter, r *http.Request) (resp *ResponseServeHttpVerifyTicket, err error) {

	// OpenTelemetry链路追踪
	ctx, span := TraceStartSpan(ctx, "ServeHttpVerifyTicket")
	defer span.End()

	var (
		query = r.URL.Query()

		wantSignature string
		haveSignature = query.Get("signature")
		timestamp     = query.Get("timestamp")
		nonce         = query.Get("nonce")

		// post
		haveMsgSignature = query.Get("msg_signature")
		encryptType      = query.Get("encrypt_type")

		// handle vars
		data            []byte
		requestHttpBody = &cipherRequestHttpBody{}
	)

	span.SetAttributes(attribute.String("http.query.signature", haveSignature))
	span.SetAttributes(attribute.String("http.query.timestamp", timestamp))
	span.SetAttributes(attribute.String("http.query.nonce", nonce))
	span.SetAttributes(attribute.String("http.query.msg_signature", haveMsgSignature))
	span.SetAttributes(attribute.String("http.query.encrypt_type", encryptType))

	if haveSignature == "" {
		err = errors.New("找不到签名参数")
		span.RecordError(err, trace.WithStackTrace(true))
		span.SetStatus(codes.Error, err.Error())
		return
	}

	if timestamp == "" {
		err = errors.New("找不到时间戳参数")
		span.RecordError(err, trace.WithStackTrace(true))
		span.SetStatus(codes.Error, err.Error())
		return resp, err
	}

	if nonce == "" {
		err = errors.New("未找到随机数参数")
		span.RecordError(err, trace.WithStackTrace(true))
		span.SetStatus(codes.Error, err.Error())
		return resp, err
	}

	wantSignature = Sign(c.GetMessageToken(), timestamp, nonce)
	if haveSignature != wantSignature {
		err = errors.New("签名错误")
		span.RecordError(err, trace.WithStackTrace(true))
		span.SetStatus(codes.Error, err.Error())
		return resp, err
	}

	// 进入事件执行
	if encryptType != "aes" {
		err = fmt.Errorf("未知的加密类型: %s", encryptType)
		span.RecordError(err, trace.WithStackTrace(true))
		span.SetStatus(codes.Error, err.Error())
		return resp, err
	}
	if haveMsgSignature == "" {
		err = errors.New("找不到签名参数")
		span.RecordError(err, trace.WithStackTrace(true))
		span.SetStatus(codes.Error, err.Error())
		return resp, err
	}

	data, err = io.ReadAll(r.Body)
	if err != nil {
		span.RecordError(err, trace.WithStackTrace(true))
		span.SetStatus(codes.Error, err.Error())
		return resp, err
	}

	xmlDecode := XmlDecode(string(data))
	if len(xmlDecode) <= 0 {
		err = fmt.Errorf("Xml解码错误：%s", xmlDecode)
		span.RecordError(err, trace.WithStackTrace(true))
		span.SetStatus(codes.Error, err.Error())
		return resp, err
	}

	err = mapstructure.Decode(xmlDecode, &requestHttpBody)
	if err != nil {
		err = fmt.Errorf("mapstructure 解码错误：%s", xmlDecode)
		span.RecordError(err, trace.WithStackTrace(true))
		span.SetStatus(codes.Error, err.Error())
		return resp, err
	}

	if requestHttpBody.Encrypt == "" {
		err = fmt.Errorf("未找到加密数据：%s", requestHttpBody)
		span.RecordError(err, trace.WithStackTrace(true))
		span.SetStatus(codes.Error, err.Error())
		return resp, err
	}

	cipherData, err := base64.StdEncoding.DecodeString(requestHttpBody.Encrypt)
	if err != nil {
		err = fmt.Errorf("encrypt 解码字符串错误：%v", err)
		span.RecordError(err, trace.WithStackTrace(true))
		span.SetStatus(codes.Error, err.Error())
		return resp, err
	}

	AesKey, err := base64.StdEncoding.DecodeString(c.GetMessageKey() + "=")
	if err != nil {
		err = fmt.Errorf("messageKey 解码字符串错误：%v", err)
		span.RecordError(err, trace.WithStackTrace(true))
		span.SetStatus(codes.Error, err.Error())
		return resp, err
	}

	msg, err := AesDecrypt(cipherData, AesKey)
	if err != nil {
		err = fmt.Errorf("AES解密错误：%v", err)
		span.RecordError(err, trace.WithStackTrace(true))
		span.SetStatus(codes.Error, err.Error())
		return resp, err
	}

	str := string(msg)

	left := strings.Index(str, "<xml>")
	if left <= 0 {
		err = fmt.Errorf("匹配不到<xml>：%v", left)
		span.RecordError(err, trace.WithStackTrace(true))
		span.SetStatus(codes.Error, err.Error())
		return resp, err
	}
	right := strings.Index(str, "</xml>")
	if right <= 0 {
		err = fmt.Errorf("匹配不到</xml>：%v", right)
		span.RecordError(err, trace.WithStackTrace(true))
		span.SetStatus(codes.Error, err.Error())
		return resp, err
	}
	msgStr := str[left:right]
	if len(msgStr) == 0 {
		err = fmt.Errorf("提取错误：%v", msgStr)
		span.RecordError(err, trace.WithStackTrace(true))
		span.SetStatus(codes.Error, err.Error())
		return resp, err
	}

	resp = &ResponseServeHttpVerifyTicket{}
	err = xml.Unmarshal([]byte(msgStr+"</xml>"), resp)
	if err != nil {
		err = fmt.Errorf("解析错误：%v", err)
		span.RecordError(err, trace.WithStackTrace(true))
		span.SetStatus(codes.Error, err.Error())
		return resp, err
	}

	return resp, nil
}
