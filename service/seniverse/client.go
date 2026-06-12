package seniverse

import (
	"go.dtapp.net/library/utils/gorequest"
)

type V3Client struct {
	key        string         // API密钥
	httpClient *gorequest.App // HTTP请求客户端
	clientIP   string         // 客户端IP
}

func NewV3Client(key string) (*V3Client, error) {
	return &V3Client{key: key, httpClient: gorequest.NewHttp()}, nil
}

type V4Client struct {
	publicKey  string
	secret     string
	httpClient *gorequest.App // HTTP请求客户端
	clientIP   string         // 客户端IP
}

func NewV4Client(publicKey string, secret string) (*V4Client, error) {
	return &V4Client{publicKey: publicKey, secret: secret, httpClient: gorequest.NewHttp()}, nil
}
