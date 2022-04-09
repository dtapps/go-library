package request

type TaobaoTbkDgCpaActivityDetailRequest struct {
	/*
	   明细类型，1：预估明细，2：结算明细 defalutValue��1    */
	QueryType *int64 `json:"query_type,omitempty" required:"false" `
	/*
	   每页条数 defalutValue��10    */
	PageSize *int64 `json:"page_size,omitempty" required:"false" `
	/*
	   页码 defalutValue��1    */
	PageNo *int64 `json:"page_no,omitempty" required:"false" `
	/*
	   CPA活动ID     */
	EventId *int64 `json:"event_id" required:"true" `
	/*
	   CPA活动奖励的统计口径，相关说明见文档：https://www.yuque.com/docs/share/7ecf8cf1-7f99-4633-a2ed-f9b6f8116af5?#     */
	IndicatorAlias *string `json:"indicator_alias,omitempty" required:"false" `
}

func (s *TaobaoTbkDgCpaActivityDetailRequest) SetQueryType(v int64) *TaobaoTbkDgCpaActivityDetailRequest {
	s.QueryType = &v
	return s
}
func (s *TaobaoTbkDgCpaActivityDetailRequest) SetPageSize(v int64) *TaobaoTbkDgCpaActivityDetailRequest {
	s.PageSize = &v
	return s
}
func (s *TaobaoTbkDgCpaActivityDetailRequest) SetPageNo(v int64) *TaobaoTbkDgCpaActivityDetailRequest {
	s.PageNo = &v
	return s
}
func (s *TaobaoTbkDgCpaActivityDetailRequest) SetEventId(v int64) *TaobaoTbkDgCpaActivityDetailRequest {
	s.EventId = &v
	return s
}
func (s *TaobaoTbkDgCpaActivityDetailRequest) SetIndicatorAlias(v string) *TaobaoTbkDgCpaActivityDetailRequest {
	s.IndicatorAlias = &v
	return s
}

func (req *TaobaoTbkDgCpaActivityDetailRequest) ToMap() map[string]interface{} {
	paramMap := make(map[string]interface{})
	if req.QueryType != nil {
		paramMap["query_type"] = *req.QueryType
	}
	if req.PageSize != nil {
		paramMap["page_size"] = *req.PageSize
	}
	if req.PageNo != nil {
		paramMap["page_no"] = *req.PageNo
	}
	if req.EventId != nil {
		paramMap["event_id"] = *req.EventId
	}
	if req.IndicatorAlias != nil {
		paramMap["indicator_alias"] = *req.IndicatorAlias
	}
	return paramMap
}

func (req *TaobaoTbkDgCpaActivityDetailRequest) ToFileMap() map[string]interface{} {
	fileMap := make(map[string]interface{})
	return fileMap
}
