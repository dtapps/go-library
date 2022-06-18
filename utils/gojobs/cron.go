package gojobs

import (
	"errors"
	"github.com/robfig/cron/v3"
	"sync"
)

// Cron 定时任务管理器
type Cron struct {
	inner *cron.Cron
	ids   map[string]cron.EntryID
	mutex sync.Mutex
}

// NewCron 创建一个定时任务管理器
func NewCron() *Cron {
	return &Cron{
		inner: cron.New(cron.WithSeconds()),
		ids:   make(map[string]cron.EntryID),
	}
}

// Start 启动任务
func (c *Cron) Start() {
	c.inner.Start()
}

// Stop 关闭任务
func (c *Cron) Stop() {
	c.inner.Stop()
}

// DelByID 删除任务
// id：唯一任务id
func (c *Cron) DelByID(id string) {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	eid, ok := c.ids[id]
	if !ok {
		return
	}
	c.inner.Remove(eid)
	delete(c.ids, id)
}

// AddJobByInterface 实现接口的方式添加定时任务
// id：唯一任务id
// spec：配置定时执行时间表达式
// cmd：需要执行的任务方法
func (c *Cron) AddJobByInterface(id string, spec string, cmd cron.Job) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if _, ok := c.ids[id]; ok {
		return errors.New("任务已存在")
	}
	eid, err := c.inner.AddJob(spec, cmd)
	if err != nil {
		return err
	}
	c.ids[id] = eid
	return nil
}

// AddJobByFunc 添加函数作为定时任务
// id：唯一任务id
// spec：配置定时执行时间表达式
// f：需要执行的任务方法
func (c *Cron) AddJobByFunc(id string, spec string, f func()) error {
	c.mutex.Lock()
	defer c.mutex.Unlock()

	if _, ok := c.ids[id]; ok {
		return errors.New("任务已存在")
	}
	eid, err := c.inner.AddFunc(spec, f)
	if err != nil {
		return err
	}
	c.ids[id] = eid
	return nil
}

// IsExistsJob 判断是否存在任务
// id：唯一任务id
func (c *Cron) IsExistsJob(id string) bool {
	_, exist := c.ids[id]
	return exist
}

// Ids ...
func (c *Cron) Ids() []string {
	c.mutex.Lock()
	defer c.mutex.Unlock()
	validIds := make([]string, 0, len(c.ids))
	invalidIds := make([]string, 0)
	for sid, eid := range c.ids {
		e := c.inner.Entry(eid)
		if e.ID != eid {
			invalidIds = append(invalidIds, sid)
			continue
		}
		validIds = append(validIds, sid)
	}
	for _, id := range validIds {
		delete(c.ids, id)
	}
	return validIds
}
