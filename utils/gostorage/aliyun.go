package gostorage

import (
	"context"
	"io"

	"github.com/aliyun/alibabacloud-oss-go-sdk-v2/oss"
	"github.com/aliyun/alibabacloud-oss-go-sdk-v2/oss/credentials"
)

// 阿里云
type Aliyun struct {
	accessKeyId     string // 账号信息
	accessKeySecret string // 账号信息
	region          string // 地域节点
	bucket          string // 存储空间名称

	client *oss.Client // 实例
}

// 初始化
// https://help.aliyun.com/document_detail/32144.html
func NewAliyun(ctx context.Context, opts ...Option) (client *Aliyun, err error) {
	options := NewOptions(opts)

	// 初始化
	client = &Aliyun{}

	client.accessKeyId = options.accessKeyId
	client.accessKeySecret = options.accessKeySecret
	client.region = options.region
	client.bucket = options.bucket

	// 加载默认配置并设置凭证提供者和区域
	cfg := oss.LoadDefaultConfig()
	cfg.WithCredentialsProvider(credentials.NewStaticCredentialsProvider(client.accessKeyId, client.accessKeySecret)) // 访问凭证
	if options.debug {
		cfg.WithLogLevel(oss.LogDebug) // 设置日志
	}
	cfg.WithRegion(client.region) // 设置区域

	// 实例
	client.client = oss.NewClient(cfg)

	return client, err
}

// PutObject
// 简单上传, 最大支持5GiB
// 支持CRC64数据校验（默认启用）
// 支持进度条
// 请求body类型为io.Reader, 当支持io.Seeker类型时，具备失败重传
func (c *Aliyun) PutObject(ctx context.Context, file io.Reader, filePath, fileName string) (resp FileInfo, err error) {
	objectKey := filePath
	if fileName != "" {
		objectKey = filePath + "/" + fileName
	}

	// 创建上传对象的请求
	request := &oss.PutObjectRequest{
		Bucket: Ptr(c.bucket),  // 存储空间名称
		Key:    Ptr(objectKey), // 对象名称
		Body:   file,           // 要上传的内容
	}

	// 执行上传对象的请求
	_, err = c.client.PutObject(ctx, request)
	resp.Path = filePath
	resp.Name = fileName
	resp.Url = objectKey
	return
}
