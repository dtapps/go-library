package gojobs

import (
	"context"
	"go.etcd.io/etcd/client/v3"
)

// Watch 监听
func (e Etcd) Watch(ctx context.Context, key string, opts ...clientv3.OpOption) clientv3.WatchChan {
	return e.Client.Watch(ctx, key, opts...)
}

// Create 创建
func (e Etcd) Create(ctx context.Context, key, val string, opts ...clientv3.OpOption) (*clientv3.PutResponse, error) {
	return e.Client.Put(ctx, key, val, opts...)
}

// Get 获取
func (e Etcd) Get(ctx context.Context, key string, opts ...clientv3.OpOption) (*clientv3.GetResponse, error) {
	return e.Client.Get(ctx, key, opts...)
}

// Update 更新
func (e Etcd) Update(ctx context.Context, key, val string, opts ...clientv3.OpOption) (*clientv3.PutResponse, error) {
	return e.Client.Put(ctx, key, val, opts...)
}

// Delete 删除
func (e Etcd) Delete(ctx context.Context, key string, opts ...clientv3.OpOption) (*clientv3.DeleteResponse, error) {
	return e.Client.Delete(ctx, key, opts...)
}
