package gorequest

import (
	"bytes"
	"encoding/json"
	"encoding/xml"
	"fmt"
	xj "github.com/basgys/goxml2json"
	"io"
	"strings"
)

func ToXml(params map[string]any) (reader io.Reader, err error) {
	buffer := bytes.NewBuffer(make([]byte, 0))

	if _, err = io.WriteString(buffer, "<xml>"); err != nil {
		return
	}

	for k, v := range params {
		switch {
		case k == "detail":
			if _, err = io.WriteString(buffer, "<detail><![CDATA["); err != nil {
				return
			}
			if _, err = io.WriteString(buffer, fmt.Sprintf("%v", v)); err != nil {
				return
			}
			if _, err = io.WriteString(buffer, "]]></detail>"); err != nil {
				return
			}
		default:
			if _, err = io.WriteString(buffer, "<"+k+">"); err != nil {
				return
			}
			if err = xml.EscapeText(buffer, []byte(fmt.Sprintf("%v", v))); err != nil {
				return
			}
			if _, err = io.WriteString(buffer, "</"+k+">"); err != nil {
				return
			}
		}
	}

	if _, err = io.WriteString(buffer, "</xml>"); err != nil {
		return
	}
	return buffer, nil
}

// XmlDecodeNoError xml字符串转结构体，不报错
func XmlDecodeNoError(b []byte) map[string]any {
	xtj := strings.NewReader(string(b))
	jtx, _ := xj.Convert(xtj)
	var data map[string]any
	_ = json.Unmarshal(jtx.Bytes(), &data)
	return data
}

// XmlEncodeNoError 结构体转json字符串，不报错
func XmlEncodeNoError(data any) string {
	return JsonEncodeNoError(data)
}
