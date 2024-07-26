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
	"strings"
	"time"
)

type TaskHelper struct {
	cfg *taskHelperConfig // 配置

	taskType string           // [任务]类型
	taskList []*GormModelTask // [任务]列表

	Ctx  context.Context // [启动]上下文
	Span trace.Span      // [启动]链路追踪
}

// NewTaskHelper 任务帮助
// ctx 链路追踪的上下文
// taskType 任务类型
// logIsDebug 日志是否启动
// traceIsFilter 链路追踪是否过滤
func NewTaskHelper(rootCtx context.Context, taskType string, opts ...TaskHelperOption) *TaskHelper {
	th := &TaskHelper{
		taskType: taskType,
	}

	// 配置
	th.cfg = newTaskHelperConfig(opts)

	if gorequest.GetRequestIDContext(rootCtx) == "" {
		rootCtx = gorequest.SetRequestIDContext(rootCtx)
	}

	// 启动OpenTelemetry链路追踪
	th.Ctx, th.Span = NewTraceStartSpan(rootCtx, th.taskType)

	th.Span.SetAttributes(attribute.String("task.help.helper", "db"))
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
func (th *TaskHelper) QueryTaskList(rootCtx context.Context, isRunCallback func(ctx context.Context, keyName string) (isUse bool, result *redis.StringCmd), listCallback func(ctx context.Context, taskType string) []*GormModelTask) (isContinue bool) {

	// 启动OpenTelemetry链路追踪
	ctx, span := NewTraceStartSpan(rootCtx, "QueryTaskList")

	// 任务列表回调函数
	if isRunCallback != nil {

		// 执行
		isRunUse, isRunResult := isRunCallback(ctx, GetRedisKeyName(th.taskType))
		if isRunUse {
			if isRunResult.Err() != nil {
				if errors.Is(isRunResult.Err(), redis.Nil) {
					err := fmt.Errorf("执行任务列表回调函数返回不存在，无法继续运行: %v@%v", GetRedisKeyName(th.taskType), isRunResult.Err().Error())
					//th.listSpan.RecordError( err, trace.WithStackTrace(true))
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
				//th.listSpan.RecordError( err, trace.WithStackTrace(true))
				span.SetStatus(codes.Error, err.Error())

				if th.cfg.logIsDebug {
					slog.DebugContext(ctx, fmt.Sprintf("执行任务列表回调函数返回错误，无法继续运行: %v", isRunResult.Err().Error()))
				}

				span.End() // 结束OpenTelemetry链路追踪
				return
			}
			if isRunResult.Val() == "" {
				err := fmt.Errorf("执行任务列表回调函数返回空，无法继续运行: %s", isRunResult.Val())
				//th.listSpan.RecordError(err, trace.WithStackTrace(true))
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

// FilterTaskList 过滤任务列表
// rootCtx 链路追踪的上下文
// isMandatoryIp 强制当前ip
// specifyIp 指定Ip
// isContinue 是否继续
func (th *TaskHelper) FilterTaskList(rootCtx context.Context, isMandatoryIp bool, specifyIp string) (isContinue bool) {

	// 启动OpenTelemetry链路追踪
	ctx, span := NewTraceStartSpan(rootCtx, "FilterTaskList")

	span.SetAttributes(attribute.String("task.help.helper", "db"))
	span.SetAttributes(attribute.String("task.help.request_id", gorequest.GetRequestIDContext(ctx)))

	if th.cfg.logIsDebug {
		slog.DebugContext(ctx, "FilterTaskList 过滤任务列表", slog.Bool("isMandatoryIp", isMandatoryIp), slog.String("specifyIp", specifyIp))
	}

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

		// 过滤
		if th.cfg.traceIsFilter {
			span.SetAttributes(attribute.String(th.cfg.traceIsFilterKeyName, th.cfg.traceIsFilterKeyValue))
		}

		span.End() // 结束OpenTelemetry链路追踪
		return
	}

	// OpenTelemetry链路追踪
	span.SetAttributes(attribute.String("task.filter.list", gojson.JsonEncodeNoError(th.taskList)))
	span.SetAttributes(attribute.Int("task.filter.count", len(th.taskList)))

	span.End() // 结束OpenTelemetry链路追踪
	return true
}

// GetTaskList 请在FilterTaskList之后获取任务列表
func (th *TaskHelper) GetTaskList() []*GormModelTask {
	return th.taskList
}

// RunMultipleTask 运行多个任务
// rootCtx 链路追踪的上下文
// wait 等待时间（秒）
// executionCallback 执行任务回调函数 返回 runCode=状态 runDesc=描述
// updateCallback 执行更新回调函数
// startCallback 开始任务回调函数
// endCallback 结束任务回调函数
func (th *TaskHelper) RunMultipleTask(rootCtx context.Context, wait int64, executionCallback func(ctx context.Context, task *GormModelTask) (runCode int, runDesc string), updateCallback func(ctx context.Context, task *GormModelTask, result *TaskHelperRunSingleTaskResponse), startCallback func(ctx context.Context, taskType string) (err error), endCallback func(ctx context.Context, taskType string)) {

	// 启动OpenTelemetry链路追踪
	ctx, span := NewTraceStartSpan(rootCtx, "RunMultipleTask")

	span.SetAttributes(attribute.String("task.help.helper", "db"))
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

	span.End() // 结束OpenTelemetry链路追踪
	return
}

type TaskHelperRunSingleTaskResponse struct {
	RunID   string `json:"run_id"`   // 运行编号
	RunCode int    `json:"run_code"` // 运行状态
	RunDesc string `json:"run_desc"` // 运行描述

	TraceID   string `json:"trace_id"`   // 追踪编号
	SpanID    string `json:"span_id"`    // 跨度编号
	RequestID string `json:"request_id"` // 请求编号
}

// RunSingleTask 运行单个任务
// rootCtx 链路追踪的上下文
// task 任务
// executionCallback 执行任务回调函数 返回 runCode=状态 runDesc=描述
// updateCallback 执行更新回调函数
func (th *TaskHelper) RunSingleTask(rootCtx context.Context, task *GormModelTask, executionCallback func(ctx context.Context, task *GormModelTask) (runCode int, runDesc string), updateCallback func(ctx context.Context, task *GormModelTask, result *TaskHelperRunSingleTaskResponse)) {

	// 启动OpenTelemetry链路追踪
	ctx, span := NewTraceStartSpan(rootCtx, "RunSingleTask "+task.CustomID)

	span.SetAttributes(attribute.String("task.help.helper", "db"))
	span.SetAttributes(attribute.String("task.help.request_id", gorequest.GetRequestIDContext(ctx)))

	span.SetAttributes(attribute.String("task.single.info", gojson.JsonEncodeNoError(task)))

	if th.cfg.logIsDebug {
		slog.DebugContext(ctx, "RunSingleTask 运行单个任务", slog.String("task", gojson.JsonEncodeNoError(task)))
	}

	// 任务回调函数
	if executionCallback != nil {

		// 需要返回的结构
		result := TaskHelperRunSingleTaskResponse{
			TraceID:   span.SpanContext().TraceID().String(),
			SpanID:    span.SpanContext().SpanID().String(),
			RequestID: gorequest.GetRequestIDContext(ctx),
		}

		// 执行
		result.RunCode, result.RunDesc = executionCallback(ctx, task)
		if result.RunCode == CodeAbnormal {
			span.SetStatus(codes.Error, result.RunDesc)
		}
		if result.RunCode == CodeSuccess {
			span.SetStatus(codes.Ok, result.RunDesc)
		}
		if result.RunCode == CodeError {
			span.RecordError(fmt.Errorf(result.RunDesc), trace.WithStackTrace(true))
			span.SetStatus(codes.Error, result.RunDesc)
		}

		// 运行编号
		result.RunID = result.TraceID
		if result.RunID == "" {
			result.RunID = result.RequestID
			if result.RunID == "" {
				span.RecordError(fmt.Errorf("上下文没有运行编号"), trace.WithStackTrace(true))
				span.SetStatus(codes.Error, "上下文没有运行编号")

				if th.cfg.logIsDebug {
					slog.ErrorContext(ctx, "RunSingleTask 上下文没有运行编号")
				}

				span.End() // 结束OpenTelemetry链路追踪
				return
			}
		}

		// OpenTelemetry链路追踪
		span.SetAttributes(attribute.Int64("task.info.id", int64(task.ID)))
		span.SetAttributes(attribute.String("task.info.status", task.Status))
		span.SetAttributes(attribute.String("task.info.params", task.Params))
		span.SetAttributes(attribute.Int64("task.info.number", task.Number))
		span.SetAttributes(attribute.Int64("task.info.max_number", task.MaxNumber))
		span.SetAttributes(attribute.String("task.info.custom_id", task.CustomID))
		span.SetAttributes(attribute.Int64("task.info.custom_sequence", task.CustomSequence))
		span.SetAttributes(attribute.String("task.info.type", task.Type))
		span.SetAttributes(attribute.String("task.run.id", result.RunID))
		span.SetAttributes(attribute.Int("task.run.code", result.RunCode))
		span.SetAttributes(attribute.String("task.run.desc", result.RunDesc))

		// 执行更新回调函数
		if updateCallback != nil {
			updateCallback(ctx, task, &result)
		}

	}

	span.End() // 结束OpenTelemetry链路追踪
	return
}

// EndRunTaskList 结束运行任务列表并停止OpenTelemetry链路追踪
func (th *TaskHelper) EndRunTaskList() {
	if th.Span != nil {
		th.Span.End()
	}
	return
}
