package gojobs

import (
	"context"
	"github.com/go-redis/redis/v9"
	"log"
)

// Publish 发布
// ctx 上下文
// channel 频道
// message 消息
func (j *JobsGorm) Publish(ctx context.Context, channel string, message interface{}) error {
	publish, err := j.redisClient.Publish(ctx, channel, message).Result()
	if j.config.debug == true {
		log.Println("gojobs.Publish", channel, message, publish, err)
	}
	return err
}

type SubscribeResult struct {
	err     error
	Message *redis.PubSub
}

// Subscribe 订阅
func (j *JobsGorm) Subscribe(ctx context.Context) SubscribeResult {
	return SubscribeResult{
		Message: j.redisClient.Subscribe(ctx, j.config.cornKeyPrefix+"_"+j.config.cornKeyCustom),
	}
}

// PSubscribe 订阅，支持通配符匹配(ch_user_*)
func (j *JobsGorm) PSubscribe(ctx context.Context) SubscribeResult {
	return SubscribeResult{
		Message: j.redisClient.PSubscribe(ctx, j.config.cornKeyPrefix+"_"+j.config.cornKeyCustom+"_*"),
	}
}
