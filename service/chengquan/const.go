package chengquan

import "fmt"

const (
	ApiUrl     = "https://api.chengquan.cn"
	ApiTestUrl = "http://test.api.chengquan.vip:11140"
)

const (
	LogTable = "chengquan"
)

const (
	version = "1.0.0"
)

func StateInfo[codeT int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | string](state codeT) string {
	stateString := fmt.Sprintf("%v", state)
	switch stateString {
	case "RECHARGE":
		return "充值中"
	case "SUCCESS":
		return "成功"
	case "FAILURE":
		return "失败"
	}
	return stateString
}

// CodeInfo 接口响应返回码
func CodeInfo[codeT int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | string](code codeT) string {
	codeString := fmt.Sprintf("%v", code)
	switch codeString {
	case "7000":
		return "请求通过"
	case "7001":
		return "请求参数错误"
	case "7002":
		return "请求超时"
	case "7003":
		return "商户账号不存在"
	case "7004":
		return "商户账号状态暂停或禁用"
	case "7005":
		return "签名错误"
	case "7006":
		return "请求IP有误"
	case "7007":
		return "查询不到手机卡号归属地"
	case "7008":
		return "查询不到产品"
	case "7009":
		return "产品已下架"
	case "7010":
		return "查询不到商户密价"
	case "7011":
		return "商户密价状态暂停"
	case "7012":
		return "请求订单号已存在"
	case "7013":
		return "请求订单入库异常"
	case "7014":
		return "商户账号余额不足"
	case "7015":
		return "商户扣款异常"
	case "7016":
		return "请求订单不存在"
	case "7017":
		return "加油卡卡号不正确"
	case "7019":
		return "产品库存不足"
	case "7777":
		return "系统异常"
	}
	return codeString
}
