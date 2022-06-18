package gojobs

import (
	"google.golang.org/grpc"
)

// ClientConfig 客户端配置
type ClientConfig struct {
	Address string // 服务端口 127.0.0.1:8888
}

// Client 定时任务
type Client struct {
	ClientConfig                  // 配置
	Conn         *grpc.ClientConn // 链接信息
}

// NewClient 创建客户端
func NewClient(config *ClientConfig) *Client {

	if config.Address == "" {
		panic("[客户端]请填写服务端口")
	}

	c := &Client{}

	c.Address = config.Address

	var err error

	// 建立连接 获取client
	c.Conn, err = grpc.Dial(c.Address, grpc.WithInsecure())
	if err != nil {
		panic("[客户端]{连接失败}" + err.Error())
	}

	return c
}
