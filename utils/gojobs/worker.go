package gojobs

import (
	"context"
	"go.dtapp.net/library/utils/gojobs/pb"
	"go.dtapp.net/library/utils/gouuid"
	"google.golang.org/grpc"
)

// WorkerConfig 工作配置
type WorkerConfig struct {
	Address  string // 服务端口 127.0.0.1:8888
	ClientIp string // 自己的ip地址
}

// Worker 工作
type Worker struct {
	WorkerConfig                  // 配置
	Pub          pb.PubSubClient  // 订阅
	Conn         *grpc.ClientConn // 链接信息
}

// NewWorker 创建工作
func NewWorker(config *WorkerConfig) *Worker {

	if config.Address == "" {
		panic("[工作线]请填写服务端口")
	}
	if config.ClientIp == "" {
		panic("[定时任务]请填写ip地址")
	}

	w := &Worker{}

	w.Address = config.Address
	w.ClientIp = config.ClientIp

	var err error

	// 建立连接 获取client
	w.Conn, err = grpc.Dial(w.Address, grpc.WithInsecure())
	if err != nil {
		panic("[工作线]{连接失败}" + err.Error())
	}

	// 新建一个客户端
	w.Pub = pb.NewPubSubClient(w.Conn)

	return w
}

// SubscribeCron 订阅服务
func (w *Worker) SubscribeCron() pb.PubSub_SubscribeClient {
	stream, err := w.Pub.Subscribe(context.Background(), &pb.SubscribeRequest{
		Id:    gouuid.GetUuId(),
		Value: prefix,
		Ip:    w.ClientIp,
	})
	if err != nil {
		panic("[工作线]{订阅服务失败}" + err.Error())
	}
	return stream
}

// StartCron 启动任务
func (w *Worker) StartCron() pb.PubSub_SubscribeClient {
	stream, err := w.Pub.Subscribe(context.Background(), &pb.SubscribeRequest{
		Id:    gouuid.GetUuId(),
		Value: prefixSprintf(w.ClientIp),
		Ip:    w.ClientIp,
	})
	if err != nil {
		panic("[工作线]{启动任务失败}" + err.Error())
	}
	return stream
}
