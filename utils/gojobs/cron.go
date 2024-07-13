package gojobs

import (
	"context"
	"fmt"
	"github.com/robfig/cron/v3"
	"go.dtapp.net/library/utils/gotime"
	"log/slog"
)

type taskList struct {
	id   cron.EntryID
	name string
}

// Cron 定时任务管理器
type Cron struct {
	inner  *cron.Cron
	list   []taskList
	option struct {
		log bool
	}
}

// NewCron 创建一个定时任务管理器
func NewCron(opts ...CronOption) *Cron {
	c := &Cron{
		inner: cron.New(),
		list:  make([]taskList, 0),
	}
	for _, opt := range opts {
		opt(c)
	}
	return c
}

func NewCronWithSeconds(opts ...CronOption) *Cron {
	c := &Cron{
		inner: cron.New(cron.WithSeconds()),
		list:  make([]taskList, 0),
	}
	for _, opt := range opts {
		opt(c)
	}
	return c
}

// Start 启动任务
func (c *Cron) Start() {
	c.inner.Start()
}

// Stop 关闭任务
func (c *Cron) Stop() context.Context {
	return c.inner.Stop()
}

// AddFunc 添加任务
func (c *Cron) AddFunc(spec string, cmd func()) (cron.EntryID, error) {
	id, err := c.inner.AddFunc(spec, cmd)
	c.list = append(c.list, taskList{
		id: id,
	})
	return id, err
}

// AddJob 添加任务
func (c *Cron) AddJob(spec string, cmd cron.Job) (cron.EntryID, error) {
	id, err := c.inner.AddJob(spec, cmd)
	c.list = append(c.list, taskList{
		id: id,
	})
	return id, err
}

// Entry 查询任务
func (c *Cron) Entry(id cron.EntryID) cron.Entry {
	return c.inner.Entry(id)
}

// Remove 删除任务
func (c *Cron) Remove(id cron.EntryID) {
	c.inner.Remove(id)
}

// List 任务列表
func (c *Cron) List() []cron.EntryID {
	ids := make([]cron.EntryID, 0)
	for _, v := range c.list {
		ids = append(ids, v.id)
	}
	return ids
}

// ListShow 任务列表
func (c *Cron) ListShow() {
	for _, v := range c.list {
		taskInfo := c.inner.Entry(v.id)
		slog.Info(fmt.Sprintf("[ID=%v][Schedule=%v][Prev=%v][Next=%v]",
			taskInfo.ID,
			taskInfo.Schedule,
			taskInfo.Prev.Format(gotime.DateTimeZhFormat),
			taskInfo.Next.Format(gotime.DateTimeZhFormat),
		))
	}
}

// RunListShow 任务列表
func (c *Cron) RunListShow(spec string) {
	_, _ = c.AddFunc(spec, func() {
		c.ListShow()
	})
}

// AddTask 添加任务
func (c *Cron) AddTask(name string, spec string, cmd func()) (cron.EntryID, error) {
	id, err := c.inner.AddFunc(spec, cmd)
	c.list = append(c.list, taskList{
		id:   id,
		name: name,
	})
	return id, err
}

// QueryTask 查询任务
func (c *Cron) QueryTask(id cron.EntryID) cron.Entry {
	return c.inner.Entry(id)
}

// QueryInfo 查询信息
func (c *Cron) QueryInfo(id cron.EntryID) (cron.EntryID, string) {
	for _, v := range c.list {
		if id == v.id {
			return v.id, v.name
		}
	}
	return 0, ""
}

// RemoveTask 删除任务
func (c *Cron) RemoveTask(id cron.EntryID) {
	c.inner.Remove(id)
	c.logTask(id, "停止成功")
}

// PrintTask 日志任务
func (c *Cron) PrintTask(id cron.EntryID, content ...string) {
	c.logTask(id, content...)
}

func (c *Cron) logTask(id cron.EntryID, content ...string) {
	if c.option.log {
		for _, v := range c.list {
			if v.id == id {
				slog.Info(fmt.Sprintf("%s [ID=%v]%s\n", v.name, id, content))
			}
		}
	}
}

// PrintNameTask 日志任务
func (c *Cron) PrintNameTask(id cron.EntryID, name string, content ...string) {
	c.logNameTask(id, name, content...)
}

func (c *Cron) logNameTask(id cron.EntryID, name string, content ...string) {
	slog.Info(fmt.Sprintf("%s [ID=%v]%s\n", name, id, content))
}

// ListTask 任务列表
func (c *Cron) ListTask() {
	for _, v := range c.list {
		taskInfo := c.inner.Entry(v.id)
		slog.Info(fmt.Sprintf("%s [ID=%v][Schedule=%v][Prev=%v][Next=%v]",
			v.name,
			taskInfo.ID,
			taskInfo.Schedule,
			taskInfo.Prev.Format(gotime.DateTimeZhFormat),
			taskInfo.Next.Format(gotime.DateTimeZhFormat),
		))
	}
}

// RunListTask 任务列表
func (c *Cron) RunListTask(spec string) {
	_, _ = c.inner.AddFunc(spec, func() {
		c.ListTask()
	})
}
