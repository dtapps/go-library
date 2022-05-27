package gorequest

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
)

// MarshalXML 结构体转xml
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
