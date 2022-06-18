package gostorage

import (
	"context"
	"fmt"
	"github.com/tencentyun/cos-go-sdk-v5"
	"io"
	"net/http"
	"net/url"
)

// Tencent 腾讯云
type Tencent struct {
	SecretID   string      // 用户的 SecretId
	SecretKey  string      // 用户的 SecretKey
	BucketName string      // 存储桶名称
	Regions    string      // 所属地域
	client     *cos.Client // 驱动
	error      error       // 错误信息
}

// NewTencent 初始化
// https://cloud.tencent.com/document/product/436/31215
func NewTencent(secretID, secretKey, regions, bucketName string) *Tencent {
	app := &Tencent{SecretID: secretID, SecretKey: secretKey, Regions: regions, BucketName: bucketName}
	u, _ := url.Parse(fmt.Sprintf("https://%s.cos.%s.myqcloud.com", bucketName, regions))
	su, _ := url.Parse(fmt.Sprintf("https://cos.%s.myqcloud.com", regions))
	b := &cos.BaseURL{BucketURL: u, ServiceURL: su}
	// 1.永久密钥
	app.client = cos.NewClient(b, &http.Client{
		Transport: &cos.AuthorizationTransport{
			SecretID:  secretID,
			SecretKey: secretKey,
		},
	})
	return app
}

// Bucket 存储空间
func (c *Tencent) Bucket(name string) *Tencent {
	return NewTencent(c.SecretID, c.SecretKey, c.Regions, name)
}

// PutObject 上传文件流
// @param file 文件流
// @param filePath 文件路径
// @param fileName 文件名称
func (c *Tencent) PutObject(file io.Reader, filePath, fileName string) (resp FileInfo, err error) {
	objectKey := filePath
	if fileName != "" {
		objectKey = filePath + "/" + fileName
	}
	_, err = c.client.Object.Put(context.Background(), objectKey, file, nil)
	resp.Path = filePath
	resp.Name = fileName
	resp.Url = objectKey
	return
}
