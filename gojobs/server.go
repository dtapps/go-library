package gojobs

import (
	"errors"
	pb2 "go.dtapp.net/library/gojobs/pb"
	"go.dtapp.net/library/gojobs/pubsub"
	"google.golang.org/grpc"
	"log"
	"net"
	"strings"
	"time"
)

// ServerConfig 服务配置
type ServerConfig struct {
	PublishTimeout time.Duration // 控制发布时最大阻塞时间
	PubBuffer      int           // 缓冲区大小，控制每个订阅者的chan缓冲区大小
	Address        string        // 服务端口 0.0.0.0:8888
}

// Server 服务
type Server struct {
	ServerConfig                   // 配置
	Pub          *pubsub.Publisher // 订阅
	Conn         *grpc.Server      // 链接信息
}

// NewServer 创建服务和注册
func NewServer(config *ServerConfig) *Server {

	if config.Address == "" {
		panic("[服务中转]请填写服务端口")
	}

	s := &Server{}

	s.PublishTimeout = config.PublishTimeout
	s.PubBuffer = config.PubBuffer
	s.Address = config.Address

	s.Pub = pubsub.NewPublisher(config.PublishTimeout, config.PubBuffer)

	// 创建gRPC服务器
	s.Conn = grpc.NewServer()

	// 注册
	pb2.RegisterPubSubServer(s.Conn, pb2.NewPubSubServerService())

	return s
}

// StartCron 启动定时任务
func (s *Server) StartCron() {
	cron := s.Pub.SubscribeTopic(func(v interface{}) bool {
		if key, ok := v.(string); ok {
			if strings.HasPrefix(key, prefix) {
				return true
			}
		}
		return false
	})

	go func() {
		log.Println("cron：topic:", <-cron)
	}()
}

// StartUp 启动服务
func (s *Server) StartUp() {

	// 监听本地端口
	lis, err := net.Listen("tcp", s.Address)
	if err != nil {
		panic(errors.New("[服务中转]{创建监听失败}" + err.Error()))
	}
	log.Println("[服务中转]{监听}", lis.Addr())

	// 启动grpc
	err = s.Conn.Serve(lis)
	if err != nil {
		panic(errors.New("[服务中转]{创建服务失败}" + err.Error()))
	}

}
