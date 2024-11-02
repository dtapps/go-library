package gojobs

import (
	"context"
	"errors"
	"fmt"
	"github.com/redis/go-redis/v9"
	"go.dtapp.net/library/utils/gojson"
	"go.dtapp.net/library/utils/gorequest"
	"log/slog"
	"time"
)

type TaskCustomHelper struct {
	cfg *taskHelperConfig // 配置

	taskType string                      // [任务]类型
	taskList []*TaskCustomHelperTaskList // [任务]列表

	Ctx context.Context // [启动]上下文
}

// NewTaskCustomHelper 任务帮助
// rootCtx 链路追踪的上下文
// taskType 任务类型
// logIsDebug 日志是否启动
// traceIsFilter 链路追踪是否过滤
func NewTaskCustomHelper(rootCtx context.Context, taskType string, opts ...TaskHelperOption) *TaskCustomHelper {
	th := &TaskCustomHelper{
		taskType: taskType,
	}

	// 配置
	th.cfg = newTaskHelperConfig(opts)

	if gorequest.GetRequestIDContext(rootCtx) == "" {
		rootCtx = gorequest.SetRequestIDContext(rootCtx)
	}

	return th
}

// QueryTaskList 通过回调函数获取任务列表
// rootCtx 链路追踪的上下文
// isRunCallback 任务列表回调函数 返回 是否使用 任务列表
// listCallback 任务回调函数 返回 任务列表
// newTaskLists 新的任务列表
// isContinue 是否继续
func (th *TaskCustomHelper) QueryTaskList(ctx context.Context, isRunCallback func(ctx context.Context, keyName string) (isUse bool, result *redis.StringCmd), listCallback func(ctx context.Context, taskType string) []*TaskCustomHelperTaskList) (isContinue bool) {

	// 任务列表回调函数
	if isRunCallback != nil {

		// 执行
		isRunUse, isRunResult := isRunCallback(ctx, GetRedisKeyName(th.taskType))
		if isRunUse {
			if isRunResult.Err() != nil {
				if errors.Is(isRunResult.Err(), redis.Nil) {
					if th.cfg.logIsDebug {
						slog.DebugContext(ctx, fmt.Sprintf("执行任务列表回调函数返回不存在，无法继续运行: %v@%v", GetRedisKeyName(th.taskType), isRunResult.Err().Error()))
					}
					return
				}
				if th.cfg.logIsDebug {
					slog.DebugContext(ctx, fmt.Sprintf("执行任务列表回调函数返回错误，无法继续运行: %v", isRunResult.Err().Error()))
				}
				return
			}
			if isRunResult.Val() == "" {
				if th.cfg.logIsDebug {
					slog.DebugContext(ctx, fmt.Sprintf("执行任务列表回调函数返回空，无法继续运行: %s", isRunResult.Val()))
				}
				return
			}
		}
	}

	// 执行任务列表回调函数
	if listCallback != nil {
		th.taskList = listCallback(ctx, th.taskType)
	}

	// 没有任务需要执行
	if len(th.taskList) <= 0 {
		if th.cfg.logIsDebug {
			slog.InfoContext(ctx, "QueryTaskList 没有任务需要执行")
		}
		return
	}

	return true
}

// GetTaskList 请在QueryTaskList之后获取任务列表
func (th *TaskCustomHelper) GetTaskList() []*TaskCustomHelperTaskList {
	return th.taskList
}

// RunMultipleTask 运行多个任务
// rootCtx 链路追踪的上下文
// wait 等待时间（秒）
// executionCallback 执行任务回调函数
// startCallback 开始任务回调函数
// endCallback 结束任务回调函数
func (th *TaskCustomHelper) RunMultipleTask(ctx context.Context, wait int64, executionCallback func(ctx context.Context, task *TaskCustomHelperTaskList) (err error), startCallback func(ctx context.Context, taskType string) (err error), endCallback func(ctx context.Context, taskType string)) {

	if th.cfg.logIsDebug {
		slog.DebugContext(ctx, "RunMultipleTask 运行多个任务", slog.Int64("wait", wait))
	}

	// 执行开始任务回调函数
	if startCallback != nil && endCallback != nil {
		err := startCallback(ctx, th.taskType)
		if err != nil {
			err = fmt.Errorf("开始任务回调函数返回错误，无法继续运行: %s", err)

			if th.cfg.logIsDebug {
				slog.DebugContext(ctx, fmt.Sprintf("开始任务回调函数返回错误，无法继续运行: %s", err))
			}

			return
		}
	}

	for _, vTask := range th.taskList {

		// 运行单个任务
		th.RunSingleTask(ctx, vTask, executionCallback)

		// 等待 wait 秒
		if wait > 0 {
			time.Sleep(time.Duration(wait) * time.Second)
		}

	}

	// 执行结束任务回调函数
	if startCallback != nil && endCallback != nil {
		endCallback(ctx, th.taskType)
	}

	return
}

// RunSingleTask 运行单个任务
// rootCtx 链路追踪的上下文
// task 任务
// executionCallback 执行任务回调函数
func (th *TaskCustomHelper) RunSingleTask(ctx context.Context, task *TaskCustomHelperTaskList, executionCallback func(ctx context.Context, task *TaskCustomHelperTaskList) (err error)) {

	if th.cfg.logIsDebug {
		slog.DebugContext(ctx, "RunSingleTask 运行单个任务", slog.String("task", gojson.JsonEncodeNoError(task)))
	}

	// 任务回调函数
	if executionCallback != nil {

		// 执行
		_ = executionCallback(ctx, task)

	}

	return
}

// EndRunTaskList 结束运行任务列表并停止OpenTelemetry链路追踪
func (th *TaskCustomHelper) EndRunTaskList() {
	return
}
