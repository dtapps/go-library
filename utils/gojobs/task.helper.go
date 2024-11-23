package gojobs

import (
	"context"
	"errors"
	"fmt"
	"github.com/redis/go-redis/v9"
	"go.dtapp.net/library/utils/gorequest"
	"log/slog"
	"strings"
	"time"
)

type TaskHelper struct {
	cfg *taskHelperConfig // 配置

	taskType string           // [任务]类型
	taskList []*GormModelTask // [任务]列表

	Ctx context.Context // [启动]上下文
}

// NewTaskHelper 任务帮助
// ctx 链路追踪的上下文
// taskType 任务类型
// logIsDebug 日志是否启动
// traceIsFilter 链路追踪是否过滤
func NewTaskHelper(ctx context.Context, taskType string, opts ...TaskHelperOption) *TaskHelper {
	th := &TaskHelper{
		taskType: taskType,
	}

	// 配置
	th.cfg = newTaskHelperConfig(opts)

	if gorequest.GetRequestIDContext(ctx) == "" {
		ctx = gorequest.SetRequestIDContext(ctx)
	}

	return th
}

// QueryTaskList 通过回调函数获取任务列表
// ctx 链路追踪的上下文
// isRunCallback 任务列表回调函数 返回 是否使用 任务列表
// listCallback 任务回调函数 返回 任务列表
// newTaskLists 新的任务列表
// isContinue 是否继续
func (th *TaskHelper) QueryTaskList(ctx context.Context, isRunCallback func(ctx context.Context, keyName string) (isUse bool, result *redis.StringCmd), listCallback func(ctx context.Context, taskType string) []*GormModelTask) (isContinue bool) {

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

	// 任务列表回调函数
	if listCallback != nil {
		// 执行
		taskLists := listCallback(ctx, th.taskType)

		// 判断任务类型是否一致
		for _, vTask := range taskLists {
			if vTask.Type == th.taskType {
				th.taskList = append(th.taskList, vTask)
			}
		}
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

// FilterTaskList 过滤任务列表
// ctx 链路追踪的上下文
// isMandatoryIp 强制当前ip
// specifyIp 指定Ip
// isContinue 是否继续
func (th *TaskHelper) FilterTaskList(ctx context.Context, isMandatoryIp bool, specifyIp string) (isContinue bool) {

	if specifyIp != "" {

		// 新的任务列表
		var newTaskLists []*GormModelTask

		// 解析指定IP
		specifyIp = gorequest.IpIs(specifyIp)

		// 循环判断 过滤指定IP
		for _, vTask := range th.taskList {

			vTask.SpecifyIP = gorequest.IpIs(vTask.SpecifyIP)

			// 强制只能是当前的IP
			if isMandatoryIp {
				// 进入强制性IP
				if vTask.SpecifyIP == specifyIp {
					// 进入强制性IP，可添加任务
					newTaskLists = append(newTaskLists, vTask)
					continue
				}
			}

			if vTask.SpecifyIP == "" {
				// 任务指定IP为空，可添加任务
				newTaskLists = append(newTaskLists, vTask)
				continue
			} else if vTask.SpecifyIP == SpecifyIpNull {
				// 任务指定Ip无限制，可添加任务
				newTaskLists = append(newTaskLists, vTask)
				continue
			} else {
				// 判断是否包含该IP
				specifyIpFind := strings.Contains(vTask.SpecifyIP, ",")
				if specifyIpFind {
					// 进入强制性多IP
					// 分割字符串
					parts := strings.Split(vTask.SpecifyIP, ",")
					for _, vv := range parts {
						if vv == specifyIp {
							// 进入强制性多IP，可添加任务
							newTaskLists = append(newTaskLists, vTask)
							continue
						}
					}
				} else {
					// 进入强制性单IP
					if vTask.SpecifyIP == specifyIp {
						// 进入强制性单IP，可添加任务
						newTaskLists = append(newTaskLists, vTask)
						continue
					}
				}
			}
		}

		// 设置任务列表
		th.taskList = newTaskLists
	}

	// 没有任务需要执行
	if len(th.taskList) <= 0 {
		if th.cfg.logIsDebug {
			slog.InfoContext(ctx, "FilterTaskList 没有任务需要执行")
		}
		return
	}

	return true
}

// GetTaskList 请在FilterTaskList之后获取任务列表
func (th *TaskHelper) GetTaskList() []*GormModelTask {
	return th.taskList
}

// RunMultipleTask 运行多个任务
// ctx 链路追踪的上下文
// wait 等待时间（秒）
// executionCallback 执行任务回调函数 返回 runCode=状态 runDesc=描述
// updateCallback 执行更新回调函数
// startCallback 开始任务回调函数
// endCallback 结束任务回调函数
func (th *TaskHelper) RunMultipleTask(ctx context.Context, wait int64, executionCallback func(ctx context.Context, task *GormModelTask) (runCode int, runDesc string), updateCallback func(ctx context.Context, task *GormModelTask, result *TaskHelperRunSingleTaskResponse), startCallback func(ctx context.Context, taskType string) (err error), endCallback func(ctx context.Context, taskType string)) {

	// 执行开始任务回调函数
	if startCallback != nil && endCallback != nil {
		err := startCallback(ctx, th.taskType)
		if err != nil {
			if th.cfg.logIsDebug {
				slog.DebugContext(ctx, fmt.Sprintf("开始任务回调函数返回错误，无法继续运行: %s", err))
			}
			return
		}
	}

	for _, vTask := range th.taskList {
		// 运行单个任务
		th.RunSingleTask(ctx, vTask, executionCallback, updateCallback)

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

type TaskHelperRunSingleTaskResponse struct {
	RunID   string `json:"run_id"`   // 运行编号
	RunCode int    `json:"run_code"` // 运行状态
	RunDesc string `json:"run_desc"` // 运行描述

	CostTime int64 `json:"cost_time"` // 消耗时长

	TraceID   string `json:"trace_id"`   // 追踪编号
	SpanID    string `json:"span_id"`    // 跨度编号
	RequestID string `json:"request_id"` // 请求编号
}

// RunSingleTask 运行单个任务
// ctx 链路追踪的上下文
// task 任务
// executionCallback 执行任务回调函数 返回 runCode=状态 runDesc=描述
// updateCallback 执行更新回调函数
func (th *TaskHelper) RunSingleTask(ctx context.Context, task *GormModelTask, executionCallback func(ctx context.Context, task *GormModelTask) (runCode int, runDesc string), updateCallback func(ctx context.Context, task *GormModelTask, result *TaskHelperRunSingleTaskResponse)) {

	// 任务回调函数
	if executionCallback != nil {

		// 需要返回的结构
		result := TaskHelperRunSingleTaskResponse{
			RequestID: gorequest.GetRequestIDContext(ctx),
		}

		// 执行
		result.RunCode, result.RunDesc = executionCallback(ctx, task)

		// 运行编号
		result.RunID = result.TraceID
		if result.RunID == "" {
			result.RunID = result.RequestID
			if result.RunID == "" {
				if th.cfg.logIsDebug {
					slog.ErrorContext(ctx, "RunSingleTask 上下文没有运行编号")
				}
				return
			}
		}

		// 执行更新回调函数
		if updateCallback != nil {
			updateCallback(ctx, task, &result)
		}

	}

	return
}

// EndRunTaskList 结束运行任务列表并停止OpenTelemetry链路追踪
func (th *TaskHelper) EndRunTaskList() {
	return
}
