package gostorage

import (
	"context"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"

	"github.com/minio/minio-go/v7"
	"github.com/minio/minio-go/v7/pkg/credentials"
)

// 天翼云
type Ctyun struct {
	acceessKey string // 账号信息
	secretKey  string // 账号信息
	endpoint   string // 地域节点
	region     string // 地域节点
	bucket     string // 存储空间名称
	secure     bool   // 是否安全连接

	client *minio.Client // 实例
}

// 初始化
func NewCtyun(ctx context.Context, opts ...Option) (client *Ctyun, err error) {
	options := NewOptions(opts)

	// 初始化
	client = &Ctyun{}

	client.acceessKey = options.acceessKey
	client.secretKey = options.secretKey
	client.endpoint = options.endpoint
	client.region = options.region
	client.bucket = options.bucket
	client.secure = options.secure

	// 实例
	client.client, err = minio.New(
		options.endpoint,
		&minio.Options{
			Creds:  credentials.NewStaticV4(client.acceessKey, client.secretKey, ""),
			Secure: options.secure,
		})
	if err != nil {
		return nil, err
	}

	if options.debug {
		// Ping 检查 MinIO 是否可用
		pingCtx, cancel := context.WithTimeout(ctx, 30*time.Second) // 设置 5 秒超时
		defer cancel()

		// 尝试列出 Bucket（轻量级操作）
		bucketInfo, err := client.client.ListBuckets(pingCtx)
		if err != nil {
			return nil, fmt.Errorf("MinIO 不可用: %v", err)
		}
		fmt.Printf("MinIO 可用，Bucket 列表: %+v\n", bucketInfo)
	}

	return client, err
}

// 上传本地文件
func (c *Ctyun) UpdateFile(ctx context.Context, filePath string, fileName string, objectName string) (info minio.UploadInfo, err error) {
	// 打开本地文件
	file, err := os.Open(filepath.Join(filePath, fileName))
	if err != nil {
		return info, fmt.Errorf("无法打开文件: %w", err)
	}
	defer file.Close()
	// 获取文件信息
	fileStat, err := file.Stat()
	if err != nil {
		return info, fmt.Errorf("获取文件信息失败: %w", err)
	}
	// 尝试从文件或扩展名中检测 MIME 类型
	contentType, err := DetectMIME(file, objectName)
	if err != nil {
		return minio.UploadInfo{}, fmt.Errorf("检测 MIME 类型失败: %w", err)
	}
	// 上传选项
	opts := minio.PutObjectOptions{
		ContentType: contentType,      // 使用检测到的 MIME 类型
		PartSize:    10 * 1024 * 1024, // 每个分块大小为 10MB（可根据需求调整）
	}
	// 上传文件
	uploadInfo, err := c.client.PutObject(ctx, c.bucket, objectName, file, fileStat.Size(), opts)
	fmt.Printf("UpdateFile: %+v opts: %+v", uploadInfo, opts)
	return uploadInfo, err
}

// 上传文件
func (c *Ctyun) PutObject(ctx context.Context, file io.Reader, filePath, fileName string) (info minio.UploadInfo, err error) {
	objectName := filePath
	if fileName != "" {
		objectName = filepath.Join(filePath, fileName)
	}
	// 尝试从文件或扩展名中检测 MIME 类型
	contentType, err := DetectMIME(file, objectName)
	if err != nil {
		return minio.UploadInfo{}, fmt.Errorf("检测 MIME 类型失败: %w", err)
	}
	// 上传选项
	opts := minio.PutObjectOptions{
		ContentType: contentType,      // 使用检测到的 MIME 类型
		PartSize:    10 * 1024 * 1024, // 每个分块大小为 10MB（可根据需求调整）
	}
	// 上传文件
	uploadInfo, err := c.client.PutObject(ctx, c.bucket, objectName, file, -1, opts)
	fmt.Printf("PutObject: %+v opts: %+v", uploadInfo, opts)
	return uploadInfo, err
}

// 上传文件
func (c *Ctyun) OriginalPutObject(ctx context.Context, objectName string, reader io.Reader, objectSize int64) (info minio.UploadInfo, err error) {
	// 尝试从文件或扩展名中检测 MIME 类型
	contentType, err := DetectMIME(reader, objectName)
	if err != nil {
		return minio.UploadInfo{}, fmt.Errorf("检测 MIME 类型失败: %w", err)
	}
	// 上传选项
	opts := minio.PutObjectOptions{
		ContentType: contentType,      // 使用检测到的 MIME 类型
		PartSize:    10 * 1024 * 1024, // 每个分块大小为 10MB（可根据需求调整）
	}
	// 上传文件
	uploadInfo, err := c.client.PutObject(ctx, c.bucket, objectName, reader, objectSize, opts)
	fmt.Printf("OriginalPutObject: %+v opts: %+v", uploadInfo, opts)
	return uploadInfo, err
}

// 上传文件
func (c *Ctyun) OriginalPutObjectWithContentType(ctx context.Context, objectName string, reader io.Reader, objectSize int64, contentType string) (info minio.UploadInfo, err error) {
	// 上传选项
	opts := minio.PutObjectOptions{
		ContentType: contentType,      // 使用检测到的 MIME 类型
		PartSize:    10 * 1024 * 1024, // 每个分块大小为 10MB（可根据需求调整）
	}
	// 上传文件
	uploadInfo, err := c.client.PutObject(ctx, c.bucket, objectName, reader, objectSize, opts)
	fmt.Printf("OriginalPutObject: %+v opts: %+v", uploadInfo, opts)
	return uploadInfo, err
}

// 删除文件
// filePath 文件路径
// fileName 文件名
func (c *Ctyun) DeleteObject(ctx context.Context, filePath string, fileName string) (err error) {
	objectName := filePath
	if fileName != "" {
		objectName = filepath.Join(filePath, fileName)
	}
	return c.client.RemoveObject(ctx, c.bucket, objectName, minio.RemoveObjectOptions{})
}

// 删除文件
// objectName 对象名称
func (c *Ctyun) RemoveObject(ctx context.Context, objectName string) (err error) {
	return c.client.RemoveObject(ctx, c.bucket, objectName, minio.RemoveObjectOptions{})
}

// 获取文件
// objectName 对象名称
func (c *Ctyun) StatObject(ctx context.Context, objectName string) (err error) {
	objInfo, err := c.client.StatObject(ctx, c.bucket, objectName, minio.StatObjectOptions{})
	fmt.Printf("StatObject: %+v", objInfo)
	return err
}

// ObjectExists 检查对象是否存在
func (c *Ctyun) ObjectExists(ctx context.Context, objectName string) (bool, error) {
	objInfo, err := c.client.StatObject(ctx, c.bucket, objectName, minio.StatObjectOptions{})
	fmt.Printf("ObjectExists: %+v", objInfo)
	if err != nil {
		if minio.ToErrorResponse(err).Code == "NoSuchKey" {
			return false, nil // 文件不存在
		}
		return false, err // 其他错误
	}
	return true, nil // 文件存在
}
