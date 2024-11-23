package gojobs

import (
	"context"
	"github.com/redis/go-redis/v9"
	"go.dtapp.net/library/utils/gojson"
	"go.dtapp.net/library/utils/gorequest"
	"go.dtapp.net/library/utils/gostring"
	"log/slog"
	"sync"
	"time"
)

type PubSubClient struct {
	client               *redis.Client // Redis客户端
	taskTypeMu           sync.Mutex    // 互斥锁，用于保护 taskTypeExecutingMap
	taskTypeExecutingMap sync.Map      // 存储正在执行的任务类型
}

func NewPubSub(ctx context.Context, client *redis.Client) *PubSubClient {
	return &PubSubClient{
		client: client,
	}
}

// Publish 发布
func (c *PubSubClient) Publish(ctx context.Context, channel string, message any) error {
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
// message 任务信息，需要json编码
// executionCallback 执行任务回调函数 返回 runCode=状态 runDesc=描述
// updateCallback 执行更新回调函数
func (c *PubSubClient) DbRunSingleTask(ctx context.Context, message string, executionCallback func(ctx context.Context, task *GormModelTask) (runCode int, runDesc string), updateCallback func(ctx context.Context, task *GormModelTask, result *TaskHelperRunSingleTaskResponse) (err error)) {

	// 解析任务
	var task GormModelTask
	err := gojson.Unmarshal([]byte(message), &task)
	if err != nil {
		slog.ErrorContext(ctx, "[运行单个任务]解析错误",
			slog.String("err_desc", err.Error()),
		)
		return
	}

	// 任务回调函数
	if executionCallback != nil {

		// 需要返回的结构
		result := TaskHelperRunSingleTaskResponse{
			RequestID: gorequest.GetRequestIDContext(ctx),
		}

		// 开始时间
		start := time.Now().UTC()

		// 执行
		result.RunCode, result.RunDesc = executionCallback(ctx, &task)

		// 结束时间
		end := time.Now().UTC()

		// 运行编号
		result.RunID = result.TraceID
		if result.RunID == "" {
			result.RunID = result.RequestID
		}
		if result.RunID == "" {
			result.RunID = gostring.GetUuId()
		}

		// 消耗时长
		result.CostTime = end.Sub(start).Milliseconds()

		// 执行更新回调函数
		if updateCallback != nil {
			err = updateCallback(ctx, &task, &result)
			if err != nil {
				slog.ErrorContext(ctx, "[运行单个任务]回调函数返回错误",
					slog.Int64("task_id", int64(task.ID)),
					slog.String("task_type", task.Type),
					slog.String("task_custom_id", task.CustomID),
					slog.String("err_desc", err.Error()),
				)
				return
			}
		}

	}

	return
}

// DbRunSingleTaskMutex 运行单个任务带互斥锁
// ctx 链路追踪的上下文
// message 任务信息，需要json编码
// executionCallback 执行任务回调函数 返回 runCode=状态 runDesc=描述
// updateCallback 执行更新回调函数
func (c *PubSubClient) DbRunSingleTaskMutex(ctx context.Context, message string, executionCallback func(ctx context.Context, task *GormModelTask) (runCode int, runDesc string), updateCallback func(ctx context.Context, task *GormModelTask, result *TaskHelperRunSingleTaskResponse) (err error), errorCallback func(ctx context.Context, task *GormModelTask, desc string)) {

	// 解析任务
	var task GormModelTask
	err := gojson.Unmarshal([]byte(message), &task)
	if err != nil {
		errorCallback(ctx, &task, "[运行单个任务带互斥锁]解析错误")
		slog.ErrorContext(ctx, "[运行单个任务带互斥锁]解析错误",
			slog.String("err_desc", err.Error()),
		)
		return
	}

	// 检查任务类型是否已经在执行
	if _, ok := c.taskTypeExecutingMap.Load(task.Type); ok {
		errorCallback(ctx, &task, "[运行单个任务带互斥锁]任务类型已经在执行")
		slog.WarnContext(ctx, "[运行单个任务带互斥锁]任务类型已经在执行",
			slog.Int64("task_id", int64(task.ID)),
			slog.String("task_type", task.Type),
			slog.String("task_custom_id", task.CustomID),
		)
		return
	}

	// 标记任务类型为正在执行
	c.taskTypeExecutingMap.Store(task.Type, struct{}{})

	// 确保任务执行完毕后清理标记
	defer c.taskTypeExecutingMap.Delete(task.Type)

	// 任务回调函数
	if executionCallback != nil {

		// 需要返回的结构
		result := TaskHelperRunSingleTaskResponse{
			RequestID: gorequest.GetRequestIDContext(ctx),
		}

		// 开始时间
		start := time.Now().UTC()

		// 执行
		result.RunCode, result.RunDesc = executionCallback(ctx, &task)

		// 结束时间
		end := time.Now().UTC()

		// 运行编号
		result.RunID = result.TraceID
		if result.RunID == "" {
			result.RunID = result.RequestID
		}
		if result.RunID == "" {
			result.RunID = gostring.GetUuId()
		}

		// 消耗时长
		result.CostTime = end.Sub(start).Milliseconds()

		// 执行更新回调函数
		if updateCallback != nil {
			err = updateCallback(ctx, &task, &result)
			if err != nil {
				slog.ErrorContext(ctx, "[运行单个任务带互斥锁]回调函数返回错误",
					slog.Int64("task_id", int64(task.ID)),
					slog.String("task_type", task.Type),
					slog.String("task_custom_id", task.CustomID),
					slog.String("err_desc", err.Error()),
				)
				return
			}
		}

	}

	return
}

// DbRunSingleTaskMutexUseID 运行单个任务带互斥锁，使用ID编号
// ctx 链路追踪的上下文
// message 任务信息，需要json编码
// executionCallback 执行任务回调函数 返回 runCode=状态 runDesc=描述
// updateCallback 执行更新回调函数
func (c *PubSubClient) DbRunSingleTaskMutexUseID(ctx context.Context, message string, executionCallback func(ctx context.Context, task *GormModelTask) (runCode int, runDesc string), updateCallback func(ctx context.Context, task *GormModelTask, result *TaskHelperRunSingleTaskResponse) (err error), errorCallback func(ctx context.Context, task *GormModelTask, desc string)) {

	// 解析任务
	var task GormModelTask
	err := gojson.Unmarshal([]byte(message), &task)
	if err != nil {
		errorCallback(ctx, &task, "[运行单个任务带互斥锁，使用ID编号]解析错误")
		slog.ErrorContext(ctx, "[运行单个任务带互斥锁，使用ID编号]解析错误",
			slog.String("err_desc", err.Error()),
		)
		return
	}

	// 检查任务类型是否已经在执行
	if _, ok := c.taskTypeExecutingMap.Load(task.ID); ok {
		errorCallback(ctx, &task, "[运行单个任务带互斥锁，使用ID编号]任务类型已经在执行")
		slog.WarnContext(ctx, "[运行单个任务带互斥锁，使用ID编号]任务类型已经在执行",
			slog.Int64("task_id", int64(task.ID)),
			slog.String("task_type", task.Type),
			slog.String("task_custom_id", task.CustomID),
		)
		return
	}

	// 标记任务类型为正在执行
	c.taskTypeExecutingMap.Store(task.ID, struct{}{})

	// 确保任务执行完毕后清理标记
	defer c.taskTypeExecutingMap.Delete(task.ID)

	// 任务回调函数
	if executionCallback != nil {

		// 需要返回的结构
		result := TaskHelperRunSingleTaskResponse{
			RequestID: gorequest.GetRequestIDContext(ctx),
		}

		// 开始时间
		start := time.Now().UTC()

		// 执行
		result.RunCode, result.RunDesc = executionCallback(ctx, &task)

		// 结束时间
		end := time.Now().UTC()

		// 运行编号
		result.RunID = result.TraceID
		if result.RunID == "" {
			result.RunID = result.RequestID
		}
		if result.RunID == "" {
			result.RunID = gostring.GetUuId()
		}

		// 消耗时长
		result.CostTime = end.Sub(start).Milliseconds()

		// 执行更新回调函数
		if updateCallback != nil {
			err = updateCallback(ctx, &task, &result)
			if err != nil {
				slog.ErrorContext(ctx, "[运行单个任务带互斥锁，使用ID编号]回调函数返回错误",
					slog.Int64("task_id", int64(task.ID)),
					slog.String("task_type", task.Type),
					slog.String("task_custom_id", task.CustomID),
					slog.String("err_desc", err.Error()),
				)
				return
			}
		}

	}

	return
}

// DbRunSingleTaskMutexUseCustomID 运行单个任务带互斥锁，使用CustomID编号
// ctx 链路追踪的上下文
// message 任务信息，需要json编码
// executionCallback 执行任务回调函数 返回 runCode=状态 runDesc=描述
// updateCallback 执行更新回调函数
func (c *PubSubClient) DbRunSingleTaskMutexUseCustomID(ctx context.Context, message string, executionCallback func(ctx context.Context, task *GormModelTask) (runCode int, runDesc string), updateCallback func(ctx context.Context, task *GormModelTask, result *TaskHelperRunSingleTaskResponse) (err error), errorCallback func(ctx context.Context, task *GormModelTask, desc string)) {

	// 解析任务
	var task GormModelTask
	err := gojson.Unmarshal([]byte(message), &task)
	if err != nil {
		errorCallback(ctx, &task, "[运行单个任务带互斥锁，使用CustomID编号]解析错误")
		slog.ErrorContext(ctx, "[运行单个任务带互斥锁，使用CustomID编号]解析错误",
			slog.String("err_desc", err.Error()),
		)
		return
	}

	// 检查任务类型是否已经在执行
	if _, ok := c.taskTypeExecutingMap.Load(task.CustomID); ok {
		errorCallback(ctx, &task, "[运行单个任务带互斥锁，使用CustomID编号]任务类型已经在执行")
		slog.WarnContext(ctx, "[运行单个任务带互斥锁，使用CustomID编号]任务类型已经在执行",
			slog.Int64("task_id", int64(task.ID)),
			slog.String("task_type", task.Type),
			slog.String("task_custom_id", task.CustomID),
		)
		return
	}

	// 标记任务类型为正在执行
	c.taskTypeExecutingMap.Store(task.CustomID, struct{}{})

	// 确保任务执行完毕后清理标记
	defer c.taskTypeExecutingMap.Delete(task.CustomID)

	// 任务回调函数
	if executionCallback != nil {

		// 需要返回的结构
		result := TaskHelperRunSingleTaskResponse{
			RequestID: gorequest.GetRequestIDContext(ctx),
		}

		// 开始时间
		start := time.Now().UTC()

		// 执行
		result.RunCode, result.RunDesc = executionCallback(ctx, &task)

		// 结束时间
		end := time.Now().UTC()

		// 运行编号
		result.RunID = result.TraceID
		if result.RunID == "" {
			result.RunID = result.RequestID
		}
		if result.RunID == "" {
			result.RunID = gostring.GetUuId()
		}

		// 消耗时长
		result.CostTime = end.Sub(start).Milliseconds()

		// 执行更新回调函数
		if updateCallback != nil {
			err = updateCallback(ctx, &task, &result)
			if err != nil {
				slog.ErrorContext(ctx, "[运行单个任务带互斥锁，使用CustomID编号]回调函数返回错误",
					slog.Int64("task_id", int64(task.ID)),
					slog.String("task_type", task.Type),
					slog.String("task_custom_id", task.CustomID),
					slog.String("err_desc", err.Error()),
				)
				return
			}
		}

	}

	return
}
