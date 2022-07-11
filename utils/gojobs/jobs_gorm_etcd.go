package gojobs

import (
	"errors"
	"fmt"
	"go.dtapp.net/library/utils/gojobs/jobs_gorm_model"
	"math/rand"
	"time"
)

// GetEtcdIssueAddress 获取ETCD下发地址
func (j *JobsGorm) GetEtcdIssueAddress(server *Etcd, v *jobs_gorm_model.Task) (address string, err error) {
	var (
		currentIp       = ""    // 当前Ip
		appointIpStatus = false // 指定Ip状态
	)

	// 赋值ip
	if v.SpecifyIp != "" {
		currentIp = v.SpecifyIp
		appointIpStatus = true
	}

	workers, err := server.ListWorkers()
	if err != nil {
		return address, errors.New(fmt.Sprintf("获取在线客户端列表失败：%s", err.Error()))
	}
	if len(workers) <= 0 {
		return address, errors.New("没有客户端在线")
	}

	// 判断是否指定某ip执行
	if len(workers) == 1 {
		if appointIpStatus == true {
			if currentIp == workers[0] {
				return fmt.Sprintf("%s/%d", server.IssueWatchKey(v.SpecifyIp), v.Id), nil
			}
			return address, errors.New("执行的客户端不在线")
		}
	}

	// 随机返回一个
	zxIp := workers[j.random(0, len(workers))]
	if zxIp == "" {
		return address, errors.New("获取执行的客户端异常")
	}
	address = fmt.Sprintf("%s/%d", server.IssueWatchKey(zxIp), v.Id)
	return address, err
}

// 随机返回一个
func (j *JobsGorm) random(min, max int) int {
	if max-min <= 0 {
		return 0
	}
	rand.Seed(time.Now().Unix())
	return rand.Intn(max-min) + min
}
