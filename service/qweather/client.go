package qweather

import (
	"go.dtapp.net/library/utils/gorequest"
)

type Client struct {
	key        string
	httpClient *gorequest.App // HTTP请求客户端
	clientIP   string         // 客户端IP
}

func NewClient(key string) (*Client, error) {
	return &Client{key: key, httpClient: gorequest.NewHttp()}, nil
}
