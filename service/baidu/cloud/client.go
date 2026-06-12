package cloud

import (
	"context"
	"go.dtapp.net/library/utils/gorequest"
)

type Client struct {
	apiKey      string
	secretKey   string
	accessToken string
	httpClient  *gorequest.App // HTTP请求客户端
	clientIP    string         // 客户端IP
}

// NewClient 创建实例化
func NewClient(ctx context.Context, apiKey string, secretKey string) (*Client, error) {
	return &Client{apiKey: apiKey, secretKey: secretKey, httpClient: gorequest.NewHttp()}, nil
}
