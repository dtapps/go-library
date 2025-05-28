package resty_extend

import (
	"encoding/json"
	"fmt"
	"io"
	"strings"

	"resty.dev/v3"
)

// Json解析
func JsonParsing(resp *resty.Response, response any) error {

	// 检查是否是 JSON 格式
	contentType := resp.Header().Get("Content-Type")
	isJSON := contentType != "" && (strings.Contains(contentType, "application/json") || strings.Contains(contentType, "+json"))
	// 如果不是 JSON，则手动解析
	if !isJSON {
		// 获取响应体字节流
		bodyBytes, err := io.ReadAll(resp.Body)
		if err != nil {
			return fmt.Errorf("读取响应体失败: %w", err)
		}

		// 手动解析 JSON
		err = json.Unmarshal(bodyBytes, response)
		if err != nil {
			return fmt.Errorf("JSON 解析失败: %w", err)
		}
	}

	return nil
}
