package gojobs

import (
	"context"
	"errors"
	"fmt"
	"go.dtapp.net/library/goip"
	"go.etcd.io/etcd/client/v3"
	"log"
	"time"
)

// NewEtcdWorker 创建 etcd Worker
func NewEtcdWorker(config *EtcdConfig) (*Etcd, error) {

	var (
		e   = &Etcd{}
		err error
	)

	e.Endpoints = config.Endpoints
	e.DialTimeout = config.DialTimeout
	if config.LocalIP == "" {
		config.LocalIP = goip.GetOutsideIp()
	}
	e.LocalIP = config.LocalIP
	e.Username = config.Username
	e.Password = config.Password
	e.CustomDirectory = config.CustomDirectory

	v3Config := clientv3.Config{
		Endpoints:   e.Endpoints,
		DialTimeout: e.DialTimeout,
	}

	// 判断有没有配置用户信息
	if e.Username != "" {
		v3Config.Username = e.Username
		v3Config.Password = e.Password
	}

	e.Client, err = clientv3.New(v3Config)
	if err != nil {
		return nil, errors.New("连接失败：" + err.Error())
	}

	// 获得kv API子集
	e.Kv = clientv3.NewKV(e.Client)

	// 创建一个lease（租约）对象
	e.Lease = clientv3.NewLease(e.Client)

	// 注册服务
	go e.RegisterWorker()

	return e, nil
}

// RegisterWorker 注册worker
func (e Etcd) RegisterWorker() {
	var (
		regKey         string
		leaseGrantResp *clientv3.LeaseGrantResponse
		err            error
		keepAliveChan  <-chan *clientv3.LeaseKeepAliveResponse
		keepAliveResp  *clientv3.LeaseKeepAliveResponse
		cancelCtx      context.Context
		cancelFunc     context.CancelFunc
	)

	for {
		// 注册路径
		regKey = getJobWorkerDir(e) + e.LocalIP
		log.Println("租约：", regKey)

		cancelFunc = nil

		// 申请一个10秒的租约
		leaseGrantResp, err = e.Lease.Grant(context.TODO(), 10)
		if err != nil {
			log.Println("申请一个10秒的租约失败", err)
			goto RETRY
		}

		// 自动永久续租
		keepAliveChan, err = e.Lease.KeepAlive(context.TODO(), leaseGrantResp.ID)
		if err != nil {
			log.Println("自动永久续租失败", err)
			goto RETRY
		}

		cancelCtx, cancelFunc = context.WithCancel(context.TODO())

		// 注册到etcd
		_, err = e.Kv.Put(cancelCtx, regKey, "", clientv3.WithLease(leaseGrantResp.ID))
		if err != nil {
			log.Println(fmt.Sprintf(" %s 服务注册失败:%s", regKey, err))
			goto RETRY
		}

		// 处理续约应答的协程
		for {
			select {
			case keepAliveResp = <-keepAliveChan:
				if keepAliveResp == nil {
					log.Println("续租失败")
					goto RETRY
				} else {
					log.Println("收到自动续租应答:", leaseGrantResp.ID)
				}
			}
		}

	RETRY:
		log.Println("异常 RETRY ", regKey)
		time.Sleep(1 * time.Second)
		if cancelFunc != nil {
			cancelFunc()
		}
	}
}
