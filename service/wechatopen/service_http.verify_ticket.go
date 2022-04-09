package wechatopen

import (
	"encoding/base64"
	"encoding/xml"
	"errors"
	"fmt"
	"github.com/mitchellh/mapstructure"
	"io/ioutil"
	"net/http"
	"strings"
)

// ResponseServeHttpVerifyTicket 验证票据推送
type ResponseServeHttpVerifyTicket struct {
	XMLName               xml.Name
	AppId                 string `xml:"AppId" json:"AppId"`                                 // 第三方平台 appid
	CreateTime            int64  `xml:"CreateTime" json:"CreateTime"`                       // 时间戳，单位：s
	InfoType              string `xml:"InfoType" json:"InfoType"`                           // 固定为："component_verify_ticket"
	ComponentVerifyTicket string `xml:"ComponentVerifyTicket" json:"ComponentVerifyTicket"` // Ticket 内容
}

type cipherRequestHttpBody struct {
	AppId   string `xml:"AppId" json:"AppId"`     // 第三方平台 appid
	Encrypt string `xml:"Encrypt" json:"Encrypt"` // 加密内容
}

// ServeHttpVerifyTicket 验证票据推送
func (app *App) ServeHttpVerifyTicket(r *http.Request) (resp *ResponseServeHttpVerifyTicket, err error) {
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
		err = errors.New("找不到签名参数")
		return
	}

	if timestamp == "" {
		return resp, errors.New("找不到时间戳参数")
	}

	if nonce == "" {
		return resp, errors.New("未找到随机数参数")
	}

	wantSignature = Sign(app.MessageToken, timestamp, nonce)
	if haveSignature != wantSignature {
		return resp, errors.New("签名错误")
	}

	// 进入事件执行
	if encryptType != "aes" {
		err = errors.New("未知的加密类型: " + encryptType)
		return
	}
	if haveMsgSignature == "" {
		err = errors.New("找不到签名参数")
		return
	}

	data, err = ioutil.ReadAll(r.Body)
	if err != nil {
		return resp, err
	}

	xmlDecode := XmlDecode(string(data))
	if len(xmlDecode) <= 0 {
		return resp, errors.New(fmt.Sprintf("Xml解码错误：%s", xmlDecode))
	}

	err = mapstructure.Decode(xmlDecode, &requestHttpBody)
	if err != nil {
		return resp, errors.New(fmt.Sprintf("mapstructure 解码错误：%s", xmlDecode))
	}

	if requestHttpBody.Encrypt == "" {
		return resp, errors.New(fmt.Sprintf("未找到加密数据：%s", requestHttpBody))
	}

	cipherData, err := base64.StdEncoding.DecodeString(requestHttpBody.Encrypt)
	if err != nil {
		return resp, errors.New(fmt.Sprintf("Encrypt 解码字符串错误：%v", err))
	}

	AesKey, err := base64.StdEncoding.DecodeString(app.MessageKey + "=")
	if err != nil {
		return resp, errors.New(fmt.Sprintf("MessageKey 解码字符串错误：%v", err))
	}

	msg, err := AesDecrypt(cipherData, AesKey)
	if err != nil {
		return resp, errors.New(fmt.Sprintf("AES解密错误：%v", err))
	}

	str := string(msg)

	left := strings.Index(str, "<xml>")
	if left <= 0 {
		return resp, errors.New(fmt.Sprintf("匹配不到<xml>：%v", left))
	}
	right := strings.Index(str, "</xml>")
	if right <= 0 {
		return resp, errors.New(fmt.Sprintf("匹配不到</xml>：%v", right))
	}
	msgStr := str[left:right]
	if len(msgStr) == 0 {
		return resp, errors.New(fmt.Sprintf("提取错误：%v", msgStr))
	}

	resp = &ResponseServeHttpVerifyTicket{}
	err = xml.Unmarshal([]byte(msgStr+"</xml>"), resp)
	if err != nil {
		return resp, errors.New(fmt.Sprintf("解析错误：%v", err))
	}

	return resp, nil
}
