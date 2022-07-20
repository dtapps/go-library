package pb

import (
	"context"
	"github.com/dtapps/go-library/utils/gojobs/pubsub"
	"log"
	"strings"
	"time"
)

type PubSubServerService struct {
	pub *pubsub.Publisher
	UnimplementedPubSubServer
}

func NewPubSubServerService() *PubSubServerService {
	return &PubSubServerService{
		// 新建一个Publisher对象
		pub: pubsub.NewPublisher(time.Millisecond*100, 10),
	}
}

// Publish 实现发布方法
func (p *PubSubServerService) Publish(ctx context.Context, req *PublishRequest) (*PublishResponse, error) {

	log.Printf("[服务中转]{发布}编号：%s 类型：%s ip地址：%s\n", req.GetId(), req.GetValue(), req.GetIp())

	// 发布消息
	p.pub.Publish(req.GetValue())
	return &PublishResponse{
		Id:    req.GetId(),
		Value: req.GetValue(),
		Ip:    req.GetIp(),
	}, nil
}

// Subscribe 实现订阅方法
func (p *PubSubServerService) Subscribe(req *SubscribeRequest, stream PubSub_SubscribeServer) error {

	// SubscribeTopic 增加一个使用函数过滤器的订阅者
	// func(v interface{}) 定义函数过滤的规则
	// SubscribeTopic 返回一个chan interface{}

	ch := p.pub.SubscribeTopic(func(v interface{}) bool {

		log.Printf("[服务中转]{订阅}主题：%v\n", v)

		// 接收数据是string，并且key是以arg为前缀的
		if key, ok := v.(string); ok {
			if strings.HasPrefix(key, req.GetValue()) {
				return true
			}
		}
		return false
	})

	log.Printf("[服务中转]{订阅}编号：%s 类型：%s 方法：%s ip地址：%s\n", req.GetId(), req.GetValue(), req.GetMethod(), req.GetIp())
	log.Println("[服务中转]{订阅}工作线：", ch)
	log.Println("[服务中转]{订阅}当前工作线数量：", p.pub.Len())

	// 服务器遍历chan，并将其中信息发送给订阅客户端
	for v := range ch {
		log.Println("[服务中转]{订阅}for ch：", ch)
		log.Println("[服务中转]{订阅}for v：", v)
		err := stream.Send(&SubscribeResponse{
			Id:     req.GetId(),
			Value:  req.GetValue(),
			Method: req.GetMethod(),
		})
		if err != nil {
			log.Println("[服务中转]{订阅}任务分配失败 ", err.Error())
			return err
		}
	}

	return nil
}
