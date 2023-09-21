package alipayopen

import (
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"github.com/dtapps/go-library/utils/golog"
	"github.com/dtapps/go-library/utils/gophp"
	"github.com/dtapps/go-library/utils/gorequest"
)

// ClientConfig 实例配置
type ClientConfig struct {
	AppId      string // 支付宝分配给开发者的应用ID
	AppKey     string // 支付宝分配给开发者的应用私钥
	AppRSA2    string // 应用公钥
	AlipayRSA2 string // 支付宝公钥
	Aes        string // 接口内容加密方式
}

// Client 实例
type Client struct {
	requestClient       *gorequest.App  // 请求服务
	requestClientStatus bool            // 请求服务状态
	privateKey          *rsa.PrivateKey // 私钥服务
	config              struct {
		appId      string // 支付宝分配给开发者的应用ID
		appKey     string // 支付宝分配给开发者的应用私钥
		appRSA2    string // 应用公钥
		alipayRSA2 string // 支付宝公钥
		aes        string // 接口内容加密方式
	}
	slog struct {
		status bool           // 状态
		client *golog.ApiSLog // 日志服务
	}
}

// NewClient 创建实例化
func NewClient(config *ClientConfig) (*Client, error) {

	var err error
	c := &Client{}

	c.config.appId = config.AppId
	c.config.appKey = config.AppKey
	c.config.appRSA2 = config.AppRSA2
	c.config.alipayRSA2 = config.AlipayRSA2
	c.config.aes = config.Aes

	// 私钥解码
	block, _ := pem.Decode([]byte("-----BEGIN RSA PRIVATE KEY-----\n" + gophp.ChunkSplit(config.AppKey, 64, "\n") + "-----END RSA PRIVATE KEY-----\n"))
	if block == nil {
		return nil, errors.New("签名私钥解码错误")
	}
	privateKey, err := x509.ParsePKCS8PrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}
	c.privateKey = privateKey.(*rsa.PrivateKey)

	return c, nil
}
