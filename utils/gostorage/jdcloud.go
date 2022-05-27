package gostorage

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/credentials"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3"
)

// Jdcloud 京东云
type Jdcloud struct {
	AccessKey  string
	SecretKey  string
	Endpoint   *string
	Regions    *string
	BucketName string
	error      error  // 错误信息
	client     *s3.S3 // 驱动
}

// NewJdcloud 初始化
// https://docs.jdcloud.com/cn/object-storage-service/sdk-go
func NewJdcloud(accessKey string, secretKey string, endpoint *string, regions *string, bucketName string) *Jdcloud {
	app := &Jdcloud{AccessKey: accessKey, SecretKey: secretKey, Endpoint: endpoint, Regions: regions, BucketName: bucketName}
	reds := credentials.NewStaticCredentials(accessKey, secretKey, "")
	_, app.error = reds.Get()
	app.client = s3.New(session.New(&aws.Config{
		Region:      app.Regions,
		Endpoint:    app.Endpoint,
		DisableSSL:  aws.Bool(false),
		Credentials: reds,
	}))
	return app
}

// Bucket 存储空间
func (c *Jdcloud) Bucket(name string) *Jdcloud {
	c.BucketName = name
	return c
}
