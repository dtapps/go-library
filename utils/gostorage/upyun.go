package gostorage

import (
	"github.com/upyun/go-sdk/v3/upyun"
	"io"
)

// Upyun 又拍云
type Upyun struct {
	Operator   string
	Password   string
	BucketName string
	client     *upyun.UpYun // 驱动
}

// NewUpyun 初始化
// https://help.upyun.com/docs/storage/
// https://github.com/upyun/go-sdk
func NewUpyun(operator string, password string, bucketName string) *Upyun {
	app := &Upyun{Operator: operator, Password: password, BucketName: bucketName}
	app.client = upyun.NewUpYun(&upyun.UpYunConfig{
		Bucket:   bucketName,
		Operator: operator,
		Password: password,
	})
	return app
}

// Bucket 存储空间
func (c *Upyun) Bucket(name string) *Upyun {
	c.BucketName = name
	return c
}

// PutObject 上传文件流
// @param file 文件流
// @param filePath 文件路径
// @param fileName 文件名称
func (c *Upyun) PutObject(file io.Reader, filePath, fileName, acl string) (resp FileInfo, err error) {
	objectKey := filePath
	if fileName != "" {
		objectKey = filePath + "/" + fileName
	}
	err = c.client.Put(&upyun.PutObjectConfig{
		Path:      "/demo.log",
		LocalPath: "/tmp/upload",
	})
	resp.Path = filePath
	resp.Name = fileName
	resp.Url = objectKey
	return
}
