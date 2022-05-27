package gostorage

import (
	"github.com/baidubce/bce-sdk-go/bce"
	"github.com/baidubce/bce-sdk-go/services/bos"
	"io"
)

// Baidu 百度云
type Baidu struct {
	AccessKey       string
	SecretAccessKey string
	Endpoint        string
	BucketName      string
	error           error       // 错误信息
	client          *bos.Client // 驱动
}

// NewBaidu 初始化
// https://cloud.baidu.com/doc/BOS/s/4jwvyry1p
func NewBaidu(accessKey string, secretAccessKey, endpoint, bucketName string) *Baidu {
	app := &Baidu{AccessKey: accessKey, SecretAccessKey: secretAccessKey, Endpoint: endpoint, BucketName: bucketName}
	clientConfig := bos.BosClientConfiguration{
		Ak:               accessKey,
		Sk:               secretAccessKey,
		Endpoint:         endpoint,
		RedirectDisabled: false,
	}
	app.client, app.error = bos.NewClientWithConfig(&clientConfig)
	return app
}

// Bucket 存储空间
func (c *Baidu) Bucket(name string) *Baidu {
	c.BucketName = name
	return c
}

// PutObject 上传文件流
// @param file 文件流
// @param filePath 文件路径
// @param fileName 文件名称
func (c *Baidu) PutObject(file io.Reader, filePath, fileName string) (resp FileInfo, err error) {
	objectKey := filePath
	if fileName != "" {
		objectKey = filePath + "/" + fileName
	}
	bodyStream, err := bce.NewBodyFromSizedReader(file, -1)
	_, err = c.client.PutObject(c.BucketName, objectKey, bodyStream, nil)
	resp.Path = filePath
	resp.Name = fileName
	resp.Url = objectKey
	return
}
