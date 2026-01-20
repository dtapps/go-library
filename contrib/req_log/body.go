package req_log

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

	ContentTypeHeader = "Content-Type"
)

// processBodyAny 处理任意类型的 Body 并转换为 json.RawMessage
func (m *LoggerMiddleware) processBodyAny(contentType string, body any) json.RawMessage {

	// 开启调试模式时
	if m.debug {
		fmt.Println("[processBodyAny] contentType:", contentType)
		fmt.Println("[processBodyAny] body:", body)
		fmt.Printf("[processBodyAny] body type: %T\n", body)
	}

	if body == nil {
		return nil
	}

	switch v := body.(type) {
	case []byte:
		return m.processBodyByte(contentType, v)
	case string:
		return m.processBodyByte(contentType, []byte(v))
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
func (m *LoggerMiddleware) processBodyByte(contentType string, data []byte) json.RawMessage {

	// 开启调试模式时
	if m.debug {
		fmt.Println("[processBodyByte] contentType:", contentType)
		fmt.Println("[processBodyByte] body:", string(data))
	}

	if len(data) == 0 {
		return nil
	}

	bodyType := m.detectBodyType(contentType, data)
	switch bodyType {
	case JSONBodyType:
		if m.isValidJSON(data) {
			return data
		}
	case XMLBodyType:
		if m.isValidXML(data) {
			xj, _ := xml2json.Convert(strings.NewReader(string(data)))
			return xj.Bytes()
		}
	}

	return nil
}

// 判断是否为 JSON 格式
func (m *LoggerMiddleware) isValidJSON(data []byte) bool {
	var js json.RawMessage
	return json.Unmarshal(data, &js) == nil
}

// 判断是否为 XML 格式
func (m *LoggerMiddleware) isValidXML(data []byte) bool {
	var v any
	return xml.Unmarshal(data, &v) == nil
}

// 根据 Content-Type 或内容判断 body 类型
func (m *LoggerMiddleware) detectBodyType(contentType string, data []byte) string {
	if strings.Contains(contentType, "application/json") || strings.HasPrefix(string(data), "{") || strings.HasPrefix(string(data), "[") {
		return JSONBodyType
	}
	if strings.Contains(contentType, "xml") || strings.Contains(contentType, "soap+xml") || strings.HasPrefix(string(data), "<") {
		return XMLBodyType
	}
	if m.isValidJSON(data) {
		return JSONBodyType
	}
	if m.isValidXML(data) {
		return XMLBodyType
	}
	return TextBodyType
}
