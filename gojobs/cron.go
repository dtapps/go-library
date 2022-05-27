package gojobs

import (
	"context"
	pb2 "go.dtapp.net/library/gojobs/pb"
	"google.golang.org/grpc"
	"log"
)

// CronConfig 定时任务配置
type CronConfig struct {
	Address string // 服务端口 127.0.0.1:8888
}

// Cron 定时任务
type Cron struct {
	CronConfig                  // 配置
	Pub        pb2.PubSubClient // 订阅
	Conn       *grpc.ClientConn // 链接信息
}

// NewCron 创建定时任务
func NewCron(config *CronConfig) *Cron {

	if config.Address == "" {
		panic("[定时任务]请填写服务端口")
	}

	c := &Cron{}

	c.Address = config.Address

	var err error

	// 建立连接 获取client
	c.Conn, err = grpc.Dial(c.Address, grpc.WithInsecure())
	if err != nil {
		panic("[定时任务]{连接失败}" + err.Error())
	}

	// 新建一个客户端
	c.Pub = pb2.NewPubSubClient(c.Conn)

	return c
}

// Send 发送
func (c *Cron) Send(in *pb2.PublishRequest) (*pb2.PublishResponse, error) {
	log.Printf("[定时任务]{广播开始}编号：%s 类型：%s ip：%s\n", in.GetId(), in.GetValue(), in.GetIp())
	stream, err := c.Pub.Publish(context.Background(), in)
	if err != nil {
		log.Printf("[定时任务]{广播失败}编号：%s %v\n", in.GetId(), err)
	}
	log.Printf("[定时任务]{广播成功}编号：%s 类型：%s ip：%s\n", in.GetId(), in.GetValue(), in.GetIp())
	return stream, err
}
