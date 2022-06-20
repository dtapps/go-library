package dorm

import (
	"encoding/json"
	"github.com/basgys/goxml2json"
	"strings"
)

// XmlDecodeNoError Xml解码，不报错
func XmlDecodeNoError(b []byte) map[string]interface{} {
	xtj := strings.NewReader(string(b))
	jtx, _ := xml2json.Convert(xtj)
	var data map[string]interface{}
	_ = json.Unmarshal(jtx.Bytes(), &data)
	return data
}
