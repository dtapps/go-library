package framework

import (
	"crypto/md5"
	"encoding/base64"
	"fmt"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

// EnsureDirExists 确保路径中的目录存在，不存在则创建
func EnsureDirExists(filePath string) error {
	dir := filepath.Dir(filePath)
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		return os.MkdirAll(dir, os.ModePerm)
	}
	return nil
}

// FileExists 判断文件是否存在
func FileExists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}

// GetFileExtFromBase64 从 base64 数据中提取 MIME 并返回对应的扩展名
func GetFileExtFromBase64(base64Str string) string {
	// 1: 提取 MIME 类型
	mimeType := "application/octet-stream"

	if strings.HasPrefix(base64Str, "data:") {
		end := strings.Index(base64Str, ",")
		if end > 0 {
			header := base64Str[:end]
			parts := strings.Split(header, ";")
			if len(parts) > 0 {
				mimeType = parts[0][5:] // 去掉 "data:" 前缀
			}
		}
	}

	// 2: 如果没提取到，则尝试用 http.DetectContentType 解析数据部分
	if mimeType == "application/octet-stream" {
		dataPart := base64Str
		if strings.Contains(base64Str, ",") {
			parts := strings.SplitN(base64Str, ",", 2)
			dataPart = parts[1]
		}

		decodedData, _ := base64.StdEncoding.DecodeString(dataPart)
		if len(decodedData) > 0 {
			mimeType = http.DetectContentType(decodedData)
		}
	}

	// 3: 映射 MIME 到扩展名
	switch mimeType {
	// 图片
	case "image/jpeg":
		return ".jpg"
	case "image/png":
		return ".png"
	case "image/gif":
		return ".gif"
	case "image/webp":
		return ".webp"
	case "image/bmp":
		return ".bmp"
	case "image/tiff":
		return ".tiff"
	case "image/svg+xml":
		return ".svg"

	// 文档
	case "application/pdf":
		return ".pdf"
	case "application/msword":
		return ".doc"
	case "application/vnd.openxmlformats-officedocument.wordprocessingml.document":
		return ".docx"
	case "application/vnd.ms-excel":
		return ".xls"
	case "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet":
		return ".xlsx"
	case "application/vnd.ms-powerpoint":
		return ".ppt"
	case "application/vnd.openxmlformats-officedocument.presentationml.presentation":
		return ".pptx"
	case "application/rtf":
		return ".rtf"
	case "application/vnd.oasis.opendocument.text":
		return ".odt"
	case "application/vnd.oasis.opendocument.spreadsheet":
		return ".ods"

	// 文本 & JSON
	case "text/plain":
		return ".txt"
	case "application/json":
		return ".json"
	case "application/xml", "text/xml":
		return ".xml"
	case "text/html":
		return ".html"
	case "text/css":
		return ".css"
	case "application/javascript", "text/javascript":
		return ".js"
	case "text/csv":
		return ".csv"

	// 音视频
	case "audio/mpeg":
		return ".mp3"
	case "audio/wav":
		return ".wav"
	case "audio/ogg":
		return ".ogg"
	case "video/mp4":
		return ".mp4"
	case "video/webm":
		return ".webm"
	case "video/ogg":
		return ".ogv"

	// 其他
	default:
		return ".bin"
	}
}

func GetMD5FromBase64(base64Str string) string {
	// 去掉 data URI 前缀
	parts := strings.Split(base64Str, ",")
	if len(parts) >= 2 {
		base64Str = parts[1]
	}

	data, err := base64.StdEncoding.DecodeString(base64Str)
	if err != nil {
		return ""
	}

	hash := md5.Sum(data)
	return fmt.Sprintf("%x", hash)
}
