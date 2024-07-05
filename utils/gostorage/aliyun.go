package gostorage

import (
	"context"
	"errors"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"io"
)

// AliYunConfig 阿里云配置
type AliYunConfig struct {
	AccessKeyId     string // 账号信息
	AccessKeySecret string // 账号信息
	Endpoint        string // 地域节点
	BucketName      string // 存储空间名称
}

// AliYun 阿里云
type AliYun struct {
	accessKeyId     string      // 账号信息
	accessKeySecret string      // 账号信息
	endpoint        string      // 地域节点
	bucketName      string      // 存储空间名称
	error           error       // 错误信息
	client          *oss.Client // 驱动
	bucket          *oss.Bucket // 存储空间
}

// NewAliYun 初始化
// https://help.aliyun.com/document_detail/32144.html
func NewAliYun(ctx context.Context, config *AliYunConfig) (*AliYun, error) {
	app := &AliYun{}
	app.accessKeyId = config.AccessKeyId
	app.accessKeySecret = config.AccessKeySecret
	app.endpoint = config.Endpoint
	app.bucketName = config.BucketName

	// 创建链接
	app.client, app.error = oss.New(app.endpoint, app.accessKeyId, app.accessKeySecret)
	if app.error != nil {
		return nil, app.error
	}
	// 填写存储空间名称
	app.bucket, app.error = app.client.Bucket(app.bucketName)
	if app.error != nil {
		return nil, app.error
	}

	// 判断存储空间是否存在
	_, app.error = app.client.IsBucketExist(app.bucketName)
	if app.error != nil {
		return nil, app.error
	}

	return app, nil
}

// PutObject 上传文件流
// @param file 文件流
// @param filePath 文件路径
// @param fileName 文件名称
func (c *AliYun) PutObject(ctx context.Context, file io.Reader, filePath, fileName string) (resp FileInfo, err error) {
	objectKey := filePath
	if fileName != "" {
		objectKey = filePath + "/" + fileName
	}
	err = c.bucket.PutObject(objectKey, file)
	resp.Path = filePath
	resp.Name = fileName
	resp.Url = objectKey
	return
}

// PutLocalFile 上传本地文件
// @param localFile 本地文件路径
// @param filePath 文件路径
// @param fileName 文件名称
func (c *AliYun) PutLocalFile(ctx context.Context, localFilePath, filePath, fileName string) (resp FileInfo, err error) {
	if localFilePath == "" {
		return FileInfo{}, errors.New("localFilePath 不能为空")
	}
	objectKey := filePath
	if fileName != "" {
		objectKey = filePath + "/" + fileName
	}
	err = c.bucket.PutObjectFromFile(objectKey, localFilePath)
	resp.Path = filePath
	resp.Name = fileName
	resp.Url = objectKey
	return
}
