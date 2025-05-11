package framework

import (
	"fmt"
	"io"
	"mime/multipart"
	"os"
)

type UploadFile struct {
	Filename string                // 文件名
	Size     int64                 // 文件大小
	File     *multipart.FileHeader // 原始文件对象
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

// SaveFile 保存上传文件
// dst: 保存路径 + 文件名
// cover: 是否覆盖
func (u *UploadFile) SaveFile(dst string, cover bool) error {
	// 确保目录存在
	if err := ensureDirExists(dst); err != nil {
		return fmt.Errorf("创建目录失败: %w", err)
	}

	// 判断文件是否存在
	if !cover {
		if fileExists(dst) {
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
