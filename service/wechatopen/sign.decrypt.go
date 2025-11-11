package wechatopen

import (
	"context"
	"encoding/base64"
	"encoding/xml"
	"fmt"
	"strings"
)

// SignDecrypt 解密
// ctx 上下文
// params 入参
// strXml 反射结构体
// resp 加密数据
// err 错误信息
func (c *Client) SignDecrypt(ctx context.Context, params SignDecryptParams, strXml interface{}) (resp []byte, err error) {

	if params.Signature == "" {
		return nil, fmt.Errorf("找不到签名参数")
	}

	if params.Timestamp == "" {
		return nil, fmt.Errorf("找不到时间戳参数")
	}

	if params.Nonce == "" {
		return nil, fmt.Errorf("未找到随机数参数")
	}

	wantSignature := Sign(c.GetMessageToken(), params.Timestamp, params.Nonce)
	if params.Signature != wantSignature {
		return nil, fmt.Errorf("签名错误")
	}

	// 进入事件执行
	if params.EncryptType != "aes" {
		return nil, fmt.Errorf("未知的加密类型: %s", params.EncryptType)
	}
	if params.Encrypt == "" {
		return nil, fmt.Errorf("找不到签名参数")
	}

	cipherData, err := base64.StdEncoding.DecodeString(params.Encrypt)
	if err != nil {
		return nil, fmt.Errorf("Encrypt 解码字符串错误：%v", err)
	}

	AesKey, err := base64.StdEncoding.DecodeString(c.GetMessageKey() + "=")
	if err != nil {
		return nil, fmt.Errorf("messageKey 解码字符串错误：%v", err)
	}

	msg, err := AesDecrypt(cipherData, AesKey)
	if err != nil {
		return nil, fmt.Errorf("AES解密错误：%v", err)
	}

	str := string(msg)

	left := strings.Index(str, "<xml>")
	if left <= 0 {
		return nil, fmt.Errorf("匹配不到<xml>：%v", left)
	}
	right := strings.Index(str, "</xml>")
	if right <= 0 {
		return nil, fmt.Errorf("匹配不到</xml>：%v", right)
	}
	msgStr := str[left:right]
	if len(msgStr) == 0 {
		return nil, fmt.Errorf("提取错误：%v", msgStr)
	}

	strByte := []byte(msgStr + "</xml>")
	err = xml.Unmarshal(strByte, strXml)
	if err != nil {
		return nil, fmt.Errorf("解析错误：%v", err)
	}

	return strByte, nil
}
