package gojobs

import (
	"context"
	"fmt"
	"github.com/dtapps/go-library/utils/goip"
	"github.com/dtapps/go-library/utils/gojobs/jobs_gorm_model"
	"github.com/dtapps/go-library/utils/gotime"
	"github.com/dtapps/go-library/utils/gotrace_id"
	"strings"
)

// Filter 过滤
// ctx 上下文
// isMandatoryIp 强制当前ip
// specifyIp 指定Ip
// tasks 过滤前的数据
// newTasks 过滤后的数据
func (c *Client) Filter(ctx context.Context, isMandatoryIp bool, specifyIp string, tasks []jobs_gorm_model.Task, isPrint bool) (newTasks []jobs_gorm_model.Task) {
	c.Println(isPrint, fmt.Sprintf("【Filter入参】是强制性Ip：%v；指定Ip：%v；任务数量：%v", isMandatoryIp, specifyIp, len(tasks)))
	if specifyIp == "" {
		specifyIp = goip.IsIp(c.GetCurrentIp())
	} else {
		specifyIp = goip.IsIp(specifyIp)
	}
	c.Println(isPrint, fmt.Sprintf("【Filter入参】指定Ip重新解析：%v", specifyIp))
	for _, v := range tasks {
		c.Println(isPrint, fmt.Sprintf("【Filter入参】任务指定Ip解析前：%v", v.SpecifyIp))
		v.SpecifyIp = goip.IsIp(v.SpecifyIp)
		c.Println(isPrint, fmt.Sprintf("【Filter入参】任务指定Ip重新解析：%v", v.SpecifyIp))
		// 强制只能是当前的ip
		if isMandatoryIp {
			c.Println(isPrint, "【Filter入参】进入强制性Ip")
			if v.SpecifyIp == specifyIp {
				c.Println(isPrint, fmt.Sprintf("【Filter入参】进入强制性Ip 添加任务：%v", v.Id))
				newTasks = append(newTasks, v)
				continue
			}
		}
		if v.SpecifyIp == "" {
			c.Println(isPrint, fmt.Sprintf("【Filter入参】任务指定Ip为空 添加任务：%v", v.Id))
			newTasks = append(newTasks, v)
			continue
		} else if v.SpecifyIp == SpecifyIpNull {
			c.Println(isPrint, fmt.Sprintf("【Filter入参】任务指定Ip无限制 添加任务：%v", v.Id))
			newTasks = append(newTasks, v)
			continue
		} else {
			// 判断是否包含该ip
			specifyIpFind := strings.Contains(v.SpecifyIp, ",")
			if specifyIpFind {
				c.Println(isPrint, "【Filter入参】进入强制性多Ip")
				// 分割字符串
				parts := strings.Split(v.SpecifyIp, ",")
				for _, vv := range parts {
					if vv == specifyIp {
						c.Println(isPrint, fmt.Sprintf("【Filter入参】进入强制性多Ip 添加任务：%v", v.Id))
						newTasks = append(newTasks, v)
						continue
					}
				}
			} else {
				c.Println(isPrint, "【Filter入参】进入强制性单Ip")
				if v.SpecifyIp == specifyIp {
					newTasks = append(newTasks, v)
					c.Println(isPrint, fmt.Sprintf("【Filter入参】进入强制性单Ip 添加任务：%v", v.Id))
					continue
				}
			}
		}
	}
	return newTasks
}

// Run 运行
func (c *Client) Run(ctx context.Context, task jobs_gorm_model.Task, taskResultCode int, taskResultDesc string) {

	runId := gotrace_id.GetTraceIdContext(ctx)
	if runId == "" {
		if c.slog.status {
			c.slog.client.WithTraceId(ctx).Error("上下文没有跟踪编号")
		}
		return
	}

	c.GormTaskLogRecord(ctx, task, runId, taskResultCode, taskResultDesc)

	switch taskResultCode {
	case 0:
		err := c.EditTask(c.gormClient.GetDb(), task.Id).
			Select("run_id", "result", "next_run_time").
			Updates(jobs_gorm_model.Task{
				RunId:       runId,
				Result:      taskResultDesc,
				NextRunTime: gotime.Current().AfterSeconds(task.Frequency).Time,
			}).Error
		if err != nil {
			if c.slog.status {
				c.slog.client.WithTraceId(ctx).Error(fmt.Sprintf("保存失败：%s", err))
			}
		}
		return
	case CodeSuccess:
		// 执行成功
		err := c.EditTask(c.gormClient.GetDb(), task.Id).
			Select("status_desc", "number", "run_id", "updated_ip", "result", "next_run_time").
			Updates(jobs_gorm_model.Task{
				StatusDesc:  "执行成功",
				Number:      task.Number + 1,
				RunId:       runId,
				UpdatedIp:   c.config.systemOutsideIp,
				Result:      taskResultDesc,
				NextRunTime: gotime.Current().AfterSeconds(task.Frequency).Time,
			}).Error
		if err != nil {
			if c.slog.status {
				c.slog.client.WithTraceId(ctx).Error(fmt.Sprintf("保存失败：%s", err))
			}
		}
	case CodeEnd:
		// 执行成功、提前结束
		err := c.EditTask(c.gormClient.GetDb(), task.Id).
			Select("status", "status_desc", "number", "updated_ip", "result", "next_run_time").
			Updates(jobs_gorm_model.Task{
				Status:      TASK_SUCCESS,
				StatusDesc:  "结束执行",
				Number:      task.Number + 1,
				UpdatedIp:   c.config.systemOutsideIp,
				Result:      taskResultDesc,
				NextRunTime: gotime.Current().Time,
			}).Error
		if err != nil {
			if c.slog.status {
				c.slog.client.WithTraceId(ctx).Error(fmt.Sprintf("保存失败：%s", err))
			}
		}
	case CodeError:
		// 执行失败
		err := c.EditTask(c.gormClient.GetDb(), task.Id).
			Select("status_desc", "number", "run_id", "updated_ip", "result", "next_run_time").
			Updates(jobs_gorm_model.Task{
				StatusDesc:  "执行失败",
				Number:      task.Number + 1,
				RunId:       runId,
				UpdatedIp:   c.config.systemOutsideIp,
				Result:      taskResultDesc,
				NextRunTime: gotime.Current().AfterSeconds(task.Frequency).Time,
			}).Error
		if err != nil {
			if c.slog.status {
				c.slog.client.WithTraceId(ctx).Error(fmt.Sprintf("保存失败：%s", err))
			}
		}
	}

	if task.MaxNumber != 0 {
		if task.Number+1 >= task.MaxNumber {
			// 关闭执行
			err := c.EditTask(c.gormClient.GetDb(), task.Id).
				Select("status").
				Updates(jobs_gorm_model.Task{
					Status: TASK_TIMEOUT,
				}).Error
			if err != nil {
				if c.slog.status {
					c.slog.client.WithTraceId(ctx).Error(fmt.Sprintf("保存失败：%s", err))
				}
			}
		}
	}
	return
}
