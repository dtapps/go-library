package gostorage

import (
	"errors"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"io"
	"sync"
)

// AliYunConfig 阿里云配置
type AliYunConfig struct {
	AccessKeyId        string // 账号信息
	AccessKeySecret    string // 账号信息
	Endpoint           string // 地域节点 外网访问 test3
	EndpointEcs        string // 地域节点 ECS 的经典网络访问（内网） test1
	EndpointAccelerate string // 地域节点 传输加速域名（全地域上传下载加速 test4
	BucketName         string // 存储空间名称
}

// AliYun 阿里云
type AliYun struct {
	accessKeyId     string      // 账号信息
	accessKeySecret string      // 账号信息
	bucketName      string      // 存储空间名称
	endpoint        string      // 地域节点
	error           error       // 错误信息
	client          *oss.Client // 驱动
	bucket          *oss.Bucket // 存储空间
	endpointEcsInfo struct {
		error    error       // 错误信息
		endpoint string      // 地域节点
		client   *oss.Client // 驱动
		bucket   *oss.Bucket // 存储空间
	} // 外网访问
	endpointInfo struct {
		error    error       // 错误信息
		endpoint string      // 地域节点
		client   *oss.Client // 驱动
		bucket   *oss.Bucket // 存储空间
	} // 内网访问
	endpointAccelerateInfo struct {
		error    error       // 错误信息
		endpoint string      // 地域节点
		client   *oss.Client // 驱动
		bucket   *oss.Bucket // 存储空间
	} // 传输加速域名访问
}

// NewAliYun 初始化
// https://help.aliyun.com/document_detail/32144.html
// https://github.com/aliyun/aliyun-oss-go-sdk
func NewAliYun(config *AliYunConfig) (*AliYun, error) {
	app := &AliYun{}
	app.accessKeyId = config.AccessKeyId
	app.accessKeySecret = config.AccessKeySecret
	app.bucketName = config.BucketName
	app.endpointEcsInfo.endpoint = config.Endpoint
	app.endpointInfo.endpoint = config.EndpointEcs
	app.endpointAccelerateInfo.endpoint = config.EndpointAccelerate

	wg := sync.WaitGroup{}
	wg.Add(3)
	go app.configEndpointEcs(&wg)
	go app.configEndpoint(&wg)
	go app.configEndpointAccelerate(&wg)
	wg.Wait()

	// 判断结果
	if app.endpointEcsInfo.error == nil {
		app.endpoint = app.endpointEcsInfo.endpoint
		app.client = app.endpointEcsInfo.client
		app.bucket = app.endpointEcsInfo.bucket
		return app, nil
	}
	if app.endpointInfo.error == nil {
		app.endpoint = app.endpointInfo.endpoint
		app.client = app.endpointInfo.client
		app.bucket = app.endpointInfo.bucket
		return app, nil
	}
	if app.endpointAccelerateInfo.error == nil {
		app.endpoint = app.endpointAccelerateInfo.endpoint
		app.client = app.endpointAccelerateInfo.client
		app.bucket = app.endpointAccelerateInfo.bucket
		return app, nil
	}

	return app, errors.New("链接失败")
}

func (c *AliYun) configEndpoint(wg *sync.WaitGroup) {
	defer wg.Done()

	if c.endpointInfo.endpoint == "" {
		c.endpointInfo.error = errors.New("没有配置")
		return
	}

	// 创建链接
	c.endpointInfo.client, c.endpointInfo.error = oss.New(c.endpointInfo.endpoint, c.accessKeyId, c.accessKeySecret)
	if c.endpointInfo.error != nil {
		return
	}

	// 填写存储空间名称
	c.endpointInfo.bucket, c.endpointInfo.error = c.endpointInfo.client.Bucket(c.bucketName)
	if c.endpointInfo.error != nil {
		return
	}

	// 判断存储空间是否存在
	_, c.endpointInfo.error = c.endpointInfo.client.IsBucketExist(c.bucketName)
	if c.endpointInfo.error != nil {
		return
	}

	c.endpointInfo.error = nil
	return
}

func (c *AliYun) configEndpointEcs(wg *sync.WaitGroup) {
	defer wg.Done()

	if c.endpointEcsInfo.endpoint == "" {
		c.endpointEcsInfo.error = errors.New("没有配置")
		return
	}

	// 创建链接
	c.endpointEcsInfo.client, c.endpointEcsInfo.error = oss.New(c.endpointEcsInfo.endpoint, c.accessKeyId, c.accessKeySecret)
	if c.endpointEcsInfo.error != nil {
		return
	}

	// 填写存储空间名称
	c.endpointEcsInfo.bucket, c.endpointEcsInfo.error = c.endpointEcsInfo.client.Bucket(c.bucketName)
	if c.endpointEcsInfo.error != nil {
		return
	}

	// 判断存储空间是否存在
	_, c.endpointEcsInfo.error = c.endpointEcsInfo.client.IsBucketExist(c.bucketName)
	if c.endpointEcsInfo.error != nil {
		return
	}

	c.endpointEcsInfo.error = nil

	return
}

func (c *AliYun) configEndpointAccelerate(wg *sync.WaitGroup) {
	defer wg.Done()

	if c.endpointAccelerateInfo.endpoint == "" {
		c.endpointAccelerateInfo.error = errors.New("没有配置")
		return
	}

	// 创建链接
	c.endpointAccelerateInfo.client, c.endpointAccelerateInfo.error = oss.New(c.endpointAccelerateInfo.endpoint, c.accessKeyId, c.accessKeySecret)
	if c.endpointAccelerateInfo.error != nil {
		return
	}

	// 填写存储空间名称
	c.endpointAccelerateInfo.bucket, c.endpointAccelerateInfo.error = c.endpointAccelerateInfo.client.Bucket(c.bucketName)
	if c.endpointAccelerateInfo.error != nil {
		return
	}

	// 判断存储空间是否存在
	_, c.endpointAccelerateInfo.error = c.endpointAccelerateInfo.client.IsBucketExist(c.bucketName)
	if c.endpointAccelerateInfo.error != nil {
		return
	}

	c.endpointAccelerateInfo.error = nil

	return
}

// PutObject 上传文件流
// @param file 文件流
// @param filePath 文件路径
// @param fileName 文件名称
func (c *AliYun) PutObject(file io.Reader, filePath, fileName string) (resp FileInfo, err error) {
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
func (c *AliYun) PutLocalFile(localFilePath, filePath, fileName string) (resp FileInfo, err error) {
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
