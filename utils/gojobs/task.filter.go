package gojobs

import (
	"context"
	"go.dtapp.net/library/utils/gorequest"
	"strings"
)

// Filter 过滤
// ctx 上下文
// isMandatoryIp 强制当前ip
// specifyIp 指定Ip
// tasks 过滤前的数据
// newTasks 过滤后的数据
func (c *Client) Filter(ctx context.Context, isMandatoryIp bool, specifyIp string, tasks []GormModelTask, isPrint bool) (newTasks []GormModelTask) {
	//c.Println(ctx, isPrint, fmt.Sprintf("【Filter入参】是强制性Ip：%v；指定Ip：%v；任务数量：%v", isMandatoryIp, specifyIp, len(tasks)))
	if specifyIp == "" {
		specifyIp = gorequest.IpIs(c.GetCurrentIp())
	} else {
		specifyIp = gorequest.IpIs(specifyIp)
	}
	//c.Println(ctx, isPrint, fmt.Sprintf("【Filter入参】指定Ip重新解析：%v", specifyIp))
	for _, v := range tasks {
		//c.Println(ctx, isPrint, fmt.Sprintf("【Filter入参】任务指定Ip解析前：%v", v.SpecifyIP))
		v.SpecifyIP = gorequest.IpIs(v.SpecifyIP)
		//c.Println(ctx, isPrint, fmt.Sprintf("【Filter入参】任务指定Ip重新解析：%v", v.SpecifyIP))
		// 强制只能是当前的ip
		if isMandatoryIp {
			//c.Println(ctx, isPrint, "【Filter入参】进入强制性Ip")
			if v.SpecifyIP == specifyIp {
				//c.Println(ctx, isPrint, fmt.Sprintf("【Filter入参】进入强制性Ip 添加任务：%v", v.ID))
				newTasks = append(newTasks, v)
				continue
			}
		}
		if v.SpecifyIP == "" {
			//c.Println(ctx, isPrint, fmt.Sprintf("【Filter入参】任务指定Ip为空 添加任务：%v", v.ID))
			newTasks = append(newTasks, v)
			continue
		} else if v.SpecifyIP == SpecifyIpNull {
			//c.Println(ctx, isPrint, fmt.Sprintf("【Filter入参】任务指定Ip无限制 添加任务：%v", v.ID))
			newTasks = append(newTasks, v)
			continue
		} else {
			// 判断是否包含该ip
			specifyIpFind := strings.Contains(v.SpecifyIP, ",")
			if specifyIpFind {
				//c.Println(ctx, isPrint, "【Filter入参】进入强制性多Ip")
				// 分割字符串
				parts := strings.Split(v.SpecifyIP, ",")
				for _, vv := range parts {
					if vv == specifyIp {
						//c.Println(ctx, isPrint, fmt.Sprintf("【Filter入参】进入强制性多Ip 添加任务：%v", v.ID))
						newTasks = append(newTasks, v)
						continue
					}
				}
			} else {
				//c.Println(ctx, isPrint, "【Filter入参】进入强制性单Ip")
				if v.SpecifyIP == specifyIp {
					newTasks = append(newTasks, v)
					//c.Println(ctx, isPrint, fmt.Sprintf("【Filter入参】进入强制性单Ip 添加任务：%v", v.ID))
					continue
				}
			}
		}
	}
	return newTasks
}
