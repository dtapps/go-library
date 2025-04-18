package ft07

import (
	"go.dtapp.net/library/utils/gorequest"
)

// Client 实例
type Client struct {
	httpClient *gorequest.App // HTTP请求客户端
	clientIP   string         // 客户端IP
	config     struct {
		sendKey string
		url     string
	}
}

// NewClient 创建实例化
func NewClient() (*Client, error) {
	c := &Client{}
	c.httpClient = gorequest.NewHttp()

	return c, nil
}
