package http_log

import (
	"bytes"
	"compress/gzip"
	"encoding/json"
	"encoding/xml"
	"fmt"
	"io"
	"net/http"
	"strings"

	xml2json "github.com/basgys/goxml2json"
)

const (
	JSONBodyType = "json"
	XMLBodyType  = "xml"
	TextBodyType = "text"

	ContentTypeHeader     = "Content-Type"
	ContentEncodingHeader = "Content-Encoding"
)

// processResponseBody 处理请求体/响应体，根据 Content-Type 转换为 JSON 或 XML
func (t *LoggingRoundTripper) processResponseBody(headers http.Header, body []byte) json.RawMessage {
	contentType := headers.Get(ContentTypeHeader)

	// 先解压（如果需要）
	if headers.Get(ContentEncodingHeader) == "gzip" {
		if gr, err := gzip.NewReader(bytes.NewReader(body)); err == nil {
			defer gr.Close()
			if decompressed, err := io.ReadAll(gr); err == nil {
				body = decompressed
			}
		}
	}

	return t.processBodyByte(contentType, body)
}

// processBodyAny 处理任意类型的 Body 并转换为 json.RawMessage
func (t *LoggingRoundTripper) processBodyAny(contentType string, body any) json.RawMessage {

	// 开启调试模式时
	if t.debug {
		fmt.Println("[processBodyAny] contentType:", contentType)
		fmt.Println("[processBodyAny] body:", body)
		fmt.Printf("[processBodyAny] body type: %T\n", body)
	}

	if body == nil {
		return nil
	}

	switch v := body.(type) {
	case []byte:
		return t.processBodyByte(contentType, v)
	case string:
		return t.processBodyByte(contentType, []byte(v))
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
func (t *LoggingRoundTripper) processBodyByte(contentType string, data []byte) json.RawMessage {

	// 开启调试模式时
	if t.debug {
		fmt.Println("[processBodyByte] contentType:", contentType)
		fmt.Println("[processBodyByte] body:", string(data))
	}

	if len(data) == 0 {
		return nil
	}

	bodyType := t.detectBodyType(contentType, data)
	switch bodyType {
	case JSONBodyType:
		if t.isValidJSON(data) {
			return data
		}
	case XMLBodyType:
		if t.isValidXML(data) {
			xj, _ := xml2json.Convert(strings.NewReader(string(data)))
			return xj.Bytes()
		}
	}

	return nil
}

// 判断是否为 JSON 格式
func (t *LoggingRoundTripper) isValidJSON(data []byte) bool {
	var js json.RawMessage
	return json.Unmarshal(data, &js) == nil
}

// 判断是否为 XML 格式
func (t *LoggingRoundTripper) isValidXML(data []byte) bool {
	var v any
	return xml.Unmarshal(data, &v) == nil
}

// 根据 Content-Type 或内容判断 body 类型
func (t *LoggingRoundTripper) detectBodyType(contentType string, data []byte) string {
	if strings.Contains(contentType, "application/json") || strings.HasPrefix(string(data), "{") || strings.HasPrefix(string(data), "[") {
		return JSONBodyType
	}
	if strings.Contains(contentType, "xml") || strings.Contains(contentType, "soap+xml") || strings.HasPrefix(string(data), "<") {
		return XMLBodyType
	}
	if t.isValidJSON(data) {
		return JSONBodyType
	}
	if t.isValidXML(data) {
		return XMLBodyType
	}
	return TextBodyType
}
