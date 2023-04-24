package ejiaofei

import "fmt"

const (
	apiUrl = "http://api.ejiaofei.net:11140"
)

const (
	LogTable = "ejiaofei"
)

// OperatorInfo 运营商描述
func OperatorInfo(operator string) string {
	switch operator {
	case "mobile":
		return "移动"
	case "unicom":
		return "联通"
	case "telecom":
		return "电信"
	}
	return fmt.Sprintf("%v", operator)
}

// StateInfo 状态描述
func StateInfo[ST int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | string](state ST) string {
	stateString := fmt.Sprintf("%v", state)
	switch stateString {
	case "0":
		return "充值中"
	case "1":
		return "充值成功"
	case "2":
		return "充值失败"
	case "8":
		return "等待扣款"
	}
	return stateString
}

// ErrorInfo 错误描述
func ErrorInfo[ET int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | string](error ET) string {
	errorString := fmt.Sprintf("%v", error)
	switch errorString {
	case "0":
		return "无错误"
	case "1003":
		return "用户ID或接口密码错误"
	case "1004":
		return "用户IP错误"
	case "1005":
		return "用户接口已关闭"
	case "1006":
		return "加密结果错误"
	case "1007":
		return "订单号不存在"
	case "1011":
		return "号码归属地未知"
	case "1013":
		return "手机对应的商品有误或者没有上架"
	case "1014":
		return "无法找到手机归属地"
	case "1015":
		return "余额不足"
	case "1016":
		return "QQ号格式错误"
	case "1017":
		return "产品未分配用户，联系商务"
	case "1018":
		return "订单生成失败"
	case "1019":
		return "充值号码与产品不匹配"
	case "1020":
		return "号码运营商未知"
	case "9998":
		return "参数有误"
	case "9999":
		return "系统错误"
	}
	return errorString
}
