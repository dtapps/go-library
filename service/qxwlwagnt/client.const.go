package qxwlwagnt

type BaseResponse struct {
	Status  string `json:"status"`  // 接口调用状态码
	Message string `json:"message"` // 状态描述
}

type CommonResponse[T any] struct {
	BaseResponse   // 嵌入公共基础结构体
	Result       T `json:"result,omitempty"` // 返回数据，操作成功时有值，操作失败时为空（具体类型由 T 决定）
}

// []string{"can_test", "can_activate", "activate", "deactivated", "invalid", "clear", "replace", "stock", "start", "pre_clear"}
// []string{"可测试",    "可激活",        "已激活",    "已停用",       "已失效",   "已清除",  "已更换",  "库存",   "开始",   "预清除"}

// SIM卡状态
func GetStatus(simStatus string) string {
	// SIM卡状态 0：正常 1：待激活 3：单向停机 4：停机 5：过户 6：可测试 7：库存 8：管控停机 9：已销户 10：已激活 11：测试期 12：沉默期 99：其它
	switch simStatus {
	case "0":
		return "activate"
	case "1":
		return "can_activate"
	case "3":
		return "deactivated"
	case "4":
		return "deactivated"
	case "5":
		return "replace"
	case "6":
		return "can_test"
	case "7":
		return "stock"
	case "8":
		return "deactivated"
	case "9":
		return "invalid"
	case "10":
		return "activate"
	case "11":
		return "can_test"
	case "12":
		return "start"
	default:
		return "pre_clear"
	}
}

// SIM卡连接状态
func GetConnectStatus(gprsStatus string) string {
	// 网络状态 Y正常；F断网
	switch gprsStatus {
	case "Y":
		return "online"
	case "F":
		return "offline"
	default:
		return "unknown"
	}
}
