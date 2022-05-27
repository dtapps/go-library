package gocache

import (
	"context"
	"errors"
	"go.etcd.io/etcd/client/v3"
	"time"
)

type EtcdConfig struct {
	Endpoints   []string      // 接口 []string{"http://127.0.0.1:2379"}
	DialTimeout time.Duration // time.Second * 5
}

type Etcd struct {
	EtcdConfig                  // 配置
	Client     *clientv3.Client // 驱动
	Kv         clientv3.KV      // kv API子集
	Lease      clientv3.Lease   // lease（租约）对象
	leaseId    clientv3.LeaseID // 租约编号
}

// NewEtcd 创建 etcd server
func NewEtcd(config *EtcdConfig) (*Etcd, error) {

	var (
		e   = &Etcd{}
		err error
	)

	e.Endpoints = config.Endpoints
	e.DialTimeout = config.DialTimeout

	e.Client, err = clientv3.New(clientv3.Config{
		Endpoints:   e.Endpoints,
		DialTimeout: e.DialTimeout,
	})
	if err != nil {
		return nil, errors.New("连接失败：" + err.Error())
	}

	// kv API子集
	e.Kv = clientv3.NewKV(e.Client)

	// 创建一个lease（租约）对象
	e.Lease = clientv3.NewLease(e.Client)

	return e, nil
}

// Set 设置
func (e Etcd) Set(ctx context.Context, key, val string, opts ...clientv3.OpOption) (*clientv3.PutResponse, error) {
	return e.Client.Put(ctx, key, val, opts...)
}

// Get 获取
func (e Etcd) Get(ctx context.Context, key string, opts ...clientv3.OpOption) (*clientv3.GetResponse, error) {
	return e.Client.Get(ctx, key, opts...)
}

// Del 删除
func (e Etcd) Del(ctx context.Context, key string, opts ...clientv3.OpOption) (*clientv3.DeleteResponse, error) {
	return e.Client.Delete(ctx, key, opts...)
}
