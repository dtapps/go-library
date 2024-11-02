package gojobs

import (
	"context"
	"github.com/redis/go-redis/v9"
	"go.dtapp.net/library/utils/gojson"
	"go.dtapp.net/library/utils/gorequest"
	"log/slog"
)

type PubSubClient struct {
	client *redis.Client
}

func NewPubSub(ctx context.Context, client *redis.Client) *PubSubClient {
	return &PubSubClient{
		client: client,
	}
}

// Publish 发布
func (c *PubSubClient) Publish(ctx context.Context, channel string, message interface{}) error {
	return c.client.Publish(ctx, channel, message).Err()
}

// Subscribe 订阅
func (c *PubSubClient) Subscribe(ctx context.Context, channel ...string) *redis.PubSub {
	return c.client.Subscribe(ctx, channel...)
}

// PSubscribe 订阅，支持通配符匹配(ch_user_*)
func (c *PubSubClient) PSubscribe(ctx context.Context, channel ...string) *redis.PubSub {
	return c.client.PSubscribe(ctx, channel...)
}

// DbRunSingleTask 运行单个任务
// ctx 链路追踪的上下文
// message 任务json编码
// executionCallback 执行任务回调函数 返回 runCode=状态 runDesc=描述
// updateCallback 执行更新回调函数
func (c *PubSubClient) DbRunSingleTask(ctx context.Context, message string, executionCallback func(ctx context.Context, task *GormModelTask) (runCode int, runDesc string), updateCallback func(ctx context.Context, task *GormModelTask, result *TaskHelperRunSingleTaskResponse) (err error)) {

	// 解析任务
	var task GormModelTask
	err := gojson.Unmarshal([]byte(message), &task)
	if err != nil {
		slog.ErrorContext(ctx, "[DbRunSingleTask] json.Unmarshal",
			slog.String("err", err.Error()),
		)
		return
	}

	// 任务回调函数
	if executionCallback != nil {

		// 需要返回的结构
		result := TaskHelperRunSingleTaskResponse{
			RequestID: gorequest.GetRequestIDContext(ctx),
		}

		// 执行
		result.RunCode, result.RunDesc = executionCallback(ctx, &task)

		// 运行编号
		result.RunID = result.TraceID
		if result.RunID == "" {
			result.RunID = result.RequestID
			if result.RunID == "" {
				slog.ErrorContext(ctx, "[DbRunSingleTask] no run_id",
					slog.String("trace_id", result.TraceID),
					slog.String("request_id", result.RequestID),
					slog.String("run_id", result.RunID),
				)
				return
			}
		}

		// 执行更新回调函数
		if updateCallback != nil {
			err = updateCallback(ctx, &task, &result)
			if err != nil {
				slog.ErrorContext(ctx, "[DbRunSingleTask] updateCallback",
					slog.String("err", err.Error()),
				)
				return
			}
		}

	}

	return
}
