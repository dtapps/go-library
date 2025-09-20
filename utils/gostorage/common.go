package gostorage

import (
	"fmt"
	"io"
	"mime"
	"os"
	"path/filepath"

	"github.com/gabriel-vasile/mimetype"
)

// FileInfo 上传文件的信息
type FileInfo struct {
	Path string `json:"path"` // 文件路径
	Name string `json:"name"` // 文件名称
	Url  string `json:"url"`  // 文件地址
}

// 根据文件扩展名获取 MIME 类型
func GetMimeTypeFromExtension(filename string) string {
	ext := filepath.Ext(filename) // 获取扩展名（如 ".jpg"）
	if ext == "" {
		return "application/octet-stream" // 默认值
	}

	// 标准库 mime.TypeByExtension
	mimeType := mime.TypeByExtension(ext)
	if mimeType == "" {
		return "application/octet-stream" // 未知扩展名
	}
	return mimeType
}

// 尝试从文件或扩展名中检测 MIME 类型
func DetectMIME(reader io.Reader, filename string) (string, error) {
	fmt.Printf("[DetectMIME] filename: %s\n", filename)

	// 如果是文件类型，尝试检测真实 MIME
	if f, ok := reader.(*os.File); ok {
		mime, err := mimetype.DetectReader(f)
		if err == nil {
			_, err = f.Seek(0, 0)
			if err != nil {
				return "", fmt.Errorf("重置文件指针失败: %w", err)
			}
			return mime.String(), nil
		}
		// MIME 检测失败时也继续回退到扩展名
	}

	// 否则回退到扩展名
	ext := filepath.Ext(filename)
	fmt.Printf("[DetectMIME] ext: %s\n", ext)
	if ext == "" {
		return "application/octet-stream", nil
	}

	mimeType := mime.TypeByExtension(ext)
	fmt.Printf("[DetectMIME] mimeType: %s\n", mimeType)
	if mimeType == "" {
		return "application/octet-stream", nil
	}
	return mimeType, nil
}
