package wechatpayapiv2

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
)

// Params 请求参数
type Params map[string]interface{}

func NewParams() Params {
	p := make(Params)
	return p
}

func (c *Client) NewParamsWith(params ...Params) Params {
	p := make(Params)
	for _, v := range params {
		p.SetParams(v)
	}
	return p
}

func (p Params) Set(key string, value interface{}) {
	p[key] = value
}

func (p Params) SetParams(params Params) {
	for key, value := range params {
		p[key] = value
	}
}

func (p Params) MarshalXML() (reader io.Reader, err error) {
	buffer := bytes.NewBuffer(make([]byte, 0))

	if _, err = io.WriteString(buffer, "<xml>"); err != nil {
		return
	}

	for k, v := range p {
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
