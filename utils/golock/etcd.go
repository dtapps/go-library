package golock

import (
	"context"
	"errors"
	"fmt"
	"go.etcd.io/etcd/client/v3"
)

type LockEtcd struct {
	etcdClient *clientv3.Client // 驱动
}

func NewLockEtcd(etcdClient *clientv3.Client) *LockEtcd {
	return &LockEtcd{etcdClient: etcdClient}
}

// Lock 上锁
// key 锁名
// val 锁内容
// ttl 锁过期时间
func (e *LockEtcd) Lock(key string, val string, ttl int64) (string, error) {
	if ttl <= 0 {
		return "", errors.New("长期请使用 LockForever 方法")
	}
	// 1、获取
	get, err := e.etcdClient.Get(context.Background(), key)
	if err != nil {
		return "", errors.New("获取异常")
	}
	if len(get.Kvs) > 0 {
		return "", errors.New("上锁失败，已存在")
	}
	// 2、申请一个lease(租约)
	lease := clientv3.NewLease(e.etcdClient)
	// 3、申请一个*秒的租约
	leaseGrantResp, err := lease.Grant(context.TODO(), ttl)
	if err != nil {
		return "", errors.New(fmt.Sprintf("申请租约失败 %s", err))
	}
	// 5、获得kv api子集
	kv := clientv3.NewKV(e.etcdClient)
	// 4、设置
	kv.Put(context.TODO(), key, val, clientv3.WithLease(leaseGrantResp.ID))
	if err != nil {
		return "", errors.New("上锁失败")
	}
	return val, nil
}

// Unlock 解锁
// key 锁名
func (e *LockEtcd) Unlock(key string) error {
	_, err := e.etcdClient.Delete(context.Background(), key)
	return err
}

// LockForever 永远上锁
// key 锁名
// val 锁内容
func (e *LockEtcd) LockForever(key string, val string) (string, error) {
	// 1、获取
	get, err := e.etcdClient.Get(context.Background(), key)
	if err != nil {
		return "", errors.New("获取异常")
	}
	if len(get.Kvs) > 0 {
		return "", errors.New("上锁失败，已存在")
	}
	// 2、设置
	e.etcdClient.Put(context.TODO(), key, val)
	if err != nil {
		return "", errors.New("上锁失败")
	}
	return val, nil
}
