package gojobs

import (
	"go.etcd.io/etcd/client/v3"
	"time"
)

// EtcdConfig etcd配置
type EtcdConfig struct {
	Endpoints   []string      // 接口 []string{"http://127.0.0.1:2379"}
	DialTimeout time.Duration // time.Second * 5
	LocalIP     string        // 本机IP
	Username    string        // 用户名
	Password    string        // 密码
}

// Etcd etcd
type Etcd struct {
	EtcdConfig                  // 配置
	Client     *clientv3.Client // 驱动
	Kv         clientv3.KV      // kv API子集
	Lease      clientv3.Lease   // lease（租约）对象
	leaseId    clientv3.LeaseID // 租约编号
}

// Close 关闭
func (e Etcd) Close() {
	e.Client.Close()
}

const (
	// JobSaveDir 定时任务任务保存目录
	JobSaveDir = "/cron/jobs/"
	// JobWorkerDir 服务注册目录
	JobWorkerDir = "/cron/workers/"
)

// GetWatchKey 监听的key
func (e Etcd) GetWatchKey() string {
	return JobSaveDir + e.LocalIP
}

// IssueWatchKey 下发的key
func (e Etcd) IssueWatchKey(ip string) string {
	return JobSaveDir + ip
}
