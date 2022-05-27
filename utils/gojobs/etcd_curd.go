package gojobs

import (
	"context"
	"go.etcd.io/etcd/client/v3"
	"log"
)

// Watch 监听
func (e Etcd) Watch(ctx context.Context, key string, opts ...clientv3.OpOption) clientv3.WatchChan {
	log.Println("监听：", key)
	return e.Client.Watch(ctx, key, opts...)
}

// Create 创建
func (e Etcd) Create(ctx context.Context, key, val string, opts ...clientv3.OpOption) (*clientv3.PutResponse, error) {
	log.Println("创建：", key, val)
	return e.Client.Put(ctx, key, val, opts...)
}

// Get 获取
func (e Etcd) Get(ctx context.Context, key string, opts ...clientv3.OpOption) (*clientv3.GetResponse, error) {
	log.Println("获取：", key)
	return e.Client.Get(ctx, key, opts...)
}

// Update 更新
func (e Etcd) Update(ctx context.Context, key, val string, opts ...clientv3.OpOption) (*clientv3.PutResponse, error) {
	log.Println("更新：", key, val)
	return e.Client.Put(ctx, key, val, opts...)
}

// Delete 删除
func (e Etcd) Delete(ctx context.Context, key string, opts ...clientv3.OpOption) (*clientv3.DeleteResponse, error) {
	log.Println("删除：", key)
	return e.Client.Delete(ctx, key, opts...)
}
