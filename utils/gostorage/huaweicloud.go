package gostorage

import (
	"github.com/huaweicloud/huaweicloud-sdk-go-obs/obs"
	"io"
)

// Huaweicloud 华为云
type Huaweicloud struct {
	AccessKey  string
	SecretKey  string
	Endpoint   string
	BucketName string
	error      error          // 错误信息
	client     *obs.ObsClient // 驱动
}

// NewHuaweicloud 初始化
// https://support.huaweicloud.com/sdk-go-devg-obs/obs_33_0001.html
func NewHuaweicloud(accessKey string, secretKey string, endpoint string, bucketName string) *Huaweicloud {
	app := &Huaweicloud{AccessKey: accessKey, SecretKey: secretKey, Endpoint: endpoint, BucketName: bucketName}
	app.client, app.error = obs.New(accessKey, secretKey, endpoint)
	if app.error == nil {
		app.client.Close() // 关闭obsClient
	}
	return app
}

// Bucket 存储空间
func (c *Huaweicloud) Bucket(name string) *Huaweicloud {
	c.BucketName = name
	return c
}

// PutObject 上传文件流
// @param file 文件流
// @param filePath 文件路径
// @param fileName 文件名称
func (c *Huaweicloud) PutObject(file io.Reader, filePath, fileName string) (resp FileInfo, err error) {
	objectKey := filePath
	if fileName != "" {
		objectKey = filePath + "/" + fileName
	}
	input := &obs.PutObjectInput{}
	input.Bucket = c.BucketName
	input.Key = objectKey
	input.Body = file
	_, err = c.client.PutObject(input)
	resp.Path = filePath
	resp.Name = fileName
	resp.Url = objectKey
	return
}
