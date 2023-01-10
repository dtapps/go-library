package gostorage

import (
	"context"
	"github.com/qiniu/go-sdk/v7/auth/qbox"
	"github.com/qiniu/go-sdk/v7/storage"
	"io"
)

// Qiniu 七牛云
type Qiniu struct {
	AccessKey  string
	SecretKey  string
	BucketName string
	client     *qbox.Mac          // 驱动
	bucket     *storage.PutPolicy // 存储空间
	upToken    string             // 上传凭证
}

// NewQiniu 初始化
// https://developer.qiniu.com/kodo/1238/go
// https://github.com/qiniu/go-sdk
func NewQiniu(accessKey string, secretKey string, bucketName string) *Qiniu {
	app := &Qiniu{AccessKey: accessKey, SecretKey: secretKey, BucketName: bucketName}
	app.client = qbox.NewMac(accessKey, secretKey)
	app.bucket.Scope = bucketName
	app.upToken = app.bucket.UploadToken(app.client)
	return app
}

// Bucket 存储空间
func (c *Qiniu) Bucket(name string) *Qiniu {
	c.BucketName = name
	return c
}

// PutObject 上传文件流
// @param file 文件流
// @param filePath 文件路径
// @param fileName 文件名称
func (c *Qiniu) PutObject(file io.Reader, filePath, fileName, acl string) (resp FileInfo, err error) {
	objectKey := filePath
	if fileName != "" {
		objectKey = filePath + "/" + fileName
	}

	cfg := storage.Config{}
	// 空间对应的机房
	cfg.Zone = &storage.ZoneHuanan
	// 是否使用https域名
	cfg.UseHTTPS = true
	// 上传是否使用CDN上传加速
	cfg.UseCdnDomains = false

	formUploader := storage.NewFormUploader(&cfg)
	ret := storage.PutRet{}
	putExtra := storage.PutExtra{}

	err = formUploader.Put(context.Background(), &ret, c.upToken, objectKey, file, -1, &putExtra)

	resp.Path = filePath
	resp.Name = fileName
	resp.Url = objectKey
	return
}
