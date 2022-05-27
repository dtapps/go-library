package jobs

var ParamsOrderType = "order"

// ParamsOrderId 订单任务
type ParamsOrderId struct {
	OrderId string `json:"order_id,omitempty"`
}

var ParamsMerchantGoldenBeanType = "merchant.golden_bean"

var ParamsNewServiceType = "new_service"

// ParamsTaskId 企业自定义任务
type ParamsTaskId struct {
	TaskId int64 `json:"task_id,omitempty"`
}

var ParamsNewServiceNextType = "new_service.next"

// ParamsTaskIdNext 企业自定义下一步任务
type ParamsTaskIdNext struct {
	TaskId         int64 `json:"task_id,omitempty"`
	MerchantUserId int64 `json:"merchant_user_id,omitempty"`
	CurrentNumber  int   `json:"current_number,omitempty"`
	MaxNumber      int   `json:"max_number,omitempty"`
}

var ParamsTeamInvType = "team.inv"
