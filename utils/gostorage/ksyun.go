package gostorage

import (
	"github.com/ks3sdklib/aws-sdk-go/aws"
	"github.com/ks3sdklib/aws-sdk-go/aws/credentials"
	"github.com/ks3sdklib/aws-sdk-go/service/s3"
	"io"
)

// Ksyun 金山云
type Ksyun struct {
	AccessKey  string
	SecretKey  string
	Endpoint   string
	Regions    string
	BucketName string
	client     *s3.S3 // 驱动
}

// NewKsyun 初始化
// https://docs.ksyun.com/documents/40487
func NewKsyun(accessKey string, secretKey string, endpoint string, regions string, bucketName string) *Ksyun {
	app := &Ksyun{AccessKey: accessKey, SecretKey: secretKey, Endpoint: endpoint, Regions: regions, BucketName: bucketName}
	var cre = credentials.NewStaticCredentials(accessKey, secretKey, "")
	app.client = s3.New(&aws.Config{
		Region:      regions,
		Credentials: cre,
		Endpoint:    endpoint,
	})
	return app
}

// Bucket 存储空间
func (c *Ksyun) Bucket(name string) *Ksyun {
	c.BucketName = name
	return c
}

// PutObject 上传文件流
// @param file 文件流
// @param filePath 文件路径
// @param fileName 文件名称
func (c *Ksyun) PutObject(file io.Reader, filePath, fileName, acl string) (resp FileInfo, err error) {
	objectKey := filePath
	if fileName != "" {
		objectKey = filePath + "/" + fileName
	}
	params := &s3.PutObjectInput{
		Bucket: aws.String(c.BucketName),
		Key:    aws.String(objectKey),
		ACL:    aws.String(acl),
		//Body:        bytes.NewReader(file),
		ContentType: aws.String("application/octet-stream"),
	}
	_, err = c.client.PutObject(params)
	resp.Path = filePath
	resp.Name = fileName
	resp.Url = objectKey
	return
}
