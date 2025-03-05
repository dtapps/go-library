package gojobs

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/redis/go-redis/v9"
	"go.dtapp.net/library/utils/gorequest"
	"go.dtapp.net/library/utils/gostring"
	"go.dtapp.net/library/utils/gotime"
	"log/slog"
	"strings"
	"sync"
	"time"
)

type PubSubClient struct {
	client               *redis.Client // Redis客户端
	taskTypeMu           sync.Mutex    // 互斥锁，用于保护 taskTypeExecutingMap
	taskTypeExecutingMap sync.Map      // 存储正在执行的任务，值是设置时间
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

	var runName = "[运行单个任务]"

	// 开始时间
	start := time.Now().UTC()

	// 解析任务
	var task GormModelTask
	err := json.Unmarshal([]byte(message), &task)
	if err != nil {
		errDesc := fmt.Sprintf("%s解析错误", runName)
		slog.ErrorContext(ctx, errDesc,
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

		// 旧的任务规则
		oldSpec := task.Spec

		// 执行
		result.RunCode, result.RunDesc = executionCallback(ctx, &task)

		// 运行编号
		result.RunID = result.TraceID
		if result.RunID == "" {
			result.RunID = result.RequestID
		}
		if result.RunID == "" {
			result.RunID = gostring.GetUuId()
		}

		// 计算执行时间
		elapsed := time.Since(start)

		// 消耗时长
		result.CostTime = elapsed.Seconds()

		// 执行更新回调函数
		if updateCallback != nil {
			// 判断更新任务规则
			if oldSpec != task.Spec {
				result.UpdateSpec = task.Spec
			}
			err = updateCallback(ctx, &task, &result)
			if err != nil {
				errDesc := fmt.Sprintf("%s回调函数返回错误", runName)
				slog.ErrorContext(ctx, errDesc,
					slog.Uint64("task_id", uint64(task.ID)),
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

	var runName = "[运行单个任务带互斥锁]"

	// 开始时间
	start := time.Now().UTC()

	// 解析任务
	var task GormModelTask
	err := json.Unmarshal([]byte(message), &task)
	if err != nil {
		errDesc := fmt.Sprintf("%s解析错误", runName)
		errorCallback(ctx, &task, errDesc)
		slog.ErrorContext(ctx, errDesc,
			slog.String("err_desc", err.Error()),
		)
		return
	}

	// 自定义任务编号
	customTaskID := task.Type

	// 检查任务类型是否已经在执行
	if value, ok := c.taskTypeExecutingMap.Load(customTaskID); ok {
		errDesc := fmt.Sprintf("%s{%v}(%v)(%v)任务已经在执行，需要等待执行完才能继续", runName, customTaskID, value, gotime.Current().Time)
		errorCallback(ctx, &task, errDesc)
		slog.WarnContext(ctx, errDesc,
			slog.Uint64("task_id", uint64(task.ID)),
			slog.String("task_type", task.Type),
			slog.String("task_custom_id", task.CustomID),
		)
		return
	}

	// 标记任务类型为正在执行
	c.taskTypeExecutingMap.Store(customTaskID, gotime.Current().Time)

	// 确保任务执行完毕后清理标记
	defer c.taskTypeExecutingMap.Delete(customTaskID)

	// 任务回调函数
	if executionCallback != nil {

		// 需要返回的结构
		result := TaskHelperRunSingleTaskResponse{
			RequestID: gorequest.GetRequestIDContext(ctx),
		}

		// 旧的任务规则
		oldSpec := task.Spec

		// 执行
		result.RunCode, result.RunDesc = executionCallback(ctx, &task)

		// 运行编号
		result.RunID = result.TraceID
		if result.RunID == "" {
			result.RunID = result.RequestID
		}
		if result.RunID == "" {
			result.RunID = gostring.GetUuId()
		}

		// 计算执行时间
		elapsed := time.Since(start)

		// 消耗时长
		result.CostTime = elapsed.Seconds()

		// 执行更新回调函数
		if updateCallback != nil {
			// 判断更新任务规则
			if oldSpec != task.Spec {
				result.UpdateSpec = task.Spec
			}
			err = updateCallback(ctx, &task, &result)
			if err != nil {
				errDesc := fmt.Sprintf("%s{%v}回调函数返回错误", runName, customTaskID)
				slog.ErrorContext(ctx, errDesc,
					slog.Uint64("task_id", uint64(task.ID)),
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

	var runName = "[运行单个任务带互斥锁，使用ID编号]"

	// 开始时间
	start := time.Now().UTC()

	// 解析任务
	var task GormModelTask
	err := json.Unmarshal([]byte(message), &task)
	if err != nil {
		errDesc := fmt.Sprintf("%s解析错误", runName)
		errorCallback(ctx, &task, errDesc)
		slog.ErrorContext(ctx, errDesc,
			slog.String("err_desc", err.Error()),
		)
		return
	}

	// 自定义任务编号
	customTaskID := task.ID

	// 检查任务类型是否已经在执行
	if value, ok := c.taskTypeExecutingMap.Load(customTaskID); ok {
		errDesc := fmt.Sprintf("%s{%v}(%v)(%v)任务已经在执行，需要等待执行完才能继续", runName, customTaskID, value, gotime.Current().Time)
		errorCallback(ctx, &task, errDesc)
		slog.WarnContext(ctx, errDesc,
			slog.Uint64("task_id", uint64(task.ID)),
			slog.String("task_type", task.Type),
			slog.String("task_custom_id", task.CustomID),
		)
		return
	}

	// 标记任务类型为正在执行
	c.taskTypeExecutingMap.Store(customTaskID, gotime.Current().Time)

	// 确保任务执行完毕后清理标记
	defer c.taskTypeExecutingMap.Delete(customTaskID)

	// 任务回调函数
	if executionCallback != nil {

		// 需要返回的结构
		result := TaskHelperRunSingleTaskResponse{
			RequestID: gorequest.GetRequestIDContext(ctx),
		}

		// 旧的任务规则
		oldSpec := task.Spec

		// 执行
		result.RunCode, result.RunDesc = executionCallback(ctx, &task)

		// 运行编号
		result.RunID = result.TraceID
		if result.RunID == "" {
			result.RunID = result.RequestID
		}
		if result.RunID == "" {
			result.RunID = gostring.GetUuId()
		}

		// 计算执行时间
		elapsed := time.Since(start)

		// 消耗时长
		result.CostTime = elapsed.Seconds()

		// 执行更新回调函数
		if updateCallback != nil {
			// 判断更新任务规则
			if oldSpec != task.Spec {
				result.UpdateSpec = task.Spec
			}
			err = updateCallback(ctx, &task, &result)
			if err != nil {
				errDesc := fmt.Sprintf("%s{%v}回调函数返回错误", runName, customTaskID)
				slog.ErrorContext(ctx, errDesc,
					slog.Uint64("task_id", uint64(task.ID)),
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

	var runName = "[运行单个任务带互斥锁，使用CustomID编号]"

	// 开始时间
	start := time.Now().UTC()

	// 解析任务
	var task GormModelTask
	err := json.Unmarshal([]byte(message), &task)
	if err != nil {
		errDesc := fmt.Sprintf("%s解析错误", runName)
		errorCallback(ctx, &task, errDesc)
		slog.ErrorContext(ctx, errDesc,
			slog.String("err_desc", err.Error()),
		)
		return
	}

	// 自定义任务编号
	var builder strings.Builder
	builder.WriteString(task.Type)
	builder.WriteString(":")
	builder.WriteString(task.CustomID)
	customTaskID := builder.String()

	// 检查任务类型是否已经在执行
	if value, ok := c.taskTypeExecutingMap.Load(customTaskID); ok {
		errDesc := fmt.Sprintf("%s{%v}(%v)(%v)任务已经在执行，需要等待执行完才能继续", runName, customTaskID, value, gotime.Current().Time)
		errorCallback(ctx, &task, errDesc)
		slog.WarnContext(ctx, errDesc,
			slog.Uint64("task_id", uint64(task.ID)),
			slog.String("task_type", task.Type),
			slog.String("task_custom_id", task.CustomID),
		)
		return
	}

	// 标记任务类型为正在执行
	c.taskTypeExecutingMap.Store(customTaskID, gotime.Current().Time)

	// 确保任务执行完毕后清理标记
	defer c.taskTypeExecutingMap.Delete(customTaskID)

	// 任务回调函数
	if executionCallback != nil {

		// 需要返回的结构
		result := TaskHelperRunSingleTaskResponse{
			RequestID: gorequest.GetRequestIDContext(ctx),
		}

		// 旧的任务规则
		oldSpec := task.Spec

		// 执行
		result.RunCode, result.RunDesc = executionCallback(ctx, &task)

		// 运行编号
		result.RunID = result.TraceID
		if result.RunID == "" {
			result.RunID = result.RequestID
		}
		if result.RunID == "" {
			result.RunID = gostring.GetUuId()
		}

		// 计算执行时间
		elapsed := time.Since(start)

		// 消耗时长
		result.CostTime = elapsed.Seconds()

		// 执行更新回调函数
		if updateCallback != nil {
			// 判断更新任务规则
			if oldSpec != task.Spec {
				result.UpdateSpec = task.Spec
			}
			err = updateCallback(ctx, &task, &result)
			if err != nil {
				errDesc := fmt.Sprintf("%s{%v}回调函数返回错误", runName, customTaskID)
				slog.ErrorContext(ctx, errDesc,
					slog.Uint64("task_id", uint64(task.ID)),
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

// DbRunSingleTaskMutexUseCustomIDOrID 运行单个任务带互斥锁，使用CustomID编号或者ID编号
// ctx 链路追踪的上下文
// message 任务信息，需要json编码
// executionCallback 执行任务回调函数 返回 runCode=状态 runDesc=描述
// updateCallback 执行更新回调函数
func (c *PubSubClient) DbRunSingleTaskMutexUseCustomIDOrID(ctx context.Context, message string, executionCallback func(ctx context.Context, task *GormModelTask) (runCode int, runDesc string), updateCallback func(ctx context.Context, task *GormModelTask, result *TaskHelperRunSingleTaskResponse) (err error), errorCallback func(ctx context.Context, task *GormModelTask, desc string)) {

	var runName = "[运行单个任务带互斥锁，使用CustomID编号或者ID编号]"

	// 开始时间
	start := time.Now().UTC()

	// 解析任务
	var task GormModelTask
	err := json.Unmarshal([]byte(message), &task)
	if err != nil {
		errDesc := fmt.Sprintf("%s解析错误", runName)
		errorCallback(ctx, &task, errDesc)
		slog.ErrorContext(ctx, errDesc,
			slog.String("err_desc", err.Error()),
		)
		return
	}

	// 自定义任务编号
	var builder strings.Builder
	builder.WriteString(task.Type)
	builder.WriteString(":")
	builder.WriteString(fmt.Sprintf("%d", task.ID))
	builder.WriteString(task.CustomID)
	builder.WriteString(":")
	customTaskID := builder.String()

	// 检查任务类型是否已经在执行
	if value, ok := c.taskTypeExecutingMap.Load(customTaskID); ok {
		errDesc := fmt.Sprintf("%s{%v}(%v)(%v)任务已经在执行，需要等待执行完才能继续", runName, customTaskID, value, gotime.Current().Time)
		errorCallback(ctx, &task, errDesc)
		slog.WarnContext(ctx, errDesc,
			slog.Uint64("task_id", uint64(task.ID)),
			slog.String("task_type", task.Type),
			slog.String("task_custom_id", task.CustomID),
		)
		return
	}

	// 标记任务类型为正在执行
	c.taskTypeExecutingMap.Store(customTaskID, gotime.Current().Time)

	// 确保任务执行完毕后清理标记
	defer c.taskTypeExecutingMap.Delete(customTaskID)

	// 任务回调函数
	if executionCallback != nil {

		// 需要返回的结构
		result := TaskHelperRunSingleTaskResponse{
			RequestID: gorequest.GetRequestIDContext(ctx),
		}

		// 旧的任务规则
		oldSpec := task.Spec

		// 执行
		result.RunCode, result.RunDesc = executionCallback(ctx, &task)

		// 运行编号
		result.RunID = result.TraceID
		if result.RunID == "" {
			result.RunID = result.RequestID
		}
		if result.RunID == "" {
			result.RunID = gostring.GetUuId()
		}

		// 计算执行时间
		elapsed := time.Since(start)

		// 消耗时长
		result.CostTime = elapsed.Seconds()

		// 执行更新回调函数
		if updateCallback != nil {
			// 判断更新任务规则
			if oldSpec != task.Spec {
				result.UpdateSpec = task.Spec
			}
			err = updateCallback(ctx, &task, &result)
			if err != nil {
				errDesc := fmt.Sprintf("%s{%v}回调函数返回错误", runName, customTaskID)
				slog.ErrorContext(ctx, errDesc,
					slog.Uint64("task_id", uint64(task.ID)),
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
