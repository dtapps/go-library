package gojobs

import (
	"context"
	"errors"
	"fmt"
	"github.com/dtapps/go-library/utils/gojobs/jobs_gorm_model"
	"github.com/dtapps/go-library/utils/gostring"
	"log"
	"math/rand"
	"time"
)

// GetIssueAddress 获取下发地址
// workers 在线列表
// v 任务信息
// ---
// address 下发地址
// err 错误信息
func (j *JobsGorm) GetIssueAddress(workers []string, v *jobs_gorm_model.Task) (address string, err error) {
	var (
		currentIp       = ""    // 当前Ip
		appointIpStatus = false // 指定Ip状态
	)

	// 赋值ip
	if v.SpecifyIp != "" {
		currentIp = v.SpecifyIp
		appointIpStatus = true
	}

	// 只有一个客户端在线
	if len(workers) == 1 {
		if appointIpStatus == true {
			// 判断是否指定某ip执行
			if gostring.Contains(workers[0], currentIp) == true {
				return j.config.cornKeyPrefix + "_" + v.SpecifyIp, nil
			}
			return address, errors.New(fmt.Sprintf("需要执行的[%s]客户端不在线", currentIp))
		}
		return j.config.cornKeyPrefix + "_" + workers[0], nil
	}

	// 优先处理指定某ip执行
	if appointIpStatus == true {
		for wk, wv := range workers {
			if gostring.Contains(wv, currentIp) == true {
				return j.config.cornKeyPrefix + "_" + workers[wk], nil
			}
		}
		return address, errors.New(fmt.Sprintf("需要执行的[%s]客户端不在线", currentIp))
	} else {
		// 随机返回一个
		zxIp := workers[j.random(0, len(workers))]
		if zxIp == "" {
			return address, errors.New("获取执行的客户端异常")
		}
		address = j.config.cornKeyPrefix + "_" + zxIp
		return address, err
	}
}

// GetSubscribeClientList 获取在线的客户端
func (j *JobsGorm) GetSubscribeClientList(ctx context.Context) ([]string, error) {

	if j.config.debug == true {
		log.Printf("获取在线的客户端：%s\n", j.config.cornKeyPrefix+"_*")
	}

	// 扫描
	values, err := j.redisClient.Keys(ctx, j.config.cornKeyPrefix+"_*").Result()
	if err != nil {
		if err == errors.New("ERR wrong number of arguments for 'mget' command") {
			return []string{}, nil
		}
		return nil, errors.New(fmt.Sprintf("获取失败：%s", err.Error()))
	}

	client := make([]string, 0, len(values))
	for _, val := range values {
		client = append(client, val.(string))
	}

	return client, nil
}

// 随机返回一个
// min最小
// max最大
func (j *JobsGorm) random(min, max int) int {
	if max-min <= 0 {
		return 0
	}
	rand.Seed(time.Now().Unix())
	return rand.Intn(max-min) + min
}
