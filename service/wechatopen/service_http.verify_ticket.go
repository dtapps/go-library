package wechatopen

import (
	"context"
	"encoding/base64"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/mitchellh/mapstructure"
	"go.dtapp.net/library/utils/gorequest"
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

	if haveSignature == "" {
		return resp, fmt.Errorf("找不到签名参数")
	}

	if timestamp == "" {
		return resp, fmt.Errorf("找不到时间戳参数")
	}

	if nonce == "" {
		return resp, fmt.Errorf("未找到随机数参数")
	}

	wantSignature = Sign(c.GetMessageToken(), timestamp, nonce)
	if haveSignature != wantSignature {
		return resp, fmt.Errorf("签名错误")
	}

	// 进入事件执行
	if encryptType != "aes" {
		return resp, fmt.Errorf("未知的加密类型: %s", encryptType)
	}
	if haveMsgSignature == "" {
		return resp, fmt.Errorf("找不到签名参数")
	}

	data, err = io.ReadAll(r.Body)
	if err != nil {
		return resp, err
	}

	xmlDecode := gorequest.XmlDecodeNoError(data)
	if len(xmlDecode) <= 0 {
		return resp, fmt.Errorf("xml解码错误：%s", xmlDecode)
	}

	err = mapstructure.Decode(xmlDecode, &requestHttpBody)
	if err != nil {
		return resp, fmt.Errorf("mapstructure 解码错误：%s", xmlDecode)
	}

	if requestHttpBody.Encrypt == "" {
		return resp, fmt.Errorf("未找到加密数据：%s", requestHttpBody)
	}

	cipherData, err := base64.StdEncoding.DecodeString(requestHttpBody.Encrypt)
	if err != nil {
		return resp, fmt.Errorf("encrypt 解码字符串错误：%v", err)
	}

	AesKey, err := base64.StdEncoding.DecodeString(c.GetMessageKey() + "=")
	if err != nil {
		return resp, fmt.Errorf("messageKey 解码字符串错误：%v", err)
	}

	msg, err := AesDecrypt(cipherData, AesKey)
	if err != nil {
		return resp, fmt.Errorf("AES解密错误：%v", err)
	}

	str := string(msg)

	left := strings.Index(str, "<xml>")
	if left <= 0 {
		return resp, fmt.Errorf("匹配不到<xml>：%v", left)
	}
	right := strings.Index(str, "</xml>")
	if right <= 0 {
		return resp, fmt.Errorf("匹配不到</xml>：%v", right)
	}
	msgStr := str[left:right]
	if len(msgStr) == 0 {
		return resp, fmt.Errorf("提取错误：%v", msgStr)
	}

	resp = &ResponseServeHttpVerifyTicket{}
	err = xml.Unmarshal([]byte(msgStr+"</xml>"), resp)
	if err != nil {
		return resp, fmt.Errorf("解析错误：%v", err)
	}

	return resp, nil
}
