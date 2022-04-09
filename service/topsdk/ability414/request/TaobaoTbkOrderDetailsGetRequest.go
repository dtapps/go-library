package request

type TaobaoTbkOrderDetailsGetRequest struct {
	/*
	   查询时间类型，1：按照订单淘客创建时间查询，2:按照订单淘客付款时间查询，3:按照订单淘客结算时间查询，4:按照订单更新时间； defalutValue��1    */
	QueryType *int64 `json:"query_type,omitempty" required:"false" `
	/*
	   位点，除第一页之外，都需要传递；前端原样返回。     */
	PositionIndex *string `json:"position_index,omitempty" required:"false" `
	/*
	   页大小，默认20，1~100 defalutValue��20    */
	PageSize *int64 `json:"page_size,omitempty" required:"false" `
	/*
	   推广者角色类型,2:二方，3:三方，不传，表示所有角色     */
	MemberType *int64 `json:"member_type,omitempty" required:"false" `
	/*
	   淘客订单状态，11-拍下未付款，12-付款，13-关闭，14-确认收货，3-结算成功;不传，表示所有状态     */
	TkStatus *int64 `json:"tk_status,omitempty" required:"false" `
	/*
	   订单查询结束时间，订单开始时间至订单结束时间，中间时间段日常要求不超过3个小时，但如618、双11、年货节等大促期间预估时间段不可超过20分钟，超过会提示错误，调用时请务必注意时间段的选择，以保证亲能正常调用！     */
	EndTime *string `json:"end_time" required:"true" `
	/*
	   订单查询开始时间     */
	StartTime *string `json:"start_time" required:"true" `
	/*
	   跳转类型，当向前或者向后翻页必须提供,-1: 向前翻页,1：向后翻页 defalutValue��1    */
	JumpType *int64 `json:"jump_type,omitempty" required:"false" `
	/*
	   第几页，默认1，1~100 defalutValue��1    */
	PageNo *int64 `json:"page_no,omitempty" required:"false" `
	/*
	   场景订单场景类型，1:常规订单，2:渠道订单，3:会员运营订单，默认为1 defalutValue��1    */
	OrderScene *int64 `json:"order_scene,omitempty" required:"false" `
}

func (s *TaobaoTbkOrderDetailsGetRequest) SetQueryType(v int64) *TaobaoTbkOrderDetailsGetRequest {
	s.QueryType = &v
	return s
}
func (s *TaobaoTbkOrderDetailsGetRequest) SetPositionIndex(v string) *TaobaoTbkOrderDetailsGetRequest {
	s.PositionIndex = &v
	return s
}
func (s *TaobaoTbkOrderDetailsGetRequest) SetPageSize(v int64) *TaobaoTbkOrderDetailsGetRequest {
	s.PageSize = &v
	return s
}
func (s *TaobaoTbkOrderDetailsGetRequest) SetMemberType(v int64) *TaobaoTbkOrderDetailsGetRequest {
	s.MemberType = &v
	return s
}
func (s *TaobaoTbkOrderDetailsGetRequest) SetTkStatus(v int64) *TaobaoTbkOrderDetailsGetRequest {
	s.TkStatus = &v
	return s
}
func (s *TaobaoTbkOrderDetailsGetRequest) SetEndTime(v string) *TaobaoTbkOrderDetailsGetRequest {
	s.EndTime = &v
	return s
}
func (s *TaobaoTbkOrderDetailsGetRequest) SetStartTime(v string) *TaobaoTbkOrderDetailsGetRequest {
	s.StartTime = &v
	return s
}
func (s *TaobaoTbkOrderDetailsGetRequest) SetJumpType(v int64) *TaobaoTbkOrderDetailsGetRequest {
	s.JumpType = &v
	return s
}
func (s *TaobaoTbkOrderDetailsGetRequest) SetPageNo(v int64) *TaobaoTbkOrderDetailsGetRequest {
	s.PageNo = &v
	return s
}
func (s *TaobaoTbkOrderDetailsGetRequest) SetOrderScene(v int64) *TaobaoTbkOrderDetailsGetRequest {
	s.OrderScene = &v
	return s
}

func (req *TaobaoTbkOrderDetailsGetRequest) ToMap() map[string]interface{} {
	paramMap := make(map[string]interface{})
	if req.QueryType != nil {
		paramMap["query_type"] = *req.QueryType
	}
	if req.PositionIndex != nil {
		paramMap["position_index"] = *req.PositionIndex
	}
	if req.PageSize != nil {
		paramMap["page_size"] = *req.PageSize
	}
	if req.MemberType != nil {
		paramMap["member_type"] = *req.MemberType
	}
	if req.TkStatus != nil {
		paramMap["tk_status"] = *req.TkStatus
	}
	if req.EndTime != nil {
		paramMap["end_time"] = *req.EndTime
	}
	if req.StartTime != nil {
		paramMap["start_time"] = *req.StartTime
	}
	if req.JumpType != nil {
		paramMap["jump_type"] = *req.JumpType
	}
	if req.PageNo != nil {
		paramMap["page_no"] = *req.PageNo
	}
	if req.OrderScene != nil {
		paramMap["order_scene"] = *req.OrderScene
	}
	return paramMap
}

func (req *TaobaoTbkOrderDetailsGetRequest) ToFileMap() map[string]interface{} {
	fileMap := make(map[string]interface{})
	return fileMap
}
