package http_log

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"strings"

	xml2json "github.com/basgys/goxml2json"
)

const (
	JSONBodyType = "json"
	XMLBodyType  = "xml"
	TextBodyType = "text"
)

// processBodyAny 处理任意类型的 Body 并转换为 json.RawMessage
func (l *LoggingRoundTripper) processBodyAny(contentType string, body any) json.RawMessage {

	// 开启调试模式时
	if l.debug {
		fmt.Println("[processBodyAny] contentType:", contentType)
		fmt.Println("[processBodyAny] body:", body)
		fmt.Printf("[processBodyAny] body type: %T\n", body)
	}

	if body == nil {
		return nil
	}

	switch v := body.(type) {
	case []byte:
		return l.processBodyByte(contentType, v)
	case string:
		return l.processBodyByte(contentType, []byte(v))
	case json.RawMessage:
		return v
	default:
		// 如果是结构体、Map 或切片，尝试直接序列化为 JSON
		data, err := json.Marshal(v)
		if err != nil {
			// 如果序列化失败，返回 nil 或根据需要记录错误
			return nil
		}
		return data
	}
}

// 处理 body，根据类型存到 JSON 或 XML 字段
func (l *LoggingRoundTripper) processBodyByte(contentType string, data []byte) json.RawMessage {

	// 开启调试模式时
	if l.debug {
		fmt.Println("[processBodyByte] contentType:", contentType)
		fmt.Println("[processBodyByte] body:", string(data))
	}

	if len(data) == 0 {
		return nil
	}

	bodyType := l.detectBodyType(contentType, data)
	switch bodyType {
	case JSONBodyType:
		if l.isValidJSON(data) {
			return data
		}
	case XMLBodyType:
		if l.isValidXML(data) {
			xj, _ := xml2json.Convert(strings.NewReader(string(data)))
			return xj.Bytes()
		}
	}

	return nil
}

// 判断是否为 JSON 格式
func (l *LoggingRoundTripper) isValidJSON(data []byte) bool {
	var js json.RawMessage
	return json.Unmarshal(data, &js) == nil
}

// 判断是否为 XML 格式
func (l *LoggingRoundTripper) isValidXML(data []byte) bool {
	var v any
	return xml.Unmarshal(data, &v) == nil
}

// 根据 Content-Type 或内容判断 body 类型
func (l *LoggingRoundTripper) detectBodyType(contentType string, data []byte) string {
	if strings.Contains(contentType, "application/json") || strings.HasPrefix(string(data), "{") || strings.HasPrefix(string(data), "[") {
		return JSONBodyType
	}
	if strings.Contains(contentType, "xml") || strings.Contains(contentType, "soap+xml") || strings.HasPrefix(string(data), "<") {
		return XMLBodyType
	}
	if l.isValidJSON(data) {
		return JSONBodyType
	}
	if l.isValidXML(data) {
		return XMLBodyType
	}
	return TextBodyType
}
