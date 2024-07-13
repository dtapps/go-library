package gojobs

import (
	"context"
	"errors"
	"fmt"
	"github.com/redis/go-redis/v9"
	"go.dtapp.net/library/utils/gojson"
	"go.dtapp.net/library/utils/gorequest"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
	"log/slog"
	"time"
)

type TaskCustomHelper struct {
	cfg *taskHelperConfig // 配置

	taskType string                      // [任务]类型
	taskList []*TaskCustomHelperTaskList // [任务]列表

	Ctx  context.Context // [启动]上下文
	Span trace.Span      // [启动]链路追踪
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

	// 启动OpenTelemetry链路追踪
	th.Ctx, th.Span = NewTraceStartSpan(rootCtx, th.taskType)

	th.Span.SetAttributes(attribute.String("task.help.helper", "custom"))
	th.Span.SetAttributes(attribute.String("task.help.request_id", gorequest.GetRequestIDContext(th.Ctx)))

	th.Span.SetAttributes(attribute.String("task.new.type", th.taskType))

	th.Span.SetAttributes(attribute.Bool("task.cfg.logIsDebug", th.cfg.logIsDebug))
	th.Span.SetAttributes(attribute.Bool("task.cfg.traceIsFilter", th.cfg.traceIsFilter))
	th.Span.SetAttributes(attribute.String("task.cfg.traceIsFilterKeyName", th.cfg.traceIsFilterKeyName))
	th.Span.SetAttributes(attribute.String("task.cfg.traceIsFilterKeyValue", th.cfg.traceIsFilterKeyValue))

	return th
}

// QueryTaskList 通过回调函数获取任务列表
// rootCtx 链路追踪的上下文
// isRunCallback 任务列表回调函数 返回 是否使用 任务列表
// listCallback 任务回调函数 返回 任务列表
// newTaskLists 新的任务列表
// isContinue 是否继续
func (th *TaskCustomHelper) QueryTaskList(rootCtx context.Context, isRunCallback func(ctx context.Context, keyName string) (isUse bool, result *redis.StringCmd), listCallback func(ctx context.Context, taskType string) []*TaskCustomHelperTaskList) (isContinue bool) {

	// 启动OpenTelemetry链路追踪
	ctx, span := NewTraceStartSpan(rootCtx, "QueryTaskList")

	span.SetAttributes(attribute.String("task.help.helper", "custom"))
	span.SetAttributes(attribute.String("task.help.request_id", gorequest.GetRequestIDContext(ctx)))

	// 任务列表回调函数
	if isRunCallback != nil {

		// 执行
		isRunUse, isRunResult := isRunCallback(ctx, GetRedisKeyName(th.taskType))
		if isRunUse {
			if isRunResult.Err() != nil {
				if errors.Is(isRunResult.Err(), redis.Nil) {
					err := fmt.Errorf("执行任务列表回调函数返回不存在，无法继续运行: %v@%v", GetRedisKeyName(th.taskType), isRunResult.Err().Error())
					span.SetStatus(codes.Error, err.Error())

					if th.cfg.logIsDebug {
						slog.DebugContext(ctx, fmt.Sprintf("执行任务列表回调函数返回不存在，无法继续运行: %v@%v", GetRedisKeyName(th.taskType), isRunResult.Err().Error()))
					}

					// 过滤
					if th.cfg.traceIsFilter {
						span.SetAttributes(attribute.String(th.cfg.traceIsFilterKeyName, th.cfg.traceIsFilterKeyValue))
					}

					span.End() // 结束OpenTelemetry链路追踪
					return
				}

				err := fmt.Errorf("执行任务列表回调函数返回错误，无法继续运行: %v", isRunResult.Err().Error())
				span.SetStatus(codes.Error, err.Error())

				if th.cfg.logIsDebug {
					slog.DebugContext(ctx, fmt.Sprintf("执行任务列表回调函数返回错误，无法继续运行: %v", isRunResult.Err().Error()))
				}

				span.End() // 结束OpenTelemetry链路追踪
				return
			}
			if isRunResult.Val() == "" {
				err := fmt.Errorf("执行任务列表回调函数返回空，无法继续运行: %s", isRunResult.Val())
				span.SetStatus(codes.Error, err.Error())

				if th.cfg.logIsDebug {
					slog.DebugContext(ctx, fmt.Sprintf("执行任务列表回调函数返回空，无法继续运行: %s", isRunResult.Val()))
				}

				// 过滤
				if th.cfg.traceIsFilter {
					span.SetAttributes(attribute.String(th.cfg.traceIsFilterKeyName, th.cfg.traceIsFilterKeyValue))
				}

				span.End() // 结束OpenTelemetry链路追踪
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

		// 过滤
		if th.cfg.traceIsFilter {
			span.SetAttributes(attribute.String(th.cfg.traceIsFilterKeyName, th.cfg.traceIsFilterKeyValue))
		}

		span.End() // 结束OpenTelemetry链路追踪
		return
	}

	// OpenTelemetry链路追踪
	span.SetAttributes(attribute.String("task.list.list", gojson.JsonEncodeNoError(th.taskList)))
	span.SetAttributes(attribute.Int("task.list.count", len(th.taskList)))

	span.End() // 结束OpenTelemetry链路追踪
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
func (th *TaskCustomHelper) RunMultipleTask(rootCtx context.Context, wait int64, executionCallback func(ctx context.Context, task *TaskCustomHelperTaskList) (err error), startCallback func(ctx context.Context, taskType string) (err error), endCallback func(ctx context.Context, taskType string)) {

	// 启动OpenTelemetry链路追踪
	ctx, span := NewTraceStartSpan(rootCtx, "RunMultipleTask")

	span.SetAttributes(attribute.String("task.help.helper", "custom"))
	span.SetAttributes(attribute.String("task.help.request_id", gorequest.GetRequestIDContext(ctx)))

	span.SetAttributes(attribute.Int64("task.multiple.wait", wait))
	span.SetAttributes(attribute.String("task.multiple.list", gojson.JsonEncodeNoError(th.taskList)))
	span.SetAttributes(attribute.Int("task.multiple.count", len(th.taskList)))

	if th.cfg.logIsDebug {
		slog.DebugContext(ctx, "RunMultipleTask 运行多个任务", slog.Int64("wait", wait))
	}

	// 执行开始任务回调函数
	if startCallback != nil && endCallback != nil {
		err := startCallback(ctx, th.taskType)
		if err != nil {
			err = fmt.Errorf("开始任务回调函数返回错误，无法继续运行: %s", err)
			span.SetStatus(codes.Error, err.Error())

			if th.cfg.logIsDebug {
				slog.DebugContext(ctx, fmt.Sprintf("开始任务回调函数返回错误，无法继续运行: %s", err))
			}

			// 过滤
			if th.cfg.traceIsFilter {
				span.SetAttributes(attribute.String(th.cfg.traceIsFilterKeyName, th.cfg.traceIsFilterKeyValue))
			}

			span.End() // 结束OpenTelemetry链路追踪
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

	span.End() // 结束OpenTelemetry链路追踪
	return
}

// RunSingleTask 运行单个任务
// rootCtx 链路追踪的上下文
// task 任务
// executionCallback 执行任务回调函数
func (th *TaskCustomHelper) RunSingleTask(rootCtx context.Context, task *TaskCustomHelperTaskList, executionCallback func(ctx context.Context, task *TaskCustomHelperTaskList) (err error)) {

	// 启动OpenTelemetry链路追踪
	ctx, span := NewTraceStartSpan(rootCtx, "RunSingleTask "+task.CustomID)

	span.SetAttributes(attribute.String("task.help.helper", "custom"))
	span.SetAttributes(attribute.String("task.help.request_id", gorequest.GetRequestIDContext(ctx)))

	span.SetAttributes(attribute.String("task.single.info", gojson.JsonEncodeNoError(task)))

	if th.cfg.logIsDebug {
		slog.DebugContext(ctx, "RunSingleTask 运行单个任务", slog.String("task", gojson.JsonEncodeNoError(task)))
	}

	// 任务回调函数
	if executionCallback != nil {

		// 执行
		err := executionCallback(ctx, task)
		if err != nil {
			span.RecordError(err, trace.WithStackTrace(true))
			span.SetStatus(codes.Error, err.Error())
		}

		// OpenTelemetry链路追踪
		span.SetAttributes(attribute.String("task.info.id", task.TaskID))
		span.SetAttributes(attribute.String("task.info.name", task.TaskName))
		span.SetAttributes(attribute.String("task.info.params", task.TaskParams))
		span.SetAttributes(attribute.String("task.info.custom_id", task.CustomID))

	}

	span.End() // 结束OpenTelemetry链路追踪
	return
}

// EndRunTaskList 结束运行任务列表并停止OpenTelemetry链路追踪
func (th *TaskCustomHelper) EndRunTaskList() {
	if th.Span != nil {
		th.Span.End()
	}
	return
}
