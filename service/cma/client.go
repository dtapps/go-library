package cma

import (
	"go.dtapp.net/library/utils/gorequest"
)

type Client struct {
	httpClient *gorequest.App // HTTP请求客户端
	clientIP   string         // 客户端IP
}

func NewClient() (*Client, error) {
	return &Client{httpClient: gorequest.NewHttp()}, nil
}
