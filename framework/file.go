package framework

import (
	"encoding/base64"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"strings"
)

type UploadFile struct {
	Filename string                // 文件名
	Size     int64                 // 文件大小
	File     *multipart.FileHeader // 原始文件对象
	Data     []byte                // Base64 解码后的字节数据
}

// FormFile 单文件
func (c *Context) FormFile(key string) (*UploadFile, error) {
	if c.IsGin() {
		file, err := c.ginCtx.FormFile(key)
		if err != nil {
			return nil, err
		}
		return &UploadFile{
			Filename: file.Filename,
			Size:     file.Size,
			File:     file,
		}, nil
	}
	if c.IsHertz() {
		file, err := c.hertzCtx.FormFile(key)
		if err != nil {
			return nil, err
		}
		return &UploadFile{
			Filename: file.Filename,
			Size:     file.Size,
			File:     file,
		}, nil
	}
	if c.IsEcho() {
		file, err := c.echoCtx.FormFile(key)
		if err != nil {
			return nil, err
		}
		return &UploadFile{
			Filename: file.Filename,
			Size:     file.Size,
			File:     file,
		}, nil
	}

	return nil, fmt.Errorf("不支持的框架上下文")
}

// FromBase64 将 base64 字符串解析为 UploadFile 对象
func (c *Context) FromBase64(base64Str, filename string) (*UploadFile, error) {

	if filename == "" {
		md5Hash := GetMD5FromBase64(base64Str)
		fileExt := GetFileExtFromBase64(base64Str)
		if md5Hash == "" {
			return nil, fmt.Errorf("无法生成文件名：MD5 计算失败")
		}
		filename = md5Hash + fileExt
	}

	// 去掉 data URI 前缀（如：data:image/png;base64,...）
	if len(base64Str) > 100 && base64Str[:5] == "data:" {
		parts := strings.Split(base64Str, ",")
		if len(parts) >= 2 {
			base64Str = parts[1]
		}
	}

	// 解码 base64 数据
	data, err := base64.StdEncoding.DecodeString(base64Str)
	if err != nil {
		return nil, fmt.Errorf("base64 解码失败: %w", err)
	}

	// 设置数据和文件名
	return &UploadFile{
		Filename: filename,
		Size:     int64(len(data)),
		File: &multipart.FileHeader{
			Filename: filename,
			Size:     int64(len(data)),
		},
		Data: data,
	}, nil
}

// SaveFile 保存上传文件
// dst: 保存路径 + 文件名
// cover: 是否覆盖
func (u *UploadFile) SaveFile(dst string, cover bool) error {
	// 确保目录存在
	if err := EnsureDirExists(dst); err != nil {
		return fmt.Errorf("创建目录失败: %w", err)
	}

	// 判断文件是否存在
	if !cover {
		if FileExists(dst) {
			return fmt.Errorf("文件已存在")
		}
	}

	// 打开上传的文件源
	src, err := u.File.Open()
	if err != nil {
		return fmt.Errorf("打开上传文件失败: %w", err)
	}
	defer src.Close()

	// 创建目标文件
	out, err := os.Create(dst)
	if err != nil {
		return fmt.Errorf("创建目标文件失败: %w", err)
	}
	defer out.Close()

	// 拷贝内容
	_, err = io.Copy(out, src)
	if err != nil {
		return fmt.Errorf("写入文件失败: %w", err)
	}

	return nil
}

// ToBytes 文件转字节流 []byte
func (u *UploadFile) ToBytes() ([]byte, error) {

	// 打开上传的文件源
	src, err := u.File.Open()
	if err != nil {
		return nil, fmt.Errorf("打开上传文件失败: %w", err)
	}
	defer src.Close()

	return io.ReadAll(src)
}

// ToBase64 文件转 base64 字符串
func (u *UploadFile) ToBase64() (string, error) {
	// 打开上传的文件源
	src, err := u.File.Open()
	if err != nil {
		return "", fmt.Errorf("打开上传文件失败: %w", err)
	}
	defer src.Close()

	// 读取文件内容到内存
	data, err := io.ReadAll(src)
	if err != nil {
		return "", fmt.Errorf("读取文件内容失败: %w", err)
	}

	// 使用 base64 编码
	base64Str := base64.StdEncoding.EncodeToString(data)
	return base64Str, nil
}
